package run

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
)

// def: `[{"name":"Int","type":"int256"},{"name":"_int","type":"int256"}]`,
// enc: "00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002",
// want: struct {
// 	Int1 *big.Int
// 	Int2 *big.Int
// }{},

func decodeType(t string, data []byte) (interface{}, error) {
	ty, err := abi.NewType(t, "", nil)
	if err != nil {
		return nil, err
	}
	decodeABI := abi.Arguments{
		{Type: ty},
	}
	values, err := decodeABI.Unpack(data)
	if err != nil {
		return nil, err
	}
	if len(values) != 1 {
		return nil, errors.Errorf("decode data %s to type %s failed", data, t)
	}
	return values[0], nil
}

// encodeType encode type and value to abi bytecodes
func encodeType(t string, value interface{}) ([]byte, error) {
	var encodeABI abi.Arguments
	ty, err := abi.NewType(t, "", nil)
	if err != nil {
		return nil, err
	}
	encodeABI = abi.Arguments{
		{Type: ty},
	}

	return encodeABI.Pack(value)
}

func decodeTypes(t string, data []byte) ([]interface{}, error) {
	types := strings.Split(t, ",")
	vals := make([]interface{}, len(types))
	// TODO, datas is split of data, how to split the data?
	datas := make([][]byte, len(types))
	for i, ty := range types {
		val, err := decodeType(ty, datas[i])
		if err != nil {
			return nil, err
		}
		vals[i] = val
	}
	return vals, nil
}

// encodeTypes encode types and values
// t is a types string, with a format of "typeA,typeB,typeC..."
// the values is required to perfect match to the types.
func encodeTypes(t string, values ...interface{}) ([]byte, error) {
	types := strings.Split(t, ",")
	if len(types) != len(values) {
		return nil, errors.Errorf("encodeTypes failed, type %s is not match to value count %d", t, len(values))
	}
	data := make([]byte, 0)
	for i, ty := range types {
		encoded, err := encodeType(ty, values[i])
		if err != nil {
			return nil, errors.Wrap(err, "encodeTypes failed")
		}
		data = append(data, encoded...)
	}
	return data, nil
}
