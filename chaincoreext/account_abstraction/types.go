package account_abstraction

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/pkg/errors"
)

var (
	entrypointABI, _ = IEntryPointMetaData.GetAbi()
)

type ABIItem interface {
	Unpack(data []byte) (interface{}, error)
}

// ReturnInfo is the return value of the entry point.
type ReturnInfo struct {
	PreOpGas         *uint256.Int `json:"preOpGas"`
	Prefund          *uint256.Int `json:"prefund"`
	SigFailed        bool         `json:"sigFailed"`
	ValidAfter       uint64       `json:"validAfter"` // Using uint64 since Go doesn't have an uint48 type.
	ValidUntil       uint64       `json:"validUntil"` // Using uint64 to accommodate uint48.
	PaymasterContext []byte       `json:"paymasterContext"`
}

// StakeInfo is the stake information of an account.
type StakeInfo struct {
	Stake           *uint256.Int `json:"stake"`
	UnstakeDelaySec *uint256.Int `json:"unstakeDelaySec"`
}

// ValidationResult is the result of the validation.
type ValidationResult struct {
	ReturnInfo    *ReturnInfo `json:"returnInfo"`
	SenderInfo    *StakeInfo  `json:"senderInfo"`
	FactoryInfo   *StakeInfo  `json:"factoryInfo"`
	PaymasterInfo *StakeInfo  `json:"paymasterInfo"`
}

func DecodeValidationResult(data []byte) (*ValidationResult, error) {
	// failed to decode as ValidationResult, try to decode as FailedOp
	validationResultABI := entrypointABI.Errors["ValidationResult"]
	return decode[ValidationResult](&validationResultABI, data)
}

// FailedOp is the failed operation error returned by aa entrypoint.
type FailedOp struct {
	OpIndex *uint256.Int `json:"opIndex"`
	Reason  string       `json:"reason"`
}

func DecodeFailedOpError(data []byte) error {
	// failed to decode as ValidationResult, try to decode as FailedOp
	failedOpABI := entrypointABI.Errors["FailedOp"]
	failedOp, err := decode[FailedOp](&failedOpABI, data)
	if err != nil {
		// decode fail means it's not a FailedOp error
		return errors.New("unknown error")
	}

	// return fail reason
	return errors.New(failedOp.Reason)
}

// ExecutionResult is the result of the aa operation execution.
type ExecutionResult struct {
	PreOpGas      *uint256.Int `json:"preOpGas"`
	Paid          *uint256.Int `json:"paid"`
	ValidAfter    uint64       `json:"validAfter"` // Using uint64 since Go doesn't have a uint48 type.
	ValidUntil    uint64       `json:"validUntil"` // Using uint64 to accommodate uint48.
	TargetSuccess bool         `json:"targetSuccess"`
	TargetResult  []byte       `json:"targetResult"`
}

func DecodeExecutionResult(data []byte) (*ExecutionResult, error) {
	// failed to decode as ValidationResult, try to decode as FailedOp
	executionResultABI := entrypointABI.Errors["ExecutionResult"]
	return decode[ExecutionResult](&executionResultABI, data)
}

func (i *UserOperation) Hash() common.Hash {
	return common.Hash{}
}

func decode[V any](decodeErrorABI *abi.Error, data []byte) (*V, error) {
	res, err := decodeErrorABI.Unpack(data)
	if err != nil {
		return nil, err
	}

	return interfaceToStruct[V](res)
}

// interfaceToStruct converts interface to struct,
// use json here for convenience, optimize later
func interfaceToStruct[T any](input interface{}) (*T, error) {
	raw, _ := json.Marshal(input)
	var output T
	err := json.Unmarshal(raw, &output)
	return &output, err
}

func PackCallData(ops []*UserOperation, beneficiary common.Address) ([]byte, error) {
	return entrypointABI.Pack("handleOps", ops, beneficiary)
}
