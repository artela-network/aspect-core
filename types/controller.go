package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	ARTELA_ADDR = "0x0000000000000000000000000000000000A27E14"
)

func IsAspectContract(to *common.Address) bool {
	if to != nil && strings.EqualFold(ARTELA_ADDR, to.Hex()) {
		// ignore contract deployment transaction & aspect op txs
		return true
	}
	return false
}

type RequestEthTxAspect struct {
	Tx          *ethtypes.Transaction
	BlockHeight int64
	BlockHash   string
	TxIndex     int64
	BaseFee     int64
	ChainId     string
}
type RequestEthMsgAspect struct {
	BlockHeight int64
	TxHash      *common.Hash
	TxIndex     uint
	To          *common.Address
	From        common.Address
	Nonce       uint64
	GasLimit    uint64
	GasPrice    *big.Int
	GasFeeCap   *big.Int
	GasTipCap   *big.Int
	Value       *big.Int
	TxType      uint64
	TxData      []byte
	AccessList  ethtypes.AccessList
	ChainId     string
}

func (msg *RequestEthMsgAspect) ToAspTx() *AspTransaction {
	// Determine the signer. For replay-protected transactions, use the most permissive
	// signer, because we assume that signers are backwards-compatible with old
	// transactions. For non-protected transactions, the homestead signer signer is used
	// because the return value of ChainId is zero for those transactions.

	result := &AspTransaction{
		ChainId:          msg.ChainId,
		Nonce:            msg.Nonce,
		GasLimit:         msg.GasLimit,
		Input:            msg.TxData,
		BlockNumber:      msg.BlockHeight,
		TransactionIndex: int64(msg.TxIndex),
		Type:             msg.TxType,
	}
	if len(msg.From) > 0 {
		result.From = msg.From.Hex()
	}
	if msg.To != nil {
		result.To = msg.To.Hex()
	}
	if msg.GasTipCap != nil {
		result.GasTipCap = msg.GasTipCap.String()
	}
	if msg.GasFeeCap != nil {
		result.GasFeeCap = msg.GasFeeCap.String()
	}
	if msg.Value != nil {
		result.Value = msg.Value.Uint64()
	}
	if msg.TxHash != nil {
		result.Hash = msg.TxHash.Bytes()
	}

	if len(msg.AccessList) > 0 {
		al := msg.AccessList
		accList := convertTuples(al)
		result.AccessList = accList
	}

	return result

}

// ResponseAspect txhash->aspectId-> AspectOutPut
type ResponseAspect struct {
	Success  bool
	Err      error
	GasInfo  *sdk.GasInfo
	AspectId string
}

func (c *ResponseAspect) HasErr() bool {
	return c.Success == false || c.Err != nil
}

func (c *ResponseAspect) WithAspectOutput(output *AspectOutput) *ResponseAspect {
	if output != nil {
		c.Success = output.Success
		if output.Success == false {
			c.Err = errors.New(output.Message)
		}
	}
	return c
}
func (c *ResponseAspect) WithErr(err error) *ResponseAspect {
	if err != nil {
		c.Err = err
		c.Success = false
	}
	return c
}
func (c *ResponseAspect) WithGas(gasWanted, gasUsed uint64) *ResponseAspect {
	if gasUsed > 0 && gasWanted > 0 {
		info := &sdk.GasInfo{
			GasWanted: gasWanted,
			GasUsed:   gasUsed,
		}
		c.GasInfo = info
	}
	return c
}
func (c *ResponseAspect) WithAspectId(aspectId string) *ResponseAspect {
	c.AspectId = aspectId
	return c
}

type RequestSdkTxAspect struct {
	Tx          sdk.Tx
	BlockHeight int64
	BlockHash   string
	TxIndex     int64
	BaseFee     int64
	ChainId     string
}

type RequestBlockAspect struct {
	BlockHeight int64
	ChainId     string
}

type ResponseBlockAspect struct {
	ResultMap map[string]*AspectOutput
}

func (c ResponseBlockAspect) With(aspectId string, output *AspectOutput) ResponseBlockAspect {
	if c.ResultMap == nil {
		c.ResultMap = make(map[string]*AspectOutput)
	}
	c.ResultMap[aspectId] = output
	return c
}

const (
	ON_TX_RECEIVE_METHOD       = "onTxReceive"
	ON_BLOCK_INITIALIZE_METHOD = "onBlockInitialize"
	ON_TX_VERIFY_METHOD        = "onTxVerify"
	ON_ACCOUNT_VERIFY_METHOD   = "onAccountVerify"
	ON_GAS_PAYMENT_METHOD      = "onGasPayment"
	PRE_TX_EXECUTE_METHOD      = "preTxExecute"
	PRE_CONTRACT_CALL_METHOD   = "preContractCall"
	POST_CONTRACT_CALL_METHOD  = "postContractCall"
	POST_TX_EXECUTE_METHOD     = "postTxExecute"
	ON_TX_COMMIT_METHOD        = "onTxCommit"
	ON_BLOCK_FINALIZE_METHOD   = "onBlockFinalize"
)

type (
	Pointcut uint8
)
