package api

import (
	"github.com/pkg/errors"
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
				errMsg, _ := proto.Marshal(errRes)
				return errMsg, nil
			}
			jitRequest := &types.JitInherentRequest{}
			if unErr := proto.Unmarshal(request, jitRequest); unErr != nil {
				errRes.ErrorMsg = "jitRequest unmarshal error"
				errMsg, _ := proto.Marshal(errRes)
				return errMsg, nil
			}
			call := hook.JITCall(r.runnerContext, jitRequest)
			if !call.Success {
				errMsg, _ := proto.Marshal(errRes)
				return errMsg, errors.New(call.ErrorMsg)
			}
			marshal, _ := proto.Marshal(call)
			return marshal, nil
		},
	}
}
