package djpm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/artela-network/aspect-core/chaincoreext/scheduler"
	"github.com/artela-network/aspect-core/djpm/run"
	"github.com/artela-network/aspect-core/types"
)

const JoinPointEnabledKey = "joinpoints"

var globalAspect *Aspect

type Aspect struct {
	provider types.AspectProvider
}

func NewAspect(provider types.AspectProvider) *Aspect {
	globalAspect = &Aspect{
		provider: provider,
	}
	scheduler.NewScheduleHost()
	types.GetScheduleHook = scheduler.GetScheduleHostApi

	return globalAspect
}

func AspectInstance() *Aspect {
	if globalAspect == nil {
		panic("aspcect instance not init,please exec NewAspect() first ")
	}
	return globalAspect
}

func (aspect Aspect) FilterTx(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.ON_TX_RECEIVE_METHOD, request)
}

func (aspect Aspect) VerifyTx(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.verification(types.ON_TX_VERIFY_METHOD, request)
}

//	func (aspect Aspect) VerifyAccount(request *types.EthTxAspect) *types.JoinPointResult {
//		return aspect.transactionAdvice(types.ON_ACCOUNT_VERIFY_METHOD, request)
//
// }
//
//	func (aspect Aspect) GetPayMaster(request *types.EthTxAspect) *types.JoinPointResult {
//		return aspect.transactionAdvice(types.ON_GAS_PAYMENT_METHOD, request)
//
// }
func (aspect Aspect) PreTxExecute(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.PRE_TX_EXECUTE_METHOD, request)
}

func (aspect Aspect) PreContractCall(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.PRE_CONTRACT_CALL_METHOD, request)
}

func (aspect Aspect) PostContractCall(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.POST_CONTRACT_CALL_METHOD, request)
}

func (aspect Aspect) PostTxExecute(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.POST_TX_EXECUTE_METHOD, request)
}

func (aspect Aspect) PostTxCommit(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.ON_TX_COMMIT_METHOD, request)
}

func (aspect Aspect) OnBlockInitialize(request *types.EthBlockAspect) *types.JoinPointResult {
	return aspect.blockAdvice(types.ON_BLOCK_INITIALIZE_METHOD, request)
}

func (aspect Aspect) OnBlockFinalize(request *types.EthBlockAspect) *types.JoinPointResult {
	return aspect.blockAdvice(types.ON_BLOCK_FINALIZE_METHOD, request)
}

func (aspect Aspect) GetSenderAndCallData(block int64, tx *types2.Transaction) (common.Address, []byte, error) {
	// transaction without a sig has different tx data encoding than the normal ethereum tx
	// the data is encoded as follows: abi.encode(validationData, callData)
	// validationData is the data that will be passed to the aspect verifier
	// callData is the data that will be passed to the contract
	validation, call, err := DecodeValidationAndCallData(tx.Data())
	if err != nil {
		return common.Address{}, nil, err
	}

	if block < 0 {
		block = aspect.provider.GetLatestBlock()
	}

	// check contract verifier
	verifiers, err := aspect.provider.GetAccountVerifiers(block, *tx.To())
	if err != nil {
		return common.Address{}, nil, err
	}

	if len(verifiers) != 1 {
		return common.Address{}, nil, errors.New(fmt.Sprintf(
			"invalid number of contract verifiers: %d",
			len(verifiers),
		))
	}

	contractVerifier := verifiers[0].AspectId

	request, err := aspect.provider.CreateTxPointRequestWithData(validation)
	if err != nil {
		return common.Address{}, nil, err
	}

	request.Tx = &types.EthTransaction{
		BlockHash:   nil,
		BlockNumber: block,
		Hash:        tx.Hash().Bytes(),
		Input:       tx.Data(),
		Nonce:       tx.Nonce(),
		To:          tx.To().Hex(),
		Value:       tx.Value().String(),
		Type:        int32(tx.Type()),
		ChainId:     tx.ChainId().String(),
	}

	// execute aspect verification
	verifyRes := aspect.VerifyTx(request)
	hasErr, err := verifyRes.HasErr()
	if hasErr {
		return common.Address{}, nil, err
	}

	resultMap := verifyRes.GetExecResultMap()
	for _, response := range resultMap {
		if response.Data == nil {
			return common.Address{}, nil, err
		}
		txResult := new(types.BytesData)
		anyData := response.Data
		if err := anyData.UnmarshalTo(txResult); err != nil {
			return common.Address{}, nil, err
		}
		if txResult.Data == nil {
			return common.Address{}, nil, err
		}

		sender := common.BytesToAddress(txResult.Data)

		// make sure sender accepts this aspect as verifier
		aspects, err := aspect.provider.GetAccountVerifiers(block, sender)
		if err != nil {
			return common.Address{}, nil, err
		}

		for _, aspect := range aspects {
			if aspect.AspectId == contractVerifier {
				return sender, call, nil
			}
		}
	}

	return common.Address{}, nil, errors.New("unable to verify tx with aspect")
}

