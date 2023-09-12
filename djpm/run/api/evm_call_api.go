package api

import (
	"github.com/artela-network/artelasdk/types"
	"google.golang.org/protobuf/proto"
)

func (r *Register) evmCallApis() interface{} {
	return map[string]interface{}{

		"staticCall": func(request []byte) []byte {

			hook, err := types.GetEvmHostHook()
			if err != nil || hook == nil {
				errRes := types.ErrCallMessageResponse(err)
				marshal, _ := proto.Marshal(errRes)
				return marshal
			}

			ethMsg := &types.EthTransaction{}
			if unErr := proto.Unmarshal(request, ethMsg); unErr != nil {
				errRes := types.ErrCallMessageResponse(unErr)
				marshal, _ := proto.Marshal(errRes)
				return marshal
			}
			call := hook.StaticCall(r.runnerContext, ethMsg)
			marshal, _ := proto.Marshal(call)
			return marshal
		},
		"jitCall": func(request []byte) []byte {
			hook, err := types.GetEvmHostHook()
			errRes := &types.JitInherentResponse{
				Success: false,
			}
			if err != nil || hook == nil {
				marshal, _ := proto.Marshal(errRes)
				return marshal
			}
			jitRequest := &types.JitInherentRequest{}
			if unErr := proto.Unmarshal(request, jitRequest); unErr != nil {
				marshal, _ := proto.Marshal(errRes)
				return marshal
			}
			call := hook.JITCall(r.runnerContext, jitRequest)
			marshal, _ := proto.Marshal(call)
			return marshal
		},
	}
}
