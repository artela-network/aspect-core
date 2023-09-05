package api

func (r *Register) stateDbApis() interface{} {
	return map[string]interface{}{
		"getBalance": func(addr string) string {
			//getBalance(key: string): string
			return ""
		},
		"getState": func(addr string, hash string) string {
			//getState(addr: string, hash: string): string
			return ""
		},
		"getRefund": func() int64 {
			//getRefund(): i64
			return 0
		},
		"getCodeHash": func(addr string) string {
			//getCodeHash(addr: string): string
			return ""
		},
		"getNonce": func(addr string) int64 {
			//getNonce(addr: string): i64
			return 0
		},
	}
}
