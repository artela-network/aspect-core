package run

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
)

func decodeParam(t string, data []byte) (interface{}, error) {
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

// decodeParams decode types and encoded data
// t is a types string, with a format of "typeA,typeB,typeC..."
func decodeParams(t string, data []byte) ([]interface{}, error) {
	types := strings.Split(t, ",")
	decodeABI := abi.Arguments{}
	for _, ele := range types {
		ty, err := abi.NewType(ele, "", nil)
		if err != nil {
			return nil, err
		}
		decodeABI = append(decodeABI, abi.Argument{Type: ty})
	}

	return decodeABI.Unpack(data)
}

// encodeParam encode type and value to abi bytecodes
func encodeParam(t string, value interface{}) ([]byte, error) {
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

// encodeParams encode types and values
// t is a types string, with a format of "typeA,typeB,typeC..."
// the values is required to perfect match to the types.
func encodeParams(t string, values ...interface{}) ([]byte, error) {
	types := strings.Split(t, ",")
	if len(types) != len(values) {
		return nil, errors.Errorf("encodeTypes failed, type %s is not match to value count %d", t, len(values))
	}
	encodeABI := abi.Arguments{}
	for _, ele := range types {
		ty, err := abi.NewType(ele, "", nil)
		if err != nil {
			return nil, err
		}
		encodeABI = append(encodeABI, abi.Argument{Type: ty})
	}

	return encodeABI.Pack(values...)
}
