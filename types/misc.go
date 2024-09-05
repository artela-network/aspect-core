package types

type AspectCode struct {
	AspectId string
	Version  uint64
	Priority int8
	Code     []byte
}

type RevertScope uint8

const (
	NotRevert RevertScope = iota
	RevertCall
	RevertTx
)

type AspectExecutionResult struct {
	Gas    uint64
	Err    error
	Ret    []byte
	Revert RevertScope
}
