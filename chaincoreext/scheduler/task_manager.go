package scheduler

import (
	"crypto/sha256"

	"github.com/artela-network/artelasdk/types"
	"github.com/pkg/errors"
)

// TxKeySize TxKey is same with the hash defined in cometbft
const TxKeySize = sha256.Size

type (
	TxKey [TxKeySize]byte

	TaskManager struct {
		// cached txs，key: txhash
		scheduleTasks map[TxKey]*types.ScheduleTask
	}
)

var globalTask *TaskManager

// TODO, need a better way call to schedule in cometbft.
func TaskInstance() *TaskManager {
	// globalTask is ensure to not empty.
	return globalTask
}

func NewTaskManager(height int64, nonce uint64, chainId string) error {
	if ScheduleManagerInstance() == nil {
		return errors.New("ScheduleManager instance not init,please exec NewScheduleManager() first")
	}
	manager := &TaskManager{
		scheduleTasks: make(map[TxKey]*types.ScheduleTask),
	}
	err := manager.genTxPool(height, nonce, chainId)

	globalTask = manager

	return err
}

func TaskManagerInstance() *ScheduleManager {
	if globalManager == nil {
		panic("task manager instance not init,please exec NewTaskManager() first ")
	}
	return globalManager
}

// genTxPool load transaction from scheduleManager and insert to pool
func (task *TaskManager) genTxPool(height int64, nonce uint64, chainId string) error {
	schedules := ScheduleManagerInstance().GetActiveSchedule()
	if len(schedules) == 0 {
		// no schedule, skip generating transactions
		return nil
	}

	for _, schedule := range schedules {
		needRetry := false

		if schedule.Status == types.ScheduleStatus_Close {
			continue
		}

		// get the retry flag
		tryTasks, err := ScheduleManagerInstance().GetScheduleTry(schedule.Id)
		if err != nil {
			return err
		}

		if len(tryTasks.TaskTxs) == int(schedule.MaxRetry) {
			// reach the max of retry count, clear the try storage
			if err := ScheduleManagerInstance().ClearScheduleTry(schedule.Id); err != nil {
				return err
			}
		}

		schedule.Tx.BlockNumber = height
		schedule.Tx.Nonce = nonce
		needRetry = tryTasks.NeedRetry

		// check if need to retry or height satisfy periodic
		if needRetry || (height >= int64(schedule.StartBlock) &&
			(height-int64(schedule.StartBlock))%int64(schedule.EveryNBlock) == 0) {
			// generate a new tx from schedule
			hash, tx, err := ScheduleManagerInstance().WrapTransition(schedule.Tx)
			if err != nil {
				return err
			}
			key := getTxKey(tx)
			scheduleTask := &types.ScheduleTask{
				Schedule:    schedule,
				BlockHeight: height,
				TxHash:      hash.String(),
				TxNonce:     nonce,
				SdkTx:       tx,
			}
			// save to task pool
			task.scheduleTasks[key] = scheduleTask
		}
	}
	return nil
}

// GetTxs return the scheduled transactions
func (task *TaskManager) GetTxs() [][]byte {
	txs := make([][]byte, len(task.scheduleTasks))
	for _, task := range task.scheduleTasks {
		txs = append(txs, task.SdkTx)
	}
	return txs
}

// Confirm return left tx
func (task *TaskManager) Confirm(txs [][]byte) ([][]byte, error) {
	// configrm all the tansactions that in block
	for _, tx := range txs {
		key := getTxKey(tx)

		// check is task tx
		_, ok := task.scheduleTasks[key]
		if !ok {
			continue
		}

		// set retry to false, clear the try task
		scheduleTask := task.scheduleTasks[key]

		if err := ScheduleManagerInstance().ClearScheduleTry(scheduleTask.Schedule.Id); err != nil {
			return nil, err
		}

		err := ScheduleManagerInstance().StoreScheduleExecResult(scheduleTask.Schedule.Id, scheduleTask.BlockHeight, scheduleTask.TxHash)
		if err != nil {
			return nil, err
		}

		//Check the number of executions,or Close schedule
		execErr := ScheduleManagerInstance().CheckClose(scheduleTask.Schedule)
		if execErr != nil {
			return nil, execErr
		}
		//clean pool
		delete(task.scheduleTasks, key)
	}

	leftTxs := make([][]byte, len(task.scheduleTasks))

	// not confirmed tasks
	for key, _ := range task.scheduleTasks {
		// check and update schedule state
		scheduleTask := task.scheduleTasks[key]

		try, err := ScheduleManagerInstance().GetScheduleTry(scheduleTask.Schedule.Id)
		if err != nil {
			return nil, err
		}
		if uint64(len(try.TaskTxs)+1) < scheduleTask.Schedule.MaxRetry {
			//try count less MaxRetry,then next  need try
			ScheduleManagerInstance().StoreScheduleTry(scheduleTask.Schedule.Id, true, scheduleTask.BlockHeight, scheduleTask.TxHash)
		} else {
			// try count more than maxRetry, Close next try
			ScheduleManagerInstance().StoreScheduleTry(scheduleTask.Schedule.Id, false, scheduleTask.BlockHeight, scheduleTask.TxHash)

			// add Fail txHash for placeholder
			err := ScheduleManagerInstance().StoreScheduleExecResult(scheduleTask.Schedule.Id, scheduleTask.BlockHeight, "F")
			if err != nil {
				return nil, err
			}

			//Check the number of executions,or Close schedule
			execErr := ScheduleManagerInstance().CheckClose(scheduleTask.Schedule)
			if execErr != nil {
				return nil, execErr
			}
		}
		leftTxs = append(leftTxs, scheduleTask.SdkTx)

	}

	// we do not need to clear the task, all the task will be renew for next proposal.
	return leftTxs, nil
}

// getTxKey use sha256 get the tx hash, which is consistent with the cosmos mempool
func getTxKey(tx []byte) TxKey {
	return sha256.Sum256(tx)
}
