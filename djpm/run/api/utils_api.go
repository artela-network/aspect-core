package api

import (
	"github.com/artela-network/aspect-runtime/types"
	"github.com/ethereum/go-ethereum/log"
)

func (r *Registry) utilAPIs() map[string]*types.HostFuncWithGasRule {
	return map[string]*types.HostFuncWithGasRule{
		"revert": {
			Func: func(msg string) {
				if r.errCallback != nil {
					r.errCallback(msg)
				}
			},
			GasRule: types.NewStaticGasRule(1),
		},
		"sLog": {
			Func: func(s string) {
				log.Info(s)
			},
			GasRule: types.NewStaticGasRule(1),
		},
	}
}
