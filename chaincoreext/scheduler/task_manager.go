package scheduler

import "github.com/artela-network/artelasdk/types"

type TaskManager struct {

	// cache schedule，key: txhash
	pool map[string]*types.ScheduleTask
}

// ADD TX CALL
func (pool *TaskManager) GetTxs(height uint64) []*types.AspTransaction {

	// ScheduleManagerInstance.GetActiveSchedule()
	// 判断
	// 生成交易
	//  检查是否有足够的费用
	// add pool
	// return txs[]
	return nil
}

// return left tx
func (pool *TaskManager) Confirm(txHash []string) []*types.AspTransaction {
	// 1. configrm

	// 2. left, update

	// 3、clear pool
	return nil
}
