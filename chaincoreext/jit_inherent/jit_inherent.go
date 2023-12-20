package jit_inherent

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
	"github.com/pkg/errors"

	aa "github.com/artela-network/aspect-core/chaincoreext/account_abstraction"
	"github.com/artela-network/aspect-core/integration"
	"github.com/artela-network/aspect-core/types"
)

// Manager manages the JIT inherent calls.
type Manager struct {
	protocol      integration.AspectProtocol
	entrypointABI *abi.ABI

	// userOpSenderLookup is used to lookup the sender Aspect of user operation.
	lookupMutex        sync.RWMutex
	userOpSenderLookup map[common.Hash]common.Address
}

// NewManager creates a new JITInherentManager instance.
func NewManager(protocol integration.AspectProtocol) *Manager {
	entrypointABI, _ := aa.IEntryPointMetaData.GetAbi()
	return &Manager{
		protocol:           protocol,
		entrypointABI:      entrypointABI,
		userOpSenderLookup: make(map[common.Hash]common.Address),
	}
}

// TODO: Refactor the code to eliminate the use of a global instance for managing JIT calls.
// After that, the protocol should not be updated.
func (m *Manager) UpdateProtocol(protocol integration.AspectProtocol) {
	m.protocol = protocol
}

// ·Submit submits a JIT inherent call. There are two types of JIT inherent calls:
//  1. JIT transaction: the JIT transaction will be submitted directly into the block proposal to guarantee the execution.
//     Please note that the JIT transaction submission could be failed if there is no space left in the block.
//  2. JIT call: the JIT call will be injected into the current evm callstack to guarantee the execution.
//     Only one JIT call can be submitted at a time.
func (m *Manager) Submit(ctx context.Context, aspect common.Address,
	gas uint64, stage integration.JoinPointStage, inherents ...*types.JitInherentRequest,
) (*types.JitInherentResponse, uint64, error) {
	if len(inherents) == 0 {
		return nil, gas, errors.New("no jit inherent to submit")
	}

	switch stage {
	case integration.TransactionExecution:
		if len(inherents) != 1 {
			return nil, gas, errors.New("only one user operation is allowed in current join point")
		}
		return m.submitJITCall(ctx, aspect, gas, inherents[0])
	default:
		return nil, gas, errors.New("cannot submit jit inherent in current join point")
	}
}

func (m *Manager) EstimateGas(aspect common.Address, inherent *types.JitInherentRequest) (
	verificationGasLimit, callGasLimit *uint256.Int, err error,
) {
	// get vm with snapshot state
	cvm, err := m.protocol.VMFromSnapshotState()
	if err != nil {
		return nil, nil, err
	}

	gas := uint256.NewInt(cvm.Msg().Gas())

	return gas, gas, nil
}

func (m *Manager) Nonce(ctx context.Context, account common.Address, key *big.Int) (nonce *big.Int, err error) {
	if key.BitLen() > 192 {
		return nil, errors.New("key is too large")
	}

	return m.getNonce(ctx, account, key)
}

// ClearLookup clears the user operation sender lookup. When current block finished, the lookup table should be cleared.
func (m *Manager) ClearLookup() {
	m.lookupMutex.Lock()
	defer m.lookupMutex.Unlock()

	m.userOpSenderLookup = make(map[common.Hash]common.Address)
}

// ClearUserOp clears the user operation sender lookup. When current call finished, the lookup table should be cleared.
func (m *Manager) ClearUserOp(userOpHash common.Hash) {
	m.lookupMutex.Lock()
	defer m.lookupMutex.Unlock()

	delete(m.userOpSenderLookup, userOpHash)
}

// SenderAspect returns the sender Aspect address of the user operation.
func (m *Manager) SenderAspect(userOpHash common.Hash) common.Address {
	m.lookupMutex.RLock()
	defer m.lookupMutex.RUnlock()

	return m.userOpSenderLookup[userOpHash]
}

