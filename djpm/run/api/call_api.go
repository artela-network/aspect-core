package api

import "github.com/artela-network/aspect-core/types"

func (r *Register) stateDbApis() interface{} {
	return map[string]interface{}{
		"getBalance": func(addr string) string {
			hook, err := types.GetStateDbHook()
			if err != nil || hook == nil {
				return ""
			}
			return hook.GetBalance(r.runnerContext, addr)
		},
		"getState": func(addr string, hash string) string {
			hook, err := types.GetStateDbHook()
			if err != nil || hook == nil {
				return ""
			}
			return hook.GetState(r.runnerContext, addr, hash)
		},
		"getRefund": func() int64 {
			hook, err := types.GetStateDbHook()
			if err != nil || hook == nil {
				return 0
			}
			return int64(hook.GetRefund(r.runnerContext))
		},
		"getCodeHash": func(addr string) string {
			hook, err := types.GetStateDbHook()
			if err != nil || hook == nil {
				return ""
			}
			return hook.GetCodeHash(r.runnerContext, addr)
		},
		"getNonce": func(addr string) int64 {
			hook, err := types.GetStateDbHook()
			if err != nil || hook == nil {
				return 0
			}
			return int64(hook.GetNonce(r.runnerContext, addr))
		},
	}
}
