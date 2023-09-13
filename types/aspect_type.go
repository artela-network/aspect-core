package types

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
)

var GetAspectContext func(contractAddr string, aspectId string, key string) string

type PointCut string

type AspectProvider interface {
	GetTxBondAspects(int64, common.Address) ([]*AspectCode, error)
	GetBlockBondAspects(int64) ([]*AspectCode, error)
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

func ErrCallMessageResponse(err error) *CallMessageResponse {
	return &CallMessageResponse{
		Result: ErrRunResult(err.Error()),
		Data:   nil,
	}
}

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
		if v.Result.Success == false {
			return true, errors.New(k + " " + v.Result.Message)
		}
	}
	return false, nil
}
func (c *JoinPointResult) WithResponse(aspectId string, output *AspectResponse) *JoinPointResult {
	if c.ExecResultMap == nil {
		c.ExecResultMap = make(map[string]*AspectResponse)
	}
	c.ExecResultMap[aspectId] = output
	return c
}
func (c *JoinPointResult) WithGas(gasWanted, gasUsed uint64) *JoinPointResult {
	if gasUsed > 0 && gasWanted > 0 {
		info := &GasInfo{
			GasWanted: gasWanted,
			GasUsed:   gasUsed,
		}
		c.GasInfo = info
	}
	return c
}
func (c *JoinPointResult) WithErr(aspectId string, err error) *JoinPointResult {
	if err == nil {
		return c
	}
	result := &RunResult{
		Success: false,
		Message: err.Error(),
	}
	response := &AspectResponse{
		Result: result,
	}
	c.WithResponse(aspectId, response)
	return c
}
