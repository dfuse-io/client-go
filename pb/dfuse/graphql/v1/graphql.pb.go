// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: dfuse/graphql/v1/graphql.proto

package pbgraphql

import (
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query         string          `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Variables     *_struct.Struct `protobuf:"bytes,2,opt,name=variables,proto3" json:"variables,omitempty"`
	OperationName string          `protobuf:"bytes,3,opt,name=operationName,proto3" json:"operationName,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_graphql_v1_graphql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_graphql_v1_graphql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_dfuse_graphql_v1_graphql_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Request) GetVariables() *_struct.Struct {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *Request) GetOperationName() string {
	if x != nil {
		return x.OperationName
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Errors []*Error `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_graphql_v1_graphql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_graphql_v1_graphql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_dfuse_graphql_v1_graphql_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Response) GetErrors() []*Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

// GraphQL Error
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Description of the error intended for the developer.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// The source location for the error.
	Locations []*SourceLocation `protobuf:"bytes,2,rep,name=locations,proto3" json:"locations,omitempty"`
	// Path to the `null` value justified by this error.
	Path *_struct.ListValue `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// Free-form extensions (starts with a map)
	Extensions *_struct.Struct `protobuf:"bytes,4,opt,name=extensions,proto3" json:"extensions,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_graphql_v1_graphql_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_graphql_v1_graphql_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_dfuse_graphql_v1_graphql_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetLocations() []*SourceLocation {
	if x != nil {
		return x.Locations
	}
	return nil
}

func (x *Error) GetPath() *_struct.ListValue {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *Error) GetExtensions() *_struct.Struct {
	if x != nil {
		return x.Extensions
	}
	return nil
}

// The source location of an error.
type SourceLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The line the error occurred at.
	Line int32 `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	// The column the error occurred at.
	Column int32 `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
}

func (x *SourceLocation) Reset() {
	*x = SourceLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_graphql_v1_graphql_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceLocation) ProtoMessage() {}

func (x *SourceLocation) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_graphql_v1_graphql_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceLocation.ProtoReflect.Descriptor instead.
func (*SourceLocation) Descriptor() ([]byte, []int) {
	return file_dfuse_graphql_v1_graphql_proto_rawDescGZIP(), []int{3}
}

func (x *SourceLocation) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *SourceLocation) GetColumn() int32 {
	if x != nil {
		return x.Column
	}
	return 0
}

type BlockCursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ver      int32  `protobuf:"varint,1,opt,name=ver,proto3" json:"ver,omitempty"`
	BlockNum uint64 `protobuf:"varint,2,opt,name=blockNum,proto3" json:"blockNum,omitempty"`
	BlockId  string `protobuf:"bytes,3,opt,name=blockId,proto3" json:"blockId,omitempty"`
}

func (x *BlockCursor) Reset() {
	*x = BlockCursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_graphql_v1_graphql_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockCursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockCursor) ProtoMessage() {}

func (x *BlockCursor) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_graphql_v1_graphql_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockCursor.ProtoReflect.Descriptor instead.
func (*BlockCursor) Descriptor() ([]byte, []int) {
	return file_dfuse_graphql_v1_graphql_proto_rawDescGZIP(), []int{4}
}

func (x *BlockCursor) GetVer() int32 {
	if x != nil {
		return x.Ver
	}
	return 0
}

func (x *BlockCursor) GetBlockNum() uint64 {
	if x != nil {
		return x.BlockNum
	}
	return 0
}

func (x *BlockCursor) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

type TransactionCursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ver              int32  `protobuf:"varint,1,opt,name=ver,proto3" json:"ver,omitempty"`
	TransactionIndex uint32 `protobuf:"varint,2,opt,name=transactionIndex,proto3" json:"transactionIndex,omitempty"`
	TransactionHash  string `protobuf:"bytes,3,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
}

func (x *TransactionCursor) Reset() {
	*x = TransactionCursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_graphql_v1_graphql_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionCursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionCursor) ProtoMessage() {}

func (x *TransactionCursor) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_graphql_v1_graphql_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionCursor.ProtoReflect.Descriptor instead.
func (*TransactionCursor) Descriptor() ([]byte, []int) {
	return file_dfuse_graphql_v1_graphql_proto_rawDescGZIP(), []int{5}
}

func (x *TransactionCursor) GetVer() int32 {
	if x != nil {
		return x.Ver
	}
	return 0
}

func (x *TransactionCursor) GetTransactionIndex() uint32 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *TransactionCursor) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

