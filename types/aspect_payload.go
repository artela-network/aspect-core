package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (input *AspectInput) ToJSON() (string, error) {
	marshal, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(marshal), err
}

func (input *AspectInput) FromJSON(data string) error {
	err := json.Unmarshal([]byte(data), input)
	return err
}

func (output *AspectOutput) ToJSON() (string, error) {
	marshal, err := json.Marshal(output)
	if err != nil {
		return "", err
	}
	return string(marshal), err
}

func (output *AspectOutput) FromJSON(data string) error {
	err := json.Unmarshal([]byte(data), output)
	return err
}

// NewTransactionFromData returns a transaction that will serialize to the RPC
// representation, with the given location metadata set (if available).
func NewTx(
	tx *ethtypes.Transaction, blockHash common.Hash, blockNumber, index int64, baseFee *big.Int,
	chainID *big.Int,
) (*AspTransaction, error) {
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

	result := &AspTransaction{
		ChainId:          tx.ChainId().String(),
		Nonce:            tx.Nonce(),
		GasTipCap:        tx.GasTipCap().String(),
		GasFeeCap:        tx.GasFeeCap().String(),
		GasLimit:         tx.Gas(),
		To:               tx.To().Hex(),
		Value:            tx.Value().Uint64(),
		Input:            tx.Data(),
		AccessList:       nil,
		BlockHash:        nil,
		BlockNumber:      0,
		From:             from.Hex(),
		Hash:             tx.Hash().Bytes(),
		TransactionIndex: 0,
		Type:             uint64(tx.Type()),
		V:                v.Bytes(),
		R:                r.Bytes(),
		S:                s.Bytes(),
	}
	if blockHash != (common.Hash{}) {
		result.BlockHash = blockHash.Bytes()
		result.BlockNumber = blockNumber
		result.TransactionIndex = index
	}
	switch tx.Type() {
	case ethtypes.AccessListTxType:
		al := tx.AccessList()
		accList := convertTuples(al)
		result.AccessList = accList
		result.ChainId = tx.ChainId().String()
	case ethtypes.DynamicFeeTxType:
		al := tx.AccessList()
		result.AccessList = convertTuples(al)
		result.ChainId = tx.ChainId().String()
		result.GasFeeCap = tx.GasFeeCap().String()
		result.GasTipCap = tx.GasTipCap().String()
		// if the transaction has been mined, compute the effective gas price
		if baseFee != nil && blockHash != (common.Hash{}) {
			// price = min(tip, gasFeeCap - baseFee) + baseFee
			price := math.BigMin(new(big.Int).Add(tx.GasTipCap(), baseFee), tx.GasFeeCap())
			result.GasPrice = price.Uint64()
		} else {
			result.GasPrice = tx.GasFeeCap().Uint64()
		}
	}
	return result, nil
}

func convertTuple(tuple ethtypes.AccessTuple) AspAccessTuple {
	address := tuple.Address
	storageKeys := tuple.StorageKeys
	storeKey := make([]string, 0)
	for _, key := range storageKeys {
		storeKey = append(storeKey, key.String())
	}
	return AspAccessTuple{
		Address:     address.String(),
		StorageKeys: storeKey,
	}
}

func convertTuples(tuples []ethtypes.AccessTuple) []*AspAccessTuple {
	store := make([]*AspAccessTuple, 0)
	for _, tuple := range tuples {
		accessTuple := convertTuple(tuple)
		store = append(store, &accessTuple)
	}
	return store
}
