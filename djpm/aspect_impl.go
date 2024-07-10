package djpm

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	runtime "github.com/artela-network/aspect-runtime/types"
	"math/big"
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
	MaxTxVerificationGas     = uint64(150000) // maximum gas allowed for customized transaction verification
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

func (aspect Aspect) VerifyTx(ctx context.Context, contract common.Address, block int64, gas uint64, request *types.TxVerifyInput) *types.AspectExecutionResult {
	return aspect.verification(ctx, contract, block, gas, request)
}

func (aspect Aspect) PreTxExecute(ctx context.Context, from, contract common.Address, input []byte, block int64, gas uint64, value *big.Int, request *types.PreTxExecuteInput, tracer types.AspectLogger) *types.AspectExecutionResult {
	return aspect.transactionAdvice(ctx, types.PRE_TX_EXECUTE_METHOD, from, contract, input, block, gas, value, request, tracer)
}

func (aspect Aspect) PreContractCall(ctx context.Context, from, contract common.Address, input []byte, block int64, gas uint64, value *big.Int, request *types.PreContractCallInput, tracer types.AspectLogger) *types.AspectExecutionResult {
	return aspect.transactionAdvice(ctx, types.PRE_CONTRACT_CALL_METHOD, from, contract, input, block, gas, value, request, tracer)
}

func (aspect Aspect) PostContractCall(ctx context.Context, from, contract common.Address, input []byte, block int64, gas uint64, value *big.Int, request *types.PostContractCallInput, tracer types.AspectLogger) *types.AspectExecutionResult {
	return aspect.transactionAdvice(ctx, types.POST_CONTRACT_CALL_METHOD, from, contract, input, block, gas, value, request, tracer)
}

func (aspect Aspect) PostTxExecute(ctx context.Context, from, contract common.Address, input []byte, block int64, gas uint64, value *big.Int, request *types.PostTxExecuteInput, tracer types.AspectLogger) *types.AspectExecutionResult {
	return aspect.transactionAdvice(ctx, types.POST_TX_EXECUTE_METHOD, from, contract, input, block, gas, value, request, tracer)
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

	if tx.To() == nil {
		aspect.logger.Error("cannot do customized verify on contract creation", "tx", tx.Hash().Hex())
		return common.Address{}, nil, errors.New("contract creation is not allowed for customized verification")
	}

	logger := aspect.logger.With("contract", tx.To().Hex(), "block", block, "tx", tx.Hash().Hex())

	// check contract verifier
	verifiers, err := aspect.provider.GetAccountVerifiers(ctx, *tx.To())
	if err != nil {
		return common.Address{}, nil, err
	}

	if len(verifiers) != 1 {
		logger.Error("contract has more than 1 verifiers", "verifiers", len(verifiers))
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
	// we cannot use tx.Gas() here, because verifier execution cannot be count into the execution gas,
	// we'll use a fixed gas here
	verifyRes := aspect.VerifyTx(ctx, *tx.To(), block, MaxTxVerificationGas, request)
	if verifyRes.Err != nil {
		logger.Error("failed to verify tx with aspect", "err", verifyRes.Err)
		return common.Address{}, nil, verifyRes.Err
	}

	sender := common.BytesToAddress(verifyRes.Ret)

	logger = logger.With("sender", "sender", sender.Hex())

	// make sure sender accepts this aspect as verifier
	aspects, err := aspect.provider.GetAccountVerifiers(ctx, sender)
	if err != nil {
		logger.Error("failed to get sender's verifiers", "err", err)
		return common.Address{}, nil, err
	}

	for _, aspect := range aspects {
		if aspect.AspectId == contractVerifier {
			logger.Debug("aspect verification passed")
			return sender, call, nil
		}
	}

	logger.Error("sender does not accept this aspect as verifier")
	return common.Address{}, nil, errors.New("unable to verify tx with aspect")
}

func (aspect Aspect) transactionAdvice(ctx context.Context, method types.PointCut, from, contract common.Address, input []byte, block int64, gas uint64, value *big.Int, request proto.Message, tracer types.AspectLogger) *types.AspectExecutionResult {
	result := &types.AspectExecutionResult{
		Gas:    gas,
		Revert: types.NotRevert,
	}

	// get binding contract address
	aspectCodes, err := aspect.provider.GetTxBondAspects(ctx, contract, method)
	if err != nil {
		result.Err = err
		result.Revert = types.RevertCall
		return result
	}
	if len(aspectCodes) == 0 {
		return result
	}

	// run aspects on received transaction
	return aspect.runAspect(ctx, method, gas, block, from, contract, input, value, request, aspectCodes, tracer)
}

func (aspect Aspect) verification(ctx context.Context, contract common.Address, block int64, gas uint64, req *types.TxVerifyInput) *types.AspectExecutionResult {
	aspectCodes, err := aspect.provider.GetAccountVerifiers(ctx, contract)
	if err != nil || len(aspectCodes) == 0 {
		return &types.AspectExecutionResult{
			Gas:    gas,
			Err:    errors.New("contract has not bound with any verifier aspect"),
			Revert: types.RevertTx,
		}
	}

	// run aspects on received transaction.
	// since tx verification is out of the execution context, we cannot really trace it with EVM tracer,
	// so for now we just set input, value and tracer to nil.
	// For tx.from, since before running the verifier, we don't know the sender yet, so we just set it to empty.
	return aspect.runAspect(ctx, types.VERIFY_TX, gas, block, common.Address{}, contract, nil, nil, req, aspectCodes, nil)
}

func (aspect Aspect) runAspect(ctx context.Context, method types.PointCut, gas uint64, blockNumber int64, from,
	contractAddr common.Address, input []byte, value *big.Int, reqData proto.Message,
	aspects []*types.AspectCode, tracer types.AspectLogger) (result *types.AspectExecutionResult) {
	result = &types.AspectExecutionResult{
		Gas:    gas,
		Revert: types.NotRevert,
	}
	defer func() {
		if err := recover(); err != nil {
			aspect.logger.Error("panic in running aspect", "err", err, "stack", debug.Stack())
			result.Err = errors.New("aspect execution crashed")
			result.Revert = types.RevertCall
		}
	}()

	for _, storedAspect := range aspects {
		jp := types.JoinPointRunType(types.JoinPointRunType_value[string(method)])
		if tracer != nil {
			tracer.CaptureAspectEnter(jp, from, contractAddr, common.HexToAddress(storedAspect.AspectId), input, gas, value, reqData)
		}
		var err error
		isCommit := types.IsCommit(ctx)
		runner, err := run.NewRunner(ctx, aspect.logger, storedAspect.AspectId, storedAspect.Version, storedAspect.Code, isCommit)
		if err != nil {
			// no need to exit tracer here, since if the logic is correct, this should never happen
			panic(err)
		}

		var ret []byte
		ret, gas, err = runner.JoinPoint(method, gas, blockNumber, contractAddr, reqData)
		runner.Return()

		result.Ret = ret
		result.Gas = gas
		result.Err = err
		if tracer != nil {
			// tracer does not care the revert scope, so we can exit tracer here before settting the revert scope
			tracer.CaptureAspectExit(jp, result)
		}

		// revert scope is not fully supported yet, so we just break the loop if any aspect fails,
		// no need to record this in the tracer for now
		if err != nil {
			result.Revert = types.RevertCall
			break
		}
	}

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
