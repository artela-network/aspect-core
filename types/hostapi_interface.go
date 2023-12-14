package types

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

type RunnerContext struct {
	Ctx          context.Context
	AspectId     *common.Address
	BlockNumber  int64
	Point        string
	ContractAddr *common.Address
	Gas          uint64
	Commit       bool
}

var (
	GetEvmHostHook     func(context.Context) (EvmHostApi, error)
	GetScheduleHook    func(context.Context) (ScheduleHostApi, error)
	GetStateDbHook     func(context.Context) (StateDbHostApi, error)
	GetRuntimeHostHook func(context.Context) (RuntimeHostApi, error)
)

type (
	RuntimeHostApi interface {
		GetContext(ctx *RunnerContext, key string) *ContextQueryResponse
		Set(ctx *RunnerContext, set *ContextSetRequest) bool
		Query(ctx *RunnerContext, query *ContextQueryRequest) *ContextQueryResponse
		Remove(ctx *RunnerContext, set *ContextRemoveRequest) bool
	}

	EvmHostApi interface {
		//	StaticCall( request CallMessageRequest) CallMessageResponse
		StaticCall(ctx *RunnerContext, request *EthMessage) *EthMessageCallResult

		// JITCall(request CallMessageRequest) *CallMessageResponse
		JITCall(ctx *RunnerContext, request *JitInherentRequest) *JitInherentResponse
	}

	ScheduleHostApi interface {
		// SubmitSchedule(sch Schedule) bool
		SubmitSchedule(ctx *RunnerContext, sch *Schedule) *RunResult
	}

	StateDbHostApi interface {
		//	GetBalance(request AddressQueryRequest) StringDataResponse
		GetBalance(ctx *RunnerContext, addressEquals string) string
		// GetState(request StateQueryRequest) StringDataResponse
		GetState(ctx *RunnerContext, addressEquals, hashEquals string) string
		// GetRefund() IntDataResponse
		GetRefund(ctx *RunnerContext) uint64
		// GetCodeHash(request AddressQueryRequest) StringDataResponse
		GetCodeHash(ctx *RunnerContext, addressEquals string) string
		// GetNonce(request AddressQueryRequest) IntDataResponse
		GetNonce(ctx *RunnerContext, addressEquals string) uint64
	}
)
