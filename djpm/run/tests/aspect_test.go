package tests

import (
	"context"
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"

	"github.com/artela-network/aspect-core/djpm/run"

	"github.com/stretchr/testify/require"

	aspectType "github.com/artela-network/aspect-core/types"
)

func TestAspect(t *testing.T) {
	raw, _ := GetTestTarget("aspect-test")
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
		CurrInnerTx: &aspectType.EthStackTransaction{
			From:          "0x1c0e4b5d5f50fe65adc4cd658cd88ae0dfdbb3ba",
			To:            "0x1c0e4b5d5f50fe65adc4cd658cd88ae0dfdbb3ba",
			Data:          []byte{0x1, 0x2, 0x3, 0x4},
			Value:         "200000",
			Gas:           "100",
			Ret:           nil,
			LeftOverGas:   1000,
			Index:         0,
			ParentIndex:   0,
			ChildrenIndex: nil,
		},
	}
	runner, err := run.NewRunner(context.Background(), "", raw)
	require.Equal(t, nil, err)
	pointcuts := []aspectType.PointCut{aspectType.FILTER_TX}

	for _, point := range pointcuts {

		output, err := runner.JoinPoint(point, 1000, 999, nil, input)
		require.Equal(t, nil, err)
		require.Equal(t, true, output.Result.Success)
		marshal, err := jsoniter.Marshal(output)
		if err != nil {
			return
		}
		fmt.Println("------" + (string(marshal)) + "------")

		data := output.Data
		strData := &aspectType.BoolData{}
		err2 := data.UnmarshalTo(strData)
		if err2 != nil {
			return
		}
		fmt.Println(strData.GetData())
	}

	pointcuts = []aspectType.PointCut{aspectType.ON_GAS_PAYMENT_METHOD}

	for _, point := range pointcuts {

		output, err := runner.JoinPoint(point, 1000, 999, nil, input)
		require.Equal(t, nil, err)
		require.Equal(t, true, output.Result.Success)
		marshal, err := jsoniter.Marshal(output)
		if err != nil {
			return
		}
		fmt.Println("------" + (string(marshal)) + "------")

		data := output.Data
		strData := &aspectType.StringData{}
		err2 := data.UnmarshalTo(strData)
		if err2 != nil {
			return
		}
		fmt.Println(strData.GetData())
	}

	pointcuts = []aspectType.PointCut{
		aspectType.PRE_CONTRACT_CALL_METHOD,
		aspectType.POST_CONTRACT_CALL_METHOD,
		aspectType.PRE_TX_EXECUTE_METHOD,
		aspectType.POST_TX_EXECUTE_METHOD,
		aspectType.POST_TX_COMMIT,
	}

	for _, point := range pointcuts {
		output, err := runner.JoinPoint(point, 1000, 999, nil, input)
		require.Equal(t, nil, err)
		require.Equal(t, true, output.Result.Success)

		marshal, outErr := jsoniter.Marshal(output)
		if outErr != nil {
			return
		}
		fmt.Println("------" + (string(marshal)) + "------")

	}

	defer runner.Return()
}

// Run "scripts/build-wasm.sh" in project root, before run this test.
func TestIsOwner(t *testing.T) {
	// cwd, _ := os.Getwd()
	raw, err := GetTestTarget("aspect-test")
	if err != nil {
		fmt.Println(err)
	}

	runner, err := run.NewRunner(context.Background(), "", raw)
	require.Equal(t, nil, err)
	ret, err := runner.IsOwner(99, 99, nil, "hello")
	require.Equal(t, nil, err)
	defer runner.Return()

	require.Equal(t, true, ret)
}
