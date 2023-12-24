package api

import "github.com/artela-network/aspect-core/types"
import "github.com/ethereum/go-ethereum/common"

func (r *Registry) stateDBAPIs() interface{} {
	return map[string]interface{}{
		"getBalance": func(addr []byte) []byte {
			hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				panic("GetStateDbHook failed")
			}
			balance := hook.GetBalance(common.BytesToAddress(addr))
			return wrapNilByte(balance.Bytes())
		},
		"getState": func(addr []byte, hash []byte) []byte {
			hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				panic("GetStateDbHook failed")
			}
			return wrapNilByte(hook.GetState(common.BytesToAddress(addr), common.BytesToHash(hash)).Bytes())
		},
		"getCodeHash": func(addr []byte) []byte {
			hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				panic("GetStateDbHook failed")
			}
			return wrapNilByte(hook.GetCodeHash(common.BytesToAddress(addr)).Bytes())
		},
		"getCodeSize": func(addr []byte) int32 {
			hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				panic("GetStateDbHook failed")
			}
			return int32(hook.GetCodeSize(common.BytesToAddress(addr)))
		},
		"getNonce": func(addr []byte) uint64 {
			hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				panic("GetStateDbHook failed")
			}
			return hook.GetNonce(common.BytesToAddress(addr))
		},
	}
}
