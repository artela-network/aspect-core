package types

import "github.com/ethereum/go-ethereum/common"

type RunnerContext struct {
	AspectId     *common.Address
	BlockNumber  int64
	Point        string
	ContractAddr *common.Address
	Gas          uint64
	Commit       bool
}

// var GetAspectRuntimeHook func() (AspectRuntimeHostApi, error)
// var GetAspectStateHook func() (AspectStateHostApi, error)
var GetEvmHostHook func() (EvmHostApi, error)
var GetScheduleHook func() (ScheduleHostApi, error)
var GetStateDbHook func() (StateDbHostApi, error)
var GetRuntimeHostHook func() (RuntimeHostApi, error)

type RuntimeHostApi interface {
	GetContext(ctx *RunnerContext, key string) *ContextQueryResponse
	Set(ctx *RunnerContext, set *ContextSetRequest) bool
	Query(ctx *RunnerContext, query *ContextQueryRequest) *ContextQueryResponse
	Remove(ctx *RunnerContext, set *ContextRemoveRequest) bool
}

//
//type AspectRuntimeHostApi interface {
//	// ContextQuery(string query ) *ContextQueryResponse
//	Get(ctx *RunnerContext, key *ContextQueryRequest) *ContextQueryResponse
//
//	// SetAspectContext(string key,string value) string
//	SetAspectContext(ctx *RunnerContext, key, value string) bool
//}
//
//type AspectStateHostApi interface {
//	//	GetAspectState( key string) string
//	GetAspectState(ctx *RunnerContext, key string) string
//	// SetAspectState( key string, value string) bool
//	SetAspectState(ctx *RunnerContext, key, value string) bool
//
//	//	RemoveAspectState( key string) bool
//	RemoveAspectState(ctx *RunnerContext, key string) bool
//	//GetProperty( key string) string
//	GetProperty(ctx *RunnerContext, key string) string
//}

type EvmHostApi interface {
	//	StaticCall( request CallMessageRequest) CallMessageResponse
	StaticCall(ctx *RunnerContext, request *EthMessage) *EthMessageCallResult

	// JITCall(request CallMessageRequest) *CallMessageResponse
	JITCall(ctx *RunnerContext, request *JitInherentRequest) *JitInherentResponse
}
type ScheduleHostApi interface {
	// SubmitSchedule(sch Schedule) bool
	SubmitSchedule(ctx *RunnerContext, sch *Schedule) *RunResult
}
type StateDbHostApi interface {

	//	GetBalance(request AddressQueryRequest) StringDataResponse
	GetBalance(ctx *RunnerContext, addressEquals string) string
	//GetState(request StateQueryRequest) StringDataResponse
	GetState(ctx *RunnerContext, addressEquals, hashEquals string) string
	//GetRefund() IntDataResponse
	GetRefund(ctx *RunnerContext) uint64
	//GetCodeHash(request AddressQueryRequest) StringDataResponse
	GetCodeHash(ctx *RunnerContext, addressEquals string) string
	//GetNonce(request AddressQueryRequest) IntDataResponse
	GetNonce(ctx *RunnerContext, addressEquals string) uint64
}
