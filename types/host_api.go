package types

var GetHostApiHook func() (HostApi, error)

type HostApi interface {
	// LocalCall calls EthCall
	//	LocalCall(req *evmtypes.EthCallRequest) (*evmtypes.MsgEthereumTxResponse, error)

	StateAt(key string) interface{}

	// TBD, if we need to return the artelamint blocks
	// LastBlock() (*coretypes.ResultBlock, error)
	// CurrentBlock() (*coretypes.ResultBlock, error)

	// LastBlock returns last ethereum block
	LastBlock() (*EthBlock, error)

	// CurrentBlock returns ethereum block built by the packing block,
	// this should only be called when a new block is generating
	CurrentBlock() (*EthBlock, error)

	// GetAppState returns the value stored in appState.
	// appState is aspect related, visit them with aspect hash
	GetAppState(hash []byte, key string) (interface{}, error)

	// GetGlobalState returns the value stored in globalState
	// globalState is shared by aspects
	GetGlobalState(key string) (interface{}, error)

	// GetProperty returns the configuration of aspect
	GetProperty(aspectID string, key string) (string, error)

	// GetStateChanges returns the state changes of fields
	GetStateChanges(addr string, variable string, key []byte) *StateChanges
	AddInherent()
	ScheduleTx(sch *Schedule) bool
	DropTx()
}
