// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/bool.proto

package cases

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type BoolNone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val bool `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *BoolNone) Reset() {
	*x = BoolNone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_bool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolNone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolNone) ProtoMessage() {}

func (x *BoolNone) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolNone.ProtoReflect.Descriptor instead.
func (*BoolNone) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_bool_proto_rawDescGZIP(), []int{0}
}

func (x *BoolNone) GetVal() bool {
	if x != nil {
		return x.Val
	}
	return false
}

type BoolConstTrue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val bool `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *BoolConstTrue) Reset() {
	*x = BoolConstTrue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_bool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolConstTrue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolConstTrue) ProtoMessage() {}

func (x *BoolConstTrue) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolConstTrue.ProtoReflect.Descriptor instead.
func (*BoolConstTrue) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_bool_proto_rawDescGZIP(), []int{1}
}

func (x *BoolConstTrue) GetVal() bool {
	if x != nil {
		return x.Val
	}
	return false
}

type BoolConstFalse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val bool `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *BoolConstFalse) Reset() {
	*x = BoolConstFalse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_bool_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolConstFalse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolConstFalse) ProtoMessage() {}

func (x *BoolConstFalse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_bool_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolConstFalse.ProtoReflect.Descriptor instead.
func (*BoolConstFalse) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_bool_proto_rawDescGZIP(), []int{2}
}

func (x *BoolConstFalse) GetVal() bool {
	if x != nil {
		return x.Val
	}
	return false
}

var File_buf_validate_conformance_cases_bool_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_bool_proto_rawDesc = []byte{
	0x0a, 0x29, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x62, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x75, 0x66,
	0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6c,
	0x4e, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x2b, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6c, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x54, 0x72, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x08, 0xfa, 0xf7, 0x18, 0x04, 0x6a, 0x02, 0x08, 0x01, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x22, 0x2c, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x46, 0x61, 0x6c, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x08, 0xfa, 0xf7, 0x18, 0x04, 0x6a, 0x02, 0x08, 0x00, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x42, 0xa6, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x42, 0x09, 0x42, 0x6f, 0x6f, 0x6c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x67,
	0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0xa2, 0x02, 0x04, 0x42, 0x56, 0x43, 0x43, 0xaa, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x73, 0xca, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x5c, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0xe2, 0x02, 0x2a, 0x42, 0x75, 0x66, 0x5c,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x3a, 0x3a, 0x43, 0x61, 0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_buf_validate_conformance_cases_bool_proto_rawDescOnce sync.Once
	file_buf_validate_conformance_cases_bool_proto_rawDescData = file_buf_validate_conformance_cases_bool_proto_rawDesc
)

func file_buf_validate_conformance_cases_bool_proto_rawDescGZIP() []byte {
	file_buf_validate_conformance_cases_bool_proto_rawDescOnce.Do(func() {
		file_buf_validate_conformance_cases_bool_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_validate_conformance_cases_bool_proto_rawDescData)
	})
	return file_buf_validate_conformance_cases_bool_proto_rawDescData
}

var file_buf_validate_conformance_cases_bool_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buf_validate_conformance_cases_bool_proto_goTypes = []interface{}{
	(*BoolNone)(nil),       // 0: buf.validate.conformance.cases.BoolNone
	(*BoolConstTrue)(nil),  // 1: buf.validate.conformance.cases.BoolConstTrue
	(*BoolConstFalse)(nil), // 2: buf.validate.conformance.cases.BoolConstFalse
}
var file_buf_validate_conformance_cases_bool_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_bool_proto_init() }
func file_buf_validate_conformance_cases_bool_proto_init() {
	if File_buf_validate_conformance_cases_bool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_validate_conformance_cases_bool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolNone); i {
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
		file_buf_validate_conformance_cases_bool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolConstTrue); i {
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
		file_buf_validate_conformance_cases_bool_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolConstFalse); i {
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
			RawDescriptor: file_buf_validate_conformance_cases_bool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_bool_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_bool_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_bool_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_bool_proto = out.File
	file_buf_validate_conformance_cases_bool_proto_rawDesc = nil
	file_buf_validate_conformance_cases_bool_proto_goTypes = nil
	file_buf_validate_conformance_cases_bool_proto_depIdxs = nil
}
