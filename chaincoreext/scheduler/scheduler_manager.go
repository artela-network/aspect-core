package scheduler

import (
	"github.com/artela-network/artelasdk/types"
)

var globalManager *ScheduleManager

type ScheduleManager struct {
	store *types.AspectStore
	// cache schedule
	pool []*types.Schedule
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

func NewScheduleManager(store *types.AspectStore) *ScheduleManager {
	manager := ScheduleManager{
		store: store,
		pool:  nil,
	}
	schedules := make([]*types.Schedule, 0)
	// cache all active item by query
	query, err := manager.Query(1)
	if err != nil {
		schedules = append(schedules, query...)
	}
	manager.pool = schedules
	return &manager
}

func (manager ScheduleManager) Submit(req *types.Schedule) error {
	return nil
}

func (manager ScheduleManager) Query(Status uint8) ([]*types.Schedule, error) {
	return nil, nil
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
