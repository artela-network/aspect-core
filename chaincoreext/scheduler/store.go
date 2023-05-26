package scheduler

import (
	"encoding/json"
	"math/big"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	"github.com/artela-network/artelasdk/types"
)

const (
	ScheduleViewKeyPrefix       = "Schedule/View/"
	ScheduleKeyPrefix           = "Schedule/Data/"
	ScheduleRetryKeyPrefix      = "Schedule/Retry/"
	ScheduleExecResultKeyPrefix = "Schedule/ExecResult/"
)

func ScheduleIdKey(
	scheduleId *types.ScheduleId,
) []byte {
	var key []byte
	marshal, _ := proto.Marshal(scheduleId)
	key = append(key, marshal...)
	key = append(key, []byte("/")...)
	return key
}

func ScheduleViewKey(
	status int64,
) []byte {
	var newInt = big.NewInt(status) // int to big Int
	var statusBytes = newInt.Bytes()
	var key []byte
	key = append(key, []byte("ScheduleView")...)
	key = append(key, []byte("/")...)
	key = append(key, statusBytes...)
	key = append(key, []byte("/")...)
	return key
}

func prefixKey(prefix string, keyData []byte) []byte {
	var key []byte
	key = append(key, []byte(prefix)...)
	key = append(key, []byte("/")...)
	key = append(key, keyData...)
	key = append(key, []byte("/")...)
	return key
}

