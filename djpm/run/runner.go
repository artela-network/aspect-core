package run

import (
	"context"
	"errors"
	runtime "github.com/artela-network/aspect-runtime/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/artela-network/aspect-core/djpm/run/api"

	"google.golang.org/protobuf/proto"

	"github.com/artela-network/aspect-core/types"
)

// ErrExecutionReverted same as EVM execution reverted error, used to indicate the execution is reverted.
var ErrExecutionReverted = errors.New("execution reverted")

type Runner struct {
	ctx      context.Context
	vmKey    string
	vm       runtime.AspectRuntime
	aspectId common.Address

	// fns      *runtime.HostAPIRegistry
	registry *api.Registry
	code     []byte
	commit   bool

	logger runtime.Logger
}

func NewRunner(ctx context.Context, logger runtime.Logger, aspID string, aspVer uint64, code []byte, commit bool) (*Runner, error) {
	aspectId := common.HexToAddress(aspID)
	registry := api.NewRegistry(ctx, aspectId, aspVer)
	key, vm, err := types.RunnerPool(commit).Runtime(ctx, logger, code, registry.HostApis())
	if err != nil {
		return nil, err
	}
	return &Runner{
		ctx:      ctx,
		logger:   logger,
		vmKey:    key,
		vm:       vm,
		registry: registry,
		aspectId: aspectId,
		code:     code,
		commit:   commit,
	}, nil
}

func (r *Runner) Return() {
	r.registry.Destroy()
	types.RunnerPool(r.commit).Return(r.vmKey, r.vm)
}

func (r *Runner) JoinPoint(name types.PointCut, gas uint64, blockNumber int64, contractAddr common.Address, txRequest proto.Message) ([]byte, uint64, error) {
	if r.vm == nil {
		panic("vm not init")
	}
	// turn inputBytes into bytes
	reqData, err := proto.Marshal(txRequest)
	if err != nil {
		return nil, gas, err
	}

	revertMsg := ""
	errorFunc := func(msg string) {
		revertMsg = msg
	}
	// for get aspect Error message
	r.registry.SetErrCallback(errorFunc)
	r.registry.SetRunnerContext(string(name), blockNumber, gas, contractAddr)

	res, leftover, err := r.vm.Call(api.APIEntrance, int64(gas), string(name), reqData)
	if err != nil {
		r.logger.Error("join point execution failed", "block", blockNumber, "contract", contractAddr.Hex(), "joinpoint", name, "gas", gas, "err", err, "revertMsg", revertMsg)
		if !strings.EqualFold(revertMsg, "") {
			// need to pack the revert message as abi, then it can be decoded by the caller
			return PackRevert(revertMsg), gas, ErrExecutionReverted
		}
		return nil, uint64(leftover), err
	}

	if res == nil {
		return nil, uint64(leftover), nil
	}

	resData, ok := res.([]byte)
	if !ok {
		return nil, gas, errors.New("read output failed, return value is not byte array")
	}

	return resData, uint64(leftover), nil
}

func (r *Runner) IsOwner(blockNumber int64, gas uint64, contractAddr common.Address, sender []byte) (bool, uint64, error) {
	if r.vm == nil {
		panic("vm not init")
	}

	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.registry.SetErrCallback(callback)
	r.registry.SetRunnerContext("isOwner", blockNumber, gas, contractAddr)

	// TODO: no gas refund for aspect for now, add later
	res, leftover, err := r.vm.Call(api.APIEntrance, int64(gas), "isOwner", sender)
	leftoverU64 := uint64(leftover)
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return false, leftoverU64, errors.New(revertMsg)
		}

		return false, leftoverU64, err
	}

	return res.(bool), leftoverU64, nil
}

func (r *Runner) Gas() uint64 {
	return r.registry.RunnerContext().Gas
}

func (r *Runner) ExecFunc(funcName string, blockNumber int64, gas uint64, contractAddr common.Address, args ...interface{}) (interface{}, uint64, error) {
	if r.vm == nil {
		panic("vm not init")
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.registry.SetErrCallback(callback)
	r.registry.SetRunnerContext(funcName, blockNumber, gas, contractAddr)

	// TODO: no gas refund for aspect for now, add later
	res, leftover, err := r.vm.Call(funcName, int64(gas), args...)
	leftoverU64 := uint64(leftover)
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return false, leftoverU64, errors.New(revertMsg)
		}
		return nil, leftoverU64, err
	}
	return res, leftoverU64, nil
}

// revertSelector is a special function selector for revert reason unpacking.
var revertSelector = crypto.Keccak256([]byte("Error(string)"))[:4]

// PackRevert packs the revert message from Aspect to make sure this message can be decoded by the caller contract.
func PackRevert(reason string) []byte {
	if len(reason) == 0 {
		return nil
	}

	selector := revertSelector

	typ, err := abi.NewType("string", "", nil)
	if err != nil {
		return nil
	}

	packed, err := (abi.Arguments{{Type: typ}}).Pack(reason)
	if err != nil {
		return nil
	}

	return append(selector, packed[:]...)
}
