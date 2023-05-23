package djpm

import (
	"github.com/artela-network/artelasdk/chaincoreext/scheduler"
	"github.com/artela-network/artelasdk/types"
)

type GetLastBlockNum func() int64

// HostApiBase implemets a part of HostApi interfaces
type HostApiBase struct {
	lastBlockNum GetLastBlockNum
}

func NewHostApiBase(lastBN GetLastBlockNum) HostApiBase {
	return HostApiBase{lastBlockNum: lastBN}
}

// ScheduledTx submit the schedule, return true if success
func (base *HostApiBase) ScheduleTx(sch *types.Schedule) bool {
	if base.lastBlockNum == nil {
		return false
	}
	sch.CreateHeight = uint64(base.lastBlockNum()) + 1
	sch.StartBlock += sch.CreateHeight
	if err := scheduler.ScheduleManagerInstance().Submit(sch); err != nil {
		return false
	}
	return true
}

func (base *HostApiBase) AddInherent() {}
func (base *HostApiBase) DropTx()      {}
