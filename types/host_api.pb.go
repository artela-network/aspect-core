// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: aspect/v2/host_api.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataSpaceType int32

const (
	DataSpaceType_TX_ASPECT_CONTEXT   DataSpaceType = 0
	DataSpaceType_TX_EXT_PROPERTIES   DataSpaceType = 1
	DataSpaceType_TX_CONTENT          DataSpaceType = 2
	DataSpaceType_TX_STATE_CHANGES    DataSpaceType = 3
	DataSpaceType_TX_CALL_TREE        DataSpaceType = 4
	DataSpaceType_TX_RECEIPT          DataSpaceType = 5
	DataSpaceType_TX_GAS_METER        DataSpaceType = 6
	DataSpaceType_ENV_CONS_PARAMS     DataSpaceType = 7
	DataSpaceType_ENV_CHAIN_CONFIG    DataSpaceType = 8
	DataSpaceType_ENV_EVM_PARAMS      DataSpaceType = 9
	DataSpaceType_ENV_BASE_INFO       DataSpaceType = 10
	DataSpaceType_BLOCK_HEADER        DataSpaceType = 11
	DataSpaceType_BLOCK_TXS           DataSpaceType = 12
	DataSpaceType_BLOCK_GAS_METER     DataSpaceType = 13
	DataSpaceType_BLOCK_MIN_GAS_PRICE DataSpaceType = 14
	DataSpaceType_BLOCK_LAST_COMMIT   DataSpaceType = 15
	DataSpaceType_BLOCK_BLOCK_ID      DataSpaceType = 16
	DataSpaceType_BLOCK_EVIDENCE      DataSpaceType = 17
)

// Enum value maps for DataSpaceType.
var (
	DataSpaceType_name = map[int32]string{
		0:  "TX_ASPECT_CONTEXT",
		1:  "TX_EXT_PROPERTIES",
		2:  "TX_CONTENT",
		3:  "TX_STATE_CHANGES",
		4:  "TX_CALL_TREE",
		5:  "TX_RECEIPT",
		6:  "TX_GAS_METER",
		7:  "ENV_CONS_PARAMS",
		8:  "ENV_CHAIN_CONFIG",
		9:  "ENV_EVM_PARAMS",
		10: "ENV_BASE_INFO",
		11: "BLOCK_HEADER",
		12: "BLOCK_TXS",
		13: "BLOCK_GAS_METER",
		14: "BLOCK_MIN_GAS_PRICE",
		15: "BLOCK_LAST_COMMIT",
		16: "BLOCK_BLOCK_ID",
		17: "BLOCK_EVIDENCE",
	}
	DataSpaceType_value = map[string]int32{
		"TX_ASPECT_CONTEXT":   0,
		"TX_EXT_PROPERTIES":   1,
		"TX_CONTENT":          2,
		"TX_STATE_CHANGES":    3,
		"TX_CALL_TREE":        4,
		"TX_RECEIPT":          5,
		"TX_GAS_METER":        6,
		"ENV_CONS_PARAMS":     7,
		"ENV_CHAIN_CONFIG":    8,
		"ENV_EVM_PARAMS":      9,
		"ENV_BASE_INFO":       10,
		"BLOCK_HEADER":        11,
		"BLOCK_TXS":           12,
		"BLOCK_GAS_METER":     13,
		"BLOCK_MIN_GAS_PRICE": 14,
		"BLOCK_LAST_COMMIT":   15,
		"BLOCK_BLOCK_ID":      16,
		"BLOCK_EVIDENCE":      17,
	}
)

func (x DataSpaceType) Enum() *DataSpaceType {
	p := new(DataSpaceType)
	*p = x
	return p
}

func (x DataSpaceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSpaceType) Descriptor() protoreflect.EnumDescriptor {
	return file_aspect_v2_host_api_proto_enumTypes[0].Descriptor()
}

func (DataSpaceType) Type() protoreflect.EnumType {
	return &file_aspect_v2_host_api_proto_enumTypes[0]
}

