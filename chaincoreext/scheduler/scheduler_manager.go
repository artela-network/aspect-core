package scheduler

import (
	"github.com/ethereum/go-ethereum/common"
)

type ScheduleTask struct {
	ScheduleName    string
	AspectId        common.Address
	ContractAddress common.Address
	CreateHeight    uint64
	//	0 close, 1 start
	Status uint8
	// 0 periodic, 1 adhoc
	RunType uint8
	TaskTx
	PeriodicType
	AdhocType
}
type TaskTx struct {
	// 按以太坊交易算
	hash common.Hash

	//使用block 的 nonce
	nonce uint64
	//rlp， as 引入abi 编程
	input []byte

	maxFeePerGas         string
	maxPriorityFeePerGas string
	BrokerAddress        common.Address
	value                string
}
type ScheduleTx struct {
	hash                 common.Hash
	blockHeight          uint64
	input                []byte
	maxFeePerGas         uint64
	maxPriorityFeePerGas uint64
	BrokerAddress        common.Address
	value                uint64
}
type PeriodicType struct {
	everyNBlocks uint64
	maxRetry     uint32
	count        uint32
}
type AdhocType struct {
	nextNBlocks uint64
	maxRetry    uint32
}

type Scheduler interface {
	SubmitScheduleTask(req *ScheduleTask) error
	CancelScheduleTask(ScheduleName string, AspectId common.Address, ContractAddress common.Address) error
	QueryScheduleTask(Status uint8) ([]ScheduleTask, error)
	CacheScheduleOpenTask() ([]ScheduleTask, error)
	GenScheduleTaskTx() ([]ScheduleTx, error)
	ConfirmScheduleTx(hash common.Hash) error
}

type ScheduleManager struct {
	MaxNewTxTotal uint32
	TaskPool      []ScheduleTask
	NewPool       []ScheduleTx
	PendingPool   []ScheduleTx
	ConfirmPool   []ScheduleTx
}

var SchedulerHook func() (Scheduler, error)
