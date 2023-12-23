package run

import (
	"context"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/artela-network/aspect-core/djpm/run/api"

	runtime "github.com/artela-network/aspect-runtime"
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
}

func NewRunner(ctx context.Context, aspID string, aspVer uint64, code []byte) (*Runner, error) {
	aspectId := common.HexToAddress(aspID)
	registry := api.NewRegistry(ctx, aspectId, aspVer)
	key, vm, err := types.Runtime(code, registry.HostApis())
	if err != nil {
		return nil, err
	}
	return &Runner{
		ctx:      ctx,
		vmKey:    key,
		vm:       vm,
		registry: registry,
		code:     code,
	}, nil
}

func (r *Runner) Return() {
	types.ReturnRuntime(r.vmKey, r.vm)
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

	res, err := r.vm.Call(api.APIEntrance, string(name), reqData)
	gas = r.registry.RunnerContext().Gas
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return []byte(revertMsg), gas, errors.New(revertMsg)
		}
		return nil, gas, err
	}

	if res == nil {
		return nil, gas, nil
	}

	resData, ok := res.([]byte)
	if !ok {
		return nil, gas, errors.New("read output failed, return value is not byte array")
	}

	return resData, gas, nil
}

func (r *Runner) IsOwner(blockNumber int64, gas uint64, contractAddr *common.Address, sender string) (bool, error) {
	if r.vm == nil {
		return false, errors.New("vm not init")
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.registry.SetErrCallback(callback)
	r.registry.SetRunnerContext("isOwner", blockNumber, gas, contractAddr)

	res, err := r.vm.Call(api.APIEntrance, "isOwner", sender)
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return false, errors.New(revertMsg)
		}

		return false, err
	}

	return res.(bool), nil
}

func (r *Runner) IsBlockLevel() (bool, error) {
	if r.vm == nil {
		return false, errors.New("not init")
	}

	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.registry.SetErrCallback(callback)
	res, err := r.vm.Call(api.CheckBlockLevel)
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return false, errors.New(revertMsg)
		}
		return false, err
	}
	return res.(bool), nil
}

func (r *Runner) IsTransactionLevel() (bool, error) {
	if r.vm == nil {
		return false, errors.New("not init")
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.registry.SetErrCallback(callback)
	res, err := r.vm.Call(api.CheckTransactionLevel)
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return false, errors.New(revertMsg)
		}
		return false, err
	}
	return res.(bool), nil
}

func (r *Runner) IsTxVerifier() (bool, error) {
	if r.vm == nil {
		return false, errors.New("not init")
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.registry.SetErrCallback(callback)
	res, err := r.vm.Call(api.CheckIsTxVerifier)
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
	res, err := r.vm.Call(funcName, args...)
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return false, errors.New(revertMsg)
		}
		return nil, err
	}
	return res, nil
}