/*
*
1、 key: "Schedule"+ Status  ， Value： [id,id,id...]
2、 key: id   ,  Value:  Schedule bytes
3、 key: id   ,  {ConfimTxs:[{blockheight,txhash},{blockheight,txhash}..]， count: 2}    // message TaskResult
4、 key: id   ,  needRetry: false ，taskTx:[{blockheight,txhash},{blockheight,txhash}..]  // message TryTask
*/
func (manager *ScheduleManager) StoreSchedule(req *types.Schedule) error {
	// add count
	key := ScheduleIdKey(req.Id)
	get := manager.Store.Get(prefixKey(ScheduleKeyPrefix, key))
	if get != nil {
		return errors.New("schedule exist id:" + req.Id.String())
	}
	reqBytes, _ := proto.Marshal(req)
	manager.Store.Set(prefixKey(ScheduleKeyPrefix, key), reqBytes)
	return nil
}
func (manager *ScheduleManager) GetSchedule(req *types.ScheduleId) (*types.Schedule, error) {
	key := ScheduleIdKey(req)
	get := manager.Store.Get(prefixKey(ScheduleKeyPrefix, key))
	if get == nil {
		return nil, errors.New("schedule exist id:" + req.String())
	}
	schedule := &types.Schedule{}
	err := proto.Unmarshal(get, schedule)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (manager *ScheduleManager) StoreScheduleView(req *types.Schedule) error {
	// add count
	key := ScheduleViewKey(int64(req.Status))
	get := manager.Store.Get(prefixKey(ScheduleViewKeyPrefix, key))
	// key=id.string , value= proto.Marshal(id)
	set := make(map[string][]byte)
	if get != nil {
		err := json.Unmarshal(get, &set)
		if err != nil {
			return err
		}
	}
	marshal, err := proto.Marshal(req.Id)
	if err != nil {
		return err
	}
	set[req.Id.String()] = marshal
	idSet, marErr := json.Marshal(set)
	if marErr != nil {
		return marErr
	}
	manager.Store.Set(prefixKey(ScheduleViewKeyPrefix, key), idSet)
	return nil
}

func (manager *ScheduleManager) DeleteScheduleView(status uint32, scheduleId *types.ScheduleId) error {
	// add count
	key := ScheduleViewKey(int64(status))
	get := manager.Store.Get(prefixKey(ScheduleViewKeyPrefix, key))
	// key=id.string , value= proto.Marshal(id)
	set := make(map[string][]byte)
	if get != nil {
		err := json.Unmarshal(get, &set)
		if err != nil {
			return err
		}
	}
	delete(set, scheduleId.String())
	idSet, marErr := json.Marshal(set)
	if marErr != nil {
		return marErr
	}
	manager.Store.Set(prefixKey(ScheduleViewKeyPrefix, key), idSet)
	return nil
}

func (manager *ScheduleManager) GetScheduleView(status int32) ([]*types.ScheduleId, error) {
	key := ScheduleViewKey(int64(status))
	get := manager.Store.Get(prefixKey(ScheduleViewKeyPrefix, key))
	set := make(map[string][]byte)
	if get != nil {
		err := json.Unmarshal(get, &set)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	ids := make([]*types.ScheduleId, 0)
	for _, v := range set {
		id := &types.ScheduleId{}
		err := proto.Unmarshal(v, id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (manager *ScheduleManager) StoreScheduleExecResult(id *types.ScheduleId, blockHeight int64, txHash string) error {
	// add count

	result, err := manager.GetScheduleExecResult(id)
	if err != nil {
		return err
	}
	//  not repeat add
	exist := false
	for _, tx := range result.ConfirmTxs {
		height := tx.BlockHeight
		hash := tx.TxHash
		if hash == txHash && height == blockHeight {
			exist = true
		}
	}
	if exist == true {
		return nil
	}
	result.Count = result.Count + 1
	tx := &types.TaskTx{
		BlockHeight: blockHeight,
		TxHash:      txHash,
	}
	result.ConfirmTxs = append(result.ConfirmTxs, tx)
	marshal, err := proto.Marshal(result)
	if err != nil {
		return err
	}
	key := ScheduleIdKey(id)
	manager.Store.Set(prefixKey(ScheduleExecResultKeyPrefix, key), marshal)
	return nil
}
func (manager *ScheduleManager) GetScheduleExecResult(id *types.ScheduleId) (*types.TaskResult, error) {
	key := ScheduleIdKey(id)
	get := manager.Store.Get(prefixKey(ScheduleExecResultKeyPrefix, key))
	result := &types.TaskResult{}
	if get != nil {
		err := proto.Unmarshal(get, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (manager *ScheduleManager) StoreScheduleTry(id *types.ScheduleId, needTry bool, blockHeight int64, txHash string) error {
	// add count

	tryTask, err := manager.GetScheduleTry(id)
	if err != nil {
		return err
	}

	tryTask.NeedRetry = needTry
	if tryTask.TaskTxs == nil {
		tryTask.TaskTxs = make([]*types.TaskTx, 0)
	}

	if blockHeight > 0 && txHash != "" {
		tx := types.TaskTx{
			BlockHeight: blockHeight,
			TxHash:      txHash,
		}
		exist := false
		for _, tx := range tryTask.TaskTxs {
			height := tx.BlockHeight
			hash := tx.TxHash
			if hash == txHash && height == blockHeight {
				exist = true
			}
		}
		if exist == false {
			tryTask.TaskTxs = append(tryTask.TaskTxs, &tx)
		}
	}

	idSet, err := proto.Marshal(tryTask)
	if err != nil {
		return err
	}
	key := ScheduleIdKey(id)
	manager.Store.Set(prefixKey(ScheduleRetryKeyPrefix, key), idSet)
	return nil
}

func (manager *ScheduleManager) GetScheduleTry(id *types.ScheduleId) (*types.TryTask, error) {
	key := ScheduleIdKey(id)
	get := manager.Store.Get(prefixKey(ScheduleRetryKeyPrefix, key))

	tryTask := &types.TryTask{}
	if get != nil {
		err := proto.Unmarshal(get, tryTask)
		if err != nil {
			return nil, err
		}
		if tryTask == nil {
			tryTask = &types.TryTask{}
		}
	}
	return tryTask, nil
}

func (manager *ScheduleManager) ClearScheduleTry(id *types.ScheduleId) error {
	tryTask, err := manager.GetScheduleTry(id)
	if err != nil {
		return err
	}
	tryTask.NeedRetry = false
	tryTask.TaskTxs = make([]*types.TaskTx, 0)

	idSet, err := proto.Marshal(tryTask)
	if err != nil {
		return err
	}
	key := ScheduleIdKey(id)
	manager.Store.Set(prefixKey(ScheduleRetryKeyPrefix, key), idSet)
	return nil
}
