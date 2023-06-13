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
func TestJoinPoint(t *testing.T) {
	aspectType.GetHostApiHook = func() (aspectType.HostApi, error) {
		return nil, errors.New("not init")
	}

	cwd, _ := os.Getwd()
	raw, _ := os.ReadFile(path.Join(cwd, "/Users/admin/mytech/go-work/src/github.com/artela-network/aspect-example/basic/wasm/build/release.wasm"))

	name := "onTxReceive"
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
	}
	runner, err := NewRunner("", raw)
	require.Equal(t, nil, err)
	output, err := runner.JoinPoint(name, input)
	require.Equal(t, nil, err)
	require.Equal(t, true, output.Success)

}

// Run "scripts/build-wasm.sh" in project root, before run this test.
func TestIsOwner(t *testing.T) {
	aspectType.GetHostApiHook = func() (aspectType.HostApi, error) {
		return nil, errors.New("not init")
	}

	//cwd, _ := os.Getwd()
	raw, _ := os.ReadFile("/Users/admin/mytech/go-work/src/github.com/artela-network/aspect-example/basic/wasm/build/release.wasm")

	runner, err := NewRunner("", raw)
	require.Equal(t, nil, err)
	ret, err := runner.IsOwner("hello")
	require.Equal(t, nil, err)
	require.Equal(t, false, ret)
}

// Run "scripts/build-wasm.sh" in project root, before run this test.
func TestOnContractBinding(t *testing.T) {
	aspectType.GetHostApiHook = func() (aspectType.HostApi, error) {
		return nil, errors.New("not init")
	}

	cwd, _ := os.Getwd()
	raw, _ := os.ReadFile(path.Join(cwd, "/Users/admin/mytech/go-work/src/github.com/artela-network/aspect-example/basic/wasm/build/release.wasm"))

	runner, err := NewRunner("", raw)
	require.Equal(t, nil, err)
	ret, err := runner.OnContractBinding("0x0000000000000000000000000000000000000001")
	require.Equal(t, nil, err)
	require.Equal(t, false, ret)
}
