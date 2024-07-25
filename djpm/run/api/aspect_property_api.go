package api

import (
	"github.com/artela-network/aspect-core/types"
	types2 "github.com/artela-network/aspect-runtime/types"
)

func (r *Registry) aspectPropertyAPIs() map[string]*types2.HostFuncWithGasRule {
	return map[string]*types2.HostFuncWithGasRule{
		"get": {
			Func: func(key string) ([]byte, error) {
				hook, err := types.GetAspectPropertyHostHook(r.runnerContext.Ctx)
				if err != nil {
					panic("failed to init aspect runtime context host api: " + err.Error())
				}
				if hook == nil {
					panic("aspect runtime context host api not found")
				}

				ret, err := hook.Get(r.runnerContext, key)
				return wrapNilByte(ret), err
			},
			GasRule: types2.NewStaticGasRule(12500),
		},
	}
}
