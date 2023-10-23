package types

import (
	"math/big"

	aa "github.com/artela-network/artelasdk/chaincoreext/account_abstraction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type UserOperation struct {
	Sender               common.Address
	Nonce                *uint256.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *uint256.Int
	VerificationGasLimit *uint256.Int
	PreVerificationGas   *uint256.Int
	MaxFeePerGas         *uint256.Int
	MaxPriorityFeePerGas *uint256.Int
	PaymasterAndData     []byte
}

func NewUserOperation(protoMsg *JitInherentRequest) *UserOperation {
	return &UserOperation{
		Sender:               common.BytesToAddress(protoMsg.Sender),
		Nonce:                uint256.NewInt(0).SetBytes(protoMsg.Nonce),
		InitCode:             protoMsg.InitCode,
		CallData:             protoMsg.CallData,
		CallGasLimit:         uint256.NewInt(0).SetBytes(protoMsg.CallGasLimit),
		VerificationGasLimit: uint256.NewInt(0).SetBytes(protoMsg.VerificationGasLimit),
		PreVerificationGas:   uint256.NewInt(10000), // Fixed gas overhead compensation for verification
		MaxFeePerGas:         uint256.NewInt(0).SetBytes(protoMsg.MaxFeePerGas),
		MaxPriorityFeePerGas: uint256.NewInt(0).SetBytes(protoMsg.MaxPriorityFeePerGas),
		PaymasterAndData:     protoMsg.PaymasterAndData,
	}
}

func NewUserOperations(protoMsg ...*JitInherentRequest) []*UserOperation {
	userOps := make([]*UserOperation, len(protoMsg))
	for i, msg := range protoMsg {
		userOps[i] = NewUserOperation(msg)
	}
	return userOps
}

func (i UserOperation) Hash() common.Hash {
	return common.Hash{}
}

func (i UserOperation) ToABIStruct() *aa.UserOperation {
	return &aa.UserOperation{
		Sender:               i.Sender,
		Nonce:                i.Nonce.ToBig(),
		InitCode:             i.InitCode,
		CallData:             i.CallData,
		CallGasLimit:         i.CallGasLimit.ToBig(),
		VerificationGasLimit: i.VerificationGasLimit.ToBig(),
		PreVerificationGas:   i.PreVerificationGas.ToBig(),
		MaxFeePerGas:         i.MaxFeePerGas.ToBig(),
		MaxPriorityFeePerGas: i.MaxPriorityFeePerGas.ToBig(),
		PaymasterAndData:     i.PaymasterAndData,
	}
}

func (i UserOperation) ToEstimateGasABIStruct(balance *big.Int) *aa.UserOperation {
	// TODO: use binary search to estimate gas cost,
	//       currently we just use one quarter of the sender balance.
	quarter := balance.Div(balance, big.NewInt(4))

	return &aa.UserOperation{
		Sender:               i.Sender,
		Nonce:                i.Nonce.ToBig(),
		InitCode:             i.InitCode,
		CallData:             i.CallData,
		CallGasLimit:         quarter,
		VerificationGasLimit: quarter,
		PreVerificationGas:   i.PreVerificationGas.ToBig(),
		MaxFeePerGas:         i.MaxFeePerGas.ToBig(),
		MaxPriorityFeePerGas: i.MaxPriorityFeePerGas.ToBig(),
		PaymasterAndData:     i.PaymasterAndData,
	}
}
