package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var GetHostApiHook func() (HostApi, error)

type HostApi interface {
	// LocalCall calls EthCall
	//	LocalCall(req *evmtypes.EthCallRequest) (*evmtypes.MsgEthereumTxResponse, error)

	// TBD, if we need to return the artelamint blocks
	// LastBlock() (*coretypes.ResultBlock, error)
	// CurrentBlock() (*coretypes.ResultBlock, error)

	// LastBlock returns last ethereum block
	LastBlock() (*EthBlock, error)

	// CurrentBlock returns ethereum block built by the packing block,
	// this should only be called when a new block is generating
	CurrentBlock() (*EthBlock, error)

	// CurrentBalance return current blance of account address
	CurrentBalance(addr common.Address) (*big.Int, error)

	// GetProperty returns the configuration of aspect
	GetProperty(aspectID string, key string) (string, error)

	// GetStateChanges returns the state changes of fields
	GetStateChanges(addr string, variable string, key []byte) *StateChanges

	SetContext(aspectID string, key, value string) error
	GetContext(aspectID string, key string) (string, error)

	SetAspectState(aspectID string, key, value string) error
	GetAspectState(aspectID string, key string) (string, error)

	AddInherent()
	ScheduleTx(sch *Schedule) bool
}

type SetDataType string

const (
	ASPECT_STATE       SetDataType = "aspect.state"
	CURRENT_TX_CONTEXT SetDataType = "current_tx.context"
)

type AspectRuntimeContextHostApiI interface {
	Select(point PointCut, dataType AspectDataAccess, aspectId common.Address, queryCriteria []interface{}) (interface{}, error)
	Set(point PointCut, data SetDataType, aspectId common.Address, key string, value any) error
	Remove(point PointCut, data SetDataType, aspectId common.Address, key string) error
}

type EvmHostApiI interface {
	LocalCall(blockNumber int64, transaction *AspTransaction) (*AspTxResponse, error)
	InnerCall(transaction *AspTransaction) (*AspTxResponse, error)
}
type ScheduleHostApiI interface {
	SubmitScheduleTx(sch *Schedule) error
}
type StateDbHostApi interface {
	GetBalance(addr common.Address) (*big.Int, error)
	GetState(addr common.Address, hash common.Hash) (common.Hash, error)
	GetRefund() (uint64, error)
	GetCodeHash(addr common.Address) (*common.Hash, error)
	GetNonce(addr common.Address) (uint64, error)
}
type AspectStateApiI interface {
	GetAspectState(aspectID string, key string) (string, error)
	SetAspectState(aspectID string, key, value string) error
	RemoveAspectState(aspectID string, key string) error
	GetProperty(aspectID string, key string) (string, error)
}

type CryptoHostApiI interface {
	Hash(hasher int32, data []byte) ([]byte, error)
}
type AbiSystemCallI interface {
	DecodeParams(t string, data []byte) ([]interface{}, error)
	EncodeParams(t string, values ...interface{}) ([]byte, error)
}
