package api

import (
	"github.com/artela-network/artelasdk/types"
	"google.golang.org/protobuf/proto"
)

func (r *Register) scheduleApis() interface{} {
	return map[string]interface{}{
		"submit": func(arg []byte) bool {
			hook, err := types.GetScheduleHook()
			if err != nil || hook == nil {
				return false
			}
			sch := &types.Schedule{}
			if err := proto.Unmarshal(arg, sch); err != nil {
				return false
			}
			schedule := hook.SubmitSchedule(r.runnerContext, sch)
			return schedule.Success
		},
	}
}
