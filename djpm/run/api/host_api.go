package api

import (
	"google.golang.org/protobuf/proto"

	"github.com/artela-network/aspect-core/types"
)

func (r *Register) hostApi() interface{} {
	return map[string]interface{}{
		"get": func(key string) []byte {
			hook, err := types.GetRuntimeHostHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				response := types.NewContextQueryResponse(false, "hook not init")
				marshal, _ := proto.Marshal(response)
				return marshal
			}
			result := hook.GetContext(r.runnerContext, key)
			marshal, _ := proto.Marshal(result)
			return marshal
		},
		"remove": func(s []byte) bool {
			hook, err := types.GetRuntimeHostHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				return false
			}
			sch := &types.ContextRemoveRequest{}
			if proErr := proto.Unmarshal(s, sch); proErr != nil {
				return false
			}
			return hook.Remove(r.runnerContext, sch)
		},
		"set": func(s []byte) bool {
			hook, err := types.GetRuntimeHostHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				return false
			}
			sch := &types.ContextSetRequest{}
			if proErr := proto.Unmarshal(s, sch); proErr != nil {
				return false
			}
			return hook.Set(r.runnerContext, sch)
		},
		"query": func(s []byte) []byte {
			hook, err := types.GetRuntimeHostHook(r.runnerContext.Ctx)
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
			result := hook.Query(r.runnerContext, sch)
			marshal, _ := proto.Marshal(result)
			return marshal
		},
		"aspectId": func() string {
			return r.runnerContext.AspectId.String()
		},
	}
}
