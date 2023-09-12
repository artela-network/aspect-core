package run

import (
	"github.com/artela-network/artelasdk/djpm/run/api"
	"strings"

	"github.com/artela-network/artelasdk/types"
	"github.com/artela-network/runtime"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

var vmPool *runtime.RuntimePool

func RuntimePool() *runtime.RuntimePool {
	if vmPool == nil {
		vmPool = runtime.NewRuntimePool(10)
	}
	return vmPool
}

type Runner struct {
	vmKey    string
	vm       runtime.AspectRuntime
	fns      *runtime.HostAPIRegistry
	register *api.Register
	code     []byte
}

func NewRunner(aspID string, code []byte) (*Runner, error) {
	register := api.NewRegister(aspID)
	key, vm, err := RuntimePool().Runtime(runtime.WASM, code, register.HostApis())
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
	RuntimePool().Return(r.vmKey, r.vm)
}

func (r *Runner) JoinPoint(name types.PointCut, gas uint64, blockNumber int64, txRequest proto.Message) (*types.AspectResponse, error) {
	if r.vm == nil {
		return nil, errors.New("runner not init")
	}
	//turn inputBytes into bytes
	reqData, err := proto.Marshal(txRequest)
	if err != nil {
		return nil, err
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	//for get aspect Error message
	r.register.SetErrCallback(callback)
	r.register.SetRunnerContext(name, blockNumber, gas)

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
		return nil, errors.Wrap(err, "unmarshal AspectOutput")
	}
	return output, nil
}

func (r *Runner) IsOwner(sender string) (bool, error) {
	if r.vm == nil {
		return false, errors.New("vm not init")
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.register.SetErrCallback(callback)
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

func (r *Runner) OnContractBinding(sender string) (bool, error) {
	if r.vm == nil {
		return false, errors.New("run not init")
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.register.SetErrCallback(callback)
	res, err := r.vm.Call(api.ApiEntrance, "onContractBinding", sender)
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

func (r *Runner) ExecFunc(funcName string, args ...interface{}) (interface{}, error) {
	if r.vm == nil {
		return false, errors.New("run not init")
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.register.SetErrCallback(callback)
	res, err := r.vm.Call(funcName, args...)
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return false, errors.New(revertMsg)
		}
		return nil, err
	}
	return res, nil
}
