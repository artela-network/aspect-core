package account_abstraction

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

var (
	// EntryPointContract defines the address of the built-in AA entry point contract.
	EntryPointContract = common.HexToAddress("0x000000000000000000000000000000000000AAEC")
)

var (
	entrypointABI, _ = IEntryPointMetaData.GetAbi()
)

var (
	userOperationPackedJJSONABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hashInitCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hashCallData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hashPaymasterAndData\",\"type\":\"bytes32\"}],\"name\":\"pack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"entrypoint\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"packHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	userOperationPackedABI      = parseABI(userOperationPackedJJSONABI)
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
	// failed to DecodeError as ValidationResult, try to DecodeError as FailedOp
	validationResultABI := entrypointABI.Errors["ValidationResult"]
	return DecodeError[ValidationResult](&validationResultABI, data)
}

// FailedOp is the failed operation error returned by aa entrypoint.
type FailedOp struct {
	OpIndex *uint256.Int `json:"opIndex"`
	Reason  string       `json:"reason"`
}

func DecodeFailedOpError(data []byte) error {
	// failed to DecodeError as ValidationResult, try to DecodeError as FailedOp
	failedOpABI := entrypointABI.Errors["FailedOp"]
	failedOp, err := DecodeError[FailedOp](&failedOpABI, data)
	if err != nil {
		// DecodeError fail means it's not a FailedOp error
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
	// failed to DecodeError as ValidationResult, try to DecodeError as FailedOp
	executionResultABI := entrypointABI.Errors["ExecutionResult"]
	return DecodeError[ExecutionResult](&executionResultABI, data)
}

func (i *UserOperation) Hash(chainId *big.Int) common.Hash {
	initCodeHash := crypto.Keccak256Hash(i.InitCode)
	callDataHash := crypto.Keccak256Hash(i.CallData)
	paymasterAndDataHash := crypto.Keccak256Hash(i.PaymasterAndData)

	packed, _ := userOperationPackedABI.Methods["pack"].Inputs.Pack(
		i.Sender,
		i.Nonce,
		initCodeHash,
		callDataHash,
		i.CallGasLimit,
		i.VerificationGasLimit,
		i.PreVerificationGas,
		i.MaxFeePerGas,
		i.MaxPriorityFeePerGas,
		paymasterAndDataHash)

	userOpHash := crypto.Keccak256Hash(packed)
	hashPacked, _ := userOperationPackedABI.Methods["packHash"].Inputs.Pack(
		userOpHash,
		EntryPointContract,
		chainId,
	)

	return crypto.Keccak256Hash(hashPacked)
}

func DecodeResponse(methodName string, data []byte) ([]interface{}, error) {
	method, ok := entrypointABI.Methods[methodName]
	if !ok {
		return nil, errors.New("method not found")
	}

	return method.Outputs.Unpack(data)
}

func DecodeError[V any](decodeErrorABI *abi.Error, data []byte) (*V, error) {
	res, err := decodeErrorABI.Unpack(data)
	if err != nil {
		return nil, err
	}

	if casted, ok := res.(*V); ok {
		return casted, nil
	}

	return InterfaceToStruct[V](res)
}

// InterfaceToStruct converts interface to struct,
// use json here for convenience, optimize later
func InterfaceToStruct[T any](input interface{}) (*T, error) {
	raw, _ := json.Marshal(input)
	var output T
	err := json.Unmarshal(raw, &output)
	return &output, err
}

func PackCallData(ops []*UserOperation, beneficiary common.Address) ([]byte, error) {
	return entrypointABI.Pack("handleOps", ops, beneficiary)
}

func parseABI(json string) *abi.ABI {
	parsedABI, _ := abi.JSON(strings.NewReader(json))
	return &parsedABI
}
