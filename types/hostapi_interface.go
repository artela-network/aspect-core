package types

type RunnerContext struct {
	AspectId     string
	BlockNumber  int64
	Point        PointCut
	ContractAddr string
}

var GetAspectRuntimeHook func() (AspectRuntimeHostApiI, error)
var GetAspectStateHook func() (AspectStateHostApiI, error)
var GetEvmHostHook func() (EvmHostApi, error)
var GetScheduleHook func() (ScheduleHostApiI, error)
var GetStateDbHook func() (StateDbHostApiI, error)
var GetCryptoHook func() (CryptoHostApiI, error)

type AspectRuntimeHostApiI interface {
	// ContextQuery(string query ) *ContextQueryResponse
	Get(ctx *RunnerContext, key *ContextQueryRequest) *ContextQueryResponse
	// SetAspectContext(string key,string value) string
	SetAspectContext(ctx *RunnerContext, request *KeyValueSetRequest) bool
}

type AspectStateHostApiI interface {
	//	GetAspectState( key string) string
	GetAspectState(ctx *RunnerContext, key string) string
	// SetAspectState( key string, value string) bool
	SetAspectState(ctx *RunnerContext, key, value string) bool

	//	RemoveAspectState( key string) bool
	RemoveAspectState(ctx *RunnerContext, key string) bool
	//GetProperty( key string) string
	GetProperty(ctx *RunnerContext, key string) string
}

type EvmHostApi interface {
	//	StaticCall( request CallMessageRequest) CallMessageResponse
	StaticCall(ctx *RunnerContext, request *EthTransaction) *CallMessageResponse
	// JITCall(request CallMessageRequest) *CallMessageResponse
	JITCall(ctx *RunnerContext, request *JitInherentRequest) *JitInherentResponse
}
type ScheduleHostApiI interface {
	// SubmitSchedule(sch Schedule) bool
	SubmitSchedule(ctx *RunnerContext, sch *Schedule) *RunResult
}
type StateDbHostApiI interface {

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

type CryptoHostApiI interface {
	//DoCrypto(request CryptoRequest) CryptoResponse
	DoCrypto(request *CryptoRequest) *CryptoResponse
}
