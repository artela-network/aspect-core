package api

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/artela-network/aspect-core/types"
)

func (r *Registry) evmCallAPIs() interface{} {
	return map[string]interface{}{
		"staticCall": func(request []byte) []byte {
			defaultResult := make([]byte, 0)
			hook, err := types.GetEvmHostHook(r.runnerContext.Ctx)
			if err != nil || hook == nil {
				return defaultResult
			}
			staticCall := &types.StaticCallRequest{}
			if unErr := proto.Unmarshal(request, staticCall); unErr != nil {
				return defaultResult
			}
			res := hook.StaticCall(r.runnerContext, staticCall)
			marshal, _ := proto.Marshal(res)
			return marshal
		},
		"jitCall": func(request []byte) ([]byte, error) {
			hook, err := types.GetEvmHostHook(r.runnerContext.Ctx)
			errRes := &types.JitInherentResponse{
				Success: false,
			}
			if err != nil || hook == nil {
				errRes.ErrorMsg = "evm host hook not init"
				errMsg, err := proto.Marshal(errRes)
				return errMsg, err
			}
			jitRequest := &types.JitInherentRequest{}
			if unErr := proto.Unmarshal(request, jitRequest); unErr != nil {
				errRes.ErrorMsg = fmt.Sprintf("jitRequest unmarshal error: %s", err.Error())
				errMsg, err := proto.Marshal(errRes)
				return errMsg, err
			}
			resp := hook.JITCall(r.runnerContext, jitRequest)
			marshal, err := proto.Marshal(resp)
			return marshal, err
		},
	}
}
