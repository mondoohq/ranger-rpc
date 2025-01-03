// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: swagger-example.proto

package swagger

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

type ListRequest_Corpus int32

const (
	ListRequest_UNIVERSAL ListRequest_Corpus = 0
	ListRequest_WEB       ListRequest_Corpus = 1
	ListRequest_IMAGES    ListRequest_Corpus = 2
	ListRequest_LOCAL     ListRequest_Corpus = 3
	ListRequest_NEWS      ListRequest_Corpus = 4
	ListRequest_PRODUCTS  ListRequest_Corpus = 5
	ListRequest_VIDEO     ListRequest_Corpus = 6
)

// Enum value maps for ListRequest_Corpus.
var (
	ListRequest_Corpus_name = map[int32]string{
		0: "UNIVERSAL",
		1: "WEB",
		2: "IMAGES",
		3: "LOCAL",
		4: "NEWS",
		5: "PRODUCTS",
		6: "VIDEO",
	}
	ListRequest_Corpus_value = map[string]int32{
		"UNIVERSAL": 0,
		"WEB":       1,
		"IMAGES":    2,
		"LOCAL":     3,
		"NEWS":      4,
		"PRODUCTS":  5,
		"VIDEO":     6,
	}
)

func (x ListRequest_Corpus) Enum() *ListRequest_Corpus {
	p := new(ListRequest_Corpus)
	*p = x
	return p
}

func (x ListRequest_Corpus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListRequest_Corpus) Descriptor() protoreflect.EnumDescriptor {
	return file_swagger_example_proto_enumTypes[0].Descriptor()
}

func (ListRequest_Corpus) Type() protoreflect.EnumType {
	return &file_swagger_example_proto_enumTypes[0]
}

func (x ListRequest_Corpus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListRequest_Corpus.Descriptor instead.
func (ListRequest_Corpus) EnumDescriptor() ([]byte, []int) {
	return file_swagger_example_proto_rawDescGZIP(), []int{6, 0}
}

// This is a status request
type StatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatsRequest) Reset() {
	*x = StatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swagger_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsRequest) ProtoMessage() {}

func (x *StatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_swagger_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsRequest.ProtoReflect.Descriptor instead.
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return file_swagger_example_proto_rawDescGZIP(), []int{0}
}

type StatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// number of items included in the store
	Items int64 `protobuf:"varint,1,opt,name=items,proto3" json:"items,omitempty"`
	// number of used items
	Used int64 `protobuf:"varint,2,opt,name=used,proto3" json:"used,omitempty"`
	// list of request failures
	Failures int64 `protobuf:"varint,3,opt,name=failures,proto3" json:"failures,omitempty"`
}

func (x *StatsResponse) Reset() {
	*x = StatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swagger_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsResponse) ProtoMessage() {}

func (x *StatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_swagger_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsResponse.ProtoReflect.Descriptor instead.
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return file_swagger_example_proto_rawDescGZIP(), []int{1}
}

func (x *StatsResponse) GetItems() int64 {
	if x != nil {
		return x.Items
	}
	return 0
}

func (x *StatsResponse) GetUsed() int64 {
	if x != nil {
		return x.Used
	}
	return 0
}

func (x *StatsResponse) GetFailures() int64 {
	if x != nil {
		return x.Failures
	}
	return 0
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	UserID  int32  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swagger_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_swagger_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_swagger_example_proto_rawDescGZIP(), []int{2}
}

func (x *AddRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *AddRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successful bool `protobuf:"varint,1,opt,name=successful,proto3" json:"successful,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swagger_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_swagger_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_swagger_example_proto_rawDescGZIP(), []int{3}
}

func (x *AddResponse) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetURL string `protobuf:"bytes,1,opt,name=targetURL,proto3" json:"targetURL,omitempty"`
	Payload   string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swagger_example_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_swagger_example_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_swagger_example_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetTargetURL() string {
	if x != nil {
		return x.TargetURL
	}
	return ""
}

func (x *DeleteRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successful bool `protobuf:"varint,1,opt,name=successful,proto3" json:"successful,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swagger_example_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_swagger_example_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_swagger_example_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteResponse) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query         string             `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber    int32              `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage int32              `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	Corpus        ListRequest_Corpus `protobuf:"varint,4,opt,name=corpus,proto3,enum=swagger.ListRequest_Corpus" json:"corpus,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swagger_example_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_swagger_example_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_swagger_example_proto_rawDescGZIP(), []int{6}
}

func (x *ListRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListRequest) GetResultPerPage() int32 {
	if x != nil {
		return x.ResultPerPage
	}
	return 0
}

func (x *ListRequest) GetCorpus() ListRequest_Corpus {
	if x != nil {
		return x.Corpus
	}
	return ListRequest_UNIVERSAL
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Result        `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Users   map[string]*User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swagger_example_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_swagger_example_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_swagger_example_proto_rawDescGZIP(), []int{7}
}

func (x *ListResponse) GetResults() []*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListResponse) GetUsers() map[string]*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Snippets []string `protobuf:"bytes,3,rep,name=snippets,proto3" json:"snippets,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swagger_example_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_swagger_example_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_swagger_example_proto_rawDescGZIP(), []int{8}
}

