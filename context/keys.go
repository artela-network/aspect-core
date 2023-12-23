package context

// IsCall Condition Context Related Keys
const (
	IsCall = "isCall"
)

// Block Context Related Keys
const (
	BlockHeaderParentHash       = "block.header.parentHash"
	BlockHeaderMiner            = "block.header.miner"
	BlockHeaderTransactionsRoot = "block.header.transactionsRoot"
	BlockHeaderNumber           = "block.header.number"
	BlockHeaderTimestamp        = "block.header.timestamp"
)

// Env Context Related Keys
const (
	EnvExtraEIPs                              = "env.extraEIPs"
	EnvEnableCreate                           = "env.enableCreate"
	EnvEnableCall                             = "env.enableCall"
	EnvAllowUnprotectedTxs                    = "env.allowUnprotectedTxs"
	EnvChainChainId                           = "env.chain.chainId"
	EnvChainHomesteadBlock                    = "env.chain.homesteadBlock"
	EnvChainDaoForkBlock                      = "env.chain.daoForkBlock"
	EnvChainDaoForkSupport                    = "env.chain.daoForkSupport"
	EnvChainEip150Block                       = "env.chain.eip150Block"
	EnvChainEip155Block                       = "env.chain.eip155Block"
	EnvChainEip158Block                       = "env.chain.eip158Block"
	EnvChainByzantiumBlock                    = "env.chain.byzantiumBlock"
	EnvChainConstantinopleBlock               = "env.chain.constantinopleBlock"
	EnvChainPetersburgBlock                   = "env.chain.petersburgBlock"
	EnvChainIstanbulBlock                     = "env.chain.istanbulBlock"
	EnvChainMuirGlacierBlock                  = "env.chain.muirGlacierBlock"
	EnvChainBerlinBlock                       = "env.chain.berlinBlock"
	EnvChainLondonBlock                       = "env.chain.londonBlock"
	EnvChainArrowGlacierBlock                 = "env.chain.arrowGlacierBlock"
	EnvChainGrayGlacierBlock                  = "env.chain.grayGlacierBlock"
	EnvChainMergeNetSplitBlock                = "env.chain.mergeNetSplitBlock"
	EnvChainShanghaiTime                      = "env.chain.shanghaiTime"
	EnvChainCancunTime                        = "env.chain.cancunTime"
	EnvChainPragueTime                        = "env.chain.pragueTime"
	EnvConsensusParamsBlockMaxGas             = "env.consensusParams.block.maxGas"
	EnvConsensusParamsBlockMaxBytes           = "env.consensusParams.block.maxBytes"
	EnvConsensusParamsEvidenceMaxAgeDuration  = "env.consensusParams.evidence.maxAgeDuration"
	EnvConsensusParamsEvidenceMaxAgeNumBlocks = "env.consensusParams.evidence.maxAgeNumBlocks"
	EnvConsensusParamsEvidenceMaxBytes        = "env.consensusParams.evidence.maxBytes"
	EnvConsensusParamsValidatorPubKeyTypes    = "env.consensusParams.validator.pubKeyTypes"
	EnvConsensusParamsAppVersion              = "env.consensusParams.appVersion"
)

// Tx Context Related Keys
const (
	TxType          = "tx.type"
	TxChainId       = "tx.chainId"
	TxAccessList    = "tx.accessList"
	TxNonce         = "tx.nonce"
	TxGasPrice      = "tx.gasPrice"
	TxGas           = "tx.gas"
	TxGasTipCap     = "tx.gasTipCap"
	TxGasFeeCap     = "tx.gasFeeCap"
	TxTo            = "tx.to"
	TxValue         = "tx.value"
	TxData          = "tx.data"
	TxBytes         = "tx.bytes"
	TxHash          = "tx.hash"
	TxUnsignedBytes = "tx.unsigned.bytes"
	TxUnsignedHash  = "tx.unsigned.hash"
	TxSigV          = "tx.sig.v"
	TxSigR          = "tx.sig.r"
	TxSigS          = "tx.sig.s"
	TxFrom          = "tx.from"
	TxIndex         = "tx.index"
)

// Aspect Context Related Keys
const (
	AspectId      = "aspect.id"
	AspectVersion = "aspect.version"
)

// Msg Context Related Keys
const (
	MsgFrom          = "msg.from"
	MsgTo            = "msg.to"
	MsgValue         = "msg.value"
	MsgGas           = "msg.gas"
	MsgInput         = "msg.input"
	MsgIndex         = "msg.index"
	MsgResultRet     = "msg.result.ret"
	MsgResultGasUsed = "msg.result.gasUsed"
	MsgResultError   = "msg.result.error"
)

// Receipt Context Related Keys
const (
	ReceiptStatus            = "receipt.status"
	ReceiptLogs              = "receipt.logs"
	ReceiptGasUsed           = "receipt.gasUsed"
	ReceiptCumulativeGasUsed = "receipt.cumulativeGasUsed"
	ReceiptBloom             = "receipt.bloom"
)
