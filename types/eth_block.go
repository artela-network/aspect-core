package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	rpctypes "github.com/evmos/ethermint/rpc/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"google.golang.org/protobuf/proto"
)

type BlockRet = BlockOutput

func NewBlockRet(success bool, errorMsg string, block *Block) *BlockRet {
	return &BlockRet{
		Res: &BlockOutput_Result{
			Success: success,
			Error:   errorMsg,
		},
		Block: block,
	}
}

func (ret *BlockRet) MarshalProto() ([]byte, error) {
	return proto.Marshal(ret)
}

func (ret *BlockRet) UnmarshalProto(data []byte) error {
	return proto.Unmarshal(data, ret)
}

// EthBlock for querying ethereum block from aspect
type Block = EthBlock

func NewEthBlock() *Block {
	return &Block{
		Header: &EthHeader{
			ReceiptHash: ethtypes.EmptyRootHash.String(),
			UncleHash:   ethtypes.EmptyUncleHash.String(),
			TxHash:      ethtypes.EmptyRootHash.String(),
		},
	}
}

// FillTendermintHeader fill a tendermint header to ethBlock
func (block *Block) FillTendermintHeader(header tmtypes.Header) {
	block.Header.Number = uint64(header.Height)
	block.Header.ParentHash = common.BytesToHash(header.LastBlockID.Hash.Bytes()).String()
	block.Header.StateRoot = common.BytesToHash(header.AppHash).String()
	block.Header.Timestamp = uint64(header.Time.Unix())
	block.DataHash = header.DataHash.Bytes()
	block.Hash = common.BytesToHash(header.Hash()).String()
}

// FillProtoHeader fill a proto header to ethBlock
func (block *EthBlock) FillProtoHeader(header tmproto.Header) {
	block.Header.Number = uint64(header.Height)
	block.Header.ParentHash = common.BytesToHash(header.LastBlockId.Hash).String()
	block.Header.StateRoot = common.BytesToHash(header.AppHash).String()
	block.Header.Timestamp = uint64(header.Time.Unix())
	block.DataHash = header.DataHash
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

func (block *Block) FillTransactions(trans []*rpctypes.RPCTransaction) {
	block.Transactions = make([]*AspTransaction, len(trans))
	if len(trans) > 0 {
		block.Header.TxHash = common.BytesToHash(block.DataHash).String()
	}

	for i, tran := range trans {
		block.Transactions[i] = &AspTransaction{
			ChainId:     tran.ChainID.String(),
			Nonce:       uint64(tran.Nonce),
			GasTipCap:   tran.GasTipCap.String(),
			GasFeeCap:   tran.GasFeeCap.String(),
			GasLimit:    0,
			GasPrice:    tran.GasPrice.ToInt().Uint64(),
			To:          tran.To.String(),
			Value:       tran.Value.ToInt().Uint64(),
			Input:       tran.Input,
			BlockHash:   tran.BlockHash.Bytes(),
			BlockNumber: tran.BlockNumber.ToInt().Uint64(),
			From:        tran.From.String(),
			Hash:        tran.Hash.Bytes(),
			Type:        uint64(tran.Type),
			V:           tran.V.ToInt().Bytes(),
			R:           tran.R.ToInt().Bytes(),
			S:           tran.S.ToInt().Bytes(),
		}
	}
}

func (block *Block) FillBloom(bloom ethtypes.Bloom) {
	// block.Header.Bloom = bloom
}

func (block *Block) FillValidatorAddr(addr common.Address) {
	block.Header.Miner = addr.String()
}

func (block *Block) FillBaseFee(baseFee *big.Int) {
	block.Header.BaseFeePerGas = baseFee.Uint64()
}

func (block *Block) MarshalProto() ([]byte, error) {
	return proto.Marshal(block)
}

func (block *Block) UnmarshalProto(data []byte) error {
	return proto.Unmarshal(data, block)
}
