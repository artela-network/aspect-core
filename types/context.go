package types

import (
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type AspectRuntimeContextI interface {
	BlockContextI
	TransactionContextI
	EnvContextI
}

type TransactionContextI interface {
	AspectContextI
	//tx_content
	TxContent() (*AspTransaction, error)
	//StateChanges
	GetStateChanges(addr string, variable string, key []byte) (*StateChanges, error)
	QueryCallStack() (*AspCallStacks, error)
	GetReceipt() (*AspReceipt, error)
	GasMeter() (*store.GasMeter, error)
}

type AspectContextI interface {
	AddContextData(aspectID string, key, value string) error
	GetContextData(aspectID string, key string) (string, error)
	RemoveContextData(aspectID string, key string) error
}

type BlockContextI interface {
	GetBlockHeader() (*EthHeader, error)
	GetBlockBody() ([][]byte, error)
	GetVoteInfo() (*AspVoteInfo, error)
	GasMeter() (*store.GasMeter, error)
}
type EnvContextI interface {
	GetChainConfig() (*ChainConfig, error)
	GetEvmParams() (*EvmParams, error)
	GetConsParams() (*ConsParams, error)
	GetEnvContent() (*EnvContent, error)
}

type executeState struct {
	txCtxMap map[common.Hash]TransactionContext
}
type TransactionContext struct {
	txContent     types.Transaction
	aspectContext map[string]map[string]string
	stateChanges  any               //*StateChanges json
	callstacks    any               //*CallStacks  json
	receipt       types.Transaction //*ethtypes.Receipt json
	extProperties map[string]interface{}
}

func (*executeState) AddTx(txBytes []byte) bool {
	//todo
	return true
}
func (*executeState) RemoveTx(txBytes []byte) bool {
	//todo
	return true
}
