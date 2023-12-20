package jit_inherent

import (
	"context"
	"math/big"
	"sync"

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
) (*types.JitInherentResponse, error) {
	if len(inherents) == 0 {
		return nil, errors.New("no jit inherent to submit")
	}

	userOps := NewUserOperations(inherents...)
	m.cacheUserOp(aspect, userOps...)

	switch stage {
	case integration.BlockInitialization:
		return m.submitJITTx(ctx, aspect, userOps...)
	case integration.TransactionExecution:
		if len(inherents) != 1 {
			return nil, errors.New("only one user operation is allowed in current join point")
		}
		return m.submitJITCall(ctx, aspect, gas, userOps[0], userOps[0].Hash(m.protocol.ChainId()))
	default:
		return nil, errors.New("cannot submit jit inherent in current join point")
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
func (m *Manager) submitJITCall(ctx context.Context, aspect common.Address, gas uint64, userOp *aa.UserOperation, userOpHash common.Hash) (
	*types.JitInherentResponse, error,
) {
	defer m.ClearUserOp(userOpHash)

	// get current evm instance with snapshot state
	evm, err := m.protocol.VMFromSnapshotState()
	if err != nil {
		return nil, err
	}

	resp := &types.JitInherentResponse{
		JitInherentHashes: [][]byte{userOpHash.Bytes()},
		Success:           false,
	}

	// FIXME: pay gas with Aspect's settlement account
	callData, err := m.entrypointABI.Pack("handleOps", []aa.UserOperation{*userOp}, userOp.Sender)
	if err != nil {
		return resp, err
	}
	ret, leftoverGas, err := evm.Call(ctx, vm.AccountRef(aspect), aa.EntryPointContract, callData, gas, big.NewInt(0))
	resp.Success = err == nil
	resp.LeftoverGas = leftoverGas
	if err == nil || (err != nil && err.Error() == vm.ErrExecutionReverted.Error()) {
		resp.Ret = ret
		resp.Success = true
		err = nil
	}
	return resp, err
}

func (m *Manager) cacheUserOp(aspect common.Address, userOps ...*aa.UserOperation) {
	m.lookupMutex.Lock()
	defer m.lookupMutex.Unlock()

	for _, userOp := range userOps {
		m.userOpSenderLookup[userOp.Hash(m.protocol.ChainId())] = aspect
	}
}

func (m *Manager) submitJITTx(ctx context.Context, aspect common.Address, userOps ...*aa.UserOperation) (
	*types.JitInherentResponse, error,
) {
	// one fails all
	userOpHashes := make([][]byte, len(userOps))
	for i, userOp := range userOps {
		userOpHashes[i] = userOp.Hash(m.protocol.ChainId()).Bytes()
		// simulate the user op validation, drop the jit tx if any of the user op failed the validation
		if err := m.simulateValidate(ctx, aspect, userOp); err != nil {
			return nil, errors.Errorf("user operation #%d validation failed, reason: %s", i, err)
		}
	}

	// convert tx with the protocol side, since the Aspect framework is not supposed to know the tx format
	callData, err := aa.PackCallData(userOps, aspect)
	if err != nil {
		return nil, err
	}

	// build aa bundled tx
	tx := &aaBundleTx{
		from: aspect,
		data: callData,
	}

	// estimate transaction gas cost
	tx.gas, err = m.protocol.EstimateGas(tx)
	if err != nil {
		return nil, err
	}

	// check out current gas price
	tx.gasPrice, err = m.protocol.GasPrice()
	if err != nil {
		return nil, err
	}

	// get last block header
	blockHeader, err := m.protocol.LastBlockHeader()
	if err != nil {
		return nil, err
	}

	// get account nonce
	tx.nonce, err = m.protocol.NonceOf(aspect)
	if err != nil {
		return nil, err
	}

	// use base fee as cap, inherent tx does not need to pay for the priority fee
	// TODO: discuss later whether priority fee should be paid or not
	tx.gasFeeCap = blockHeader.BaseFee()
	tx.gasTipCap = blockHeader.BaseFee()

	// convert to underlying protocol tx
	protocolTx, err := m.protocol.ConvertProtocolTx(tx)
	if err != nil {
		return nil, err
	}

	// submit tx to current proposal, this should be handled by the protocol side
	if err := m.protocol.SubmitTxToCurrentProposal(protocolTx); err != nil {
		return nil, err
	}

	return &types.JitInherentResponse{
		JitInherentHashes: userOpHashes,
		TxHash:            protocolTx.Hash(),
		Success:           true,
		Ret:               nil,
	}, nil
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

func NewUserOperation(protoMsg *types.JitInherentRequest) *aa.UserOperation {
	return &aa.UserOperation{
		Sender:               common.BytesToAddress(protoMsg.Sender),
		Nonce:                big.NewInt(0).SetBytes(protoMsg.Nonce),
		InitCode:             protoMsg.InitCode,
		CallData:             protoMsg.CallData,
		CallGasLimit:         big.NewInt(0).SetBytes(protoMsg.CallGasLimit),
		VerificationGasLimit: big.NewInt(0).SetBytes(protoMsg.VerificationGasLimit),
		PreVerificationGas:   big.NewInt(10000), // Fixed gas overhead compensation for verification
		MaxFeePerGas:         big.NewInt(0).SetBytes(protoMsg.MaxFeePerGas),
		MaxPriorityFeePerGas: big.NewInt(0).SetBytes(protoMsg.MaxPriorityFeePerGas),
		PaymasterAndData:     protoMsg.PaymasterAndData,
	}
}

func NewUserOperations(protoMsg ...*types.JitInherentRequest) []*aa.UserOperation {
	userOps := make([]*aa.UserOperation, len(protoMsg))
	for i, msg := range protoMsg {
		userOps[i] = NewUserOperation(msg)
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
