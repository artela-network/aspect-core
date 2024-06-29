package api

import (
	"github.com/artela-network/aspect-core/types"
	"google.golang.org/protobuf/proto"
)

func (r *Registry) traceAPIs() interface{} {
	return map[string]interface{}{
		"queryStateChange": func(rawQuery []byte) []byte {
			hook, err := types.GetAspectTraceHostHook(r.runnerContext.Ctx)
			if err != nil {
				panic("failed to init aspect runtime context host api: " + err.Error())
			}
			if hook == nil {
				panic("aspect runtime context host api not found")
			}

			query := &types.StateChangeQuery{}
			if err := proto.Unmarshal(rawQuery, query); err != nil {
				panic("failed to unmarshal query: " + err.Error())
			}

			return wrapNilByte(hook.QueryStateChange(r.runnerContext, query))
		},
		"queryCallTree": func(rawQuery []byte) []byte {
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
			//	panic("failed to unmarshal query: " + err.Error())
			//}

			return wrapNilByte(nil)
		},
	}
}