func (x DataSpaceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSpaceType.Descriptor instead.
func (DataSpaceType) EnumDescriptor() ([]byte, []int) {
	return file_aspect_v2_host_api_proto_rawDescGZIP(), []int{0}
}

type ContextQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataSpace  DataSpaceType `protobuf:"varint,1,opt,name=data_space,json=dataSpace,proto3,enum=aspect.v2.DataSpaceType" json:"data_space,omitempty"`
	Conditions []string      `protobuf:"bytes,2,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *ContextQueryRequest) Reset() {
	*x = ContextQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspect_v2_host_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextQueryRequest) ProtoMessage() {}

func (x *ContextQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aspect_v2_host_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextQueryRequest.ProtoReflect.Descriptor instead.
func (*ContextQueryRequest) Descriptor() ([]byte, []int) {
	return file_aspect_v2_host_api_proto_rawDescGZIP(), []int{0}
}

func (x *ContextQueryRequest) GetDataSpace() DataSpaceType {
	if x != nil {
		return x.DataSpace
	}
	return DataSpaceType_TX_ASPECT_CONTEXT
}

func (x *ContextQueryRequest) GetConditions() []string {
	if x != nil {
		return x.Conditions
	}
	return nil
}

type ContextQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result          *RunResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	DataMessageType string     `protobuf:"bytes,2,opt,name=data_message_type,json=dataMessageType,proto3" json:"data_message_type,omitempty"`
	Data            *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ContextQueryResponse) Reset() {
	*x = ContextQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspect_v2_host_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextQueryResponse) ProtoMessage() {}

func (x *ContextQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aspect_v2_host_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextQueryResponse.ProtoReflect.Descriptor instead.
func (*ContextQueryResponse) Descriptor() ([]byte, []int) {
	return file_aspect_v2_host_api_proto_rawDescGZIP(), []int{1}
}

func (x *ContextQueryResponse) GetResult() *RunResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *ContextQueryResponse) GetDataMessageType() string {
	if x != nil {
		return x.DataMessageType
	}
	return ""
}

func (x *ContextQueryResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type CallMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *EthMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CallMessageRequest) Reset() {
	*x = CallMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspect_v2_host_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallMessageRequest) ProtoMessage() {}

func (x *CallMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aspect_v2_host_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallMessageRequest.ProtoReflect.Descriptor instead.
func (*CallMessageRequest) Descriptor() ([]byte, []int) {
	return file_aspect_v2_host_api_proto_rawDescGZIP(), []int{2}
}

func (x *CallMessageRequest) GetMessage() *EthMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type CallMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *RunResult            `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Data   *EthMessageCallResult `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CallMessageResponse) Reset() {
	*x = CallMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspect_v2_host_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallMessageResponse) ProtoMessage() {}

func (x *CallMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aspect_v2_host_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallMessageResponse.ProtoReflect.Descriptor instead.
func (*CallMessageResponse) Descriptor() ([]byte, []int) {
	return file_aspect_v2_host_api_proto_rawDescGZIP(), []int{3}
}

func (x *CallMessageResponse) GetResult() *RunResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *CallMessageResponse) GetData() *EthMessageCallResult {
	if x != nil {
		return x.Data
	}
	return nil
}

// MsgEthereumTxResponse defines the Msg/EthereumTx response type.
type EthMessageCallResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hash of the ethereum transaction in hex format. This hash differs from the
	// Tendermint sha256 hash of the transaction bytes. See
	// https://github.com/tendermint/tendermint/issues/6539 for reference
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// logs contains the transaction hash and the proto-compatible ethereum
	// logs.
	Logs []*EthLog `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
	// ret is the returned data from evm function (result or data supplied with revert
	// opcode)
	Ret []byte `protobuf:"bytes,3,opt,name=ret,proto3" json:"ret,omitempty"`
	// vm_error is the error returned by vm execution
	VmError string `protobuf:"bytes,4,opt,name=vm_error,json=vmError,proto3" json:"vm_error,omitempty"`
	// gas_used specifies how much gas was consumed by the transaction
	GasUsed uint64 `protobuf:"varint,5,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
}

func (x *EthMessageCallResult) Reset() {
	*x = EthMessageCallResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspect_v2_host_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthMessageCallResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthMessageCallResult) ProtoMessage() {}

func (x *EthMessageCallResult) ProtoReflect() protoreflect.Message {
	mi := &file_aspect_v2_host_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthMessageCallResult.ProtoReflect.Descriptor instead.
func (*EthMessageCallResult) Descriptor() ([]byte, []int) {
	return file_aspect_v2_host_api_proto_rawDescGZIP(), []int{4}
}

func (x *EthMessageCallResult) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *EthMessageCallResult) GetLogs() []*EthLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *EthMessageCallResult) GetRet() []byte {
	if x != nil {
		return x.Ret
	}
	return nil
}

func (x *EthMessageCallResult) GetVmError() string {
	if x != nil {
		return x.VmError
	}
	return ""
}

func (x *EthMessageCallResult) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

// Valida
type EthMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To        string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	GasTipCap string `protobuf:"bytes,3,opt,name=gas_tip_cap,json=gasTipCap,proto3" json:"gas_tip_cap,omitempty"`
	// gas_fee_cap defines the max value for the gas fee
	GasFeeCap string `protobuf:"bytes,4,opt,name=gas_fee_cap,json=gasFeeCap,proto3" json:"gas_fee_cap,omitempty"`
	// gas defines the gas limit defined for the transaction.
	Gas      uint64 `protobuf:"varint,5,opt,name=gas,proto3" json:"gas,omitempty"`
	GasPrice string `protobuf:"bytes,6,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	Value    string `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	Nonce    uint64 `protobuf:"varint,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// data is the data payload bytes of the transaction.
	Input  []byte `protobuf:"bytes,9,opt,name=input,proto3" json:"input,omitempty"`
	IsFake bool   `protobuf:"varint,10,opt,name=is_fake,json=isFake,proto3" json:"is_fake,omitempty"`
}

