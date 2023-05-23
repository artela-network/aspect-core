package scheduler

import "github.com/artela-network/artelasdk/types"

type TaskManager struct {

	// cache schedule，key: txhash
	pool map[string]*types.ScheduleTask
}

var globalTask *TaskManager

func NewTaskManager(height int64, nonce uint64, chainId string) {
	manager := TaskManager{}
	manager.genTxPool(height, nonce, chainId)

}

// ADD TX CALL
func (task *TaskManager) genTxPool(height int64, nonce uint64, chainId string) []*types.AspTransaction {

	// ScheduleManagerInstance.GetActiveSchedule()
	schedule := globalManager.GetActiveSchedule()

	transactions := make([]*types.AspTransaction, 0)
	if schedule == nil || len(schedule) == 0 {
		task.pool = nil
		return transactions
	}
	//todo 没有判断高度
	//todo 没有判断重试

	taskMap := make(map[string]*types.ScheduleTask, 0)
	for _, s := range schedule {
		s.Tx.BlockNumber = height
		s.Tx.Nonce = nonce
		hash := s.Tx.TxHash()
		transactions = append(transactions, s.Tx)
		task := &types.ScheduleTask{
			Id:          s.Id,
			BlockHeight: height,
			TxHash:      hash,
			Nonce:       nonce,
			Tx:          s.Tx,
		}
		taskMap[hash] = task
		globalManager.ExecRecord(s.Id, height, hash)
	}

	task.pool = taskMap

	//  检查是否有足够的费用
	return transactions
}

// return left tx
func (task *TaskManager) Confirm(txHash []string) []*types.AspTransaction {
	// 1. configrm
	for _, hash := range txHash {
		delete(task.pool, hash)
	}
	transactions := make([]*types.AspTransaction, 0)
	if len(task.pool) == 0 {
		return transactions
	}

	// left task,fail
	for _, scheduleTask := range task.pool {
		globalManager.StoreScheduleTry(scheduleTask.Id, true, 0, "")
		transactions = append(transactions, scheduleTask.Tx)
		delete(task.pool, scheduleTask.TxHash)
	}
	return transactions
}
