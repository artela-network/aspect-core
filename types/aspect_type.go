package types

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// for jit-inherent
var (
	GetAspectContext func(ctx context.Context, aspectId common.Address, key string) ([]byte, error)
	SetAspectContext func(ctx context.Context, aspectId common.Address, key string, value []byte) error
)

type PointCut string

type AspectProvider interface {
	GetTxBondAspects(context.Context, common.Address, PointCut) ([]*AspectCode, error)
	GetAccountVerifiers(context.Context, common.Address) ([]*AspectCode, error)
	GetLatestBlock() int64
}

const (
	FILTER_TX                 PointCut = "filterTx"
	VERIFY_TX                 PointCut = "verifyTx"
	ON_GAS_PAYMENT_METHOD     PointCut = "onGasPayment"
	PRE_TX_EXECUTE_METHOD     PointCut = "preTxExecute"
	PRE_CONTRACT_CALL_METHOD  PointCut = "preContractCall"
	POST_CONTRACT_CALL_METHOD PointCut = "postContractCall"
	POST_TX_EXECUTE_METHOD    PointCut = "postTxExecute"
	POST_TX_COMMIT            PointCut = "postTxCommit"

	INIT_METHOD      PointCut = "init"
	OPERATION_METHOD PointCut = "operation"
	IS_OWNER_METHOD  PointCut = "isOwner"
)

type JoinPointRunType int64

func (j JoinPointRunType) String() string {
	return joinPointName[j]
}

func (j JoinPointRunType) IsPreCall() bool {
	return j == JoinPointRunType_PreTxExecute || j == JoinPointRunType_PreContractCall
}

const (
	JoinPointRunType_Unknown          JoinPointRunType = 0
	JoinPointRunType_VerifyTx         JoinPointRunType = 1
	JoinPointRunType_PreTxExecute     JoinPointRunType = 2
	JoinPointRunType_PreContractCall  JoinPointRunType = 4
	JoinPointRunType_PostContractCall JoinPointRunType = 8
	JoinPointRunType_PostTxExecute    JoinPointRunType = 16

	TransactionLevelJP = int64(JoinPointRunType_PreTxExecute) + int64(JoinPointRunType_PreContractCall) + int64(JoinPointRunType_PostContractCall) + int64(JoinPointRunType_PostTxExecute)
)

// Enum value maps for JoinPointRunType.
var (
	JoinPointRunType_value = map[string]int64{
		string(VERIFY_TX):                 int64(JoinPointRunType_VerifyTx),
		string(PRE_TX_EXECUTE_METHOD):     int64(JoinPointRunType_PreTxExecute),
		string(PRE_CONTRACT_CALL_METHOD):  int64(JoinPointRunType_PreContractCall),
		string(POST_CONTRACT_CALL_METHOD): int64(JoinPointRunType_PostContractCall),
		string(POST_TX_EXECUTE_METHOD):    int64(JoinPointRunType_PostTxExecute),
	}
)

var (
	joinPointName = map[JoinPointRunType]string{
		JoinPointRunType_VerifyTx:         string(VERIFY_TX),
		JoinPointRunType_PreTxExecute:     string(PRE_TX_EXECUTE_METHOD),
		JoinPointRunType_PreContractCall:  string(PRE_CONTRACT_CALL_METHOD),
		JoinPointRunType_PostContractCall: string(POST_CONTRACT_CALL_METHOD),
		JoinPointRunType_PostTxExecute:    string(POST_TX_EXECUTE_METHOD),
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

func CheckIsTransactionLevel(runJPs int64) bool {
	return runJPs&TransactionLevelJP > 0
}
func CheckIsTxVerifier(runJPs int64) bool {
	return runJPs&(int64(JoinPointRunType_VerifyTx)) == int64(JoinPointRunType_VerifyTx)
}
