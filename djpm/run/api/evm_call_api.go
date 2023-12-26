package api

import (
	"google.golang.org/protobuf/proto"

	"github.com/artela-network/aspect-core/types"
)

func (r *Registry) evmCallAPIs() interface{} {
	return map[string]interface{}{
		"staticCall": func(request []byte) []byte {
			hook, err := types.GetEvmHostHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				panic("failed to init evm host hook")
			}
			staticCall := &types.StaticCallRequest{}
			if err := proto.Unmarshal(request, staticCall); err != nil {
				panic("failed to unmarshal static call request, err: " + err.Error())
			}
			res := hook.StaticCall(r.runnerContext, staticCall)
			marshal, err := proto.Marshal(res)
			if err != nil {
				panic("failed to marshal static call response, err: " + err.Error())
			}
			return marshal
		},
		"jitCall": func(request []byte) ([]byte, error) {
			hook, err := types.GetEvmHostHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				return nil, err
			}
			jitRequest := &types.JitInherentRequest{}
			if err := proto.Unmarshal(request, jitRequest); err != nil {
				return nil, err
			}
			resp := hook.JITCall(r.runnerContext, jitRequest)
			marshal, err := proto.Marshal(resp)
			return marshal, err
		},
	}
}
