// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: boolean.proto

package grpcService

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BooleanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *BooleanResponse) Reset() {
	*x = BooleanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boolean_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooleanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooleanResponse) ProtoMessage() {}

func (x *BooleanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boolean_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooleanResponse.ProtoReflect.Descriptor instead.
func (*BooleanResponse) Descriptor() ([]byte, []int) {
	return file_boolean_proto_rawDescGZIP(), []int{0}
}

func (x *BooleanResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type WeightedBooleanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FalseWeight int32 `protobuf:"varint,1,opt,name=FalseWeight,proto3" json:"FalseWeight,omitempty"`
	TrueWeight  int32 `protobuf:"varint,2,opt,name=TrueWeight,proto3" json:"TrueWeight,omitempty"`
}

func (x *WeightedBooleanRequest) Reset() {
	*x = WeightedBooleanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boolean_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeightedBooleanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeightedBooleanRequest) ProtoMessage() {}

func (x *WeightedBooleanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boolean_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeightedBooleanRequest.ProtoReflect.Descriptor instead.
func (*WeightedBooleanRequest) Descriptor() ([]byte, []int) {
	return file_boolean_proto_rawDescGZIP(), []int{1}
}

func (x *WeightedBooleanRequest) GetFalseWeight() int32 {
	if x != nil {
		return x.FalseWeight
	}
	return 0
}

func (x *WeightedBooleanRequest) GetTrueWeight() int32 {
	if x != nil {
		return x.TrueWeight
	}
	return 0
}

var File_boolean_proto protoreflect.FileDescriptor

var file_boolean_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x29, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x5a, 0x0a, 0x16, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x61, 0x6c, 0x73, 0x65, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x46, 0x61, 0x6c, 0x73, 0x65,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x72, 0x75, 0x65, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x54, 0x72, 0x75, 0x65,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x1b, 0x5a, 0x19, 0x72, 0x61, 0x6e, 0x64, 0x41, 0x50,
	0x49, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_boolean_proto_rawDescOnce sync.Once
	file_boolean_proto_rawDescData = file_boolean_proto_rawDesc
)

func file_boolean_proto_rawDescGZIP() []byte {
	file_boolean_proto_rawDescOnce.Do(func() {
		file_boolean_proto_rawDescData = protoimpl.X.CompressGZIP(file_boolean_proto_rawDescData)
	})
	return file_boolean_proto_rawDescData
}

var file_boolean_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_boolean_proto_goTypes = []interface{}{
	(*BooleanResponse)(nil),        // 0: BooleanResponse
	(*WeightedBooleanRequest)(nil), // 1: WeightedBooleanRequest
}
var file_boolean_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_boolean_proto_init() }
func file_boolean_proto_init() {
	if File_boolean_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_boolean_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooleanResponse); i {
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
		file_boolean_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeightedBooleanRequest); i {
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
			RawDescriptor: file_boolean_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_boolean_proto_goTypes,
		DependencyIndexes: file_boolean_proto_depIdxs,
		MessageInfos:      file_boolean_proto_msgTypes,
	}.Build()
	File_boolean_proto = out.File
	file_boolean_proto_rawDesc = nil
	file_boolean_proto_goTypes = nil
	file_boolean_proto_depIdxs = nil
}
