package types

import (
	"context"
	runtime "github.com/artela-network/aspect-runtime"
	"github.com/artela-network/aspect-runtime/types"
)

const DefaultAspectPoolSize = 10

var (
	globalQueryPool *globalPool
	globalMsgPool   *globalPool
)

// InitRuntimePool init runtime pool with given capacity.
func InitRuntimePool(ctx context.Context, logger types.Logger, msgPoolCapacity, queryPoolCapacity int32) {
	globalMsgPool = newGlobalPool(ctx, logger, msgPoolCapacity)
	globalQueryPool = newGlobalPool(ctx, logger, queryPoolCapacity)
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

func newGlobalPool(ctx context.Context, logger types.Logger, capacity int32) *globalPool {
	return &globalPool{
		enable: capacity > 0,
		vmPool: runtime.NewRuntimePool(ctx, logger, int(capacity)),
	}
}

// Runtime returns a aspect-runtime instance from the pool or creating a new one.
func (p *globalPool) Runtime(ctx context.Context, logger types.Logger, code []byte, registry *types.HostAPIRegistry) (string, types.AspectRuntime, error) {
	if !p.enable {
		vm, err := runtime.NewAspectRuntime(ctx, logger, runtime.WASM, code, registry)
		return "", vm, err
	}

	return p.vmPool.Runtime(runtime.WASM, code, registry)
}

// ReturnRuntime returns the the runtime instance to the pool, is the pool is enabled.
func (p *globalPool) Return(key string, vm types.AspectRuntime) {
	if !p.enable {
		// release the host functions and memorys in Destory
		vm.Destroy()
		return
	}
	if p.vmPool == nil {
		p = newGlobalPool(vm.Context(), vm.Logger(), DefaultAspectPoolSize)
	}
	p.vmPool.Return(key, vm)
}
