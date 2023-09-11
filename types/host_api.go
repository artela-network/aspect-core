package types

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"reflect"
)

func NewStringResponse(condition bool, data string, errMsg string) *StringDataResponse {
	message := "success"
	if !condition {
		message = errMsg
	}
	return &StringDataResponse{
		Result: &RunResult{
			Success: condition,
			Message: message,
		},
		Data: data,
	}
}

func NewIntDataResponse(condition bool, data int64, errMsg string) *IntDataResponse {
	message := "success"
	if condition {
		message = errMsg
	}
	return &IntDataResponse{
		Result: &RunResult{
			Success: condition,
			Message: message,
		},
		Data: data,
	}
}

func NewContextQueryResponse(condition bool, errMsg string) *ContextQueryResponse {
	message := "success"
	if condition {
		message = errMsg
	}
	return &ContextQueryResponse{
		Result: &RunResult{
			Success: condition,
			Message: message,
		},
	}
}
func (c *ContextQueryResponse) SetData(message proto.Message) {
	if message == nil {
		return
	}
	anyData, _ := anypb.New(message)
	c.Data = anyData
	messageType := reflect.TypeOf(message)
	c.DataMessageType = messageType.Name()
}

type SetDataType string

const (
	ASPECT_STATE       SetDataType = "aspect.state"
	CURRENT_TX_CONTEXT SetDataType = "current_tx.context"
)

type AspectRuntimeContextHostApiI interface {
	Select(point PointCut, dataType AspectDataAccess, aspectId common.Address, queryCriteria []interface{}) (interface{}, error)
	Set(point PointCut, data SetDataType, aspectId common.Address, key string, value any) error
	Remove(point PointCut, data SetDataType, aspectId common.Address, key string) error
}

type EvmHostApiI interface {
	LocalCall(blockNumber int64, transaction *AspTransaction) (*AspTxResponse, error)
	InnerCall(transaction *AspTransaction) (*AspTxResponse, error)
}
type ScheduleHostApiI interface {
	SubmitScheduleTx(sch *Schedule) error
}
type StateDbHostApi interface {
	GetBalance(addr common.Address) (*big.Int, error)
	GetState(addr common.Address, hash common.Hash) (common.Hash, error)
	GetRefund() (uint64, error)
	GetCodeHash(addr common.Address) (*common.Hash, error)
	GetNonce(addr common.Address) (uint64, error)
}
type AspectStateApiI interface {
	GetAspectState(aspectID string, key string) (string, error)
	SetAspectState(aspectID string, key, value string) error
	RemoveAspectState(aspectID string, key string) error
	GetProperty(aspectID string, key string) (string, error)
}

type CryptoHostApiI interface {
	Hash(hasher int32, data []byte) ([]byte, error)
}
type AbiSystemCallI interface {
	DecodeParams(t string, data []byte) ([]interface{}, error)
	EncodeParams(t string, values ...interface{}) ([]byte, error)
}
