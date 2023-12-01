package types

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

// for jit-inherent
var (
	GetAspectContext func(aspectId string, key string) string
	SetAspectContext func(aspectId string, key string, value string)
)

var GetAspectPaymaster func(blockNum int64, aspectId common.Address) (*common.Address, error)

type PointCut string

type AspectProvider interface {
	GetTxBondAspects(int64, common.Address) ([]*AspectCode, error)
	GetAccountVerifiers(int64, common.Address) ([]*AspectCode, error)
	GetBlockBondAspects(int64) ([]*AspectCode, error)
	GetLatestBlock() int64
	CreateTxPointRequestWithData(data []byte) (*EthTxAspect, error)
}

const (
	ON_TX_RECEIVE_METHOD       PointCut = "onTxReceive"
	ON_BLOCK_INITIALIZE_METHOD PointCut = "onBlockInitialize"
	ON_TX_VERIFY_METHOD        PointCut = "onTxVerify"
	ON_ACCOUNT_VERIFY_METHOD   PointCut = "onAccountVerify"
	ON_GAS_PAYMENT_METHOD      PointCut = "onGasPayment"
	PRE_TX_EXECUTE_METHOD      PointCut = "preTxExecute"
	PRE_CONTRACT_CALL_METHOD   PointCut = "preContractCall"
	POST_CONTRACT_CALL_METHOD  PointCut = "postContractCall"
	POST_TX_EXECUTE_METHOD     PointCut = "postTxExecute"
	ON_TX_COMMIT_METHOD        PointCut = "onTxCommit"
	ON_BLOCK_FINALIZE_METHOD   PointCut = "onBlockFinalize"
	OPERATION_METHOD           PointCut = "operation"
	IS_OWNER_METHOD            PointCut = "isOwner"
	ON_CONTRACT_BINDING_METHOD PointCut = "onContractBinding"
)

const DefaultKey = "-"

func ErrRunResult(message string) *RunResult {
	return &RunResult{
		Success: false,
		Message: message,
	}
}

func DefRunResult() *RunResult {
	return &RunResult{
		Success: true,
		Message: "success",
	}
}

func ErrJoinPointResult(message string) *JoinPointResult {
	response := &AspectResponse{
		Result: &RunResult{
			Success: false,
			Message: message,
		},
		DataMessageType: DefaultKey,
	}
	m := make(map[string]*AspectResponse, 0)
	m[DefaultKey] = response
	return &JoinPointResult{
		ExecResultMap: m,
	}
}

func DefJoinPointResult(message string) *JoinPointResult {
	response := &AspectResponse{
		Result: &RunResult{
			Success: true,
			Message: message,
		},
		DataMessageType: DefaultKey,
	}
	m := make(map[string]*AspectResponse, 0)
	m[DefaultKey] = response
	return &JoinPointResult{
		ExecResultMap: m,
	}
}

func (c *JoinPointResult) HasErr() (bool, error) {
	if c.ExecResultMap == nil {
		return false, nil
	}
	for k, v := range c.ExecResultMap {
		if !v.Result.Success {
			return true, errors.New(k + " " + v.Result.Message)
		}
	}
	return false, nil
}

func (c *JoinPointResult) WithResponse(aspectId string, output *AspectResponse) *JoinPointResult {
	if c.ExecResultMap == nil {
		c.ExecResultMap = make(map[string]*AspectResponse)
	}
	if aspectId != "" && output != nil {
		c.ExecResultMap[aspectId] = output
	}
	return c
}

func (c *JoinPointResult) WithGas(gasWanted, gasUsed, gasLeft uint64) *JoinPointResult {
	c.GasInfo = &GasInfo{
		GasWanted: gasWanted,
		GasUsed:   gasUsed,
		Gas:       gasLeft,
	}
	return c
}

func (c *JoinPointResult) WithErr(aspectId string, err error) *JoinPointResult {
	errMsg := "fail"
	if err != nil {
		errMsg = err.Error()
	}
	result := &RunResult{
		Success: false,
		Message: errMsg,
	}
	response := &AspectResponse{
		Result: result,
	}
	c.WithResponse(aspectId, response)
	return c
}
