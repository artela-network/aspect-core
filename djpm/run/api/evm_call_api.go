package api

import (
	types2 "github.com/artela-network/aspect-runtime/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	"github.com/artela-network/aspect-core/types"
)

func (r *Registry) evmCallAPIs() map[string]*types2.HostFuncWithGasRule {
	return map[string]*types2.HostFuncWithGasRule{
		"staticCall": {
			Func: func(request []byte) ([]byte, error) {
				hook, err := types.GetEvmHostHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					panic("failed to init evm host hook")
				}
				staticCall := &types.StaticCallRequest{}
				if err := proto.Unmarshal(request, staticCall); err != nil {
					return nil, errors.Wrap(err, "failed to unmarshal static call request")
				}
				res, err := hook.StaticCall(r.runnerContext, staticCall)
				if err != nil {
					return nil, err
				}

				marshal, err := proto.Marshal(res)
				if err != nil {
					panic("failed to marshal static call response, err: " + err.Error())
				}
				return marshal, nil
			},
			GasRule: types2.NewDynamicGasRule(30000, 37500),
		},
		"jitCall": {
			Func: func(request []byte) ([]byte, error) {
				hook, err := types.GetEvmHostHook(r.runnerContext.Ctx)
				if err != nil || hook == nil {
					return nil, err
				}
				jitRequest := &types.JitInherentRequest{}
				if err := proto.Unmarshal(request, jitRequest); err != nil {
					return nil, errors.Wrap(err, "failed to unmarshal jit call request")
				}
				resp, err := hook.JITCall(r.runnerContext, jitRequest)
				if err != nil {
					return nil, err
				}

				marshal, err := proto.Marshal(resp)
				if err != nil {
					panic("failed to marshal jit call response, err: " + err.Error())
				}

				return marshal, err
			},
			GasRule: types2.NewDynamicGasRule(30000, 75000),
		},
	}
}
