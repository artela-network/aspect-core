package api

import "github.com/artela-network/artelasdk/types"

func (r *Register) stateApis() interface{} {
	return map[string]interface{}{

		"setAspectState": func(key string, value string) bool {
			hook, err := types.GetAspectStateHook()
			if err != nil || hook == nil {
				return false
			}
			return hook.SetAspectState(r.runnerContext, key, value)
		},
		"getAspectState": func(key string) string {
			hook, err := types.GetAspectStateHook()
			if err != nil || hook == nil {
				return ""
			}
			return hook.GetAspectState(r.runnerContext, key)
		},
		"removeAspectState": func(key string) bool {
			hook, err := types.GetAspectStateHook()
			if err != nil || hook == nil {
				return false
			}
			return hook.RemoveAspectState(r.runnerContext, key)
		},
		"getProperty": func(key string) string {
			hook, err := types.GetAspectStateHook()
			if err != nil || hook == nil {
				return ""
			}
			return hook.GetProperty(r.runnerContext, key)
		},
	}
}
