package run

import (
	"context"
	"errors"
	runtime "github.com/artela-network/aspect-runtime/types"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/artela-network/aspect-core/djpm/run/api"

	"google.golang.org/protobuf/proto"

	"github.com/artela-network/aspect-core/types"
)

type Runner struct {
	ctx   context.Context
	vmKey string
	vm    runtime.AspectRuntime
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
		code:     code,
		commit:   commit,
	}, nil
}

func (r *Runner) Return() {
	types.RunnerPool(r.commit).Return(r.vmKey, r.vm)
}

func (r *Runner) JoinPoint(name types.PointCut, gas uint64, blockNumber int64, contractAddr *common.Address, txRequest proto.Message) ([]byte, uint64, error) {
	if r.vm == nil {
		return nil, gas, errors.New("runner not init")
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
		if !strings.EqualFold(revertMsg, "") {
			return []byte(revertMsg), gas, errors.New(revertMsg)
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

func (r *Runner) IsOwner(blockNumber int64, gas uint64, contractAddr *common.Address, sender []byte) (bool, error) {
	if r.vm == nil {
		return false, errors.New("vm not init")
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.registry.SetErrCallback(callback)
	r.registry.SetRunnerContext("isOwner", blockNumber, gas, contractAddr)

	// TODO: no gas refund for aspect for now, add later
	res, _, err := r.vm.Call(api.APIEntrance, int64(gas), "isOwner", sender)
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return false, errors.New(revertMsg)
		}

		return false, err
	}

	return res.(bool), nil
}

func (r *Runner) Gas() uint64 {
	return r.registry.RunnerContext().Gas
}

func (r *Runner) ExecFunc(funcName string, blockNumber int64, gas uint64, contractAddr *common.Address, args ...interface{}) (interface{}, error) {
	if r.vm == nil {
		return false, errors.New("run not init")
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.registry.SetErrCallback(callback)
	r.registry.SetRunnerContext(funcName, blockNumber, gas, contractAddr)

	// TODO: no gas refund for aspect for now, add later
	res, _, err := r.vm.Call(funcName, int64(gas), args...)
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return false, errors.New(revertMsg)
		}
		return nil, err
	}
	return res, nil
}