func (x *Result) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Result) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Result) GetSnippets() []string {
	if x != nil {
		return x.Snippets
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  int32  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swagger_example_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_swagger_example_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_swagger_example_proto_rawDescGZIP(), []int{9}
}

func (x *User) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_swagger_example_proto protoreflect.FileDescriptor

var file_swagger_example_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x22, 0x34, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x0f, 0x10, 0x10, 0x4a, 0x04, 0x08, 0x09,
	0x10, 0x0c, 0x4a, 0x08, 0x08, 0x28, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x52, 0x03, 0x46, 0x4f,
	0x4f, 0x52, 0x03, 0x42, 0x41, 0x52, 0x22, 0x55, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x22, 0x3e, 0x0a,
	0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2d, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x22, 0x47, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x30, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x22, 0xfd, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x0f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50,
	0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x72,
	0x70, 0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x22, 0x5a, 0x0a, 0x06, 0x43,
	0x6f, 0x72, 0x70, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x49, 0x56, 0x45, 0x52, 0x53,
	0x41, 0x4c, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x53, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x43,
	0x41, 0x4c, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x45, 0x57, 0x53, 0x10, 0x04, 0x12, 0x0c,
	0x0a, 0x08, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x53, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05,
	0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x06, 0x22, 0xba, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x47, 0x0a, 0x0a, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x77, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x4c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65,
	0x74, 0x73, 0x22, 0x2c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x32, 0xef, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x13, 0x2e, 0x73, 0x77, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x16, 0x2e, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x33, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x77, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x12, 0x15, 0x2e, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x77, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x6f, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_swagger_example_proto_rawDescOnce sync.Once
	file_swagger_example_proto_rawDescData = file_swagger_example_proto_rawDesc
)

func file_swagger_example_proto_rawDescGZIP() []byte {
	file_swagger_example_proto_rawDescOnce.Do(func() {
		file_swagger_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_swagger_example_proto_rawDescData)
	})
	return file_swagger_example_proto_rawDescData
}

var file_swagger_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_swagger_example_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_swagger_example_proto_goTypes = []any{
	(ListRequest_Corpus)(0), // 0: swagger.ListRequest.Corpus
	(*StatsRequest)(nil),    // 1: swagger.StatsRequest
	(*StatsResponse)(nil),   // 2: swagger.StatsResponse
	(*AddRequest)(nil),      // 3: swagger.AddRequest
	(*AddResponse)(nil),     // 4: swagger.AddResponse
	(*DeleteRequest)(nil),   // 5: swagger.DeleteRequest
	(*DeleteResponse)(nil),  // 6: swagger.DeleteResponse
	(*ListRequest)(nil),     // 7: swagger.ListRequest
	(*ListResponse)(nil),    // 8: swagger.ListResponse
	(*Result)(nil),          // 9: swagger.Result
	(*User)(nil),            // 10: swagger.User
	nil,                     // 11: swagger.ListResponse.UsersEntry
}
var file_swagger_example_proto_depIdxs = []int32{
	0,  // 0: swagger.ListRequest.corpus:type_name -> swagger.ListRequest.Corpus
	9,  // 1: swagger.ListResponse.results:type_name -> swagger.Result
	11, // 2: swagger.ListResponse.users:type_name -> swagger.ListResponse.UsersEntry
	10, // 3: swagger.ListResponse.UsersEntry.value:type_name -> swagger.User
	3,  // 4: swagger.ExampleService.Add:input_type -> swagger.AddRequest
	5,  // 5: swagger.ExampleService.Delete:input_type -> swagger.DeleteRequest
	7,  // 6: swagger.ExampleService.List:input_type -> swagger.ListRequest
	1,  // 7: swagger.ExampleService.Statistics:input_type -> swagger.StatsRequest
	4,  // 8: swagger.ExampleService.Add:output_type -> swagger.AddResponse
	6,  // 9: swagger.ExampleService.Delete:output_type -> swagger.DeleteResponse
	8,  // 10: swagger.ExampleService.List:output_type -> swagger.ListResponse
	2,  // 11: swagger.ExampleService.Statistics:output_type -> swagger.StatsResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_swagger_example_proto_init() }
func file_swagger_example_proto_init() {
	if File_swagger_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_swagger_example_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StatsRequest); i {
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
		file_swagger_example_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StatsResponse); i {
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
		file_swagger_example_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AddRequest); i {
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
		file_swagger_example_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AddResponse); i {
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
		file_swagger_example_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRequest); i {
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
		file_swagger_example_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteResponse); i {
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
		file_swagger_example_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListRequest); i {
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
		file_swagger_example_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListResponse); i {
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
		file_swagger_example_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Result); i {
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
		file_swagger_example_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
			RawDescriptor: file_swagger_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_swagger_example_proto_goTypes,
		DependencyIndexes: file_swagger_example_proto_depIdxs,
		EnumInfos:         file_swagger_example_proto_enumTypes,
		MessageInfos:      file_swagger_example_proto_msgTypes,
	}.Build()
	File_swagger_example_proto = out.File
	file_swagger_example_proto_rawDesc = nil
	file_swagger_example_proto_goTypes = nil
	file_swagger_example_proto_depIdxs = nil
}
