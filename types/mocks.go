package types

import types2 "github.com/artela-network/aspect-runtime/types"

type NoOpsLogger struct {
}

func (n NoOpsLogger) Debug(_ string, _ ...interface{}) {
}

func (n NoOpsLogger) Info(_ string, _ ...interface{}) {
}

func (n NoOpsLogger) Error(_ string, _ ...interface{}) {
}

func (n NoOpsLogger) With(_ ...interface{}) types2.Logger {
	return n
}
