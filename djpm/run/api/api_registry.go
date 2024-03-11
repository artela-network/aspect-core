package api

import (
	"context"
	rttypes "github.com/artela-network/aspect-runtime/types"
	"github.com/artela-network/aspect-runtime/wasmtime"

	"github.com/ethereum/go-ethereum/common"

	"github.com/artela-network/aspect-core/types"
)

const (
	// module of hostapis
	moduleUtils                  = "util-api"
	moduleCrypto                 = "crypto-api"
	moduleStateDB                = "statedb-api"
	moduleRuntimeContext         = "runtime-api"
	moduleEvmCall                = "evm-call-api"
	moduleAspectState            = "aspect-state-api"
	moduleAspectProperty         = "aspect-property-api"
	moduleAspectTransientStorage = "aspect-transient-storage-api"
	moduleTrace                  = "trace-api"

	// namespace of hostapis
	nsUtils                  = "__UtilApi__"
	nsCrypto                 = "__CryptoApi__"
	nsStateDB                = "__StateDbApi__"
	nsRuntimeContext         = "__RuntimeContextApi__"
	nsEvmCall                = "__EvmCallApi__"
	nsTrace                  = "__TraceApi__"
	nsAspectState            = "__AspectStateApi__"
	nsAspectProperty         = "__AspectPropertyApi__"
	nsAspectTransientStorage = "__AspectTransientStorageApi__"

	// entrance of api functions
	APIEntrance           = "execute"
	CheckBlockLevel       = "isBlockLevel"
	CheckTransactionLevel = "isTransactionLevel"
	CheckIsTxVerifier     = "isTransactionVerifier"
)

// Registry keeps the properity owned by current
type Registry struct {
	runnerContext *types.RunnerContext
	collection    *rttypes.HostAPIRegistry
	errCallback   func(message string)
}

func NewRegistry(ctx context.Context, aspectID common.Address, aspVer uint64) *Registry {
	return &Registry{
		runnerContext: &types.RunnerContext{
			Ctx:           ctx,
			AspectId:      aspectID,
			AspectVersion: aspVer,
		},
		collection: rttypes.NewHostAPIRegistry(wasmtime.Wrap),
	}
}

// HostApis return the collection of aspect runtime host apis
func (r *Registry) HostApis() *rttypes.HostAPIRegistry {
	r.registerApis(moduleStateDB, nsStateDB, r.stateDBAPIs())
	r.registerApis(moduleUtils, nsUtils, r.utilAPIs())
	r.registerApis(moduleCrypto, nsCrypto, r.cryptoAPIs())
	r.registerApis(moduleEvmCall, nsEvmCall, r.evmCallAPIs())
	r.registerApis(moduleRuntimeContext, nsRuntimeContext, r.runtimeContextAPIs())
	r.registerApis(moduleAspectProperty, nsAspectProperty, r.aspectPropertyAPIs())
	r.registerApis(moduleAspectState, nsAspectState, r.aspectStateAPIs())
	r.registerApis(moduleAspectTransientStorage, nsAspectTransientStorage, r.transientStorageAPIs())
	r.registerApis(moduleTrace, nsTrace, r.traceAPIs())

	return r.collection
}

func (r *Registry) registerApis(module, namespace string, apis map[string]*rttypes.HostFuncWithGasRule) {
	for method, hostFunc := range apis {
		// Here we cannot make new variable function to call fn in it,
		// and to pass it into AddApi in loop instead pass fn directly.
		hostFunc.HostContext = r.runnerContext
		err := r.collection.AddAPI(rttypes.Module(module), rttypes.NameSpace(namespace), rttypes.MethodName(method), hostFunc)
		if err != nil {
			panic("add host api failed" + err.Error())
		}
	}
}

func (r *Registry) SetRunnerContext(name string, blockNum int64, gas uint64, contractAddr *common.Address) {
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
		r.runnerContext.ContractAddr = *contractAddr
	}
}

func (r *Registry) RunnerContext() *types.RunnerContext {
	return r.runnerContext
}

func (r *Registry) SetErrCallback(errFunc func(message string)) {
	r.errCallback = errFunc
}
