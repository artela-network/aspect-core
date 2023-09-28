package types

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"strings"
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

var ContextKeys = [...]string{TxAspectContext, TxContent, TxStateChanges, TxExtProperties,
	TxCallTree, TxReceipt, TxGasMeter, EnvConsensusParams, EnvChainConfig, EnvEvmParams,
	EnvBaseInfo, BlockHeader, BlockGasMeter, BlockMinGasPrice, BlockLastCommit, BlockTxs}

// keypath match
// -- bool： match success
// -- string： context key
// -- []string： params
func HasContextKey(key string) (bool, string, []string) {
	for _, contextKey := range ContextKeys {
		if strings.HasPrefix(key, contextKey) {
			sdata := key[len(contextKey)-1:]
			split := make([]string, 0)
			if sdata != "" {
				split = strings.Split(sdata, DefConnector)
			}
			return true, contextKey, split
		}
	}
	return false, "", nil
}
