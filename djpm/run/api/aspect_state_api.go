package api

import "github.com/artela-network/aspect-core/types"

func (r *Registry) aspectStateAPIs() interface{} {
	return map[string]interface{}{
		"get": func(aspectId []byte, key string) []byte {
			hook, err := types.GetAspectStateHostHook(r.runnerContext.Ctx)
			if err != nil {
				panic("failed to init aspect runtime context host api: " + err.Error())
			}
			if hook == nil {
				panic("aspect runtime context host api not found")
			}
			return hook.Get(r.runnerContext, key)
		},
		"set": func(aspectId []byte, key string, val []byte) {
			hook, err := types.GetAspectStateHostHook(r.runnerContext.Ctx)
			if err != nil {
				panic("failed to init aspect runtime context host api: " + err.Error())
			}
			if hook == nil {
				panic("aspect runtime context host api not found")
			}
			hook.Set(r.runnerContext, key, val)
		},
	}
}