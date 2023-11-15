package types

import runtime "github.com/artela-network/aspect-runtime"

const DefaultAspectPoolSize = 10

var vmPool *runtime.RuntimePool

func InitRuntimePool(capacity int32) {
	vmPool = runtime.NewRuntimePool(int(capacity))
}

func RuntimePool() *runtime.RuntimePool {
	if vmPool == nil {
		InitRuntimePool(DefaultAspectPoolSize)
	}
	return vmPool
}
