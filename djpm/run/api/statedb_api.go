package api

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/artela-network/aspect-core/types"
)

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
		"getCodeSize": func(addr []byte) uint64 {
			hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				panic("GetStateDbHook failed")
			}
			address := common.BytesToAddress(addr)
			size := hook.GetCodeSize(address)
			return uint64(size)
		},
		"getNonce": func(addr []byte) uint64 {
			hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				panic("GetStateDbHook failed")
			}
			return hook.GetNonce(common.BytesToAddress(addr))
		},
		"hasSuicided": func(addr []byte) bool {
			hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				panic("GetStateDbHook failed")
			}
			return hook.HasSuicided(common.BytesToAddress(addr))
		},
	}
}
