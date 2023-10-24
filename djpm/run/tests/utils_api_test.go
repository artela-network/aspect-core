package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/artela-network/aspect-core/djpm/run"
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
