package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/proto"
)

// EthBlock for querying ethereum block from aspect
type Block = EthBlock

func NewEthBlock() *EthBlock {
	return &EthBlock{
		Header: &EthBlockHeader{
			ReceiptHash:      ethtypes.EmptyRootHash.String(),
			UncleHash:        ethtypes.EmptyUncleHash.String(),
			TransactionsRoot: ethtypes.EmptyRootHash.String(),
		},
	}
}

func (block *Block) FillHeader(header *EthBlockHeader, dataHash []byte, hash string) {
	block.Header = header
	block.DataHash = dataHash
	block.Hash = hash
}

func (block *Block) FillSize(size uint64) {
	block.Size = size
}

func (block *Block) FillGasLimit(gasLimit uint64) {
	block.Header.GasLimit = gasLimit
}

func (block *Block) FillGasUsed(gasUsed uint64) {
	block.Header.GasUsed = gasUsed
}

func (block *Block) FillTransactions(trans []*EthTransaction) {
	if len(trans) > 0 {
		block.Header.TransactionsRoot = common.BytesToHash(block.DataHash).String()
	}
	block.Transactions = trans
}

func (block *Block) FillBloom(bloom ethtypes.Bloom) {
	// block.Header.Bloom = bloom
}

func (block *Block) FillValidatorAddr(addr common.Address) {
	block.Header.Coinbase = addr.String()
}

func (block *Block) FillBaseFee(baseFee *big.Int) {
	block.Header.BaseFee = baseFee.Uint64()
}

func (block *Block) MarshalProto() ([]byte, error) {
	return proto.Marshal(block)
}

func (block *Block) UnmarshalProto(data []byte) error {
	return proto.Unmarshal(data, block)
}
