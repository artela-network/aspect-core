package scheduler

import (
	"github.com/artela-network/artelasdk/types"
	"github.com/ethereum/go-ethereum/common"
)

var globalManager *ScheduleManager

type ScheduleManager struct {
	Store types.AspectStore
	// cache schedule
	Pool   []*types.Schedule
	WrapTx func(tx *types.EthTransaction) (common.Hash, []byte, error)
}

/**
1、 key: "Schedule"+ Status  ， Value： [id,id,id...]
2、 key: id                 ,  Value:  Schedule bytes
3、 key: id                 ,   {ConfimTxs:[{blockheight,txhash},{blockheight,txhash}..]， count: 2}    // exec result
4、 key: id                 ,  needRetry: false ，tashTx:[{blockheight,txhash},{blockheight,txhash}..]
*/

func ScheduleManagerInstance() *ScheduleManager {
	if globalManager == nil {
		panic(" ScheduleManager instance not init,please exec NewScheduleManager() first ")
	}
	return globalManager
}

func NewScheduleManager(store types.AspectStore, wrapTx func(tx *types.EthTransaction) (common.Hash, []byte, error)) error {
	manager := ScheduleManager{
		Store:  store,
		Pool:   nil,
		WrapTx: wrapTx,
	}
	// cache all active item by query
	err := manager.initPool()
	if err != nil {
		return err
	}
	globalManager = &manager
	return nil
}

func (manager *ScheduleManager) Submit(req *types.Schedule) error {
	storeErr := manager.StoreScheduleView(req)
	if storeErr != nil {
		return storeErr
	}
	err := manager.StoreSchedule(req)
	if err != nil {
		return err
	}
	if req.Status == types.ScheduleStatus_Open {
		manager.addPool(req)
	}
	return nil
}

func (manager *ScheduleManager) Query(status types.ScheduleStatus) ([]*types.Schedule, error) {
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

func (manager *ScheduleManager) CheckClose(schedule *types.Schedule) error {
	//Check the number of executions
	result, execErr := ScheduleManagerInstance().GetScheduleExecResult(schedule.Id)
	if execErr != nil {
		return execErr
	}
	// When the number of executions is sufficient, Close schedule
	if uint64(len(result.GetConfirmTxs())) == schedule.Count {
		manager.Close(schedule.Id)
	}
	return nil
}

// begin block call
func (manager ScheduleManager) GetActiveSchedule() []*types.Schedule {
	return manager.Pool
}

func (manager *ScheduleManager) Close(scheduleId *types.ScheduleId) error {
	schedule, err := manager.GetSchedule(scheduleId)
	if err != nil {
		return err
	}
	if schedule.Status == types.ScheduleStatus_Close {
		return nil
	}

	delErr := manager.DeleteScheduleView(uint32(types.ScheduleStatus_Open), scheduleId)
	if delErr != nil {
		return delErr
	}
	schedule.Status = types.ScheduleStatus_Close
	storeErr := manager.StoreSchedule(schedule)
	if storeErr != nil {
		return storeErr
	}
	storeViewErr := manager.StoreScheduleView(schedule)
	if storeViewErr != nil {
		return storeViewErr
	}
	manager.rmPool(scheduleId)
	return nil
}

func (manager *ScheduleManager) initPool() error {
	query, err := manager.Query(types.ScheduleStatus_Open)
	if err != nil {
		return err
	}
	manager.Pool = query
	return nil
}

func (manager *ScheduleManager) addPool(schedule *types.Schedule) {
	if schedule == nil {
		return
	}
	manager.Pool = append(manager.Pool, schedule)
}

func (manager *ScheduleManager) rmPool(id *types.ScheduleId) {
	for i, schedule := range manager.Pool {
		if schedule.Id.String() == id.String() {
			manager.Pool = append(manager.Pool[0:i], manager.Pool[i+1:len(manager.Pool)]...)
		}
	}
}
func (manager *ScheduleManager) WrapTransition(tx *types.EthTransaction) (common.Hash, []byte, error) {
	return manager.WrapTx(tx)
}
