package api

func (r *Register) evmCallApis() interface{} {
	return map[string]interface{}{

		"staticCall": func(request []byte) []byte {

			// staticCall(request: CallMessageRequest): CallMessageResponse  | null

			return nil
		},
		"jitCall": func(request []byte) []byte {
			// jitCall(request: CallMessageRequest): CallMessageResponse
			return nil
		},
	}
}
