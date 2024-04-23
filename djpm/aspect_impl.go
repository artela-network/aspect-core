package djpm

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	runtime "github.com/artela-network/aspect-runtime/types"
	"runtime/debug"

	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/artela-network/aspect-core/djpm/run"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"

	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/protobuf/proto"

	"github.com/artela-network/aspect-core/types"
)

var (
	CustomVerificationPrefix = hexutil.MustDecode("0xCAFECAFE")
)

var globalAspect *Aspect

type Aspect struct {
	provider types.AspectProvider
	logger   runtime.Logger
}

func NewAspect(provider types.AspectProvider, logger runtime.Logger) *Aspect {
	globalAspect = &Aspect{
		provider: provider,
		logger:   logger,
	}
	return globalAspect
}

func AspectInstance() *Aspect {
	if globalAspect == nil {
		panic("aspcect instance not init,please exec NewAspect() first ")
	}
	return globalAspect
}

func (aspect Aspect) VerifyTx(ctx context.Context, contract *common.Address, block int64, gas uint64, request *types.TxVerifyInput) *types.AspectExecutionResult {
	return aspect.verification(ctx, contract, block, gas, request)
}

func (aspect Aspect) PreTxExecute(ctx context.Context, contract *common.Address, block int64, gas uint64, request *types.PreTxExecuteInput) *types.AspectExecutionResult {
	return aspect.transactionAdvice(ctx, types.PRE_TX_EXECUTE_METHOD, contract, block, gas, request)
}

func (aspect Aspect) PreContractCall(ctx context.Context, contract *common.Address, block int64, gas uint64, request *types.PreContractCallInput) *types.AspectExecutionResult {
	return aspect.transactionAdvice(ctx, types.PRE_CONTRACT_CALL_METHOD, contract, block, gas, request)
}

func (aspect Aspect) PostContractCall(ctx context.Context, contract *common.Address, block int64, gas uint64, request *types.PostContractCallInput) *types.AspectExecutionResult {
	return aspect.transactionAdvice(ctx, types.POST_CONTRACT_CALL_METHOD, contract, block, gas, request)
}

func (aspect Aspect) PostTxExecute(ctx context.Context, contract *common.Address, block int64, gas uint64, request *types.PostTxExecuteInput) *types.AspectExecutionResult {
	return aspect.transactionAdvice(ctx, types.POST_TX_EXECUTE_METHOD, contract, block, gas, request)
}

func (aspect Aspect) GetSenderAndCallData(ctx context.Context, block int64, tx *ethtypes.Transaction) (common.Address, []byte, error) {
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
	verifiers, err := aspect.provider.GetAccountVerifiers(ctx, *tx.To())
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
	uintBlock := uint64(block)
	request := &types.TxVerifyInput{
		Tx: &types.NoFromTxInput{
			Hash: tx.Hash().Bytes(),
			To:   tx.To().Bytes(),
		},
		Block:          &types.BlockInput{Number: &uintBlock},
		ValidationData: validation,
		CallData:       call,
	}

	// execute aspect verification
	verifyRes := aspect.VerifyTx(ctx, tx.To(), block, tx.Gas(), request)
	if verifyRes.Err != nil {
		return common.Address{}, nil, verifyRes.Err
	}

	sender := common.BytesToAddress(verifyRes.Ret)

	// make sure sender accepts this aspect as verifier
	aspects, err := aspect.provider.GetAccountVerifiers(ctx, sender)
	if err != nil {
		return common.Address{}, nil, err
	}

	for _, aspect := range aspects {
		if aspect.AspectId == contractVerifier {
			return sender, call, nil
		}
	}

	return common.Address{}, nil, errors.New("unable to verify tx with aspect")
}