var File_dfuse_graphql_v1_graphql_proto protoreflect.FileDescriptor

var file_dfuse_graphql_v1_graphql_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7c, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x35, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x09, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4f,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2f,
	0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22,
	0xca, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x37, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3c, 0x0a, 0x0e,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x22, 0x55, 0x0a, 0x0b, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x64, 0x22, 0x7b, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x76, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x32, 0x4f,
	0x0a, 0x07, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x12, 0x44, 0x0a, 0x07, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x66,
	0x75, 0x73, 0x65, 0x2d, 0x69, 0x6f, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x2f, 0x64, 0x66, 0x75, 0x73,
	0x65, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dfuse_graphql_v1_graphql_proto_rawDescOnce sync.Once
	file_dfuse_graphql_v1_graphql_proto_rawDescData = file_dfuse_graphql_v1_graphql_proto_rawDesc
)

func file_dfuse_graphql_v1_graphql_proto_rawDescGZIP() []byte {
	file_dfuse_graphql_v1_graphql_proto_rawDescOnce.Do(func() {
		file_dfuse_graphql_v1_graphql_proto_rawDescData = protoimpl.X.CompressGZIP(file_dfuse_graphql_v1_graphql_proto_rawDescData)
	})
	return file_dfuse_graphql_v1_graphql_proto_rawDescData
}

var file_dfuse_graphql_v1_graphql_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dfuse_graphql_v1_graphql_proto_goTypes = []interface{}{
	(*Request)(nil),           // 0: dfuse.graphql.v1.Request
	(*Response)(nil),          // 1: dfuse.graphql.v1.Response
	(*Error)(nil),             // 2: dfuse.graphql.v1.Error
	(*SourceLocation)(nil),    // 3: dfuse.graphql.v1.SourceLocation
	(*BlockCursor)(nil),       // 4: dfuse.graphql.v1.BlockCursor
	(*TransactionCursor)(nil), // 5: dfuse.graphql.v1.TransactionCursor
	(*_struct.Struct)(nil),    // 6: google.protobuf.Struct
	(*_struct.ListValue)(nil), // 7: google.protobuf.ListValue
}
var file_dfuse_graphql_v1_graphql_proto_depIdxs = []int32{
	6, // 0: dfuse.graphql.v1.Request.variables:type_name -> google.protobuf.Struct
	2, // 1: dfuse.graphql.v1.Response.errors:type_name -> dfuse.graphql.v1.Error
	3, // 2: dfuse.graphql.v1.Error.locations:type_name -> dfuse.graphql.v1.SourceLocation
	7, // 3: dfuse.graphql.v1.Error.path:type_name -> google.protobuf.ListValue
	6, // 4: dfuse.graphql.v1.Error.extensions:type_name -> google.protobuf.Struct
	0, // 5: dfuse.graphql.v1.GraphQL.Execute:input_type -> dfuse.graphql.v1.Request
	1, // 6: dfuse.graphql.v1.GraphQL.Execute:output_type -> dfuse.graphql.v1.Response
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_dfuse_graphql_v1_graphql_proto_init() }
func file_dfuse_graphql_v1_graphql_proto_init() {
	if File_dfuse_graphql_v1_graphql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dfuse_graphql_v1_graphql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_dfuse_graphql_v1_graphql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_dfuse_graphql_v1_graphql_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_dfuse_graphql_v1_graphql_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceLocation); i {
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
		file_dfuse_graphql_v1_graphql_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockCursor); i {
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
		file_dfuse_graphql_v1_graphql_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionCursor); i {
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
			RawDescriptor: file_dfuse_graphql_v1_graphql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dfuse_graphql_v1_graphql_proto_goTypes,
		DependencyIndexes: file_dfuse_graphql_v1_graphql_proto_depIdxs,
		MessageInfos:      file_dfuse_graphql_v1_graphql_proto_msgTypes,
	}.Build()
	File_dfuse_graphql_v1_graphql_proto = out.File
	file_dfuse_graphql_v1_graphql_proto_rawDesc = nil
	file_dfuse_graphql_v1_graphql_proto_goTypes = nil
	file_dfuse_graphql_v1_graphql_proto_depIdxs = nil
}