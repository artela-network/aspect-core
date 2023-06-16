package djpm

import (
	"fmt"

	"github.com/artela-network/artelasdk/djpm/run"
	"github.com/artela-network/artelasdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var globalAspect *Aspect

type Aspect struct {
	GetBondAspects      func(int64, common.Address) ([]*types.AspectCode, error)
	GetBondBlockAspects func(int64) ([]*types.AspectCode, error)
	IsEthTx             func(tx sdk.Msg) bool
	ConvertEthTx        func(tx sdk.Msg) *ethtypes.Transaction
}

func NewAspect(
	getFunc func(int64, common.Address) ([]*types.AspectCode, error),
	getBlockAspectsFunc func(int64) ([]*types.AspectCode, error),
	checkTxFunc func(tx sdk.Msg) bool,
	convertTxFunc func(tx sdk.Msg) *ethtypes.Transaction) *Aspect {
	globalAspect = &Aspect{
		GetBondAspects:      getFunc,
		GetBondBlockAspects: getBlockAspectsFunc,
		IsEthTx:             checkTxFunc,
		ConvertEthTx:        convertTxFunc,
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
	result := &types.ResponseAspect{
		Success: true,
	}
	if req.Tx == nil || len(req.Tx.GetMsgs()) == 0 {
		return result
	}
	for _, msg := range req.Tx.GetMsgs() {
		ok := aspect.IsEthTx(msg)
		if !ok {
			// ignore cosmos tx
			continue
		}
		ethTx := aspect.ConvertEthTx(msg)
		if ethTx.To() == nil || types.IsAspectContract(ethTx.To()) {
			continue
		}

		var signer ethtypes.Signer
		if ethTx.Protected() {
			signer = ethtypes.LatestSignerForChainID(ethTx.ChainId())
		} else {
			signer = ethtypes.HomesteadSigner{}
		}
		from, _ := ethtypes.Sender(signer, ethTx)
		hash := ethTx.Hash()
		txAspect := &types.RequestEthMsgAspect{
			BlockHeight: req.BlockHeight,
			TxHash:      &hash,
			TxIndex:     uint(req.TxIndex),
			To:          ethTx.To(),
			From:        from,
			Nonce:       ethTx.Nonce(),
			GasLimit:    ethTx.Gas(),
			GasPrice:    ethTx.GasPrice(),
			GasFeeCap:   ethTx.GasFeeCap(),
			GasTipCap:   ethTx.GasTipCap(),
			Value:       ethTx.Value(),
			TxType:      0,
			TxData:      ethTx.Data(),
			AccessList:  ethTx.AccessList(),
			ChainId:     req.ChainId,
		}

		result = aspect.execAspectByEthMsg(methodName, txAspect)
		if result.HasErr() {
			return result
		}
	}
	return result

}

func (aspect Aspect) execAspectBlock(methodName string, req *types.RequestBlockAspect) *types.ResponseAspect {
	response := &types.ResponseAspect{Success: true}

	if req == nil {
		return response
	}
	boundAspects, err := aspect.GetBondBlockAspects(req.BlockHeight)
	// load aspects
	if err != nil || len(boundAspects) == 0 {
		return response
	}
	aspectInput := &types.AspectInput{
		BlockHeight: req.BlockHeight,
	}
	//todo gas
	response.WithGas(10000, 10000)
	// run aspects on received transaction
	for _, aspect := range boundAspects {
		res := &types.AspectOutput{}
		runner, err := run.NewRunner(aspect.AspectId, aspect.Code)
		if err != nil {
			response.WithErr(err)
		} else {
			res, err = runner.JoinPoint(methodName, aspectInput)
			response.WithErr(err).WithAspectOutput(res)

		}
		id := aspect.AspectId
		response.WithAspectId(id)
		if response.HasErr() {
			// short-circuit Aspect call
			return response
		}
	}
	return response
}

func (aspect Aspect) execAspectByEthMsg(methodName string, req *types.RequestEthMsgAspect) *types.ResponseAspect {
	response := &types.ResponseAspect{
		Success: true,
	}
	if req.To == nil || types.IsAspectContract(req.To) {
		return response
	}
	boundAspects, err := aspect.GetBondAspects(req.BlockHeight, *req.To)
	// load aspects
	if err != nil || len(boundAspects) == 0 {
		return response
	}

	transaction := req.ToAspTx()
	aspectInput := &types.AspectInput{
		BlockHeight: req.BlockHeight,
		Tx:          transaction,
	}

	return runAspect(methodName, boundAspects, aspectInput)
}

func runAspect(methodName string, boundAspects []*types.AspectCode, aspectInput *types.AspectInput) *types.ResponseAspect {
	errCode := int32(0)
	revertMsg := ""
	callback := func(code int32, msg string) {
		errCode = code
		revertMsg = msg
	}

	response := &types.ResponseAspect{
		Success: true,
	}
	response.WithGas(10000, 10000)
	// run aspects on received transaction
	for _, aspect := range boundAspects {
		res := &types.AspectOutput{}
		response.WithAspectId(aspect.AspectId)
		runner, err := run.NewRunnerWithCallBack(aspect.AspectId, aspect.Code, run.CallBackRevertFunc(callback))
		if err != nil {
			switch errCode {
			case 0:
				fmt.Println("run error", revertMsg)
			case 1:
				fmt.Println("transaction reverted", revertMsg)
				// TOOD revert tx.
			default:
				fmt.Println(revertMsg)
			}
			response.WithErr(err)
		} else {
			res, err = runner.JoinPoint(methodName, aspectInput)
			response.WithErr(err).WithAspectOutput(res)
		}
		if response.HasErr() {
			// short-circuit Aspect call
			return response
		}
	}
	return response
}

func (aspect Aspect) OnTxReceive(req *types.RequestSdkTxAspect) *types.ResponseAspect {
	tx := aspect.execAspectBySdkTx(types.ON_TX_RECEIVE_METHOD, req)
	return tx
}

func (aspect Aspect) OnBlockInitialize(req *types.RequestBlockAspect) *types.ResponseAspect {
	return aspect.execAspectBlock(types.ON_BLOCK_INITIALIZE_METHOD, req)
}
func (aspect Aspect) OnTxVerify(req *types.RequestEthMsgAspect) *types.ResponseAspect {
	return aspect.execAspectByEthMsg(types.ON_TX_VERIFY_METHOD, req)

}

func (aspect Aspect) OnAccountVerify(req *types.RequestEthMsgAspect) *types.ResponseAspect {
	return aspect.execAspectByEthMsg(types.ON_ACCOUNT_VERIFY_METHOD, req)
}
func (aspect Aspect) OnGasPayment(req *types.RequestEthMsgAspect) *types.ResponseAspect {
	return nil

}
func (aspect Aspect) PreTxExecute(req *types.RequestEthMsgAspect) *types.ResponseAspect {
	return aspect.execAspectByEthMsg(types.PRE_TX_EXECUTE_METHOD, req)

}
func (aspect Aspect) PreContractCall(req *types.RequestEthMsgAspect) *types.ResponseAspect {
	return aspect.execAspectByEthMsg(types.PRE_CONTRACT_CALL_METHOD, req)

}
func (aspect Aspect) PostContractCall(req *types.RequestEthMsgAspect) *types.ResponseAspect {
	return aspect.execAspectByEthMsg(types.POST_CONTRACT_CALL_METHOD, req)

}
func (aspect Aspect) PostTxExecute(req *types.RequestEthMsgAspect) *types.ResponseAspect {
	return aspect.execAspectByEthMsg(types.POST_TX_EXECUTE_METHOD, req)

}
func (aspect Aspect) OnTxCommit(req *types.RequestEthMsgAspect) *types.ResponseAspect {
	return aspect.execAspectByEthMsg(types.ON_TX_COMMIT_METHOD, req)

}

func (aspect Aspect) OnBlockFinalize(req *types.RequestBlockAspect) *types.ResponseAspect {
	return aspect.execAspectBlock(types.ON_BLOCK_FINALIZE_METHOD, req)
}
