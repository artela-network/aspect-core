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
	GetCallStack() *InnerTransactions

	SetContext(aspectID string, key, value string) error
	GetContext(aspectID string, key string) (string, error)

	SetAspectState(aspectID string, key, value string) error
	GetAspectState(aspectID string, key string) (string, error)

	AddInherent()
	ScheduleTx(sch *Schedule) bool
}
