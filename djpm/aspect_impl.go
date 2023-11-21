package djpm

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/protobuf/proto"

	"github.com/artela-network/aspect-core/chaincoreext/scheduler"
	"github.com/artela-network/aspect-core/djpm/run"
	"github.com/artela-network/aspect-core/types"
)

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
	return aspect.transactionAdvice(types.ON_TX_VERIFY_METHOD, request)
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
