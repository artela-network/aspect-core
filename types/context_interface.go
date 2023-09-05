package types

type AspectRuntimeContextI interface {
	BlockContextI
	EthTxContextI
	EnvConfigI
	AspectContextI
}

type EthTxContextI interface {
	TxContent() (*EthTransaction, error)
	GetStateChanges(addr string, variable string, key []byte) (*EthStateChanges, error)
	QueryCallStack(query *EthCallStackQuery) (*EthCallStacks, error)
	GetReceipt() (*EthReceipt, error)
	GasMeter() (*CosmosGasMeter, error)
}

type AspectContextI interface {
	Add(txHash string, aspectID string, key, value string) error
	Get(txHash string, aspectID string, key string) (string, error)
	Remove(txHash string, aspectID string, key string) error
	Clear(txHash string) error
}

type BlockContextI interface {
	GetBlockHeader() (*EthBlockHeader, error)
	GetBlockBody(query *EthTxQuery) ([]*EthTransaction, error)
	GetLastCommitInfo() (*LastCommitInfo, error)
	GasMeter() (*CosmosGasMeter, error)
	GetEvidences() ([]*Evidence, error)
	GetBlockId() ([]*BlockID, error)
}

type EnvConfigI interface {
	GetChainConfig() (*ChainConfig, error)
	GetEvmParams() (*EvmParams, error)
	GetConsParams() (*ConsParams, error)
	GetEnvContent() (*EnvContent, error)
}
