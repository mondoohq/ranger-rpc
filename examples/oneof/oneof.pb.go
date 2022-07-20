// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: oneof.proto

package oneof

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

type OneOfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Options:
	//	*OneOfRequest_Text
	//	*OneOfRequest_Number
	Options isOneOfRequest_Options `protobuf_oneof:"options"`
}

func (x *OneOfRequest) Reset() {
	*x = OneOfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneOfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneOfRequest) ProtoMessage() {}

func (x *OneOfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneOfRequest.ProtoReflect.Descriptor instead.
func (*OneOfRequest) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{0}
}

func (m *OneOfRequest) GetOptions() isOneOfRequest_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (x *OneOfRequest) GetText() string {
	if x, ok := x.GetOptions().(*OneOfRequest_Text); ok {
		return x.Text
	}
	return ""
}

func (x *OneOfRequest) GetNumber() int64 {
	if x, ok := x.GetOptions().(*OneOfRequest_Number); ok {
		return x.Number
	}
	return 0
}

type isOneOfRequest_Options interface {
	isOneOfRequest_Options()
}

type OneOfRequest_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type OneOfRequest_Number struct {
	Number int64 `protobuf:"varint,2,opt,name=number,proto3,oneof"`
}

func (*OneOfRequest_Text) isOneOfRequest_Options() {}

func (*OneOfRequest_Number) isOneOfRequest_Options() {}

type OneOfReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Options:
	//	*OneOfReply_Text
	//	*OneOfReply_Number
	Options isOneOfReply_Options `protobuf_oneof:"options"`
}

func (x *OneOfReply) Reset() {
	*x = OneOfReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneOfReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneOfReply) ProtoMessage() {}

func (x *OneOfReply) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneOfReply.ProtoReflect.Descriptor instead.
func (*OneOfReply) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{1}
}

func (m *OneOfReply) GetOptions() isOneOfReply_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (x *OneOfReply) GetText() string {
	if x, ok := x.GetOptions().(*OneOfReply_Text); ok {
		return x.Text
	}
	return ""
}

func (x *OneOfReply) GetNumber() int64 {
	if x, ok := x.GetOptions().(*OneOfReply_Number); ok {
		return x.Number
	}
	return 0
}

type isOneOfReply_Options interface {
	isOneOfReply_Options()
}

type OneOfReply_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type OneOfReply_Number struct {
	Number int64 `protobuf:"varint,2,opt,name=number,proto3,oneof"`
}

func (*OneOfReply_Text) isOneOfReply_Options() {}

func (*OneOfReply_Number) isOneOfReply_Options() {}

var File_oneof_proto protoreflect.FileDescriptor

var file_oneof_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x22, 0x49, 0x0a, 0x0c, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x47, 0x0a, 0x0a, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x09, 0x0a,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x37, 0x0a, 0x05, 0x4f, 0x6e, 0x65, 0x4f,
	0x66, 0x12, 0x2e, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x13, 0x2e, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x6f, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_oneof_proto_rawDescOnce sync.Once
	file_oneof_proto_rawDescData = file_oneof_proto_rawDesc
)

func file_oneof_proto_rawDescGZIP() []byte {
	file_oneof_proto_rawDescOnce.Do(func() {
		file_oneof_proto_rawDescData = protoimpl.X.CompressGZIP(file_oneof_proto_rawDescData)
	})
	return file_oneof_proto_rawDescData
}

var file_oneof_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_oneof_proto_goTypes = []interface{}{
	(*OneOfRequest)(nil), // 0: oneof.OneOfRequest
	(*OneOfReply)(nil),   // 1: oneof.OneOfReply
}
var file_oneof_proto_depIdxs = []int32{
	0, // 0: oneof.OneOf.Echo:input_type -> oneof.OneOfRequest
	1, // 1: oneof.OneOf.Echo:output_type -> oneof.OneOfReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oneof_proto_init() }
func file_oneof_proto_init() {
	if File_oneof_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oneof_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneOfRequest); i {
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
		file_oneof_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneOfReply); i {
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
	file_oneof_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OneOfRequest_Text)(nil),
		(*OneOfRequest_Number)(nil),
	}
	file_oneof_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*OneOfReply_Text)(nil),
		(*OneOfReply_Number)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oneof_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oneof_proto_goTypes,
		DependencyIndexes: file_oneof_proto_depIdxs,
		MessageInfos:      file_oneof_proto_msgTypes,
	}.Build()
	File_oneof_proto = out.File
	file_oneof_proto_rawDesc = nil
	file_oneof_proto_goTypes = nil
	file_oneof_proto_depIdxs = nil
}
