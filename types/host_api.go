package types

import (
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func ErrEthMessageCallResult(err error) *EthMessageCallResult {
	return &EthMessageCallResult{
		Hash:    "",
		Logs:    nil,
		Ret:     nil,
		VmError: err.Error(),
		GasUsed: 0,
	}
}

func NewContextQueryResponse(condition bool, errMsg string) *ContextQueryResponse {
	message := "success"
	if condition {
		message = errMsg
	}
	return &ContextQueryResponse{
		Result: &RunResult{
			Success: condition,
			Message: message,
		},
	}
}

func (c *ContextQueryResponse) SetData(message proto.Message) {
	if message == nil {
		return
	}
	anyData, _ := anypb.New(message)
	c.Data = anyData
}

const (
	TxAspectContext = "tx^context"
	TxContent       = "tx^content"
	TxStateChanges  = "tx^stateChanges"
	TxExtProperties = "tx^extProperties"
	TxCallTree      = "tx^callTree"
	TxReceipt       = "tx^receipt"
	TxGasMeter      = "tx^gasMeter"
	TxMsgHash       = "tx^msgHash"

	EnvConsensusParams = "env^consensusParams"
	EnvChainConfig     = "env^chainConfig"
	EnvEvmParams       = "env^evmParams"
	EnvBaseInfo        = "env^baseFee"

	BlockHeader      = "block^header"
	BlockGasMeter    = "block^gasMeter"
	BlockMinGasPrice = "block^minGasPrice"
	BlockLastCommit  = "block^lastCommit"
	BlockTxs         = "block^txs"
)

var DefConnector = "^"

var ContextKeys = [...]string{
	TxAspectContext, TxContent, TxStateChanges, TxExtProperties, TxMsgHash,
	TxCallTree, TxReceipt, TxGasMeter, EnvConsensusParams, EnvChainConfig, EnvEvmParams,
	EnvBaseInfo, BlockHeader, BlockGasMeter, BlockMinGasPrice, BlockLastCommit, BlockTxs,
}

// keypath match
// -- bool： match success
// -- string： context key
// -- []string： params
func HasContextKey(key string) (bool, string, []string) {
	split := make([]string, 0)
	for _, contextKey := range ContextKeys {
		if strings.HasPrefix(key, contextKey) {
			if key == contextKey {
				return true, contextKey, split
			}
			strData := key[len(contextKey)+1:]
			if strData != "" {
				split = strings.Split(strData, DefConnector)
			}
			return true, contextKey, split
		}
	}
	return false, "", split
}
