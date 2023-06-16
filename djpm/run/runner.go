package run

import (
	"github.com/artela-network/runtime"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	"github.com/artela-network/artelasdk/types"
)

var vmPool *runtime.RuntimePool

func RuntimePool() *runtime.RuntimePool {
	if vmPool == nil {
		vmPool = runtime.NewRuntimePool(10)
	}
	return vmPool
}

type Runner struct {
	vm   runtime.AspectRuntime
	fns  *runtime.HostAPIRegistry
	code []byte
}

func NewRunner(aspID string, code []byte) (*Runner, error) {
	return NewRunnerWithCallBack(aspID, code, nil)
}

func NewRunnerWithCallBack(aspID string, code []byte, cbRevert CallBackRevertFunc) (*Runner, error) {
	register := NewRegister(aspID, cbRevert)
	_, vm, err := RuntimePool().Get(runtime.WASM, code, register.HostApis(), false)
	if err != nil {
		return nil, err
	}
	return &Runner{
		vm:   vm,
		code: code,
	}, nil
}

func (r *Runner) JoinPoint(name string, input *types.AspectInput) (*types.AspectOutput, error) {
	if r.vm == nil {
		return nil, errors.New("not init")
	}
	// turn input into bytes
	reqData, err := proto.Marshal(input)
	if err != nil {
		return nil, err
	}
	res, err := r.vm.Call(ApiEntrance, name, reqData)
	if err != nil {
		call, err := r.vm.Call(GetErr)
		if err == nil && call != nil {
			message := call.(string)
			return nil, errors.New(message)
		}
		return nil, err
	}

	resData, ok := res.([]byte)
	if !ok {
		return nil, errors.New("read output failed, return value is not byte array")
	}

	output := &types.AspectOutput{}
	if err := proto.Unmarshal(resData, output); err != nil {
		return nil, errors.Wrap(err, "unmarshal AspectOutput")
	}

	return output, nil
}

func (r *Runner) IsOwner(sender string) (bool, error) {
	if r.vm == nil {
		return false, errors.New("not init")
	}

	res, err := r.vm.Call(ApiEntrance, "isOwner", sender)
	if err != nil {
		call, err := r.vm.Call(GetErr)
		if err == nil && call != nil {
			message := call.(string)
			return false, errors.New(message)
		}

		return false, err
	}

	return res.(bool), nil
}

func (r *Runner) IsBlockLevel() (bool, error) {
	if r.vm == nil {
		return false, errors.New("not init")
	}
	res, err := r.vm.Call(CheckBlockLevel)
	if err != nil {
		call, err := r.vm.Call(GetErr)
		if err == nil && call != nil {
			message := call.(string)
			return false, errors.New(message)
		}
		return false, err
	}
	return res.(bool), nil
}

func (r *Runner) OnContractBinding(sender string) (bool, error) {
	if r.vm == nil {
		return false, errors.New("not init")
	}

	res, err := r.vm.Call(ApiEntrance, "onContractBinding", sender)
	if err != nil {
		call, err := r.vm.Call(GetErr)
		if err == nil && call != nil {
			message := call.(string)
			return false, errors.New(message)
		}
		return false, err
	}

	return res.(bool), nil
}

func (r *Runner) IsTransactionLevel() (bool, error) {
	if r.vm == nil {
		return false, errors.New("not init")
	}
	res, err := r.vm.Call(CheckTransactionLevel)
	if err != nil {
		call, err := r.vm.Call(GetErr)
		if err == nil && call != nil {
			message := call.(string)
			return false, errors.New(message)
		}
		return false, err
	}
	return res.(bool), nil
}
