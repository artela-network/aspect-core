package types

import runtime "github.com/artela-network/aspect-runtime"

const DefaultAspectPoolSize = 10

var (
	enable bool
	vmPool *runtime.RuntimePool
)

// InitRuntimePool init runtime pool with given capacity.
func InitRuntimePool(capacity int32) {
	enable = capacity > 0
	if !enable {
		return
	}
	vmPool = runtime.NewRuntimePool(int(capacity))
}

// Runtime returns a aspect-runtime instance from the pool or creating a new one.
func Runtime(code []byte, registry *runtime.HostAPIRegistry) (string, runtime.AspectRuntime, error) {
	if !enable {
		vm, err := runtime.NewAspectRuntime(runtime.WASM, code, registry)
		return "", vm, err
	}

	return vmPool.Runtime(runtime.WASM, code, registry)
}

// ReturnRuntime returns the the runtime instance to the pool, is the pool is enabled.
func ReturnRuntime(key string, instance runtime.AspectRuntime) {
	if !enable {
		return
	}
	if vmPool == nil {
		InitRuntimePool(DefaultAspectPoolSize)
	}
	vmPool.Return(key, instance)
}
