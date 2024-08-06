package context

var PostTxCtxKeys = []interface{}{
	IsCall,

	BlockHeaderParentHash,
	BlockHeaderMiner,
	BlockHeaderNumber,
	BlockHeaderTimestamp,

	EnvExtraEIPs,
	EnvEnableCreate,
	EnvEnableCall,
	EnvAllowUnprotectedTxs,
	EnvChainChainId,
	EnvChainHomesteadBlock,
	EnvChainDaoForkBlock,
	EnvChainDaoForkSupport,
	EnvChainEip150Block,
	EnvChainEip155Block,
	EnvChainEip158Block,
	EnvChainByzantiumBlock,
	EnvChainConstantinopleBlock,
	EnvChainPetersburgBlock,
	EnvChainIstanbulBlock,
	EnvChainMuirGlacierBlock,
	EnvChainBerlinBlock,
	EnvChainLondonBlock,
	EnvChainArrowGlacierBlock,
	EnvChainGrayGlacierBlock,
	EnvChainMergeNetSplitBlock,
	EnvChainShanghaiTime,
	EnvChainCancunTime,
	EnvChainPragueTime,
	EnvConsensusParamsBlockMaxGas,
	EnvConsensusParamsBlockMaxBytes,
	EnvConsensusParamsEvidenceMaxAgeDuration,
	EnvConsensusParamsEvidenceMaxAgeNumBlocks,
	EnvConsensusParamsEvidenceMaxBytes,
	EnvConsensusParamsValidatorPubKeyTypes,
	EnvConsensusParamsAppVersion,

	TxType,
	TxChainId,
	TxAccessList,
	TxNonce,
	TxGasPrice,
	TxGas,
	TxGasTipCap,
	TxGasFeeCap,
	TxTo,
	TxValue,
	TxData,
	TxBytes,
	TxHash,
	TxUnsignedBytes,
	TxUnsignedHash,
	TxSigV,
	TxSigR,
	TxSigS,
	TxFrom,
	TxIndex,

	AspectId,
	AspectVersion,

	ReceiptStatus,
	ReceiptLogs,
	ReceiptGasUsed,
	ReceiptCumulativeGasUsed,
	ReceiptBloom,
}
