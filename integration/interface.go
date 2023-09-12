package integration

import (
	artevm "github.com/artela-network/evm/vm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"math/big"
)

type SystemContractType int

const (
	Native SystemContractType = iota
	Solidity
)

// JoinPointStage defines the stage of join point.
type JoinPointStage int

const (
	BlockInitialization JoinPointStage = iota
	PreTransactionExecution
	TransactionExecution
	PostTransactionExecution
	BlockFinalization
	UNKNOW
)

// TxData defines the interface of a transaction data.
type TxData interface {
	TxType() byte // returns the type ID
	From() common.Address
	To() common.Address
	Data() []byte
	Gas() uint64
	GasPrice() *big.Int
	GasTipCap() *big.Int
	GasFeeCap() *big.Int
	Value() *big.Int
	Nonce() uint64
	Extra() map[string]interface{}
}

type BlockHeader interface {
	ParentHash() common.Hash
	Coinbase() common.Address
	Root() common.Hash
	TxHash() common.Hash
	ReceiptHash() common.Hash
	Number() *big.Int
	GasLimit() uint64
	GasUsed() uint64
	Time() uint64
	Extra() []byte
	MixDigest() common.Hash
	BaseFee() *big.Int
}

type BaseLayerTx interface {
	Bytes() []byte
	Hash() []byte
	Sender() []byte
	Recipient() []byte
}

// AspectProtocol is the core interface for integrating Aspect Programming into an existing protocol
type AspectProtocol interface {
	VMFromSnapshotState() (VM, error)
	VMFromCanonicalState() (VM, error)
	ConvertProtocolTx(txData TxData) (BaseLayerTx, error)
	EstimateGas(txData TxData) (uint64, error)
	GasPrice() (*big.Int, error)
	LastBlockHeader() (BlockHeader, error)
	NonceOf(address common.Address) (uint64, error)
	SubmitTxToCurrentProposal(tx BaseLayerTx) error
	InitSystemContract(addr common.Address, code []byte,
		storage map[common.Hash][]byte, contractType SystemContractType) error
	BalanceOf(address common.Address) *big.Int
}

// VM defines the interface to interact with VM.
type VM interface {
	// Msg returns the current vm message
	Msg() artevm.Message

	// Call executes the contract call using the given input.
	Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error)
}
