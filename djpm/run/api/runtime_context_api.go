package api

import (
	"github.com/artela-network/artelasdk/types"
	"google.golang.org/protobuf/proto"
)

func (r *Register) contextCallApis() interface{} {
	return map[string]interface{}{
		"get": func(s []byte) []byte {
			hook, err := types.GetAspectRuntimeHook()
			if err != nil || hook == nil {
				response := types.NewContextQueryResponse(false, "hook not init")
				marshal, _ := proto.Marshal(response)
				return marshal
			}
			sch := &types.ContextQueryRequest{}
			if err := proto.Unmarshal(s, sch); err != nil {
				response := types.NewContextQueryResponse(false, "Unmarshal failed.please check input")
				marshal, _ := proto.Marshal(response)
				return marshal
			}
			result := hook.Get(r.runnerContext, sch)
			marshal, _ := proto.Marshal(result)
			return marshal
		},
		"setAspectContext": func(key string, value string) bool {
			hook, err := types.GetAspectRuntimeHook()
			if err != nil || hook == nil {
				return false
			}

			return hook.SetAspectContext(r.runnerContext, key, value)
		},
		"aspectId": func() string {
			return r.runnerContext.AspectId.String()
		},
	}
}
