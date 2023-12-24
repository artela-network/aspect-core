package api

import (
	"github.com/artela-network/aspect-core/types"
)

func (r *Registry) transientStorageAPIs() interface{} {
	return map[string]interface{}{
		"get": func(aspectId []byte, key string) []byte {
			hook, err := types.GetAspectTransientStorageHostHook(r.runnerContext.Ctx)
			if err != nil {
				panic("failed to init aspect runtime context host api: " + err.Error())
			}
			if hook == nil {
				panic("aspect runtime context host api not found")
			}
			return wrapNilByte(hook.Get(r.runnerContext, aspectId, key))
		},
		"set": func(key string, val []byte) {
			hook, err := types.GetAspectTransientStorageHostHook(r.runnerContext.Ctx)
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
