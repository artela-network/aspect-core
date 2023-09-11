package scheduler

import (
	"github.com/artela-network/artelasdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

var (
	globalHostApi types.ScheduleHostApiI = (*scheduleHost)(nil)
)

func GetScheduleHostApi() (types.ScheduleHostApiI, error) {
	if globalHostApi == nil {
		return nil, errors.New("scheduleHost is not init")
	}
	return globalHostApi, nil
}

type scheduleHost struct {
}

func NewScheduleHost() types.ScheduleHostApiI {
	return &scheduleHost{}
}

func (base *scheduleHost) SubmitSchedule(ctx *types.RunnerContext, sch *types.Schedule) *types.RunResult {
	defResult := &types.RunResult{
		Success: false,
	}
	if ctx.BlockNumber <= 0 {
		defResult.Message = "Get Last BlockNumber is nil"
		return defResult
	}

	sch.CreateHeight = uint64(ctx.BlockNumber + 1)
	sch.StartBlock += sch.CreateHeight

	input := sch.Tx.Input
	decode, err := hexutil.Decode(string(input))
	if err != nil {
		defResult.Message = "decode error[" + err.Error() + "]"
		return defResult
	}
	sch.Tx.Input = decode
	// TODO get gaslimt from user contract.
	sch.Tx.Gas = 2000000
	if err := ScheduleManagerInstance().Submit(sch); err != nil {
		defResult.Message = "submit error[" + err.Error() + "]"
		return defResult
	}
	defResult.Success = true
	return defResult
}
