package types

import (
	"github.com/ethereum/go-ethereum/core"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func Ternary[T string | []byte | uint64](condition bool, trueValue func() T, falseValue T) T {
	if condition {
		return trueValue()
	}
	return falseValue
}

func NewInnerTransaction(
	from string,
	to string,
	data []byte,
	value string,
	gas string,
	ret []byte,
	leftOverGas uint64,
	index uint64,
	parentIndex uint64,
	childrenIndex []uint64,
) *EthStackTransaction {
	return &EthStackTransaction{
		From:          from,
		To:            to,
		Data:          data,
		Value:         value,
		Gas:           gas,
		Ret:           ret,
		LeftOverGas:   leftOverGas,
		Index:         index,
		ParentIndex:   parentIndex,
		ChildrenIndex: childrenIndex,
	}
}

func NewEthTransactionByMessage(message core.Message, txHash common.Hash, chainId string, blockHash common.Hash, blockHeight int64, txType uint8) *EthTransaction {
	return &EthTransaction{
		ChainId:     chainId,
		Nonce:       message.Nonce(),
		GasTipCap:   Ternary(message.GasTipCap() != nil, func() string { return message.GasTipCap().String() }, "0"),
		GasFeeCap:   Ternary(message.GasFeeCap() != nil, func() string { return message.GasFeeCap().String() }, "0"),
		Gas:         message.Gas(),
		GasPrice:    Ternary(message.GasPrice() != nil, func() string { return message.GasPrice().String() }, "0"),
		To:          Ternary(message.To() != nil, func() string { return message.To().Hex() }, ""),
		Value:       Ternary(message.Value() != nil, func() string { return message.Value().String() }, "0"),
		Input:       message.Data(),
		AccessList:  ConvertTuples(message.AccessList()),
		BlockHash:   blockHash.Bytes(),
		BlockNumber: blockHeight,
		From:        Ternary(message.From() != common.Address{}, func() string { return message.From().Hex() }, ""),
		Hash:        txHash.Bytes(),
		Type:        int32(txType),
	}
}

// NewTransactionFromData returns a transaction that will serialize to the RPC
// representation, with the given location metadata set (if available).
func NewEthTransaction(
	tx *ethtypes.Transaction, blockHash common.Hash, blockNumber, index int64, baseFee *big.Int,
	chainID string,
) (*EthTransaction, error) {
	// Determine the signer. For replay-protected transactions, use the most permissive
	// signer, because we assume that signers are backwards-compatible with old
	// transactions. For non-protected transactions, the homestead signer signer is used
	// because the return value of ChainId is zero for those transactions.
	var signer ethtypes.Signer
	if tx.Protected() {
		signer = ethtypes.LatestSignerForChainID(tx.ChainId())
	} else {
		signer = ethtypes.HomesteadSigner{}
	}
	from, _ := ethtypes.Sender(signer, tx)
	v, r, s := tx.RawSignatureValues()

	result := &EthTransaction{
		ChainId:          chainID,
		TransactionIndex: index,
		To:               Ternary(tx.To() != nil, func() string { return tx.To().Hex() }, ""),
		Value:            Ternary(tx.Value() != nil, func() string { return tx.Value().String() }, "0"),
		AccessList:       ConvertTuples(tx.AccessList()),
		BlockHash:        Ternary(blockHash != common.Hash{}, func() []byte { return blockHash.Bytes() }, []byte{0}),
		BlockNumber:      blockNumber,
		From:             Ternary(from != common.Address{}, func() string { return from.Hex() }, ""),
		Hash:             Ternary(tx.Hash() != common.Hash{}, func() []byte { return tx.Hash().Bytes() }, []byte{0}),
		Nonce:            tx.Nonce(),
		GasTipCap:        Ternary(tx.GasTipCap() != nil, func() string { return tx.GasTipCap().String() }, "0"),
		GasFeeCap:        Ternary(tx.GasFeeCap() != nil, func() string { return tx.GasFeeCap().String() }, "0"),
		Gas:              tx.Gas(),
		Input:            tx.Data(),
		Type:             int32(tx.Type()),
		V:                Ternary(v != nil, func() []byte { return v.Bytes() }, []byte{0}),
		R:                Ternary(r != nil, func() []byte { return r.Bytes() }, []byte{0}),
		S:                Ternary(s != nil, func() []byte { return s.Bytes() }, []byte{0}),
	}
	switch tx.Type() {
	case ethtypes.AccessListTxType:
		al := tx.AccessList()
		accList := ConvertTuples(al)
		result.AccessList = accList
		result.ChainId = tx.ChainId().String()
	case ethtypes.DynamicFeeTxType:
		al := tx.AccessList()
		result.AccessList = ConvertTuples(al)
		result.ChainId = tx.ChainId().String()
		result.GasFeeCap = tx.GasFeeCap().String()
		result.GasTipCap = tx.GasTipCap().String()
		// if the transaction has been mined, compute the effective gas price
		if baseFee != nil && blockHash != (common.Hash{}) {
			// price = min(tip, gasFeeCap - baseFee) + baseFee
			price := math.BigMin(new(big.Int).Add(tx.GasTipCap(), baseFee), tx.GasFeeCap())
			result.GasPrice = price.String()
		} else {
			result.GasPrice = tx.GasFeeCap().String()
		}
	}
	return result, nil
}

func ConvertTuple(tuple ethtypes.AccessTuple) EthAccessTuple {
	address := tuple.Address
	storageKeys := tuple.StorageKeys
	storeKey := make([]string, 0)
	for _, key := range storageKeys {
		storeKey = append(storeKey, key.String())
	}
	return EthAccessTuple{
		Address:     address.String(),
		StorageKeys: storeKey,
	}
}

func ConvertTuples(tuples []ethtypes.AccessTuple) []*EthAccessTuple {
	if tuples == nil || len(tuples) == 0 {
		return nil
	}
	store := make([]*EthAccessTuple, 0, len(tuples))
	for _, tuple := range tuples {
		accessTuple := ConvertTuple(tuple)
		store = append(store, &accessTuple)
	}
	return store
}
