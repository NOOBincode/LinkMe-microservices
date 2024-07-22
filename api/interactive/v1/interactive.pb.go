// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: api/interactive/v1/interactive.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetInteractiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int64 `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`
}

func (x *GetInteractiveRequest) Reset() {
	*x = GetInteractiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_interactive_v1_interactive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInteractiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInteractiveRequest) ProtoMessage() {}

func (x *GetInteractiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_interactive_v1_interactive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInteractiveRequest.ProtoReflect.Descriptor instead.
func (*GetInteractiveRequest) Descriptor() ([]byte, []int) {
	return file_api_interactive_v1_interactive_proto_rawDescGZIP(), []int{0}
}

func (x *GetInteractiveRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type GetInteractiveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32                 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *GetOrListInteractive `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetInteractiveReply) Reset() {
	*x = GetInteractiveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_interactive_v1_interactive_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInteractiveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInteractiveReply) ProtoMessage() {}

func (x *GetInteractiveReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_interactive_v1_interactive_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInteractiveReply.ProtoReflect.Descriptor instead.
func (*GetInteractiveReply) Descriptor() ([]byte, []int) {
	return file_api_interactive_v1_interactive_proto_rawDescGZIP(), []int{1}
}

func (x *GetInteractiveReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetInteractiveReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetInteractiveReply) GetData() *GetOrListInteractive {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListInteractiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListInteractiveRequest) Reset() {
	*x = ListInteractiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_interactive_v1_interactive_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInteractiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInteractiveRequest) ProtoMessage() {}

func (x *ListInteractiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_interactive_v1_interactive_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInteractiveRequest.ProtoReflect.Descriptor instead.
func (*ListInteractiveRequest) Descriptor() ([]byte, []int) {
	return file_api_interactive_v1_interactive_proto_rawDescGZIP(), []int{2}
}

func (x *ListInteractiveRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListInteractiveRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListInteractiveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32                   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []*GetOrListInteractive `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListInteractiveReply) Reset() {
	*x = ListInteractiveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_interactive_v1_interactive_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInteractiveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInteractiveReply) ProtoMessage() {}

func (x *ListInteractiveReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_interactive_v1_interactive_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInteractiveReply.ProtoReflect.Descriptor instead.
func (*ListInteractiveReply) Descriptor() ([]byte, []int) {
	return file_api_interactive_v1_interactive_proto_rawDescGZIP(), []int{3}
}

func (x *ListInteractiveReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListInteractiveReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ListInteractiveReply) GetData() []*GetOrListInteractive {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId  int64  `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`
	BizName string `protobuf:"bytes,2,opt,name=bizName,proto3" json:"bizName,omitempty"`
}

func (x *AddCountRequest) Reset() {
	*x = AddCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_interactive_v1_interactive_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCountRequest) ProtoMessage() {}

func (x *AddCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_interactive_v1_interactive_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCountRequest.ProtoReflect.Descriptor instead.
func (*AddCountRequest) Descriptor() ([]byte, []int) {
	return file_api_interactive_v1_interactive_proto_rawDescGZIP(), []int{4}
}

func (x *AddCountRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *AddCountRequest) GetBizName() string {
	if x != nil {
		return x.BizName
	}
	return ""
}

type AddCountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AddCountReply) Reset() {
	*x = AddCountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_interactive_v1_interactive_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCountReply) ProtoMessage() {}

func (x *AddCountReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_interactive_v1_interactive_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCountReply.ProtoReflect.Descriptor instead.
func (*AddCountReply) Descriptor() ([]byte, []int) {
	return file_api_interactive_v1_interactive_proto_rawDescGZIP(), []int{5}
}

func (x *AddCountReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AddCountReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetOrListInteractive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	BizID        int64  `protobuf:"varint,2,opt,name=BizID,proto3" json:"BizID,omitempty"`
	BizName      string `protobuf:"bytes,3,opt,name=BizName,proto3" json:"BizName,omitempty"`
	ReadCount    int64  `protobuf:"varint,4,opt,name=ReadCount,proto3" json:"ReadCount,omitempty"`
	LikeCount    int64  `protobuf:"varint,5,opt,name=LikeCount,proto3" json:"LikeCount,omitempty"`
	CollectCount int64  `protobuf:"varint,6,opt,name=CollectCount,proto3" json:"CollectCount,omitempty"`
	UpdateTime   int64  `protobuf:"varint,7,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	CreateTime   int64  `protobuf:"varint,8,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	PostId       int64  `protobuf:"varint,9,opt,name=PostId,proto3" json:"PostId,omitempty"`
}

func (x *GetOrListInteractive) Reset() {
	*x = GetOrListInteractive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_interactive_v1_interactive_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrListInteractive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrListInteractive) ProtoMessage() {}

func (x *GetOrListInteractive) ProtoReflect() protoreflect.Message {
	mi := &file_api_interactive_v1_interactive_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrListInteractive.ProtoReflect.Descriptor instead.
func (*GetOrListInteractive) Descriptor() ([]byte, []int) {
	return file_api_interactive_v1_interactive_proto_rawDescGZIP(), []int{6}
}

func (x *GetOrListInteractive) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GetOrListInteractive) GetBizID() int64 {
	if x != nil {
		return x.BizID
	}
	return 0
}

func (x *GetOrListInteractive) GetBizName() string {
	if x != nil {
		return x.BizName
	}
	return ""
}

func (x *GetOrListInteractive) GetReadCount() int64 {
	if x != nil {
		return x.ReadCount
	}
	return 0
}

func (x *GetOrListInteractive) GetLikeCount() int64 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *GetOrListInteractive) GetCollectCount() int64 {
	if x != nil {
		return x.CollectCount
	}
	return 0
}

func (x *GetOrListInteractive) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *GetOrListInteractive) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *GetOrListInteractive) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

