package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	ARTELA_ADDR = "0x0000000000000000000000000000000000A27E14"
)

type RequestEthTxAspect struct {
	Tx          *ethtypes.Transaction
	Context     map[string]string
	BlockHeight int64
	BlockHash   string
	TxIndex     int64
	BaseFee     int64
	ChainId     string
}

// ResponseAspect txhash->aspectId-> AspectOutPut
type ResponseAspect struct {
	ResultMap map[string]map[string]*AspectOutput
}

func (c ResponseAspect) With(txHash string, aspectId string, output *AspectOutput) ResponseAspect {
	c.ResultMap[txHash][aspectId] = output
	return c
}
func (c ResponseAspect) GetAspectResult(txHash string, aspectId string) *AspectOutput {
	return c.ResultMap[txHash][aspectId]
}
func (c ResponseAspect) GetTXResult(txHash string) []*AspectOutput {
	m := c.ResultMap[txHash]
	outputs := make([]*AspectOutput, 0)
	for _, output := range m {
		outputs = append(outputs, output)
	}
	return outputs
}
func (c ResponseAspect) Merge(out *ResponseAspect) {
	if out == nil {
		return
	}
	for tx, m := range out.ResultMap {
		for k, v := range m {
			c.With(tx, k, v)
		}
	}
}

type RequestSdkTxAspect struct {
	Tx sdk.Tx

	Context     map[string]string
	BlockHeight int64
	BlockHash   string
	TxIndex     int64
	BaseFee     int64
	ChainId     string
}

type RequestBlockAspect struct {
	BlockHeight int64
	ChainId     string
	Context     map[string]string
}

type ResponseBlockAspect struct {
	Result AspectOutput
}

type SdkTxEndPoint interface {
	IsEthTx(tx sdk.Tx) bool
	ConvertEthTx(tx sdk.Tx) ethtypes.Transaction
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
