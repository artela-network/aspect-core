package api

func (r *Register) scheduleApis() interface{} {
	return map[string]interface{}{
		"submit": func(arg []byte) bool {
			//submitSchedule(sch: ScheduleMsg): bool
			return false
		},
	}
}
