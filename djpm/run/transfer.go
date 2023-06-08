package run

import (
	"errors"
	"math/big"
	"reflect"

	"github.com/artela-network/artelasdk/types"
)

type TypeValue struct {
	value *types.Value
}

func (tv *TypeValue) FromUint32(val uint32) {
	buf := make([]byte, 4)
	for i := 0; i < 4; i++ {
		var b = uint8(val & 0xff)
		buf[i] = b
		val = val >> 8
	}
	tv.value = &types.Value{
		Kind: types.ValueKind_INT,
		Data: buf[:],
	}
}

func (tv *TypeValue) FromInt32(val int32) {
	buf := make([]byte, 4)
	for i := 0; i < 4; i++ {
		var b = uint8(val & 0xff)
		buf[i] = b
		val = val >> 8
	}
	tv.value = &types.Value{
		Kind: types.ValueKind_INT,
		Data: buf[:],
	}
}

func (tv *TypeValue) FromString(val string) {
	tv.value = &types.Value{
		Kind: types.ValueKind_STRING,
		Data: []byte(val),
	}
}

func (tv *TypeValue) ToUint32() uint32 {
	if tv.value.Kind != types.ValueKind_INT {
		return 0
	}

	return 0
}

func (tv *TypeValue) ToUint256() *big.Int {
	return big.NewInt(0)
}

func (tv *TypeValue) ToString() string {
	return string(tv.value.Data)
}

func (tv *TypeValue) GetValue() interface{} {
	switch tv.value.Kind {
	case types.ValueKind_STRING:
		return tv.ToString()
	default:
		return errors.New("not valid")
	}
}

func (tv *TypeValue) SetValue(val interface{}) {
	switch reflect.TypeOf(val) {
	case reflect.TypeOf(string("")):
		tv.FromString(val.(string))
	case reflect.TypeOf(uint32(0)):
		tv.FromUint32(val.(uint32))
	case reflect.TypeOf(int32(0)):
		tv.FromInt32(val.(int32))
	}
}
