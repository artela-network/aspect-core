package api

import (
	"github.com/ethereum/go-ethereum/log"
)

func (r *Registry) utilAPIs() interface{} {
	return map[string]interface{}{
		"revert": func(msg string) {
			if r.errCallback != nil {
				r.errCallback(msg)
			}
		},
		"sLog": func(s string) {
			log.Info(s)
		},
	}
}