func (aspect Aspect) blockAdvice(method types.PointCut, req *types.EthBlockAspect) *types.JoinPointResult {
	if req == nil || method == "" {
		return types.DefJoinPointResult("blockAdvice input is empty.")
	}
	aspectCodes, err := aspect.provider.GetBlockBondAspects(int64(req.Header.Number) - 1)
	if err != nil {
		return types.DefJoinPointResult("blockAdvice GetBlockBondAspects error." + err.Error())
	}
	// load aspects
	if len(aspectCodes) == 0 {
		return types.DefJoinPointResult("not bond aspects.")
	}
	// run aspects on received transaction

	return aspect.runAspect(method, req.GasInfo.Gas, int64(req.Header.Number), nil, req, aspectCodes)
}

func (aspect Aspect) transactionAdvice(method types.PointCut, req *types.EthTxAspect) *types.JoinPointResult {
	if req == nil || req.Tx == nil || types.IsAspectContract(req.Tx.To) {
		result := types.DefJoinPointResult("transactionAdvice invalid input.")
		result.GasInfo = req.GasInfo
		return result
	}
	if req.Tx.To == "" {
		result := types.DefJoinPointResult("it is create tx.")
		result.GasInfo = req.GasInfo
		return result
	}
	if len(req.Tx.Hash) != 0 {
		// skip scheduleTx
		txHash := common.BytesToHash(req.Tx.Hash)
		if nil != scheduler.TaskInstance() && scheduler.TaskInstance().IsScheduleTx(txHash) {
			result := types.DefJoinPointResult("it is schedule tx.")
			result.GasInfo = req.GasInfo
			return result
		}
	}
	// get binding contract address
	contractAddr := common.HexToAddress(req.Tx.To)
	if req.CurrInnerTx != nil && req.CurrInnerTx.To != "" {
		contractAddr = common.HexToAddress(req.CurrInnerTx.To)
	}
	aspectCodes, err := aspect.provider.GetTxBondAspects(req.GetTx().BlockNumber-1, contractAddr)
	if err != nil {
		result := types.DefJoinPointResult("transactionAdvice GetTxBondAspects error." + err.Error())
		result.GasInfo = req.GasInfo
		return result
	}
	if len(aspectCodes) == 0 {
		result := types.DefJoinPointResult("not bond aspects.")
		result.GasInfo = req.GasInfo
		return result
	}

	// run aspects on received transaction
	runAspect := aspect.runAspect(method, req.GasInfo.Gas, req.GetTx().BlockNumber, &contractAddr, req, aspectCodes)
	if len(req.Tx.Hash) != 0 {
		runAspect.TxHash = common.Bytes2Hex(req.Tx.Hash)
	}
	return runAspect
}

func (aspect Aspect) verification(method types.PointCut, req *types.EthTxAspect) *types.JoinPointResult {
	if req == nil || req.Tx == nil || types.IsAspectContract(req.Tx.To) {
		result := types.DefJoinPointResult("verification invalid input.")
		result.GasInfo = req.GasInfo
		return result
	}
	if req.Tx.To == "" {
		result := types.DefJoinPointResult("it is create tx.")
		result.GasInfo = req.GasInfo
		return result
	}
	if len(req.Tx.Hash) != 0 {
		// skip scheduleTx
		txHash := common.BytesToHash(req.Tx.Hash)
		if nil != scheduler.TaskInstance() && scheduler.TaskInstance().IsScheduleTx(txHash) {
			result := types.DefJoinPointResult("it is schedule tx.")
			result.GasInfo = req.GasInfo
			return result
		}
	}
	// get binding contract address
	contractAddr := common.HexToAddress(req.Tx.To)
	if req.CurrInnerTx != nil && req.CurrInnerTx.To != "" {
		contractAddr = common.HexToAddress(req.CurrInnerTx.To)
	}
	aspectCodes, err := aspect.provider.GetTxBondAspects(req.GetTx().BlockNumber-1, contractAddr)
	if err != nil {
		result := types.DefJoinPointResult("transactionAdvice GetTxBondAspects error." + err.Error())
		result.GasInfo = req.GasInfo
		return result
	}
	if len(aspectCodes) == 0 {
		result := types.DefJoinPointResult("not bond aspects.")
		result.GasInfo = req.GasInfo
		return result
	}

	// run aspects on received transaction
	runAspect := aspect.runAspect(method, req.GasInfo.Gas, req.GetTx().BlockNumber, &contractAddr, req, aspectCodes)
	if len(req.Tx.Hash) != 0 {
		runAspect.TxHash = common.Bytes2Hex(req.Tx.Hash)
	}
	return runAspect
}

