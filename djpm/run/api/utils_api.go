package api

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/log"
	"strings"
)

func (r *Register) utilApis() interface{} {
	return map[string]interface{}{
		"fromHexString": func(s string) []byte {
			s = strings.TrimPrefix(s, "0x")
			data, err := hex.DecodeString(s)
			if err != nil {
				return []byte{}
			}
			return data
		},
		"toHexString": func(data []byte) string {
			return hex.EncodeToString(data)
		},
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
