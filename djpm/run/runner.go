package run

import (
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
	register *Register
	code     []byte
}

func NewRunner(aspID string, code []byte) (*Runner, error) {
	register := NewRegister(aspID)
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

func (r *Runner) PutBack() {
	RuntimePool().PutBack(r.vmKey, r.vm)
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

	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.register.SetErrCallback(callback)

	res, err := r.vm.Call(ApiEntrance, name, reqData)
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
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.register.SetErrCallback(callback)
	res, err := r.vm.Call(ApiEntrance, "isOwner", sender)
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
	res, err := r.vm.Call(CheckBlockLevel)
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
		return false, errors.New("not init")
	}
	revertMsg := ""
	callback := func(msg string) {
		revertMsg = msg
	}
	r.register.SetErrCallback(callback)
	res, err := r.vm.Call(ApiEntrance, "onContractBinding", sender)
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
	res, err := r.vm.Call(CheckTransactionLevel)
	if err != nil {
		if !strings.EqualFold(revertMsg, "") {
			return false, errors.New(revertMsg)
		}
		return false, err
	}
	return res.(bool), nil
}