func (aspect Aspect) runAspect(method types.PointCut, gas uint64, blockNumber int64, contractAddr *common.Address, reqData proto.Message, req []*types.AspectCode) (response *types.JoinPointResult) {
	aspectId := ""
	defer func() {
		if err := recover(); err != nil {
			// TODO log.Error(running aspect failed")
			response.WithErr(aspectId, errors.New("fatal: panic in running aspect"))
		}
	}()

	response = &types.JoinPointResult{}

	gasLeft := gas
	for _, aspect := range req {
		aspectId = aspect.AspectId
		if !isJoinPointEnabled(aspectId, blockNumber, method) {
			response.WithGas(0, 0, gasLeft)
			return response
		}

		runner, err := run.NewRunner(aspectId, aspect.Code)
		if err != nil {
			return response.WithErr(aspectId, err)
		}

		if res, callErr := runner.JoinPoint(method, gasLeft, blockNumber, contractAddr, reqData); callErr != nil {
			response.WithErr(aspectId, callErr)
		} else {
			response.WithResponse(aspectId, res)
		}

		gasLeft = runner.Gas()

		runner.Return()

		if hasErr, _ := response.HasErr(); hasErr {
			// short-circuit Aspect call
			totalGasUsed := gas - gasLeft
			response.WithGas(totalGasUsed, totalGasUsed, gasLeft)
			return response
		}
	}

	totalGasUsed := gas - gasLeft
	response.WithGas(totalGasUsed, totalGasUsed, gasLeft)

	return response
}

func DecodeValidationAndCallData(txData []byte) (validationData, callData []byte, err error) {
	validationData, err = loadParamBytes(txData, 0)
	if err != nil {
		return
	}

	callData, err = loadParamBytes(txData, 1)
	return
}

func loadParamBytes(input []byte, index int) ([]byte, error) {
	dataOffset, overflow := uint256.NewInt(0).SetBytes32(input[index*32 : (index+1)*32]).Uint64WithOverflow()
	if overflow {
		return nil, errors.New("invalid offset")
	}

	start := dataOffset + 32
	if start > uint64(len(input)) {
		return nil, errors.New("invalid param length")
	}

	dataLen, overflow := uint256.NewInt(0).SetBytes32(input[dataOffset:start]).Uint64WithOverflow()
	if overflow {
		return nil, errors.New("invalid length")
	}

	end := start + dataLen
	if end > uint64(len(input)) {
		return nil, errors.New("invalid param length")
	}

	return input[start:end], nil
}

func isJoinPointEnabled(aspectID string, blockNum int64, method types.PointCut) bool {
	hook, err := types.GetRuntimeHostHook()
	if err != nil || hook == nil {
		return false
	}

	// Convert aspectID to address
	aspID := common.HexToAddress(aspectID)

	// Set query data of property
	valueData := &types.StringData{Data: JoinPointEnabledKey}
	anyData, _ := anypb.New(valueData)

	// Build the requst struct
	request := &types.ContextQueryRequest{
		NameSpace: types.QueryNameSpace_QueryAspectProperty,
		Query:     anyData,
	}
	res := hook.Query(&types.RunnerContext{
		AspectId:    &aspID,
		BlockNumber: blockNum,
	}, request)

	// enables, err := anypb.UnmarshalNew(res.Data, )
	enableMsg, err := res.Data.UnmarshalNew()
	if err != nil {
		return false
	}

	enableStr := enableMsg.(*types.StringData).Data
	// If no property is set to control the enabled joinpoints, all joinpoints are considered open.
	// This mechanism ensures compatibility with aspects deployed before enabling this property.
	// Enabling all joinpoints obviously leads to the execution of blank joinpoints, resulting in higher gas consumption.
	// If there is no need to enable all aspects, set it to enable the specified aspects and separate them with commas.
	if len(strings.Trim(enableStr, " ")) == 0 {
		return true
	}

	return slices.Contains(strings.Split(enableStr, ","), string(method))
}
