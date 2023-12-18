package types

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// for jit-inherent
var (
	GetAspectContext func(aspectId string, key string) string
	SetAspectContext func(aspectId string, key string, value string)
)

var GetAspectPaymaster func(blockNum int64, aspectId common.Address) (*common.Address, error)

type PointCut string

type AspectProvider interface {
	GetTxBondAspects(int64, common.Address, PointCut) ([]*AspectCode, error)
	GetAccountVerifiers(int64, common.Address) ([]*AspectCode, error)
	GetBlockBondAspects(int64) ([]*AspectCode, error)
	GetLatestBlock() int64
	CreateTxPointRequestWithData(data []byte) (*EthTxAspect, error)
}

const (
	FILTER_TX                  PointCut = "filterTx"
	ON_BLOCK_INITIALIZE_METHOD PointCut = "onBlockInitialize"
	VERIFY_TX                  PointCut = "verifyTx"
	ON_ACCOUNT_VERIFY_METHOD   PointCut = "onAccountVerify"
	ON_GAS_PAYMENT_METHOD      PointCut = "onGasPayment"
	PRE_TX_EXECUTE_METHOD      PointCut = "preTxExecute"
	PRE_CONTRACT_CALL_METHOD   PointCut = "preContractCall"
	POST_CONTRACT_CALL_METHOD  PointCut = "postContractCall"
	POST_TX_EXECUTE_METHOD     PointCut = "postTxExecute"
	POST_TX_COMMIT             PointCut = "postTxCommit"
	ON_BLOCK_FINALIZE_METHOD   PointCut = "onBlockFinalize"
	OPERATION_METHOD           PointCut = "operation"
	IS_OWNER_METHOD            PointCut = "isOwner"
	ON_CONTRACT_BINDING_METHOD PointCut = "onContractBinding"
)

const DefaultKey = "-"

type JoinPointRunType int64

const (
	JoinPointRunType_VerifyTx          JoinPointRunType = 1
	JoinPointRunType_PreTxExecute      JoinPointRunType = 2
	JoinPointRunType_PreContractCall   JoinPointRunType = 4
	JoinPointRunType_PostContractCall  JoinPointRunType = 8
	JoinPointRunType_PostTxExecute     JoinPointRunType = 16
	JoinPointRunType_PostTxCommit      JoinPointRunType = 32
	JoinPointRunType_OnBlockInitialize JoinPointRunType = 64
	JoinPointRunType_OnBlockFinalize   JoinPointRunType = 128

	BlockLevelJP = int64(JoinPointRunType_OnBlockInitialize) + int64(JoinPointRunType_OnBlockFinalize)

	TransactionLevelJP = int64(JoinPointRunType_PreTxExecute) + int64(JoinPointRunType_PreContractCall) + int64(JoinPointRunType_PostContractCall) + int64(JoinPointRunType_PostTxExecute) + int64(JoinPointRunType_PostTxCommit)
)

// Enum value maps for JoinPointRunType.
var (
	JoinPointRunType_value = map[string]int64{
		string(VERIFY_TX):                  int64(JoinPointRunType_VerifyTx),
		string(PRE_TX_EXECUTE_METHOD):      int64(JoinPointRunType_PreTxExecute),
		string(PRE_CONTRACT_CALL_METHOD):   int64(JoinPointRunType_PreContractCall),
		string(POST_CONTRACT_CALL_METHOD):  int64(JoinPointRunType_PostContractCall),
		string(POST_TX_EXECUTE_METHOD):     int64(JoinPointRunType_PostTxExecute),
		string(POST_TX_COMMIT):             int64(JoinPointRunType_PostTxCommit),
		string(ON_BLOCK_INITIALIZE_METHOD): int64(JoinPointRunType_OnBlockInitialize),
		string(ON_BLOCK_FINALIZE_METHOD):   int64(JoinPointRunType_OnBlockFinalize),
	}
)

func CheckIsJoinPoint(runJPs *big.Int) (map[int64]string, bool) {
	if runJPs == nil {
		return nil, false
	}
	runValue := runJPs.Int64()
	jpMap := make(map[int64]string)
	if runValue <= 0 {
		return jpMap, false
	}
	for k, v := range JoinPointRunType_value {
		// verify with & to see if it is included jp value，like:  5&1==1
		if runValue&v == v {
			jpMap[v] = k
		}
	}
	return jpMap, len(jpMap) > 0
}
func CanExecPoint(runJPs int64, cut PointCut) bool {
	if value, exit := JoinPointRunType_value[string(cut)]; exit {
		return runJPs&value == value
	}
	return false
}

func CheckIsBlockLevel(runJPs int64) bool {
	return runJPs&BlockLevelJP > 0
}
func CheckIsTransactionLevel(runJPs int64) bool {
	return runJPs&TransactionLevelJP > 0
}
func CheckIsTxVerifier(runJPs int64) bool {
	return runJPs&(int64(JoinPointRunType_VerifyTx)) == int64(JoinPointRunType_VerifyTx)
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
