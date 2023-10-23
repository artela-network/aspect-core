package scheduler

import (
	"crypto/sha256"

	"github.com/ethereum/go-ethereum/common"

	"github.com/pkg/errors"

	"github.com/artela-network/artelasdk/types"
)

// TxKeySize TxKey is same with the hash defined in cometbft
const TxKeySize = sha256.Size

type (
	TxKey [TxKeySize]byte

	TaskManager struct {
		// cached txs，key: txhash
		scheduleTasks map[TxKey]*types.ScheduleTask
		ethTxIndexMap map[string]TxKey
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
	if globalTask != nil && len(globalTask.scheduleTasks) > 0 {
		err := globalTask.genTxPool(height, nonce, chainId)
		return err
	}
	manager := &TaskManager{
		scheduleTasks: make(map[TxKey]*types.ScheduleTask),
		ethTxIndexMap: make(map[string]TxKey),
	}
	err := manager.genTxPool(height, nonce, chainId)

	globalTask = manager

	return err
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
			task.ethTxIndexMap[hash.String()] = key
		}
	}
	return nil
}

// GetTxs return the scheduled transactions
func (task *TaskManager) GetTxs() [][]byte {
	txs := make([][]byte, 0, len(task.scheduleTasks))
	for _, task := range task.scheduleTasks {
		txs = append(txs, task.SdkTx)
	}
	return txs
}

// GetTxs return the scheduled transactions
func (task *TaskManager) IsScheduleTx(hash common.Hash) bool {
	_, ok := task.ethTxIndexMap[hash.String()]
	return ok
}

func (task *TaskManager) GetFromAddr(hash common.Hash) string {
	key := task.ethTxIndexMap[hash.String()]
	return task.scheduleTasks[key].Schedule.Tx.From
}

func (task *TaskManager) Remove(tx []byte) error {
	if task == nil || task.scheduleTasks == nil || len(task.scheduleTasks) == 0 {
		return nil
	}
	key := getTxKey(tx)

	// check is task tx
	scheduleTask, ok := task.scheduleTasks[key]
	if !ok {
		return nil
	}
	if err := ScheduleManagerInstance().ClearScheduleTry(scheduleTask.Schedule.Id); err != nil {
		return err
	}

	err := ScheduleManagerInstance().StoreScheduleExecResult(scheduleTask.Schedule.Id, scheduleTask.BlockHeight, scheduleTask.TxHash)
	if err != nil {
		return err
	}

	// Check the number of executions,or Close schedule
	execErr := ScheduleManagerInstance().CheckClose(scheduleTask.Schedule)
	if execErr != nil {
		return execErr
	}
	// clean pool
	delete(task.scheduleTasks, key)
	delete(task.ethTxIndexMap, scheduleTask.TxHash)
	return nil
}

func (task *TaskManager) Check() ([][]byte, error) {
	leftTxs := make([][]byte, len(task.scheduleTasks))
	// not confirmed tasks
	for key := range task.scheduleTasks {
		// check and update schedule state
		scheduleTask := task.scheduleTasks[key]

		try, err := ScheduleManagerInstance().GetScheduleTry(scheduleTask.Schedule.Id)
		if err != nil {
			return nil, err
		}
		if uint64(len(try.TaskTxs)+1) < scheduleTask.Schedule.MaxRetry {
			// try count less MaxRetry,then next  need try
			err := ScheduleManagerInstance().StoreScheduleTry(scheduleTask.Schedule.Id, true, scheduleTask.BlockHeight, scheduleTask.TxHash)
			if err != nil {
				return nil, err
			}
		} else {
			// try count more than maxRetry, Close next try
			err := ScheduleManagerInstance().StoreScheduleTry(scheduleTask.Schedule.Id, false, scheduleTask.BlockHeight, scheduleTask.TxHash)
			if err != nil {
				return nil, err
			}

			// add Fail txHash for placeholder
			errExec := ScheduleManagerInstance().StoreScheduleExecResult(scheduleTask.Schedule.Id, scheduleTask.BlockHeight, "F")
			if errExec != nil {
				return nil, errExec
			}

			// Check the number of executions,or Close schedule
			execErr := ScheduleManagerInstance().CheckClose(scheduleTask.Schedule)
			if execErr != nil {
				return nil, execErr
			}
		}
		delete(task.scheduleTasks, key)
		delete(task.ethTxIndexMap, scheduleTask.TxHash)
		leftTxs = append(leftTxs, scheduleTask.SdkTx)
	}

	// we do not need to clear the task, all the task will be renew for next proposal.
	return leftTxs, nil
}

// Confirm return left tx
func (task *TaskManager) Confirm(txs [][]byte) ([][]byte, error) {
	// confirm all the tansactions that in block
	if task == nil || task.scheduleTasks == nil || len(task.scheduleTasks) == 0 {
		return nil, nil
	}
	for _, tx := range txs {
		key := getTxKey(tx)

		// check is task tx
		scheduleTask, ok := task.scheduleTasks[key]
		if !ok {
			continue
		}

		if err := ScheduleManagerInstance().ClearScheduleTry(scheduleTask.Schedule.Id); err != nil {
			return nil, err
		}

		err := ScheduleManagerInstance().StoreScheduleExecResult(scheduleTask.Schedule.Id, scheduleTask.BlockHeight, scheduleTask.TxHash)
		if err != nil {
			return nil, err
		}

		// Check the number of executions,or Close schedule
		execErr := ScheduleManagerInstance().CheckClose(scheduleTask.Schedule)
		if execErr != nil {
			return nil, execErr
		}
		// clean pool
		delete(task.scheduleTasks, key)
		delete(task.ethTxIndexMap, scheduleTask.TxHash)
	}

	leftTxs := make([][]byte, len(task.scheduleTasks))

	// not confirmed tasks
	for key := range task.scheduleTasks {
		// check and update schedule state
		scheduleTask := task.scheduleTasks[key]

		try, err := ScheduleManagerInstance().GetScheduleTry(scheduleTask.Schedule.Id)
		if err != nil {
			return nil, err
		}
		if uint64(len(try.TaskTxs)+1) < scheduleTask.Schedule.MaxRetry {
			// try count less MaxRetry,then next  need try
			storeErr := ScheduleManagerInstance().StoreScheduleTry(scheduleTask.Schedule.Id, true, scheduleTask.BlockHeight, scheduleTask.TxHash)
			if storeErr != nil {
				return nil, storeErr
			}
		} else {
			// try count more than maxRetry, Close next try
			tryErr := ScheduleManagerInstance().StoreScheduleTry(scheduleTask.Schedule.Id, false, scheduleTask.BlockHeight, scheduleTask.TxHash)
			if tryErr != nil {
				return nil, tryErr
			}

			// add Fail txHash for placeholder
			err := ScheduleManagerInstance().StoreScheduleExecResult(scheduleTask.Schedule.Id, scheduleTask.BlockHeight, "F")
			if err != nil {
				return nil, err
			}

			// Check the number of executions,or Close schedule
			execErr := ScheduleManagerInstance().CheckClose(scheduleTask.Schedule)
			if execErr != nil {
				return nil, execErr
			}
		}
		delete(task.scheduleTasks, key)
		delete(task.ethTxIndexMap, scheduleTask.TxHash)

		leftTxs = append(leftTxs, scheduleTask.SdkTx)

	}

	// we do not need to clear the task, all the task will be renew for next proposal.
	return leftTxs, nil
}

// getTxKey use sha256 get the tx hash, which is consistent with the cosmos mempool
func getTxKey(tx []byte) TxKey {
	return sha256.Sum256(tx)
}
