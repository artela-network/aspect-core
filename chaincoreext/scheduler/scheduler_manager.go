package scheduler

import (
	"github.com/artela-network/artelasdk/types"
)

var globalManager *ScheduleManager

type ScheduleManager struct {
	Store types.AspectStore
	// cache schedule
	Pool []*types.Schedule
}

/**
1、 key: "Schedule"+ Status  ， Value： [id,id,id...]
2、 key: id                 ,  Value:  Schedule bytes
3、 key: id                 ,   {ConfimTxs:[{blockheight,txhash},{blockheight,txhash}..]， count: 2}    // exec result
4、 key: id                 ,  needRetry: false ，tryCount: 1, startblockheight：100
*/

func ScheduleManagerInstance() *ScheduleManager {
	if globalManager == nil {
		panic("aspcect instance not init,please exec NewAspect() first ")
	}
	return globalManager
}

func NewScheduleManager(store types.AspectStore) *ScheduleManager {
	manager := ScheduleManager{
		Store: store,
		Pool:  nil,
	}
	schedules := make([]*types.Schedule, 0)
	// cache all active item by query
	query, err := manager.Query(1)
	if err != nil {
		schedules = append(schedules, query...)
	}
	manager.Pool = schedules
	return &manager
}

func (manager ScheduleManager) Submit(req *types.Schedule) error {

	return nil
}

func (manager ScheduleManager) Query(status types.ScheduleStatus) ([]*types.Schedule, error) {
	ids, err := manager.GetScheduleView(int32(status))
	if err != nil {
		return nil, err
	}
	schedules := make([]*types.Schedule, 0)
	for i := range ids {
		id := ids[i]
		schedule, getErr := manager.GetSchedule(id)
		if getErr != nil {
			return nil, getErr
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (manager ScheduleManager) Close(req *types.ScheduleId) bool {
	return false
}

// begin block call
func (manager ScheduleManager) GetActiveSchedule() ([]*types.Schedule, error) {
	// new taskpool
	return nil, nil
}

func (manager ScheduleManager) ExecRecord(id *types.ScheduleId, blockHeight uint64, txHash string) {
	// add count
}
