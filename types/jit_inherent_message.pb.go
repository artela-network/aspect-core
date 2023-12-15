// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: jitinherent/v1/jit_inherent_message.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JitInherentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender               []byte `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Nonce                uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	NonceKey             []byte `protobuf:"bytes,3,opt,name=nonce_key,json=nonceKey,proto3" json:"nonce_key,omitempty"`
	InitCode             []byte `protobuf:"bytes,4,opt,name=init_code,json=initCode,proto3" json:"init_code,omitempty"`
	CallData             []byte `protobuf:"bytes,5,opt,name=call_data,json=callData,proto3" json:"call_data,omitempty"`
	CallGasLimit         uint64 `protobuf:"varint,6,opt,name=call_gas_limit,json=callGasLimit,proto3" json:"call_gas_limit,omitempty"`
	VerificationGasLimit uint64 `protobuf:"varint,7,opt,name=verification_gas_limit,json=verificationGasLimit,proto3" json:"verification_gas_limit,omitempty"`
	PaymasterAndData     []byte `protobuf:"bytes,8,opt,name=paymaster_and_data,json=paymasterAndData,proto3" json:"paymaster_and_data,omitempty"`
}

func (x *JitInherentRequest) Reset() {
	*x = JitInherentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jitinherent_v1_jit_inherent_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JitInherentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JitInherentRequest) ProtoMessage() {}

func (x *JitInherentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jitinherent_v1_jit_inherent_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JitInherentRequest.ProtoReflect.Descriptor instead.
func (*JitInherentRequest) Descriptor() ([]byte, []int) {
	return file_jitinherent_v1_jit_inherent_message_proto_rawDescGZIP(), []int{0}
}

func (x *JitInherentRequest) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *JitInherentRequest) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *JitInherentRequest) GetNonceKey() []byte {
	if x != nil {
		return x.NonceKey
	}
	return nil
}

func (x *JitInherentRequest) GetInitCode() []byte {
	if x != nil {
		return x.InitCode
	}
	return nil
}

func (x *JitInherentRequest) GetCallData() []byte {
	if x != nil {
		return x.CallData
	}
	return nil
}

func (x *JitInherentRequest) GetCallGasLimit() uint64 {
	if x != nil {
		return x.CallGasLimit
	}
	return 0
}

func (x *JitInherentRequest) GetVerificationGasLimit() uint64 {
	if x != nil {
		return x.VerificationGasLimit
	}
	return 0
}

func (x *JitInherentRequest) GetPaymasterAndData() []byte {
	if x != nil {
		return x.PaymasterAndData
	}
	return nil
}

type JitInherentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JitInherentHashes [][]byte `protobuf:"bytes,1,rep,name=jit_inherent_hashes,json=jitInherentHashes,proto3" json:"jit_inherent_hashes,omitempty"`
	TxHash            []byte   `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Success           bool     `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	Ret               []byte   `protobuf:"bytes,4,opt,name=ret,proto3" json:"ret,omitempty"`
	LeftoverGas       uint64   `protobuf:"varint,5,opt,name=leftover_gas,json=leftoverGas,proto3" json:"leftover_gas,omitempty"`
}

func (x *JitInherentResponse) Reset() {
	*x = JitInherentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jitinherent_v1_jit_inherent_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JitInherentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JitInherentResponse) ProtoMessage() {}

func (x *JitInherentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jitinherent_v1_jit_inherent_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JitInherentResponse.ProtoReflect.Descriptor instead.
func (*JitInherentResponse) Descriptor() ([]byte, []int) {
	return file_jitinherent_v1_jit_inherent_message_proto_rawDescGZIP(), []int{1}
}

func (x *JitInherentResponse) GetJitInherentHashes() [][]byte {
	if x != nil {
		return x.JitInherentHashes
	}
	return nil
}

func (x *JitInherentResponse) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *JitInherentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *JitInherentResponse) GetRet() []byte {
	if x != nil {
		return x.Ret
	}
	return nil
}

func (x *JitInherentResponse) GetLeftoverGas() uint64 {
	if x != nil {
		return x.LeftoverGas
	}
	return 0
}

