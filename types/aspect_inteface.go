package types

type JoinPointAdviceI interface {
	FilterTx(request *EthTxAspect) *JoinPointResult
	VerifyTx(request *EthTxAspect) *JoinPointResult
	VerifyAccount(request *EthTxAspect) *JoinPointResult
	GetPayMaster(request *EthTxAspect) *JoinPointResult
	PreTxExecute(request *EthTxAspect) *JoinPointResult
	PreContractCall(request *EthTxAspect) *JoinPointResult
	PostContractCall(request *EthTxAspect) *JoinPointResult
	PostTxExecute(request *EthTxAspect) *JoinPointResult
	PostTxCommit(request *EthTxAspect) *JoinPointResult

	OnBlockInitialize(request *EthBlockAspect) *JoinPointResult
	OnBlockFinalize(request *EthBlockAspect) *JoinPointResult
}