// submitJITCall submits a JIT call to the current EVM callstack.
func (m *Manager) submitJITCall(ctx context.Context, aspect common.Address, gas uint64, request *types.JitInherentRequest) (
	*types.JitInherentResponse, uint64, error,
) {
	baseLayerVM, err := m.protocol.VMFromSnapshotState()
	if err != nil {
		log.Error("failed to get vm from snapshot state", "err", err)
		return nil, gas, err
	}

	msg := baseLayerVM.Msg()
	maxFeePerGas, maxPriorityFeePerGas := msg.GasFeeCap().Uint64(), msg.GasTipCap().Uint64()
	userOp := NewUserOperations(gas, maxFeePerGas, maxPriorityFeePerGas, request)[0]

	// get nonce from entrypoint and set it if not provided
	if userOp.Nonce.Cmp(big.NewInt(0)) == 0 {
		userOp.Nonce, gas, err = m.getAAWalletNonce(ctx, baseLayerVM,
			common.BytesToAddress(request.Sender),
			uint256.NewInt(0).SetBytes(request.NonceKey),
			gas)
		if err != nil {
			log.Error("failed to get nonce", "err", err)
			return nil, gas, err
		}
	}

	userOpHashes := m.cacheUserOp(aspect, userOp)
	defer m.ClearUserOp(userOpHashes[0])

	resp := &types.JitInherentResponse{
		JitInherentHashes: [][]byte{userOpHashes[0].Bytes()},
		Success:           false,
	}

	callData, err := m.entrypointABI.Pack("handleOps", []aa.UserOperation{*userOp}, userOp.Sender)
	if err != nil {
		return nil, gas, err
	}
	ret, gas, err := baseLayerVM.Call(ctx, vm.AccountRef(userOp.Sender), aa.EntryPointContract, callData, gas, big.NewInt(0))
	resp.Success = err == nil
	if err == nil || err.Error() == vm.ErrExecutionReverted.Error() {
		// ignore the reverted error
		resp.Ret = ret
	}

	return resp, gas, err
}

func (m *Manager) getAAWalletNonce(ctx context.Context, baseLayerVM integration.VM, address common.Address, nonceKey *uint256.Int, gas uint64) (*big.Int, uint64, error) {
	// call entrypoint's getNonce method to retrieve the nonce
	callData, err := m.entrypointABI.Pack("getNonce", address, nonceKey.ToBig())
	if err != nil {
		return nil, gas, err
	}

	ret, leftoverGas, err := baseLayerVM.Call(ctx, vm.AccountRef(address), aa.EntryPointContract, callData, gas, big.NewInt(0))
	if err != nil {
		return nil, leftoverGas, err
	}

	return uint256.NewInt(0).SetBytes(ret).ToBig(), leftoverGas, nil
}

func (m *Manager) cacheUserOp(aspect common.Address, userOps ...*aa.UserOperation) []common.Hash {
	m.lookupMutex.Lock()
	defer m.lookupMutex.Unlock()

	res := make([]common.Hash, len(userOps))
	for i, userOp := range userOps {
		hash := userOp.Hash(m.protocol.ChainId())
		m.userOpSenderLookup[hash] = aspect
		res[i] = hash
	}

	return res
}

func (m *Manager) simulateValidate(ctx context.Context, aspect common.Address, userOp *aa.UserOperation) error {
	// get vm with canonical state
	cvm, err := m.protocol.VMFromCanonicalState()
	if err != nil {
		return err
	}

	// call simulateValidation method of entry point contract to validate the operation
	calldata, err := m.entrypointABI.Pack("simulateValidation", userOp)
	if err != nil {
		return err
	}

	ret, _, err := cvm.Call(ctx, vm.AccountRef(aspect), aa.EntryPointContract,
		calldata, userOp.CallGasLimit.Uint64(), big.NewInt(0))
	if err != nil && !errors.Is(err, vm.ErrExecutionReverted) {
		return err
	}

	res, err := aa.DecodeValidationResult(ret)
	if err != nil {
		// return fail reason
		return aa.DecodeFailedOpError(ret)
	}

	if res.ReturnInfo != nil && res.ReturnInfo.SigFailed {
		// this should not happen, since the JIT inherent does not have a signature
		return errors.New("signature verification failed")
	}

	return nil
}

