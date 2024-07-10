package tests

import (
	"context"
	"github.com/artela-network/aspect-core/types"
	"github.com/ethereum/go-ethereum/common"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/artela-network/aspect-core/djpm/run"
)

func TestSlog(t *testing.T) {
	raw, _ := GetTestTarget("utilapi-test")
	runner, err := run.NewRunner(context.Background(), &types.NoOpsLogger{}, "", 0, raw, false)
	require.Equal(t, nil, err)
	ret, _, err := runner.ExecFunc("TestSlog", 99, 99, common.Address{})
	require.Equal(t, nil, err)
	require.Equal(t, true, ret)
	defer runner.Return()
}
