package jitinherent

import (
	"github.com/artela-network/artelasdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
	"github.com/pkg/errors"
	"math/big"
	"sync"
)

type JoinPointStage int

const (
	BlockInitialization JoinPointStage = iota
	PreTransactionExecution
	TransactionExecution
	PostTransactionExecution
	BlockFinalization
)

type EVMCaller func(caller vm.ContractRef, addr common.Address, aspect common.Address, input []byte,
	gas uint64, value *big.Int) (txHash common.Hash, ret []byte, leftOverGas uint64, err error)

type EVMLoader interface {
	GetCurrent() EVM
}

type MemPoolLoader interface {
	Get() MemPool
}

type MemPool interface {
	Submit(aspect common.Address, inherent *types.JitInherentRequest, evmCall EVMCaller, joinPoint JoinPoint) *types.JitInherentResponse
}

type EVM interface {
	CallFromAspect(caller vm.ContractRef, addr common.Address, aspect common.Address, input []byte,
		gas uint64, value *big.Int) (txHash common.Hash, ret []byte, leftOverGas uint64, err error)
}

type innerJITInherent struct {
	Sender               common.Address
	Nonce                uint64
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *uint256.Int
	VerificationGasLimit *uint256.Int
	PreVerificationGas   *uint256.Int
	MaxFeePerGas         *uint256.Int
	MaxPriorityFeePerGas *uint256.Int
	PaymasterAndData     []byte
}

type JoinPoint interface {
	Stage() JoinPointStage
}

func newInnerJITInherentFromProto(protoMsg *types.JitInherentRequest) *innerJITInherent {
	return &innerJITInherent{
		Sender:               common.BytesToAddress(protoMsg.Sender),
		Nonce:                protoMsg.Nonce,
		InitCode:             protoMsg.InitCode,
		CallData:             protoMsg.CallData,
		CallGasLimit:         uint256.NewInt(0).SetBytes(protoMsg.CallGasLimit),
		VerificationGasLimit: uint256.NewInt(0).SetBytes(protoMsg.VerificationGasLimit),
		PreVerificationGas:   uint256.NewInt(0).SetBytes(protoMsg.PreVerificationGas),
		MaxFeePerGas:         uint256.NewInt(0).SetBytes(protoMsg.MaxFeePerGas),
		MaxPriorityFeePerGas: uint256.NewInt(0).SetBytes(protoMsg.MaxPriorityFeePerGas),
		PaymasterAndData:     protoMsg.PaymasterAndData,
	}
}

func (i innerJITInherent) Hash() common.Hash {
	return common.Hash{}
}

var (
	JITInherentEntryPoint = common.HexToAddress("0x000000000000000000000000000000000000AAEC")
)

var (
	// global jit inherent instance
	instance *Manager

	// lock for instance
	lock sync.Mutex
)

// Get returns the global JITInherentManager instance.
func Get() *Manager {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = newManager()
	}

	return instance
}

type Manager struct {
}

func newManager() *Manager {
	return &Manager{}
}

func (m *Manager) Submit(aspect common.Address, inherent *types.JitInherentRequest, evmCall EVMCaller, joinPoint JoinPoint) (*types.JitInherentResponse, error) {
	userOp := newInnerJITInherentFromProto(inherent)
	hash := userOp.Hash()
	resp := &types.JitInherentResponse{
		JitInherentHash: hash.Bytes(),
	}

	switch joinPoint.Stage() {
	case BlockInitialization:
		return nil, nil
	case TransactionExecution:
		txHash, ret, _, err := evmCall(vm.AccountRef(userOp.Sender), JITInherentEntryPoint, aspect,
			userOp.CallData, userOp.CallGasLimit.Uint64(), big.NewInt(0))
		resp.TxHash = txHash.Bytes()

		if err == nil {
			resp.Success = true
			return resp, nil
		} else if errors.Is(err, vm.ErrExecutionReverted) {
			resp.Success = false
			resp.Ret = ret
		} else {
			resp.Success = false
		}

		return resp, err
	case PostTransactionExecution:
		return nil, nil
	case BlockFinalization:
		return nil, nil
	default:
		return nil, errors.New("cannot submit jit inherent in current join point")
	}
}
