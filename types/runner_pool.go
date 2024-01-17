package types

import (
	runtime "github.com/artela-network/aspect-runtime"
)

const DefaultAspectPoolSize = 10

var (
	globalQueryPool *globalPool
	globalMsgPool   *globalPool
)

// InitRuntimePool init runtime pool with given capacity.
func InitRuntimePool(msgPoolCapacity, queryPoolCapacity int32) {
	globalMsgPool = newGlobalPool(msgPoolCapacity)
	globalQueryPool = newGlobalPool(queryPoolCapacity)
}

func RunnerPool(commit bool) *globalPool {
	if commit {
		return globalMsgPool
	}
	return globalQueryPool
}

type globalPool struct {
	enable bool
	vmPool *runtime.RuntimePool
}

func newGlobalPool(capacity int32) *globalPool {
	return &globalPool{
		enable: capacity > 0,
		vmPool: runtime.NewRuntimePool(int(capacity)),
	}
}

// Runtime returns a aspect-runtime instance from the pool or creating a new one.
func (p *globalPool) Runtime(code []byte, registry *runtime.HostAPIRegistry) (string, runtime.AspectRuntime, error) {
	if !p.enable {
		vm, err := runtime.NewAspectRuntime(runtime.WASM, code, registry)
		return "", vm, err
	}

	return p.vmPool.Runtime(runtime.WASM, code, registry)
}

// ReturnRuntime returns the the runtime instance to the pool, is the pool is enabled.
func (p *globalPool) Return(key string, vm runtime.AspectRuntime) {
	if !p.enable {
		// release the host functions and memorys in Destory
		vm.Destroy()
		return
	}
	if p.vmPool == nil {
		p = newGlobalPool(DefaultAspectPoolSize)
	}
	p.vmPool.Return(key, vm)
}
