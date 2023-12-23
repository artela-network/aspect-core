package api

import (
	"github.com/artela-network/aspect-core/types"
)

func (r *Registry) aspectPropertyAPIs() interface{} {
	return map[string]interface{}{
		"get": func(aspectId []byte, key string) []byte {
			hook, err := types.GetAspectPropertyHostHook(r.runnerContext.Ctx)
			if err != nil {
				panic("failed to init aspect runtime context host api: " + err.Error())
			}
			if hook == nil {
				panic("aspect runtime context host api not found")
			}
			return hook.Get(r.runnerContext, key)
		},
	}
}
