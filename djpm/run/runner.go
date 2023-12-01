package run

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/artela-network/aspect-core/djpm/run/api"

	runtime "github.com/artela-network/aspect-runtime"
	"google.golang.org/protobuf/proto"

	"github.com/artela-network/aspect-core/types"
)

type Runner struct {
	vmKey string
	vm    runtime.AspectRuntime
	// fns      *runtime.HostAPIRegistry
	register *api.Register
	code     []byte
}

func NewRunner(aspID string, code []byte) (*Runner, error) {
	aspectId := common.HexToAddress(aspID)
	register := api.NewRegister(&aspectId)
	key, vm, err := types.Runtime(code, register.HostApis())
	if err != nil {
		return nil, err
	}
	return &Runner{
		vmKey:    key,
		vm:       vm,
		register: register,
		code:     code,
	}, nil
}

func (r *Runner) Return() {
	types.ReturnRuntime(r.vmKey, r.vm)
}

func (r *Runner) JoinPoint(name types.PointCut, gas uint64, blockNumber int64, contractAddr *common.Address, txRequest proto.Message) (*types.AspectResponse, error) {
	if r.vm == nil {
		return nil, errors.New("runner not init")
	}
	// turn inputBytes into bytes
	reqData, err := proto.Marshal(txRequest)
	if err != nil {
		return nil, err
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	// for get aspect Error message
	r.register.SetErrCallback(callback)
	r.register.SetRunnerContext(string(name), blockNumber, gas, contractAddr)

	res, err := r.vm.Call(api.ApiEntrance, string(name), reqData)
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return nil, errors.New(revertMsg)
		}
		return nil, err
	}
	resData, ok := res.([]byte)
	if !ok {
		return nil, errors.New("read output failed, return value is not byte array")
	}
	output := &types.AspectResponse{}
	if err := proto.Unmarshal(resData, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (r *Runner) IsOwner(blockNumber int64, gas uint64, contractAddr *common.Address, sender string) (bool, error) {
	if r.vm == nil {
		return false, errors.New("vm not init")
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.register.SetErrCallback(callback)
	r.register.SetRunnerContext("isOwner", blockNumber, gas, contractAddr)

	res, err := r.vm.Call(api.ApiEntrance, "isOwner", sender)
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
	r.register.SetErrCallback(callback)
	res, err := r.vm.Call(api.CheckBlockLevel)
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return false, errors.New(revertMsg)
		}
		return false, err
	}
	return res.(bool), nil
}

func (r *Runner) OnContractBinding(blockNumber int64, gas uint64, contractAddr *common.Address, sender string) (bool, error) {
	if r.vm == nil {
		return false, errors.New("run not init")
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.register.SetErrCallback(callback)
	r.register.SetRunnerContext(string(types.ON_CONTRACT_BINDING_METHOD), blockNumber, gas, contractAddr)
	res, err := r.vm.Call(api.ApiEntrance, string(types.ON_CONTRACT_BINDING_METHOD), sender)
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
	r.register.SetErrCallback(callback)
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
	r.register.SetErrCallback(callback)
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
	return r.register.RunnerContext().Gas
}

func (r *Runner) ExecFunc(funcName string, blockNumber int64, gas uint64, contractAddr *common.Address, args ...interface{}) (interface{}, error) {
	if r.vm == nil {
		return false, errors.New("run not init")
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.register.SetErrCallback(callback)
	r.register.SetRunnerContext(funcName, blockNumber, gas, contractAddr)
	res, err := r.vm.Call(funcName, args...)
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return false, errors.New(revertMsg)
		}
		return nil, err
	}
	return res, nil
}
