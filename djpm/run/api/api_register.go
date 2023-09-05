package api

import (
	"github.com/artela-network/artelasdk/types"
	"github.com/artela-network/runtime"
)

const (
	// module of hostapis
	moduleUtils          = "util-api"
	moduleCrypto         = "crypto-api"
	moduleStateDb        = "statedb-api"
	moduleSchedule       = "schedule-api"
	moduleRuntimeContext = "runtime-context-api"
	moduleEvmCall        = "evm-call-api"
	moduleAspectState    = "aspect-state-api"

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

func NewRegister(aspectID string) *Register {
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
	r.registerApis(moduleRuntimeContext, nsRuntimeContextApi, r.contextCallApis())
	r.registerApis(moduleEvmCall, nsEvmCallApi, r.evmCallApis())
	r.registerApis(moduleAspectState, nsAspectStateApi, r.stateApis())
	return r.collection
}

func (r *Register) registerApis(module, namespace string, apis interface{}) {
	for method, fn := range apis.(map[string]interface{}) {
		// Here we cannot make new variable function to call fn in it,
		// and to pass it into AddApi in loop instead pass fn directly.
		r.collection.AddApi(runtime.Module(module), runtime.Namesapce(namespace), runtime.MethodName(method), fn)
	}
}

func (r *Register) SetRunnerContext(name types.PointCut, blockNum int64) {
	r.runnerContext.Point = name
	r.runnerContext.BlockNumber = blockNum
}
func (r *Register) SetErrCallback(errfunc func(message string)) {
	r.errCallback = errfunc
}
