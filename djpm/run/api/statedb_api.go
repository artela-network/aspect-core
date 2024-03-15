package api

import (
	types2 "github.com/artela-network/aspect-runtime/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/artela-network/aspect-core/types"
)

func (r *Registry) stateDBAPIs() map[string]*types2.HostFuncWithGasRule {
	return map[string]*types2.HostFuncWithGasRule{
		"getBalance": {
			Func: func(addr []byte) ([]byte, error) {
				hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					panic("GetStateDbHook failed")
				}
				balance := hook.GetBalance(common.BytesToAddress(addr))
				return wrapNilByte(balance.Bytes()), nil
			},
			GasRule: types2.NewStaticGasRule(2000),
		},
		"getState": {
			Func: func(addr []byte, hash []byte) ([]byte, error) {
				hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					panic("GetStateDbHook failed")
				}
				return wrapNilByte(hook.GetState(common.BytesToAddress(addr), common.BytesToHash(hash)).Bytes()), nil
			},
			GasRule: types2.NewStaticGasRule(5000),
		},
		"getCodeHash": {
			Func: func(addr []byte) ([]byte, error) {
				hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					panic("GetStateDbHook failed")
				}
				return wrapNilByte(hook.GetCodeHash(common.BytesToAddress(addr)).Bytes()), nil
			},
			GasRule: types2.NewStaticGasRule(40000),
		},
		"getCodeSize": {
			Func: func(addr []byte) (uint64, error) {
				hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					panic("GetStateDbHook failed")
				}
				return uint64(hook.GetCodeSize(common.BytesToAddress(addr))), nil
			},
			GasRule: types2.NewStaticGasRule(40000),
		},
		"getNonce": {
			Func: func(addr []byte) (uint64, error) {
				hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					panic("GetStateDbHook failed")
				}
				return hook.GetNonce(common.BytesToAddress(addr)), nil
			},
			GasRule: types2.NewStaticGasRule(40000),
		},
		"hasSuicided": {
			Func: func(addr []byte) (bool, error) {
				hook, err := types.GetStateDbHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					panic("GetStateDbHook failed")
				}
				return hook.HasSuicided(common.BytesToAddress(addr)), nil
			},
			GasRule: types2.NewStaticGasRule(40000),
		},
	}
}
