package api

import (
	"google.golang.org/protobuf/proto"

	"github.com/artela-network/aspect-core/types"
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
