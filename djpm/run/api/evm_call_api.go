package api

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/artela-network/aspect-core/types"
)

func (r *Register) evmCallApis() interface{} {
	return map[string]interface{}{
		"staticCall": func(request []byte) []byte {
			defaultResult := make([]byte, 0)
			hook, err := types.GetEvmHostHook()
			if err != nil || hook == nil {
				return defaultResult
			}
			ethMsg := &types.EthMessage{}
			if unErr := proto.Unmarshal(request, ethMsg); unErr != nil {
				return defaultResult
			}
			call := hook.StaticCall(r.runnerContext, ethMsg)
			marshal, _ := proto.Marshal(call)
			return marshal
		},
		"jitCall": func(request []byte) ([]byte, error) {
			hook, err := types.GetEvmHostHook()
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
