// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: hello.proto

package rangerguard

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

type HelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *HelloReq) Reset() {
	*x = HelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq) ProtoMessage() {}

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq.ProtoReflect.Descriptor instead.
func (*HelloReq) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type HelloResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *HelloResp) Reset() {
	*x = HelloResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResp) ProtoMessage() {}

func (x *HelloResp) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResp.ProtoReflect.Descriptor instead.
func (*HelloResp) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResp) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[2]
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
	return file_hello_proto_rawDescGZIP(), []int{2}
}

type Tags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags map[string]string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Tags) Reset() {
	*x = Tags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tags) ProtoMessage() {}

func (x *Tags) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tags.ProtoReflect.Descriptor instead.
func (*Tags) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{3}
}

func (x *Tags) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_hello_proto protoreflect.FileDescriptor

var file_hello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x69,
	0x6f, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x6e, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x22, 0x24, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x1f, 0x0a, 0x09, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x88, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x47, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x69, 0x6f, 0x2e, 0x6d,
	0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x54, 0x61, 0x67, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xd3, 0x01,
	0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x66, 0x0a, 0x05,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x2d, 0x2e, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f,
	0x6f, 0x2e, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x1a, 0x2e, 0x2e, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f,
	0x2e, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x5d, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x2e, 0x69,
	0x6f, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x6e, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x29, 0x2e, 0x69, 0x6f, 0x2e, 0x6d, 0x6f,
	0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x66, 0x61, 0x6c, 0x63, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x54,
	0x61, 0x67, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x6f, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2d, 0x72, 0x70, 0x63, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x67,
	0x75, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_proto_rawDescOnce sync.Once
	file_hello_proto_rawDescData = file_hello_proto_rawDesc
)

func file_hello_proto_rawDescGZIP() []byte {
	file_hello_proto_rawDescOnce.Do(func() {
		file_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_proto_rawDescData)
	})
	return file_hello_proto_rawDescData
}

var file_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_hello_proto_goTypes = []any{
	(*HelloReq)(nil),  // 0: io.mondoo.falcon.example.helloworld.HelloReq
	(*HelloResp)(nil), // 1: io.mondoo.falcon.example.helloworld.HelloResp
	(*Empty)(nil),     // 2: io.mondoo.falcon.example.helloworld.Empty
	(*Tags)(nil),      // 3: io.mondoo.falcon.example.helloworld.Tags
	nil,               // 4: io.mondoo.falcon.example.helloworld.Tags.TagsEntry
}
var file_hello_proto_depIdxs = []int32{
	4, // 0: io.mondoo.falcon.example.helloworld.Tags.tags:type_name -> io.mondoo.falcon.example.helloworld.Tags.TagsEntry
	0, // 1: io.mondoo.falcon.example.helloworld.HelloWorld.Hello:input_type -> io.mondoo.falcon.example.helloworld.HelloReq
	2, // 2: io.mondoo.falcon.example.helloworld.HelloWorld.Info:input_type -> io.mondoo.falcon.example.helloworld.Empty
	1, // 3: io.mondoo.falcon.example.helloworld.HelloWorld.Hello:output_type -> io.mondoo.falcon.example.helloworld.HelloResp
	3, // 4: io.mondoo.falcon.example.helloworld.HelloWorld.Info:output_type -> io.mondoo.falcon.example.helloworld.Tags
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hello_proto_init() }
func file_hello_proto_init() {
	if File_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HelloReq); i {
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
		file_hello_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HelloResp); i {
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
		file_hello_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_hello_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Tags); i {
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
			RawDescriptor: file_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_proto_goTypes,
		DependencyIndexes: file_hello_proto_depIdxs,
		MessageInfos:      file_hello_proto_msgTypes,
	}.Build()
	File_hello_proto = out.File
	file_hello_proto_rawDesc = nil
	file_hello_proto_goTypes = nil
	file_hello_proto_depIdxs = nil
}
