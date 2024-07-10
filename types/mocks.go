package types

type NoOpsLogger struct {
}

func (n NoOpsLogger) Debug(_ string, _ ...interface{}) {
}

func (n NoOpsLogger) Info(_ string, _ ...interface{}) {
}

func (n NoOpsLogger) Error(_ string, _ ...interface{}) {
}
