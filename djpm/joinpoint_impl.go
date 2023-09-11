package djpm

import (
	"github.com/artela-network/artelasdk/chaincoreext/scheduler"
	"github.com/artela-network/artelasdk/djpm/run"
	"github.com/artela-network/artelasdk/types"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/protobuf/proto"
)

var globalAspect *Aspect

type Aspect struct {
	provider types.AspectProvider
}

func NewAspect(provider types.AspectProvider) *Aspect {
	globalAspect = &Aspect{
		provider: provider,
	}
	return globalAspect
}
func AspectInstance() *Aspect {
	if globalAspect == nil {
		panic("aspcect instance not init,please exec NewAspect() first ")
	}
	return globalAspect
}
func (aspect Aspect) FilterTx(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.ON_TX_RECEIVE_METHOD, request)
}
func (aspect Aspect) VerifyTx(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.ON_TX_VERIFY_METHOD, request)

}
func (aspect Aspect) VerifyAccount(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.ON_ACCOUNT_VERIFY_METHOD, request)

}
func (aspect Aspect) GetPayMaster(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.ON_GAS_PAYMENT_METHOD, request)

}
func (aspect Aspect) PreTxExecute(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.PRE_TX_EXECUTE_METHOD, request)

}
func (aspect Aspect) PreContractCall(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.PRE_CONTRACT_CALL_METHOD, request)

}
func (aspect Aspect) PostContractCall(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.POST_CONTRACT_CALL_METHOD, request)

}
func (aspect Aspect) PostTxExecute(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.POST_TX_EXECUTE_METHOD, request)

}
func (aspect Aspect) PostTxCommit(request *types.EthTxAspect) *types.JoinPointResult {
	return aspect.transactionAdvice(types.ON_TX_COMMIT_METHOD, request)
}

func (aspect Aspect) OnBlockInitialize(request *types.EthBlockAspect) *types.JoinPointResult {
	return aspect.blockAdvice(types.ON_BLOCK_INITIALIZE_METHOD, request)
}
func (aspect Aspect) OnBlockFinalize(request *types.EthBlockAspect) *types.JoinPointResult {
	return aspect.blockAdvice(types.ON_BLOCK_FINALIZE_METHOD, request)
}

func (aspect Aspect) blockAdvice(method types.PointCut, req *types.EthBlockAspect) *types.JoinPointResult {
	if req == nil || method == "" {
		return types.DefJoinPointResult("blockAdvice input is empty.")
	}
	aspectCodes, err := aspect.provider.GetBlockBondAspects(int64(req.GetHeader().Number))
	if err != nil {
		return types.DefJoinPointResult("blockAdvice GetBlockBondAspects error." + err.Error())
	}
	// load aspects
	if aspectCodes == nil || len(aspectCodes) == 0 {
		return types.DefJoinPointResult("not bond aspects.")
	}
	// run aspects on received transaction
	return aspect.runAspect(method, int64(req.GetHeader().Number), req, aspectCodes)

}

func (aspect Aspect) transactionAdvice(method types.PointCut, req *types.EthTxAspect) *types.JoinPointResult {
	if req == nil || req.Tx == nil || len(req.Tx.Hash) == 0 || types.IsAspectContract(req.Tx.To) {
		return types.DefJoinPointResult("transactionAdvice invalid input.")
	}
	if req.Tx.To == "" {
		return types.DefJoinPointResult("it is create tx.")
	}
	//skip scheduleTx
	hash := req.Tx.Hash
	if scheduler.TaskInstance() != nil && hash != nil && scheduler.TaskInstance().IsScheduleTx(common.BytesToHash(hash)) {
		return types.DefJoinPointResult("it is task tx.")
	}

	aspectCodes, err := aspect.provider.GetTxBondAspects(req.GetTx().BlockNumber, common.HexToAddress(req.GetTx().To))
	if err != nil {
		return types.DefJoinPointResult("transactionAdvice GetTxBondAspects error." + err.Error())
	}
	if aspectCodes == nil || len(aspectCodes) == 0 {
		return types.DefJoinPointResult("not bond aspects.")
	}

	// run aspects on received transaction
	return aspect.runAspect(method, req.GetTx().BlockNumber, req, aspectCodes)

}

func (aspect Aspect) runAspect(method types.PointCut, blockNumber int64, reqData proto.Message, req []*types.AspectCode) *types.JoinPointResult {
	response := types.DefJoinPointResult("success.")
	//todo gas
	response.WithGas(10000, 10000)
	for _, aspect := range req {
		aspectId := aspect.AspectId
		runner, err := run.NewRunner(aspectId, aspect.Code)

		if err != nil {
			return response.WithErr(aspectId, err)
		} else {
			res, runErr := runner.JoinPoint(method, blockNumber, reqData)
			response.WithErr(aspectId, runErr).WithResponse(aspectId, res)
			runner.Return()
		}
		if hasErr, _ := response.HasErr(); hasErr {
			// short-circuit Aspect call
			return response
		}
	}
	return response
}
