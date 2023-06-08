package run

import (
	"math"
	"math/big"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeDecode(t *testing.T) {
	bigIntMax := big.NewInt(0)
	bigIntMax.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
	tests := []struct {
		types string
		value interface{}
	}{
		{
			types: "uint32",
			value: uint32(1000),
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
		data, err := encodeType(test.types, test.value)
		require.Equal(t, nil, err)

		val, err := decodeType(test.types, data)
		require.Equal(t, nil, err)
		require.Equal(t, true, reflect.DeepEqual(test.value, val))
	}
}
