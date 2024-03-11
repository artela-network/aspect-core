package api

import (
	"github.com/artela-network/aspect-core/types"
	types2 "github.com/artela-network/aspect-runtime/types"
)

func (r *Registry) aspectStateAPIs() map[string]*types2.HostFuncWithGasRule {
	return map[string]*types2.HostFuncWithGasRule{
		"get": {
			Func: func(key string) ([]byte, error) {
				hook, err := types.GetAspectStateHostHook(r.runnerContext.Ctx)
				if err != nil {
					panic("failed to init aspect runtime context host api: " + err.Error())
				}
				if hook == nil {
					panic("aspect runtime context host api not found")
				}

				return wrapNilByte(hook.Get(r.runnerContext, key)), nil
			},
			GasRule: types2.NewStaticGasRule(1),
		},
		"set": {
			Func: func(key string, val []byte) error {
				hook, err := types.GetAspectStateHostHook(r.runnerContext.Ctx)
				if err != nil {
					panic("failed to init aspect runtime context host api: " + err.Error())
				}
				if hook == nil {
					panic("aspect runtime context host api not found")
				}
				return hook.Set(r.runnerContext, key, val)
			},
			GasRule: types2.NewStaticGasRule(1),
		},
	}
}
