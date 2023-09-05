package api

func (r *Register) contextCallApis() interface{} {
	return map[string]interface{}{
		"get": func(s string) []byte {
			//contextQuery(query: string): ContextQueryResponse | null
			return nil
		},
		"setAspectContext": func(key string, value string) bool {
			// setAspectContext(key: string, value: string): bool
			return true
		},
	}
}
