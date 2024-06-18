package types

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type RunnerContext struct {
	Ctx           context.Context
	AspectId      common.Address
	AspectVersion uint64
	BlockNumber   int64
	Point         string
	ContractAddr  common.Address
	Gas           uint64
	Commit        bool
}

func (r *RunnerContext) RemainingGas() uint64 {
	return r.Gas
}

func (r *RunnerContext) SetGas(gas uint64) {
	r.Gas = gas
}

var (
	GetEvmHostHook                    func(context.Context) (EVMHostAPI, error)
	GetStateDbHook                    func(context.Context) (StateDBHostAPI, error)
	GetAspectRuntimeContextHostHook   func(context.Context) (RuntimeContextHostAPI, error)
	GetAspectStateHostHook            func(context.Context) (AspectStateHostAPI, error)
	GetAspectPropertyHostHook         func(context.Context) (AspectPropertyHostAPI, error)
	GetAspectTransientStorageHostHook func(context.Context) (AspectTransientStorageHostAPI, error)
	GetAspectTraceHostHook            func(context.Context) (AspectTraceHostAPI, error)

	// JITSenderAspectByContext returns the sender Aspect address of the user operation
	JITSenderAspectByContext func(ctx context.Context, userOpHash common.Hash) (common.Address, error)
	IsCommit                 func(ctx context.Context) bool
)

type (
	RuntimeContextHostAPI interface {
		Get(ctx *RunnerContext, key string) ([]byte, error)
	}

	AspectStateHostAPI interface {
		Get(ctx *RunnerContext, key string) []byte
		Set(ctx *RunnerContext, key string, value []byte) error
	}

	AspectPropertyHostAPI interface {
		Get(ctx *RunnerContext, key string) ([]byte, error)
	}

	AspectTransientStorageHostAPI interface {
		Get(ctx *RunnerContext, aspectId []byte, key string) ([]byte, error)
		Set(ctx *RunnerContext, key string, value []byte) error
	}

	AspectTraceHostAPI interface {
		QueryStateChange(ctx *RunnerContext, query *StateChangeQuery) ([]byte, error)
		QueryCallTree(ctx *RunnerContext, query *CallTreeQuery) ([]byte, error)
	}

	EVMHostAPI interface {
		StaticCall(ctx *RunnerContext, request *StaticCallRequest) (*StaticCallResult, error)
		JITCall(ctx *RunnerContext, request *JitInherentRequest) (*JitInherentResponse, error)
	}

	StateDBHostAPI interface {
		GetBalance(address common.Address) *big.Int
		GetState(address common.Address, key common.Hash) common.Hash
		GetCodeHash(address common.Address) common.Hash
		GetCodeSize(address common.Address) int
		GetNonce(address common.Address) uint64
		HasSuicided(address common.Address) bool
	}
)
