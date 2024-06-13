package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

type precompiledTest struct {
	Input, Expected string
	Gas             uint64
	Name            string
	NoBenchmark     bool // Benchmark primarily the worst-cases
}

func loadJson(name string) ([]precompiledTest, error) {
	data, err := os.ReadFile(fmt.Sprintf("testdata/%v.json", name))
	if err != nil {
		return nil, err
	}
	var testcases []precompiledTest
	err = json.Unmarshal(data, &testcases)
	return testcases, err
}

func getData(data []byte, start uint64, size uint64) []byte {
	length := uint64(len(data))
	if start > length {
		start = length
	}
	end := start + size
	if end > length {
		end = length
	}
	return common.RightPadBytes(data[start:end], int(size))
}

func TestModExp(t *testing.T) {
	r := NewRegistry(context.Background(), common.Address{}, 1)
	apis := r.cryptoAPIs()

	api, ok := apis["bigModExp"]
	require.Equal(t, true, ok)

	fn, ok := api.Func.(func(b, e, m []byte) []byte)
	require.Equal(t, true, ok)

	testCases, err := loadJson("modexp")
	require.Equal(t, nil, err)
	for _, testCase := range testCases {
		input, err := hex.DecodeString(testCase.Input)
		require.Equal(t, nil, err)

		baseLen := new(big.Int).SetBytes(getData(input, 0, 32)).Uint64()
		expLen := new(big.Int).SetBytes(getData(input, 32, 32)).Uint64()
		modLen := new(big.Int).SetBytes(getData(input, 64, 32)).Uint64()

		if len(input) > 96 {
			input = input[96:]
		} else {
			input = input[:0]
		}

		b := getData(input, 0, baseLen)
		e := getData(input, baseLen, expLen)
		m := getData(input, baseLen+expLen, modLen)

		res := fn(b, e, m)
		require.Equal(t, testCase.Expected, hex.EncodeToString(res))
	}
}
