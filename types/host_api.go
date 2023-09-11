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
