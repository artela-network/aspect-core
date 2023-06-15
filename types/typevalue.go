package types

import (
	"errors"
	"math/big"
	"reflect"
)

type TypeValue struct {
	value *Value
}

func NewTypeValue(value *Value) *TypeValue {
	return &TypeValue{
		value: value,
	}
}

func (tv *TypeValue) Value() *Value {
	return tv.value
}

func (tv *TypeValue) FromUint32(val uint32) {
	buf := make([]byte, 4)
	for i := 0; i < 4; i++ {
		var b = uint8(val & 0xff)
		buf[i] = b
		val = val >> 8
	}
	tv.value = &Value{
		Kind: ValueKind_UINT32,
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
	tv.value = &Value{
		Kind: ValueKind_INT32,
		Data: buf[:],
	}
}

func (tv *TypeValue) FromUint256(val *big.Int) {
	tv.value = &Value{
		Kind: ValueKind_UINT256,
		Data: val.Bytes(),
	}
}

func (tv *TypeValue) FromString(val string) {
	tv.value = &Value{
		Kind: ValueKind_STRING,
		Data: []byte(val),
	}
}

func (tv *TypeValue) ToInt32() int32 {
	val := big.NewInt(0).SetBytes(tv.value.Data).Int64()
	return int32(val)
}

func (tv *TypeValue) ToInt64() int64 {
	val := big.NewInt(0).SetBytes(tv.value.Data).Int64()
	return val
}

func (tv *TypeValue) ToUint32() uint32 {
	val := big.NewInt(0).SetBytes(tv.value.Data).Uint64()
	return uint32(val)
}

func (tv *TypeValue) ToUint64() uint64 {
	val := big.NewInt(0).SetBytes(tv.value.Data).Uint64()
	return val
}

func (tv *TypeValue) ToUint256() *big.Int {
	return big.NewInt(0).SetBytes(tv.value.Data)
}

func (tv *TypeValue) ToBool() bool {
	return tv.value.Data[0] != 0
}

func (tv *TypeValue) ToString() string {
	return string(tv.value.Data)
}

func (tv *TypeValue) GetValue() interface{} {
	switch tv.value.Kind {
	case ValueKind_STRING:
		return tv.ToString()
	case ValueKind_INT32:
		return tv.ToInt32()
	case ValueKind_INT64:
		return tv.ToInt64()
	case ValueKind_UINT32:
		return tv.ToUint32()
	case ValueKind_UINT64:
		return tv.ToUint64()
	case ValueKind_UINT256:
		return tv.ToUint256()
	case ValueKind_BOOL:
		return tv.ToBool()
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
	case reflect.TypeOf(big.NewInt(0)):
		tv.FromUint256(val.(*big.Int))
	}
}
