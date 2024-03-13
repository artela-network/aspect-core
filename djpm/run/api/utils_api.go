package api

import (
	"github.com/artela-network/aspect-runtime/types"
	"github.com/ethereum/go-ethereum/log"
)

func (r *Registry) utilAPIs() map[string]*types.HostFuncWithGasRule {
	return map[string]*types.HostFuncWithGasRule{
		"revert": {
			Func: func(msg string) error {
				if r.errCallback != nil {
					r.errCallback(msg)
				}
				return nil
			},
			GasRule: types.NewStaticGasRule(1),
		},
		"sLog": {
			Func: func(s string) error {
				log.Info(s)
				return nil
			},
			GasRule: types.NewDynamicGasRule(1000, 3750),
		},
	}
}
