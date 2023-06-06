package run

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
)

func decodeInt32(data []byte) int32 {
	// def := fmt.Sprintf(`[{ "name" : "method", "type": "function", "outputs": %s}]`, `[{"type": "int32"}]`)
	// abi, err := abi.JSON(strings.NewReader(def))
	// if err != nil {
	// 	return 0
	// }

	// outptr := reflect.New(reflect.TypeOf(int32(0)))
	// err = abi.UnpackIntoInterface(outptr.Interface(), "method", data)
	// if err != nil {
	// 	return 0
	// }
	// out := outptr.Elem().Interface()
	out, err := decodeType("int32", data)
	if err != nil {
		return -1
	}
	val, ok := out.(int32)
	if !ok {
		return -1
	}
	return val
}

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
