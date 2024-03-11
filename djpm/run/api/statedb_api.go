package api

import (
	types2 "github.com/artela-network/aspect-runtime/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/artela-network/aspect-core/types"
)

func (r *Registry) stateDBAPIs() map[string]*types2.HostFuncWithGasRule {
	return map[string]*types2.HostFuncWithGasRule{
		"getBalance": {
			Func: func(addr []byte) []byte {
				hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					panic("GetStateDbHook failed")
				}
				balance := hook.GetBalance(common.BytesToAddress(addr))
				return wrapNilByte(balance.Bytes())
			},
			GasRule: types2.NewStaticGasRule(1),
		},
		"getState": {
			Func: func(addr []byte, hash []byte) []byte {
				hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					panic("GetStateDbHook failed")
				}
				return wrapNilByte(hook.GetState(common.BytesToAddress(addr), common.BytesToHash(hash)).Bytes())
			},
			GasRule: types2.NewStaticGasRule(1),
		},
		"getCodeHash": {
			Func: func(addr []byte) []byte {
				hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					panic("GetStateDbHook failed")
				}
				return wrapNilByte(hook.GetCodeHash(common.BytesToAddress(addr)).Bytes())
			},
			GasRule: types2.NewStaticGasRule(1),
		},
		"getCodeSize": {
			Func: func(addr []byte) uint64 {
				hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					panic("GetStateDbHook failed")
				}
				return uint64(hook.GetCodeSize(common.BytesToAddress(addr)))
			},
			GasRule: types2.NewStaticGasRule(1),
		},
		"getNonce": {
			Func: func(addr []byte) uint64 {
				hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					panic("GetStateDbHook failed")
				}
				return hook.GetNonce(common.BytesToAddress(addr))
			},
			GasRule: types2.NewStaticGasRule(1),
		},
		"hasSuicided": {
			Func: func(addr []byte) bool {
				hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					panic("GetStateDbHook failed")
				}
				return hook.HasSuicided(common.BytesToAddress(addr))
			},
			GasRule: types2.NewStaticGasRule(1),
		},
	}
}