func (x *EthMessage) Reset() {
	*x = EthMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspect_v2_host_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthMessage) ProtoMessage() {}

func (x *EthMessage) ProtoReflect() protoreflect.Message {
	mi := &file_aspect_v2_host_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthMessage.ProtoReflect.Descriptor instead.
func (*EthMessage) Descriptor() ([]byte, []int) {
	return file_aspect_v2_host_api_proto_rawDescGZIP(), []int{5}
}

func (x *EthMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *EthMessage) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *EthMessage) GetGasTipCap() string {
	if x != nil {
		return x.GasTipCap
	}
	return ""
}

func (x *EthMessage) GetGasFeeCap() string {
	if x != nil {
		return x.GasFeeCap
	}
	return ""
}

func (x *EthMessage) GetGas() uint64 {
	if x != nil {
		return x.Gas
	}
	return 0
}

func (x *EthMessage) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

func (x *EthMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *EthMessage) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *EthMessage) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *EthMessage) GetIsFake() bool {
	if x != nil {
		return x.IsFake
	}
	return false
}

type StringData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StringData) Reset() {
	*x = StringData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspect_v2_host_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringData) ProtoMessage() {}

func (x *StringData) ProtoReflect() protoreflect.Message {
	mi := &file_aspect_v2_host_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringData.ProtoReflect.Descriptor instead.
func (*StringData) Descriptor() ([]byte, []int) {
	return file_aspect_v2_host_api_proto_rawDescGZIP(), []int{6}
}

func (x *StringData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type IntData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data int64 `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *IntData) Reset() {
	*x = IntData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspect_v2_host_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntData) ProtoMessage() {}

func (x *IntData) ProtoReflect() protoreflect.Message {
	mi := &file_aspect_v2_host_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntData.ProtoReflect.Descriptor instead.
func (*IntData) Descriptor() ([]byte, []int) {
	return file_aspect_v2_host_api_proto_rawDescGZIP(), []int{7}
}

func (x *IntData) GetData() int64 {
	if x != nil {
		return x.Data
	}
	return 0
}

type BoolData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data bool `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BoolData) Reset() {
	*x = BoolData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspect_v2_host_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolData) ProtoMessage() {}

func (x *BoolData) ProtoReflect() protoreflect.Message {
	mi := &file_aspect_v2_host_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolData.ProtoReflect.Descriptor instead.
func (*BoolData) Descriptor() ([]byte, []int) {
	return file_aspect_v2_host_api_proto_rawDescGZIP(), []int{8}
}

func (x *BoolData) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

type BytesArrayData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BytesArrayData) Reset() {
	*x = BytesArrayData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspect_v2_host_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesArrayData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesArrayData) ProtoMessage() {}

func (x *BytesArrayData) ProtoReflect() protoreflect.Message {
	mi := &file_aspect_v2_host_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesArrayData.ProtoReflect.Descriptor instead.
func (*BytesArrayData) Descriptor() ([]byte, []int) {
	return file_aspect_v2_host_api_proto_rawDescGZIP(), []int{9}
}

func (x *BytesArrayData) GetData() [][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BytesData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BytesData) Reset() {
	*x = BytesData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspect_v2_host_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesData) ProtoMessage() {}

func (x *BytesData) ProtoReflect() protoreflect.Message {
	mi := &file_aspect_v2_host_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesData.ProtoReflect.Descriptor instead.
func (*BytesData) Descriptor() ([]byte, []int) {
	return file_aspect_v2_host_api_proto_rawDescGZIP(), []int{10}
}

func (x *BytesData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_aspect_v2_host_api_proto protoreflect.FileDescriptor

var file_aspect_v2_host_api_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x68, 0x6f, 0x73, 0x74,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x61, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e,
	0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x61, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x70, 0x61, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9a,
	0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a, 0x12, 0x43,
	0x61, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x45,
	0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x78, 0x0a, 0x13, 0x43, 0x61, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x45, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x99, 0x01, 0x0a,
	0x14, 0x45, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x25, 0x0a, 0x04, 0x6c, 0x6f, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x45, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72,
	0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x6d, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x19, 0x0a,
	0x08, 0x67, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x22, 0xfa, 0x01, 0x0a, 0x0a, 0x45, 0x74, 0x68,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0b, 0x67,
	0x61, 0x73, 0x5f, 0x74, 0x69, 0x70, 0x5f, 0x63, 0x61, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x67, 0x61, 0x73, 0x54, 0x69, 0x70, 0x43, 0x61, 0x70, 0x12, 0x1e, 0x0a, 0x0b, 0x67,
	0x61, 0x73, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x63, 0x61, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x67, 0x61, 0x73, 0x46, 0x65, 0x65, 0x43, 0x61, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x67,
	0x61, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x73, 0x5f, 0x66, 0x61, 0x6b, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69,
	0x73, 0x46, 0x61, 0x6b, 0x65, 0x22, 0x20, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1d, 0x0a, 0x07, 0x49, 0x6e, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1e, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x0e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1f, 0x0a, 0x09,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0xf7, 0x02,
	0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x53, 0x70, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x54, 0x58, 0x5f, 0x41, 0x53, 0x50, 0x45, 0x43, 0x54, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x45, 0x58, 0x54, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x58, 0x5f, 0x45, 0x58, 0x54,
	0x5f, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x49, 0x45, 0x53, 0x10, 0x01, 0x12, 0x0e, 0x0a,
	0x0a, 0x54, 0x58, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x14, 0x0a,
	0x10, 0x54, 0x58, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x53, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x58, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x54,
	0x52, 0x45, 0x45, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x58, 0x5f, 0x52, 0x45, 0x43, 0x45,
	0x49, 0x50, 0x54, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x58, 0x5f, 0x47, 0x41, 0x53, 0x5f,
	0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x56, 0x5f, 0x43,
	0x4f, 0x4e, 0x53, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x53, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10,
	0x45, 0x4e, 0x56, 0x5f, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47,
	0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x4e, 0x56, 0x5f, 0x45, 0x56, 0x4d, 0x5f, 0x50, 0x41,
	0x52, 0x41, 0x4d, 0x53, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x4e, 0x56, 0x5f, 0x42, 0x41,
	0x53, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x4c, 0x4f,
	0x43, 0x4b, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x0b, 0x12, 0x0d, 0x0a, 0x09, 0x42,
	0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x54, 0x58, 0x53, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x4c,
	0x4f, 0x43, 0x4b, 0x5f, 0x47, 0x41, 0x53, 0x5f, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x0d, 0x12,
	0x17, 0x0a, 0x13, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x4d, 0x49, 0x4e, 0x5f, 0x47, 0x41, 0x53,
	0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x0e, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x5f, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10, 0x0f, 0x12,
	0x12, 0x0a, 0x0e, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x49,
	0x44, 0x10, 0x10, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x45, 0x56, 0x49,
	0x44, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x11, 0x42, 0x8d, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x32, 0x42, 0x0c, 0x48, 0x6f, 0x73, 0x74, 0x41,
	0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x74, 0x65, 0x6c, 0x61, 0x2d, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x72, 0x74, 0x65, 0x6c, 0x61, 0x73, 0x64, 0x6b, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x41, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x09, 0x41, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5c,
	0x56, 0x32, 0xe2, 0x02, 0x15, 0x41, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5c, 0x56, 0x32, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x41, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aspect_v2_host_api_proto_rawDescOnce sync.Once
	file_aspect_v2_host_api_proto_rawDescData = file_aspect_v2_host_api_proto_rawDesc
)

func file_aspect_v2_host_api_proto_rawDescGZIP() []byte {
	file_aspect_v2_host_api_proto_rawDescOnce.Do(func() {
		file_aspect_v2_host_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_aspect_v2_host_api_proto_rawDescData)
	})
	return file_aspect_v2_host_api_proto_rawDescData
}

var file_aspect_v2_host_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aspect_v2_host_api_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_aspect_v2_host_api_proto_goTypes = []interface{}{
	(DataSpaceType)(0),           // 0: aspect.v2.DataSpaceType
	(*ContextQueryRequest)(nil),  // 1: aspect.v2.ContextQueryRequest
	(*ContextQueryResponse)(nil), // 2: aspect.v2.ContextQueryResponse
	(*CallMessageRequest)(nil),   // 3: aspect.v2.CallMessageRequest
	(*CallMessageResponse)(nil),  // 4: aspect.v2.CallMessageResponse
	(*EthMessageCallResult)(nil), // 5: aspect.v2.EthMessageCallResult
	(*EthMessage)(nil),           // 6: aspect.v2.EthMessage
	(*StringData)(nil),           // 7: aspect.v2.StringData
	(*IntData)(nil),              // 8: aspect.v2.IntData
	(*BoolData)(nil),             // 9: aspect.v2.BoolData
	(*BytesArrayData)(nil),       // 10: aspect.v2.BytesArrayData
	(*BytesData)(nil),            // 11: aspect.v2.BytesData
	(*RunResult)(nil),            // 12: aspect.v2.RunResult
	(*anypb.Any)(nil),            // 13: google.protobuf.Any
	(*EthLog)(nil),               // 14: aspect.v2.EthLog
}
var file_aspect_v2_host_api_proto_depIdxs = []int32{
	0,  // 0: aspect.v2.ContextQueryRequest.data_space:type_name -> aspect.v2.DataSpaceType
	12, // 1: aspect.v2.ContextQueryResponse.result:type_name -> aspect.v2.RunResult
	13, // 2: aspect.v2.ContextQueryResponse.data:type_name -> google.protobuf.Any
	6,  // 3: aspect.v2.CallMessageRequest.message:type_name -> aspect.v2.EthMessage
	12, // 4: aspect.v2.CallMessageResponse.result:type_name -> aspect.v2.RunResult
	5,  // 5: aspect.v2.CallMessageResponse.data:type_name -> aspect.v2.EthMessageCallResult
	14, // 6: aspect.v2.EthMessageCallResult.logs:type_name -> aspect.v2.EthLog
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_aspect_v2_host_api_proto_init() }
func file_aspect_v2_host_api_proto_init() {
	if File_aspect_v2_host_api_proto != nil {
		return
	}
	file_aspect_v2_base_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_aspect_v2_host_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContextQueryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aspect_v2_host_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContextQueryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aspect_v2_host_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallMessageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aspect_v2_host_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallMessageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aspect_v2_host_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthMessageCallResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aspect_v2_host_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aspect_v2_host_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aspect_v2_host_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aspect_v2_host_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aspect_v2_host_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesArrayData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aspect_v2_host_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aspect_v2_host_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aspect_v2_host_api_proto_goTypes,
		DependencyIndexes: file_aspect_v2_host_api_proto_depIdxs,
		EnumInfos:         file_aspect_v2_host_api_proto_enumTypes,
		MessageInfos:      file_aspect_v2_host_api_proto_msgTypes,
	}.Build()
	File_aspect_v2_host_api_proto = out.File
	file_aspect_v2_host_api_proto_rawDesc = nil
	file_aspect_v2_host_api_proto_goTypes = nil
	file_aspect_v2_host_api_proto_depIdxs = nil
}
