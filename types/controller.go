package types

import (
	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	ARTELA_ADDR = "0x0000000000000000000000000000000000A27E14"
)

type RequestEthTxAspect struct {
	Tx          *ethtypes.Transaction
	BlockHeight int64
	BlockHash   string
	TxIndex     int64
	BaseFee     int64
	ChainId     string
}

// ResponseAspect txhash->aspectId-> AspectOutPut
type ResponseAspect struct {
	Success  bool
	Err      error
	GasInfo  *sdk.GasInfo
	AspectId string
}

func (c ResponseAspect) HasErr() bool {
	return c.Success == false || c.Err != nil
}

func (c ResponseAspect) WithAspectOutput(output *AspectOutput) ResponseAspect {
	if output != nil {
		c.Success = output.Success
		if output.Success == false {
			c.Err = errors.New(output.Message)
		}
	}
	return c
}
func (c ResponseAspect) WithErr(err error) ResponseAspect {
	if err != nil {
		c.Err = err
		c.Success = false
	}
	return c
}
func (c ResponseAspect) WithGas(gasWanted, gasUsed uint64) ResponseAspect {
	if gasUsed > 0 && gasWanted > 0 {
		info := &sdk.GasInfo{
			GasWanted: gasWanted,
			GasUsed:   gasUsed,
		}
		c.GasInfo = info
	}
	return c
}
func (c ResponseAspect) WithAspectId(aspectId string) ResponseAspect {
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
