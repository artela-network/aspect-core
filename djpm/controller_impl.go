package djpm

import (
	"math/big"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/artela-network/artelasdk/djpm/run"
	"github.com/artela-network/artelasdk/types"
)

var globalAspect *Aspect

type Aspect struct {
	GetBondAspects func(int64, *ethtypes.Transaction) ([]*types.AspectCode, error)
	IsEthTx        func(tx sdk.Msg) bool
	ConvertEthTx   func(tx sdk.Msg) *ethtypes.Transaction
}

func NewAspect(
	getFunc func(int64, *ethtypes.Transaction) ([]*types.AspectCode, error),
	checkTxFunc func(tx sdk.Msg) bool,
	convertTxFunc func(tx sdk.Msg) *ethtypes.Transaction) *Aspect {
	globalAspect = &Aspect{
		GetBondAspects: getFunc,
		IsEthTx:        checkTxFunc,
		ConvertEthTx:   convertTxFunc,
	}
	return globalAspect
}

func AspectInstance() *Aspect {
	if globalAspect == nil {
		panic("aspcect instance not init,please exec NewAspect() first ")
	}
	return globalAspect
}

func (aspect Aspect) execAspectBySdkTx(methodName string, req *types.RequestSdkTxAspect) *types.ResponseAspect {

	if req.Tx == nil || len(req.Tx.GetMsgs()) == 0 {
		return nil
	}
	result := &types.ResponseAspect{}
	for _, msg := range req.Tx.GetMsgs() {
		ok := aspect.IsEthTx(msg)
		if !ok {
			// ignore cosmos tx
			continue
		}
		ethTx := aspect.ConvertEthTx(msg)
		txAspect := types.RequestEthTxAspect{
			Tx:          ethTx,
			Context:     req.Context,
			BlockHeight: req.BlockHeight,
			BlockHash:   req.BlockHash,
			TxIndex:     req.TxIndex,
			BaseFee:     req.BaseFee,
			ChainId:     req.ChainId,
		}
		out := aspect.execAspectByEthTx(methodName, &txAspect)
		result.Merge(out)
	}
	return result

}

func (aspect Aspect) execAspectByEthTx(methodName string, req *types.RequestEthTxAspect) *types.ResponseAspect {
	if req.Tx == nil {
		return nil
	}
	to := req.Tx.To()
	if to == nil || strings.EqualFold(types.ARTELA_ADDR, to.Hex()) {
		// ignore contract deployment transaction & aspect op txs
		return nil
	}
	boundAspects, err := aspect.GetBondAspects(req.BlockHeight, req.Tx)
	// load aspects
	if err != nil || len(boundAspects) == 0 {
		return nil
	}

	chain, _ := strconv.ParseInt(req.ChainId, 10, 64)
	transaction, newErr := types.NewTx(req.Tx, common.HexToHash(req.BlockHash), req.BlockHeight, req.TxIndex, big.NewInt(req.BaseFee), big.NewInt(chain))
	if newErr != nil {
		return nil
	}
	aspectInput := &types.AspectInput{
		BlockHeight: req.BlockHeight,
		Tx:          transaction,
		Context:     req.Context,
	}
	response := &types.ResponseAspect{}
	txHash := common.BytesToHash(transaction.Hash).String()
	// run aspects on received transaction
	for _, aspect := range boundAspects {
		var res *types.AspectOutput
		runner, err := run.NewRunner(aspect.AspectId, aspect.Code)
		if err != nil {
			res = &types.AspectOutput{
				Success: false,
				Message: err.Error(),
				Context: nil,
			}
		} else {
			res, err = runner.JoinPoint(methodName, aspectInput)
			if err != nil {
				res = &types.AspectOutput{
					Success: false,
					Message: err.Error(),
					Context: nil,
				}
			}
		}
		id := aspect.AspectId

		response.With(txHash, id, res)
	}
	return response
}
func (aspect Aspect) OnTxReceive(req *types.RequestSdkTxAspect) *types.ResponseAspect {
	return aspect.execAspectBySdkTx(types.ON_TX_RECEIVE_METHOD, req)
}

func (aspect Aspect) OnBlockInitialize(req *types.RequestBlockAspect) *types.ResponseBlockAspect {
	return nil

}
func (aspect Aspect) OnTxVerify(req *types.RequestEthTxAspect) *types.ResponseAspect {

	return aspect.execAspectByEthTx(types.ON_TX_VERIFY_METHOD, req)

}

func (aspect Aspect) OnAccountVerify(req *types.RequestEthTxAspect) *types.ResponseAspect {
	return aspect.execAspectByEthTx(types.ON_ACCOUNT_VERIFY_METHOD, req)
}
func (aspect Aspect) OnGasPayment(req *types.RequestEthTxAspect) *types.ResponseAspect {
	return nil

}
func (aspect Aspect) PreTxExecute(req *types.RequestSdkTxAspect) *types.ResponseAspect {
	return aspect.execAspectBySdkTx(types.PRE_TX_EXECUTE_METHOD, req)

}
func (aspect Aspect) PreContractCall(req *types.RequestEthTxAspect) *types.ResponseAspect {
	return aspect.execAspectByEthTx(types.PRE_CONTRACT_CALL_METHOD, req)

}
func (aspect Aspect) PostContractCall(req *types.RequestEthTxAspect) *types.ResponseAspect {
	return aspect.execAspectByEthTx(types.POST_CONTRACT_CALL_METHOD, req)

}
func (aspect Aspect) PostTxExecute(req *types.RequestSdkTxAspect) *types.ResponseAspect {
	return aspect.execAspectBySdkTx(types.POST_TX_EXECUTE_METHOD, req)

}
func (aspect Aspect) OnTxCommit(req *types.RequestEthTxAspect) *types.ResponseAspect {
	return aspect.execAspectByEthTx(types.ON_TX_COMMIT_METHOD, req)

}

func (aspect Aspect) OnBlockFinalize(req *types.RequestBlockAspect) *types.ResponseBlockAspect {
	return nil

}
