package api

func (r *Register) stateApis() interface{} {
	return map[string]interface{}{

		"setAspectState": func(key string, value string) bool {
			// setAspectState(key: string, value: string): bool
			return true
		},
		"getAspectState": func(key string) string {
			//getAspectState(key: string): string
			return ""
		},
		"removeAspectState": func(key string) bool {
			//getAspectState(key: string): string
			return false
		},
		"getProperty": func(key string) string {
			//getProperty(key: string): string
			return ""
		},
	}
}
