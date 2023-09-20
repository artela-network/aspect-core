package api

import (
	"github.com/artela-network/artelasdk/types"
	"github.com/pkg/errors"
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