var File_jitinherent_v1_jit_inherent_message_proto protoreflect.FileDescriptor

var file_jitinherent_v1_jit_inherent_message_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6a, 0x69, 0x74, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x6a, 0x69, 0x74, 0x5f, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6a, 0x69, 0x74,
	0x69, 0x6e, 0x68, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0xa3, 0x02, 0x0a, 0x12,
	0x4a, 0x69, 0x74, 0x49, 0x6e, 0x68, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x69, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63,
	0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x47, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x34, 0x0a,
	0x16, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x61,
	0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x61, 0x73, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x61, 0x79, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x10, 0x70, 0x61, 0x79, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x22, 0xad, 0x01, 0x0a, 0x13, 0x4a, 0x69, 0x74, 0x49, 0x6e, 0x68, 0x65, 0x72, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x6a, 0x69, 0x74,
	0x5f, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x11, 0x6a, 0x69, 0x74, 0x49, 0x6e, 0x68, 0x65, 0x72,
	0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x6c, 0x65, 0x66, 0x74, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6c, 0x65, 0x66, 0x74, 0x6f, 0x76, 0x65, 0x72, 0x47, 0x61,
	0x73, 0x42, 0xb3, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6a, 0x69, 0x74, 0x69, 0x6e, 0x68,
	0x65, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x4a, 0x69, 0x74, 0x49, 0x6e, 0x68,
	0x65, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x72, 0x74, 0x65, 0x6c, 0x61, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61,
	0x73, 0x70, 0x65, 0x63, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x4a, 0x58, 0x58, 0xaa, 0x02, 0x0e, 0x4a, 0x69, 0x74, 0x69, 0x6e, 0x68, 0x65,
	0x72, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x4a, 0x69, 0x74, 0x69, 0x6e, 0x68,
	0x65, 0x72, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x4a, 0x69, 0x74, 0x69, 0x6e,
	0x68, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x4a, 0x69, 0x74, 0x69, 0x6e, 0x68, 0x65, 0x72,
	0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jitinherent_v1_jit_inherent_message_proto_rawDescOnce sync.Once
	file_jitinherent_v1_jit_inherent_message_proto_rawDescData = file_jitinherent_v1_jit_inherent_message_proto_rawDesc
)

func file_jitinherent_v1_jit_inherent_message_proto_rawDescGZIP() []byte {
	file_jitinherent_v1_jit_inherent_message_proto_rawDescOnce.Do(func() {
		file_jitinherent_v1_jit_inherent_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_jitinherent_v1_jit_inherent_message_proto_rawDescData)
	})
	return file_jitinherent_v1_jit_inherent_message_proto_rawDescData
}

var file_jitinherent_v1_jit_inherent_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_jitinherent_v1_jit_inherent_message_proto_goTypes = []interface{}{
	(*JitInherentRequest)(nil),  // 0: jitinherent.v1.JitInherentRequest
	(*JitInherentResponse)(nil), // 1: jitinherent.v1.JitInherentResponse
}
var file_jitinherent_v1_jit_inherent_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_jitinherent_v1_jit_inherent_message_proto_init() }
func file_jitinherent_v1_jit_inherent_message_proto_init() {
	if File_jitinherent_v1_jit_inherent_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jitinherent_v1_jit_inherent_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JitInherentRequest); i {
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
		file_jitinherent_v1_jit_inherent_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JitInherentResponse); i {
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
			RawDescriptor: file_jitinherent_v1_jit_inherent_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jitinherent_v1_jit_inherent_message_proto_goTypes,
		DependencyIndexes: file_jitinherent_v1_jit_inherent_message_proto_depIdxs,
		MessageInfos:      file_jitinherent_v1_jit_inherent_message_proto_msgTypes,
	}.Build()
	File_jitinherent_v1_jit_inherent_message_proto = out.File
	file_jitinherent_v1_jit_inherent_message_proto_rawDesc = nil
	file_jitinherent_v1_jit_inherent_message_proto_goTypes = nil
	file_jitinherent_v1_jit_inherent_message_proto_depIdxs = nil
}
