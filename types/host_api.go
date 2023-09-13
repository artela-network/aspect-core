package types

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"reflect"
)

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
