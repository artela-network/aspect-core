package api

import (
	runtime "github.com/artela-network/aspect-runtime"
	"github.com/ethereum/go-ethereum/common"

	"github.com/artela-network/aspect-core/types"
)

const (
	// module of hostapis
	moduleUtils          = "util-api"
	moduleCrypto         = "crypto-api"
	moduleStateDb        = "statedb-api"
	moduleSchedule       = "schedule-api"
	moduleRuntimeContext = "runtime-api"
	moduleEvmCall        = "evm-call-api"
	moduleAspectState    = "aspect-state-api"
	moduleHost           = "aspect-state-api"

	// namespace of hostapis
	nsUtils             = "__UtilApi__"
	nsCryptoApi         = "__CryptoApi__"
	nsStateDbApi        = "__StateDbApi__"
	nsScheduleApi       = "__ScheduleApi__"
	nsRuntimeContextApi = "__RuntimeContextApi__"
	nsEvmCallApi        = "__EvmCallApi__"
	nsAspectStateApi    = "__AspectStateApi__"

	// entrance of api functions
	ApiEntrance           = "execute"
	CheckBlockLevel       = "isBlockLevel"
	CheckTransactionLevel = "isTransactionLevel"
	CheckIsTxVerifier     = "isTransactionVerifier"
)

type HostFunc interface {
	FuncRegister() *runtime.HostAPIRegistry
}

// Register keeps the properity owned by current
type Register struct {
	runnerContext *types.RunnerContext
	collection    *runtime.HostAPIRegistry
	errCallback   func(message string)
}

func NewRegister(aspectID *common.Address) *Register {
	return &Register{
		runnerContext: &types.RunnerContext{
			AspectId: aspectID,
		},
		collection: runtime.NewHostAPIRegistry(),
	}
}

// HostApis return the collection of aspect runtime host apis
func (r *Register) HostApis() *runtime.HostAPIRegistry {
	r.registerApis(moduleStateDb, nsStateDbApi, r.stateDbApis())
	r.registerApis(moduleUtils, nsUtils, r.utilApis())
	r.registerApis(moduleCrypto, nsCryptoApi, r.cryptoApis())
	r.registerApis(moduleSchedule, nsScheduleApi, r.scheduleApis())
	r.registerApis(moduleEvmCall, nsEvmCallApi, r.evmCallApis())
	r.registerApis(moduleRuntimeContext, nsRuntimeContextApi, r.hostApi())

	return r.collection
}

func (r *Register) registerApis(module, namespace string, apis interface{}) {
	for method, fn := range apis.(map[string]interface{}) {
		// Here we cannot make new variable function to call fn in it,
		// and to pass it into AddApi in loop instead pass fn directly.
		_ = r.collection.AddAPI(runtime.Module(module), runtime.NameSpace(namespace), runtime.MethodName(method), fn)
	}
}

func (r *Register) SetRunnerContext(name string, blockNum int64, gas uint64, contractAddr *common.Address) {
	if name != "" {
		r.runnerContext.Point = name
	}
	if blockNum > 0 {
		r.runnerContext.BlockNumber = blockNum
	}
	if gas > 0 {
		r.runnerContext.Gas = gas
	}
	if contractAddr != nil {
		r.runnerContext.ContractAddr = contractAddr
	}
}

func (r *Register) RunnerContext() *types.RunnerContext {
	return r.runnerContext
}

func (r *Register) SetErrCallback(errfunc func(message string)) {
	r.errCallback = errfunc
}
