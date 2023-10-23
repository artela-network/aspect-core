package tests

import (
	"testing"

	"github.com/artela-network/artelasdk/djpm/run"
	"github.com/stretchr/testify/require"
)

func TestSlog(t *testing.T) {
	raw, _ := GetTestTarget("utilapi-test")
	runner, err := run.NewRunner("", raw)
	require.Equal(t, nil, err)
	ret, err := runner.ExecFunc("TestSlog", 99, 99, nil)
	require.Equal(t, nil, err)
	require.Equal(t, true, ret)
	defer runner.Return()
}
