package scheduler

import (
	"crypto/sha256"

	"github.com/artela-network/artelasdk/types"
)

// TxKey is same with the hash defined in cometbft
const TxKeySize = sha256.Size

type (
	TxKey [TxKeySize]byte

	TaskManager struct {
		// cached txs，key: txhash
		txs       map[TxKey][]byte
		schedules map[TxKey]*types.ScheduleId
	}
)

var globalTask *TaskManager

func NewTaskManager(height int64, nonce uint64, chainId string) error {
	manager := TaskManager{
		txs:       make(map[TxKey][]byte),
		schedules: make(map[TxKey]*types.ScheduleId),
	}
	return manager.genTxPool(height, nonce, chainId)
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
		} else {
			schedule.Tx.BlockNumber = height
			schedule.Tx.Nonce = nonce
			needRetry = true
		}

		// check if need to retry or height satisfy periodic
		if needRetry || (height >= int64(schedule.StartBlock) &&
			(height-int64(schedule.StartBlock))%int64(schedule.EveryNBlock) == 0) {
			// generate a new tx from schedule
			_, tx, err := ScheduleManagerInstance().WrapTransition(schedule.Tx)
			if err != nil {
				return err
			}

			key := getTxKey(tx)
			// save to task pool
			task.txs[key] = tx
			task.schedules[key] = schedule.Id
		}
	}
	return nil
}

// GetTxs return the scheduled transactions
func (task *TaskManager) GetTxs() [][]byte {
	txs := make([][]byte, len(task.txs))
	for _, tx := range task.txs {
		txs = append(txs, tx)
	}
	return txs
}

// return left tx
func (task *TaskManager) Confirm(txs [][]byte) error {
	// configrm all the tansactions that in block
	for _, tx := range txs {
		key := getTxKey(tx)
		delete(task.txs, key)

		// set retry to false, clear the try task
		schID := task.schedules[key]
		if err := ScheduleManagerInstance().ClearScheduleTry(schID); err != nil {
			return err
		}
		delete(task.schedules, key)
	}

	// not confirmed tasks
	for key, _ := range task.txs {
		// check and update schedule state
		schID := task.schedules[key]
		globalManager.StoreScheduleTry(schID, true, 0, "")
	}

	// we do not need to clear the task, all the task will be renew for next proposal.
	return nil
}

// getTxKey use sha256 get the tx hash, which is consistent with the cosmos mempool
func getTxKey(tx []byte) TxKey {
	return sha256.Sum256(tx)
}