func (m *Manager) getNonce(ctx context.Context, address common.Address, key *big.Int) (*big.Int, error) {
	// FIXME: get vm with canonical state, this is just a temporary solution
	cvm, err := m.protocol.VMFromSnapshotState()
	if err != nil {
		return nil, err
	}

	// call simulateValidation method of entry point contract to validate the operation
	calldata, err := m.entrypointABI.Pack("getNonce", address, key)
	if err != nil {
		return nil, err
	}

	// FIXME: use a fixed gas limit for now
	ret, _, err := cvm.Call(ctx, vm.AccountRef(address), aa.EntryPointContract,
		calldata, 100000, big.NewInt(0))
	if err != nil && !errors.Is(err, vm.ErrExecutionReverted) {
		return nil, err
	}

	decoded, err := aa.DecodeResponse("getNonce", ret)
	if err != nil {
		return nil, err
	}

	return decoded[0].(*big.Int), nil
}

func NewUserOperation(leftoverGas uint64, maxFeePerGas uint64, maxPriorityFeePerGas uint64, protoMsg *types.JitInherentRequest) *aa.UserOperation {
	zero := new(big.Int)
	callGasLimit := new(big.Int).SetUint64(protoMsg.CallGasLimit)
	verificationGasLimit := new(big.Int).SetUint64(protoMsg.VerificationGasLimit)
	if verificationGasLimit.Cmp(zero) <= 0 {
		// by default use 1/5 remaining gas for verification
		verificationGasLimit.SetUint64(leftoverGas / 5)
	}
	if callGasLimit.Cmp(zero) <= 0 {
		// by default use 4/5 remaining gas for call
		callGasLimit.SetUint64(leftoverGas * 3 / 5)
	}

	nonceKey := uint256.NewInt(0).SetBytes(protoMsg.NonceKey)
	nonceKey.Lsh(nonceKey, 64)

	userOp := &aa.UserOperation{
		Sender:               common.BytesToAddress(protoMsg.Sender),
		Nonce:                nonceKey.Add(nonceKey, uint256.NewInt(0).SetUint64(protoMsg.Nonce)).ToBig(),
		InitCode:             protoMsg.InitCode,
		CallData:             protoMsg.CallData,
		CallGasLimit:         callGasLimit,
		VerificationGasLimit: verificationGasLimit,
		PreVerificationGas:   big.NewInt(21000), // Use this fixed value for now, unless we came up a more proper one
		MaxFeePerGas:         new(big.Int).SetUint64(maxFeePerGas),
		MaxPriorityFeePerGas: new(big.Int).SetUint64(maxPriorityFeePerGas),
		PaymasterAndData:     protoMsg.PaymasterAndData,
	}

	return userOp
}

func NewUserOperations(leftoverGas uint64, maxFeePerGas uint64, maxPriorityFeePerGas uint64, protoMsg ...*types.JitInherentRequest) []*aa.UserOperation {
	userOps := make([]*aa.UserOperation, len(protoMsg))
	for i, msg := range protoMsg {
		userOps[i] = NewUserOperation(leftoverGas, maxFeePerGas, maxPriorityFeePerGas, msg)
	}
	return userOps
}

type aaBundleTx struct {
	from      common.Address
	data      []byte
	gas       uint64
	gasPrice  *big.Int
	gasTipCap *big.Int
	gasFeeCap *big.Int
	nonce     uint64
	extra     map[string]interface{}
}

func (t *aaBundleTx) TxType() byte {
	return types2.DynamicFeeTxType
}

func (t *aaBundleTx) From() common.Address {
	return t.from
}

func (t *aaBundleTx) To() common.Address {
	return aa.EntryPointContract
}

func (t *aaBundleTx) Data() []byte {
	return t.data
}

func (t *aaBundleTx) Gas() uint64 {
	return t.gas
}

func (t *aaBundleTx) GasPrice() *big.Int {
	return t.gasPrice
}

func (t *aaBundleTx) GasTipCap() *big.Int {
	return t.gasTipCap
}

func (t *aaBundleTx) GasFeeCap() *big.Int {
	return t.gasFeeCap
}

func (t *aaBundleTx) Value() *big.Int {
	return big.NewInt(1)
}

func (t *aaBundleTx) Nonce() uint64 {
	return t.nonce
}

func (t *aaBundleTx) Extra() map[string]interface{} {
	return t.extra
}
