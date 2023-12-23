package run

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"

	aspectType "github.com/artela-network/aspect-core/types"
)

// Run "scripts/build-wasm.sh" in project root, before run this test.
func TestJoinPoint(t *testing.T) {
	raw, _ := os.ReadFile("/Users/admin/mytech/go-work/src/github.com/artela-network/aspect-tooling/packages/libs-test/build/release.wasm")

	name := aspectType.FILTER_TX
	input := &aspectType.EthTxAspect{
		Tx: &aspectType.EthTransaction{
			ChainId:          "9000-artela",
			Nonce:            123456789,
			GasTipCap:        "GasTipCap-value",
			GasFeeCap:        "GasFeeCap-value",
			Gas:              1000000000,
			GasPrice:         "998",
			To:               "0x1c0e4b5d5f50fe65adc4cd658cd88ae0dfdbb3ba",
			Value:            "9998",
			Input:            []byte{0x1, 0x2, 0x3, 0x4},
			AccessList:       []*aspectType.EthAccessTuple{},
			BlockHash:        []byte{},
			BlockNumber:      999,
			From:             "",
			Hash:             []byte{},
			TransactionIndex: 0,
			Type:             0,
			V:                []byte{},
			R:                []byte{},
			S:                []byte{},
		},
	}
	runner, err := NewRunner(context.Background(), "", raw)
	require.Equal(t, nil, err)
	address := common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4")

	output, err := runner.JoinPoint(name, 99, 999, &address, input)
	require.Equal(t, nil, err)
	defer runner.Return()

	data := output.Data
	strData := &aspectType.StringData{}
	err2 := data.UnmarshalTo(strData)
	if err2 != nil {
		return
	}
	fmt.Println(strData.GetData() + "------")

	require.Equal(t, true, output.Result.Success)
}

// Run "scripts/build-wasm.sh" in project root, before run this test.
func TestIsOwner(t *testing.T) {
	// cwd, _ := os.Getwd()
	raw, err := os.ReadFile("/Users/admin/mytech/go-work/src/github.com/artela-network/aspect-example/new_test/build/release.wasm")
	if err != nil {
		fmt.Println(err)
	}

	runner, err := NewRunner(context.Background(), "", raw)
	require.Equal(t, nil, err)
	address := common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4")

	ret, err := runner.IsOwner(99, 99, &address, "hello")
	require.Equal(t, nil, err)
	defer runner.Return()

	require.Equal(t, true, ret)
}
