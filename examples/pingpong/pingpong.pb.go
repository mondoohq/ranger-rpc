// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: pingpong.proto

package pingpong

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pingpong_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pingpong_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pingpong_proto_rawDescGZIP(), []int{0}
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pingpong_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pingpong_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_pingpong_proto_rawDescGZIP(), []int{1}
}

func (x *PingRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

type PongReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PongReply) Reset() {
	*x = PongReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pingpong_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongReply) ProtoMessage() {}

func (x *PongReply) ProtoReflect() protoreflect.Message {
	mi := &file_pingpong_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongReply.ProtoReflect.Descriptor instead.
func (*PongReply) Descriptor() ([]byte, []int) {
	return file_pingpong_proto_rawDescGZIP(), []int{2}
}

func (x *PongReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pingpong_proto protoreflect.FileDescriptor

var file_pingpong_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x25, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x09, 0x50, 0x6f,
	0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x6e, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x32, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
	0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2e, 0x0a, 0x06, 0x4e, 0x6f, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x70, 0x69,
	0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x70,
	0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x6f, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pingpong_proto_rawDescOnce sync.Once
	file_pingpong_proto_rawDescData = file_pingpong_proto_rawDesc
)

func file_pingpong_proto_rawDescGZIP() []byte {
	file_pingpong_proto_rawDescOnce.Do(func() {
		file_pingpong_proto_rawDescData = protoimpl.X.CompressGZIP(file_pingpong_proto_rawDescData)
	})
	return file_pingpong_proto_rawDescData
}

var file_pingpong_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pingpong_proto_goTypes = []interface{}{
	(*Empty)(nil),       // 0: pingpong.Empty
	(*PingRequest)(nil), // 1: pingpong.PingRequest
	(*PongReply)(nil),   // 2: pingpong.PongReply
}
var file_pingpong_proto_depIdxs = []int32{
	1, // 0: pingpong.PingPong.Ping:input_type -> pingpong.PingRequest
	0, // 1: pingpong.PingPong.NoPing:input_type -> pingpong.Empty
	2, // 2: pingpong.PingPong.Ping:output_type -> pingpong.PongReply
	2, // 3: pingpong.PingPong.NoPing:output_type -> pingpong.PongReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pingpong_proto_init() }
func file_pingpong_proto_init() {
	if File_pingpong_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pingpong_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_pingpong_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_pingpong_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongReply); i {
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
			RawDescriptor: file_pingpong_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pingpong_proto_goTypes,
		DependencyIndexes: file_pingpong_proto_depIdxs,
		MessageInfos:      file_pingpong_proto_msgTypes,
	}.Build()
	File_pingpong_proto = out.File
	file_pingpong_proto_rawDesc = nil
	file_pingpong_proto_goTypes = nil
	file_pingpong_proto_depIdxs = nil
}