var File_api_interactive_v1_interactive_proto protoreflect.FileDescriptor

var file_api_interactive_v1_interactive_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x3c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x40, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x7a, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x3c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x43, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x69, 0x7a, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x69, 0x7a, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x8e,
	0x02, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x69, 0x7a, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x42, 0x69, 0x7a, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x42, 0x69, 0x7a, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x42, 0x69, 0x7a, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x52, 0x65, 0x61, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x32,
	0x90, 0x04, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x7b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f,
	0x67, 0x65, 0x74, 0x2f, 0x7b, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x7d, 0x12, 0x79, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x3a, 0x01, 0x2a,
	0x22, 0x05, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x56, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x59, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x47, 0x6f, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x4c, 0x69,
	0x6e, 0x6b, 0x4d, 0x65, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_interactive_v1_interactive_proto_rawDescOnce sync.Once
	file_api_interactive_v1_interactive_proto_rawDescData = file_api_interactive_v1_interactive_proto_rawDesc
)

func file_api_interactive_v1_interactive_proto_rawDescGZIP() []byte {
	file_api_interactive_v1_interactive_proto_rawDescOnce.Do(func() {
		file_api_interactive_v1_interactive_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_interactive_v1_interactive_proto_rawDescData)
	})
	return file_api_interactive_v1_interactive_proto_rawDescData
}

var file_api_interactive_v1_interactive_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_interactive_v1_interactive_proto_goTypes = []any{
	(*GetInteractiveRequest)(nil),  // 0: api.interactive.v1.GetInteractiveRequest
	(*GetInteractiveReply)(nil),    // 1: api.interactive.v1.GetInteractiveReply
	(*ListInteractiveRequest)(nil), // 2: api.interactive.v1.ListInteractiveRequest
	(*ListInteractiveReply)(nil),   // 3: api.interactive.v1.ListInteractiveReply
	(*AddCountRequest)(nil),        // 4: api.interactive.v1.AddCountRequest
	(*AddCountReply)(nil),          // 5: api.interactive.v1.AddCountReply
	(*GetOrListInteractive)(nil),   // 6: api.interactive.v1.GetOrListInteractive
}
var file_api_interactive_v1_interactive_proto_depIdxs = []int32{
	6, // 0: api.interactive.v1.GetInteractiveReply.data:type_name -> api.interactive.v1.GetOrListInteractive
	6, // 1: api.interactive.v1.ListInteractiveReply.data:type_name -> api.interactive.v1.GetOrListInteractive
	0, // 2: api.interactive.v1.Interactive.GetInteractive:input_type -> api.interactive.v1.GetInteractiveRequest
	2, // 3: api.interactive.v1.Interactive.ListInteractive:input_type -> api.interactive.v1.ListInteractiveRequest
	4, // 4: api.interactive.v1.Interactive.AddReadCount:input_type -> api.interactive.v1.AddCountRequest
	4, // 5: api.interactive.v1.Interactive.AddLikeCount:input_type -> api.interactive.v1.AddCountRequest
	4, // 6: api.interactive.v1.Interactive.AddCollectCount:input_type -> api.interactive.v1.AddCountRequest
	1, // 7: api.interactive.v1.Interactive.GetInteractive:output_type -> api.interactive.v1.GetInteractiveReply
	3, // 8: api.interactive.v1.Interactive.ListInteractive:output_type -> api.interactive.v1.ListInteractiveReply
	5, // 9: api.interactive.v1.Interactive.AddReadCount:output_type -> api.interactive.v1.AddCountReply
	5, // 10: api.interactive.v1.Interactive.AddLikeCount:output_type -> api.interactive.v1.AddCountReply
	5, // 11: api.interactive.v1.Interactive.AddCollectCount:output_type -> api.interactive.v1.AddCountReply
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_interactive_v1_interactive_proto_init() }
func file_api_interactive_v1_interactive_proto_init() {
	if File_api_interactive_v1_interactive_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_interactive_v1_interactive_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetInteractiveRequest); i {
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
		file_api_interactive_v1_interactive_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetInteractiveReply); i {
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
		file_api_interactive_v1_interactive_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListInteractiveRequest); i {
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
		file_api_interactive_v1_interactive_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListInteractiveReply); i {
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
		file_api_interactive_v1_interactive_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AddCountRequest); i {
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
		file_api_interactive_v1_interactive_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AddCountReply); i {
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
		file_api_interactive_v1_interactive_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetOrListInteractive); i {
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
			RawDescriptor: file_api_interactive_v1_interactive_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_interactive_v1_interactive_proto_goTypes,
		DependencyIndexes: file_api_interactive_v1_interactive_proto_depIdxs,
		MessageInfos:      file_api_interactive_v1_interactive_proto_msgTypes,
	}.Build()
	File_api_interactive_v1_interactive_proto = out.File
	file_api_interactive_v1_interactive_proto_rawDesc = nil
	file_api_interactive_v1_interactive_proto_goTypes = nil
	file_api_interactive_v1_interactive_proto_depIdxs = nil
}