func (aspect Aspect) transactionAdvice(ctx context.Context, method types.PointCut, contract *common.Address, block int64, gas uint64, request proto.Message) *types.AspectExecutionResult {
	result := &types.AspectExecutionResult{
		Gas:    gas,
		Revert: types.NotRevert,
	}

	if contract == nil {
		// pass on contract creation call
		return result
	}

	// get binding contract address
	aspectCodes, err := aspect.provider.GetTxBondAspects(ctx, *contract, method)
	if err != nil {
		result.Err = err
		result.Revert = types.RevertCall
		return result
	}
	if len(aspectCodes) == 0 {
		return result
	}

	// run aspects on received transaction
	return aspect.runAspect(ctx, method, gas, block, contract, request, aspectCodes)
}

func (aspect Aspect) verification(ctx context.Context, contract *common.Address, block int64, gas uint64, req *types.TxVerifyInput) *types.AspectExecutionResult {
	if contract == nil {
		// not able to verify contract creation tx
		return &types.AspectExecutionResult{
			Gas:    gas,
			Err:    errors.New("not able to verify contract creation tx"),
			Revert: types.RevertTx,
		}
	}

	aspectCodes, err := aspect.provider.GetAccountVerifiers(ctx, *contract)
	if err != nil || len(aspectCodes) == 0 {
		return &types.AspectExecutionResult{
			Gas:    gas,
			Err:    errors.New("contract has not bound with any verifier aspect"),
			Revert: types.RevertTx,
		}
	}

	// run aspects on received transaction
	return aspect.runAspect(ctx, types.VERIFY_TX, gas, block, contract, req, aspectCodes)
}

func (aspect Aspect) runAspect(ctx context.Context, method types.PointCut, gas uint64, blockNumber int64, contractAddr *common.Address, reqData proto.Message, aspects []*types.AspectCode) (result *types.AspectExecutionResult) {
	result = &types.AspectExecutionResult{
		Gas:    gas,
		Revert: types.NotRevert,
	}
	defer func() {
		if err := recover(); err != nil {
			aspect.logger.Error("panic in running aspect", "err", err, "stack", debug.Stack())
			result.Err = errors.New("fatal: panic in running aspect: " + fmt.Sprintln(err))
			result.Revert = types.RevertCall
		}
	}()

	for _, storedAspect := range aspects {
		var err error
		isCommit := types.IsCommit(ctx)
		runner, err := run.NewRunner(ctx, aspect.logger, storedAspect.AspectId, storedAspect.Version, storedAspect.Code, isCommit)
		if err != nil {
			panic(err)
		}

		var ret []byte
		ret, gas, err = runner.JoinPoint(method, gas, blockNumber, contractAddr, reqData)
		runner.Return()

		result.Ret = ret
		if err != nil {
			result.Err = err
			result.Revert = types.RevertCall
			break
		}
	}

	result.Gas = gas

	return result
}

func DecodeValidationAndCallData(txData []byte) (validationData, callData []byte, err error) {
	// the customized data layout will be [4B Header][4B Checksum][NB ABI.Encode(ValidationData, CallData)]
	if len(txData) < 8 {
		return nil, nil, errors.New("invalid validation data")
	}

	// check header
	header := txData[:4]
	if bytes.Compare(header, CustomVerificationPrefix) != 0 {
		return nil, nil, errors.New("invalid validation data header")
	}

	// check checksum
	checksum := txData[4:8]
	dataHash := crypto.Keccak256(txData[8:])
	if bytes.Compare(checksum, dataHash[:4]) != 0 {
		return nil, nil, errors.New("invalid validation data checksum")
	}

	// decode payload
	payload := txData[8:]
	validationData, err = loadParamBytes(payload, 0)
	if err != nil {
		return
	}

	callData, err = loadParamBytes(payload, 1)
	return
}

func loadParamBytes(input []byte, index int) ([]byte, error) {
	offsetLowerBound := index * 32
	offsetUpperbound := offsetLowerBound + 32
	if len(input) < offsetUpperbound {
		return nil, errors.New("invalid input data length")
	}

	dataOffset, overflow := uint256.NewInt(0).SetBytes32(input[offsetLowerBound:offsetUpperbound]).Uint64WithOverflow()
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
