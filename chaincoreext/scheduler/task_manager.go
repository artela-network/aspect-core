package scheduler

import "github.com/artela-network/artelasdk/types"

type TaskManager struct {
	store *types.AspectStore
	// cache schedule
	pool map[string]*types.Task
}

func (manager TaskManager) GenTaskTx(heigh int64) ([]*types.Task, error) {
	return nil, nil
}

func (manager TaskManager) ConfirmTx(txHash string) error {
	return nil
}

func (manager TaskManager) UpdateTx(txHash string) error {
	return nil
}
