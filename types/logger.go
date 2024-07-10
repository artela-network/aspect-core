package types

import (
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/protobuf/proto"
	"math/big"
)

// AspectLogger is used to collect execution traces from when an aspect gets triggered
type AspectLogger interface {
	CaptureAspectEnter(joinpoint JoinPointRunType, from, to, aspectId common.Address, input []byte, gas uint64, value *big.Int, execCtx proto.Message)
	CaptureAspectExit(joinpoint JoinPointRunType, result *AspectExecutionResult)
}

type NoOpsAspectLogger struct{}

func (n NoOpsAspectLogger) CaptureAspectEnter(joinpoint JoinPointRunType, from, to, aspectId common.Address, input []byte, gas uint64, value *big.Int, execCtx proto.Message) {
}

func (n NoOpsAspectLogger) CaptureAspectExit(joinpoint JoinPointRunType, result *AspectExecutionResult) {
}
