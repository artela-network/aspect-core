package run

import (
	"encoding/hex"
	"math"
	"math/big"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeDecodeParam(t *testing.T) {
	bigIntMax := big.NewInt(0)
	bigIntMax.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
	tests := []struct {
		types string
		value interface{}
	}{
		{
			types: "uint32",
			value: uint32(123),
		},
		{
			types: "uint256",
			value: big.NewInt(123),
		},
		{
			types: "int64",
			value: int64(123),
		},
		{
			types: "int32",
			value: int32(123),
		},
		{
			types: "uint32",
			value: uint32(math.MaxUint32),
		},
		{
			types: "uint32[]",
			value: []uint32{1001, 1002, 1003},
		},
		{
			types: "string",
			value: "asdfghjkl;'",
		},
		{
			types: "uint256",
			value: bigIntMax,
		},
	}

	for _, test := range tests {
		data, err := encodeParam(test.types, test.value)
		require.Equal(t, nil, err)
		// fmt.Println(hex.EncodeToString(data))

		val, err := decodeParam(test.types, data)
		require.Equal(t, nil, err)
		require.Equal(t, true, reflect.DeepEqual(test.value, val))
	}
}

func TestEncodeDecodeParams(t *testing.T) {
	testCases := []struct {
		expectedHex string
		types       string
		values      []interface{}
	}{
		{
			expectedHex: "00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000477616e670000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000026964000000000000000000000000000000000000000000000000000000000000",
			types:       "string,string",
			values: []interface{}{
				"wang",
				"id",
			},
		},
		{
			expectedHex: "000000000000000000000000000000000000000000000000000000000000000900000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000003616161000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000036262620000000000000000000000000000000000000000000000000000000000",
			types:       "uint256,string,string,int32,int32",
			values: []interface{}{
				big.NewInt(9),
				"aaa",
				"bbb",
				int32(7),
				int32(8),
			},
		},
		{
			expectedHex: "000000000000000000000000000000000000000000000000000000000000000900000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000008",
			types:       "uint32,uint32,uint256",
			values: []interface{}{
				uint32(9),
				uint32(5),
				big.NewInt(8),
			},
		},
	}

	for _, testCase := range testCases {
		data, err := encodeParams(testCase.types, testCase.values...)
		require.Equal(t, nil, err)
		hexData := hex.EncodeToString(data)
		require.Equal(t, testCase.expectedHex, hexData)

		values, err := decodeParams(testCase.types, data)
		require.Equal(t, nil, err)
		require.Equal(t, len(testCase.values), len(values))

		for i, value := range values {
			require.Equal(t, true, reflect.DeepEqual(testCase.values[i], value))
		}
	}
}
