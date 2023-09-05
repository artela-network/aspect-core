package types

type RunnerContext struct {
	AspectId     string
	BlockNumber  int64
	Point        PointCut
	ContractAddr string
}

var GetAspectRuntimeHook func() (AspectRuntimeHostApiI, error)
var GetAspectStateHook func() (AspectStateHostApiI, error)
var GetEvmHostHook func() (EvmHostApiI, error)
var GetScheduleHook func() (ScheduleHostApiI, error)
var GetStateDbHook func() (StateDbHostApiI, error)
var GetCryptoHook func() (CryptoHostApiI, error)

type AspectRuntimeHostApiI interface {
	// ContextQuery(string query ) *ContextQueryResponse
	Get(ctx *RunnerContext, key *ContextQueryRequest) *ContextQueryResponse
	// SetAspectContext(string key,string value) string
	SetAspectContext(ctx *RunnerContext, request *KeyValueSetRequest) *StringDataResponse
}

type AspectStateHostApiI interface {
	//	GetAspectState( key string) string
	GetAspectState(ctx *RunnerContext, request *KeyGetRequest) *StringDataResponse
	// SetAspectState( key string, value string) bool
	SetAspectState(ctx *RunnerContext, request *KeyValueSetRequest) *RunResult

	//	RemoveAspectState( key string) bool
	RemoveAspectState(ctx *RunnerContext, request *KeyGetRequest) *RunResult
	//GetProperty( key string) string
	GetProperty(ctx *RunnerContext, request *KeyGetRequest) *StringDataResponse
}

type EvmHostApiI interface {
	//	StaticCall( request CallMessageRequest) CallMessageResponse
	StaticCall(ctx *RunnerContext, request *CallMessageRequest) *CallMessageResponse
	// JITCall(request CallMessageRequest) *CallMessageResponse
	JITCall(ctx *RunnerContext, request *CallMessageRequest) *CallMessageResponse
}
type ScheduleHostApiI interface {
	// SubmitSchedule(sch Schedule) bool
	SubmitSchedule(ctx *RunnerContext, sch *Schedule) *RunResult
}
type StateDbHostApiI interface {
	//	GetBalance(request AddressQueryRequest) StringDataResponse
	GetBalance(ctx *RunnerContext, request *AddressQueryRequest) *StringDataResponse
	//GetState(request StateQueryRequest) StringDataResponse
	GetState(ctx *RunnerContext, request *StateQueryRequest) *StringDataResponse
	//GetRefund() IntDataResponse
	GetRefund(ctx *RunnerContext) *IntDataResponse
	//GetCodeHash(request AddressQueryRequest) StringDataResponse
	GetCodeHash(ctx *RunnerContext, request *AddressQueryRequest) *StringDataResponse
	//GetNonce(request AddressQueryRequest) IntDataResponse
	GetNonce(ctx *RunnerContext, request *AddressQueryRequest) *IntDataResponse
}

type CryptoHostApiI interface {
	//DoCrypto(request CryptoRequest) CryptoResponse
	DoCrypto(request *CryptoRequest) *CryptoResponse
}
