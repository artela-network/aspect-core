package run

import (
	"errors"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"

	aspectType "github.com/artela-network/artelasdk/types"
)

// Run "scripts/build-wasm.sh" in project root, before run this test.
func TestRunAspect(t *testing.T) {
	aspectType.GetHostApiHook = func() (aspectType.HostApi, error) {
		return nil, errors.New("not init")
	}

	cwd, _ := os.Getwd()
	raw, _ := os.ReadFile(path.Join(cwd, "./testdata/build/release.wasm"))

	method := "onTxReceive"
	input := &aspectType.AspectInput{
		BlockHeight: 999,
		Tx: &aspectType.AspTransaction{
			ChainId:          "9000-artela",
			Nonce:            123456789,
			GasTipCap:        "GasTipCap-value",
			GasFeeCap:        "GasFeeCap-value",
			GasLimit:         1000000000,
			GasPrice:         998,
			To:               "0x1c0e4b5d5f50fe65adc4cd658cd88ae0dfdbb3ba",
			Value:            9998,
			Input:            []byte{0x1, 0x2, 0x3, 0x4},
			AccessList:       []*aspectType.AspAccessTuple{},
			BlockHash:        []byte{},
			BlockNumber:      0,
			From:             "",
			Hash:             []byte{},
			TransactionIndex: 0,
			Type:             0,
			V:                []byte{},
			R:                []byte{},
			S:                []byte{},
		},
		Context: map[string]string{
			"111": "abc",
			"222": "def",
			"333": "ghi",
		},
	}
	output, err := RunAspect(raw, method, input)
	require.Equal(t, nil, err)
	require.Equal(t, true, output.Success)

	// verify the context
	ctx := output.Context
	require.Equal(t, 6, len(ctx))
	expected := map[string]string{
		"111":          "abc",
		"222":          "def",
		"333":          "ghi",
		"k1":           "v1",
		"k2":           "v2",
		"lastBlockNum": "not found",
	}
	require.Equal(t, expected["111"], ctx["111"])
	require.Equal(t, expected["222"], ctx["222"])
	require.Equal(t, expected["333"], ctx["333"])
	require.Equal(t, expected["k1"], ctx["k1"])
	require.Equal(t, expected["k2"], ctx["k2"])
	require.Equal(t, expected["lastBlockNum"], ctx["lastBlockNum"])
}
