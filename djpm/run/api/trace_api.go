package api

import (
	"github.com/artela-network/aspect-core/types"
	types2 "github.com/artela-network/aspect-runtime/types"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func (r *Registry) traceAPIs() map[string]*types2.HostFuncWithGasRule {
	return map[string]*types2.HostFuncWithGasRule{
		"queryStateChange": {
			Func: func(rawQuery []byte) ([]byte, error) {
				hook, err := types.GetAspectTraceHostHook(r.runnerContext.Ctx)
				if err != nil {
					panic("failed to init aspect runtime context host api: " + err.Error())
				}
				if hook == nil {
					panic("aspect runtime context host api not found")
				}

				query := &types.StateChangeQuery{}
				if err := proto.Unmarshal(rawQuery, query); err != nil {
					return nil, errors.Wrap(err, "failed to unmarshal state change query")
				}

				res, err := hook.QueryStateChange(r.runnerContext, query)
				if err != nil {
					return nil, err
				}

				return wrapNilByte(res), nil
			},
			GasRule: types2.NewStaticGasRule(10000000),
		},
		"queryCallTree": {
			Func: func(rawQuery []byte) ([]byte, error) {
				//hook, err := types.GetAspectTraceHostHook(r.runnerContext.Ctx)
				//if err != nil {
				//	panic("failed to init aspect runtime context host api: " + err.Error())
				//}
				//if hook == nil {
				//	panic("aspect runtime context host api not found")
				//}
				//
				//query := &types.CallTreeQuery{}
				//if err := proto.Unmarshal(rawQuery, query); err != nil {
				//	return nil, errors.Wrap(err, "failed to unmarshal call tree query")
				//}
				//
				//res, err := hook.QueryCallTree(r.runnerContext, query)
				//if err != nil {
				//	return nil, err
				//}

				return wrapNilByte(nil), nil
			},
			GasRule: types2.NewStaticGasRule(10000000),
		},
	}
}
