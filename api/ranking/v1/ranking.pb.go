// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: api/ranking/v1/ranking.proto

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

// TopN 请求消息
type TopNRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"` // 要获取的排名数量
}

func (x *TopNRequest) Reset() {
	*x = TopNRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ranking_v1_ranking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopNRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopNRequest) ProtoMessage() {}

func (x *TopNRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ranking_v1_ranking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopNRequest.ProtoReflect.Descriptor instead.
func (*TopNRequest) Descriptor() ([]byte, []int) {
	return file_api_ranking_v1_ranking_proto_rawDescGZIP(), []int{0}
}

func (x *TopNRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

// TopN 响应消息
type TopNReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranking []*Rank `protobuf:"bytes,1,rep,name=ranking,proto3" json:"ranking,omitempty"` // 排名列表
}

func (x *TopNReply) Reset() {
	*x = TopNReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ranking_v1_ranking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopNReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopNReply) ProtoMessage() {}

func (x *TopNReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_ranking_v1_ranking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopNReply.ProtoReflect.Descriptor instead.
func (*TopNReply) Descriptor() ([]byte, []int) {
	return file_api_ranking_v1_ranking_proto_rawDescGZIP(), []int{1}
}

func (x *TopNReply) GetRanking() []*Rank {
	if x != nil {
		return x.Ranking
	}
	return nil
}

// CreateRanking 请求消息
type CreateRankingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranking *Rank `protobuf:"bytes,1,opt,name=ranking,proto3" json:"ranking,omitempty"` // 要创建的排名信息
}

func (x *CreateRankingRequest) Reset() {
	*x = CreateRankingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ranking_v1_ranking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRankingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRankingRequest) ProtoMessage() {}

func (x *CreateRankingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ranking_v1_ranking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRankingRequest.ProtoReflect.Descriptor instead.
func (*CreateRankingRequest) Descriptor() ([]byte, []int) {
	return file_api_ranking_v1_ranking_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRankingRequest) GetRanking() *Rank {
	if x != nil {
		return x.Ranking
	}
	return nil
}

// CreateRanking 响应消息
type CreateRankingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranking *Rank `protobuf:"bytes,1,opt,name=ranking,proto3" json:"ranking,omitempty"` // 创建的排名信息
}

func (x *CreateRankingReply) Reset() {
	*x = CreateRankingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ranking_v1_ranking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRankingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRankingReply) ProtoMessage() {}

func (x *CreateRankingReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_ranking_v1_ranking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRankingReply.ProtoReflect.Descriptor instead.
func (*CreateRankingReply) Descriptor() ([]byte, []int) {
	return file_api_ranking_v1_ranking_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRankingReply) GetRanking() *Rank {
	if x != nil {
		return x.Ranking
	}
	return nil
}

// UpdateRanking 请求消息
type UpdateRankingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ranking *Rank `protobuf:"bytes,2,opt,name=ranking,proto3" json:"ranking,omitempty"` // 要更新的排名信息
}

func (x *UpdateRankingRequest) Reset() {
	*x = UpdateRankingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ranking_v1_ranking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRankingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRankingRequest) ProtoMessage() {}

func (x *UpdateRankingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ranking_v1_ranking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRankingRequest.ProtoReflect.Descriptor instead.
func (*UpdateRankingRequest) Descriptor() ([]byte, []int) {
	return file_api_ranking_v1_ranking_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRankingRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRankingRequest) GetRanking() *Rank {
	if x != nil {
		return x.Ranking
	}
	return nil
}

// UpdateRanking 响应消息
type UpdateRankingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranking *Rank `protobuf:"bytes,1,opt,name=ranking,proto3" json:"ranking,omitempty"` // 更新后的排名信息
}

func (x *UpdateRankingReply) Reset() {
	*x = UpdateRankingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ranking_v1_ranking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRankingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRankingReply) ProtoMessage() {}

func (x *UpdateRankingReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_ranking_v1_ranking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRankingReply.ProtoReflect.Descriptor instead.
func (*UpdateRankingReply) Descriptor() ([]byte, []int) {
	return file_api_ranking_v1_ranking_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRankingReply) GetRanking() *Rank {
	if x != nil {
		return x.Ranking
	}
	return nil
}

// DeleteRanking 请求消息
type DeleteRankingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 要删除的排名 ID
}

func (x *DeleteRankingRequest) Reset() {
	*x = DeleteRankingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ranking_v1_ranking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRankingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRankingRequest) ProtoMessage() {}

func (x *DeleteRankingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ranking_v1_ranking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRankingRequest.ProtoReflect.Descriptor instead.
func (*DeleteRankingRequest) Descriptor() ([]byte, []int) {
	return file_api_ranking_v1_ranking_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRankingRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// DeleteRanking 响应消息
type DeleteRankingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRankingReply) Reset() {
	*x = DeleteRankingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ranking_v1_ranking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRankingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRankingReply) ProtoMessage() {}

func (x *DeleteRankingReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_ranking_v1_ranking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRankingReply.ProtoReflect.Descriptor instead.
func (*DeleteRankingReply) Descriptor() ([]byte, []int) {
	return file_api_ranking_v1_ranking_proto_rawDescGZIP(), []int{7}
}

// GetRanking 请求消息
type GetRankingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 要获取的排名 ID
}

func (x *GetRankingRequest) Reset() {
	*x = GetRankingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ranking_v1_ranking_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankingRequest) ProtoMessage() {}

func (x *GetRankingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ranking_v1_ranking_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankingRequest.ProtoReflect.Descriptor instead.
func (*GetRankingRequest) Descriptor() ([]byte, []int) {
	return file_api_ranking_v1_ranking_proto_rawDescGZIP(), []int{8}
}

func (x *GetRankingRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// GetRanking 响应消息
type GetRankingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranking *Rank `protobuf:"bytes,1,opt,name=ranking,proto3" json:"ranking,omitempty"` // 获取的排名信息
}

func (x *GetRankingReply) Reset() {
	*x = GetRankingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ranking_v1_ranking_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankingReply) ProtoMessage() {}

func (x *GetRankingReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_ranking_v1_ranking_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankingReply.ProtoReflect.Descriptor instead.
func (*GetRankingReply) Descriptor() ([]byte, []int) {
	return file_api_ranking_v1_ranking_proto_rawDescGZIP(), []int{9}
}

func (x *GetRankingReply) GetRanking() *Rank {
	if x != nil {
		return x.Ranking
	}
	return nil
}

// ListRanking 请求消息
type ListRankingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"` // 页码
	Size int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"` // 每页大小
}

func (x *ListRankingRequest) Reset() {
	*x = ListRankingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ranking_v1_ranking_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRankingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRankingRequest) ProtoMessage() {}

func (x *ListRankingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ranking_v1_ranking_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRankingRequest.ProtoReflect.Descriptor instead.
func (*ListRankingRequest) Descriptor() ([]byte, []int) {
	return file_api_ranking_v1_ranking_proto_rawDescGZIP(), []int{10}
}

func (x *ListRankingRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRankingRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

// ListRanking 响应消息
type ListRankingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rankings []*Rank `protobuf:"bytes,1,rep,name=rankings,proto3" json:"rankings,omitempty"` // 排名列表
}

func (x *ListRankingReply) Reset() {
	*x = ListRankingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ranking_v1_ranking_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRankingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRankingReply) ProtoMessage() {}

func (x *ListRankingReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_ranking_v1_ranking_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRankingReply.ProtoReflect.Descriptor instead.
func (*ListRankingReply) Descriptor() ([]byte, []int) {
	return file_api_ranking_v1_ranking_proto_rawDescGZIP(), []int{11}
}

func (x *ListRankingReply) GetRankings() []*Rank {
	if x != nil {
		return x.Rankings
	}
	return nil
}

// Rank 消息
type Rank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                               // 排名 ID
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`                          // 排名标题
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`              // 排名描述
	Score       int32    `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`                         // 排名分数
	Category    string   `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`                    // 排名类别
	Tags        []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`                            // 排名标签
	CreatedAt   string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // 创建时间
	UpdatedAt   string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // 更新时间
}

func (x *Rank) Reset() {
	*x = Rank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ranking_v1_ranking_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rank) ProtoMessage() {}

func (x *Rank) ProtoReflect() protoreflect.Message {
	mi := &file_api_ranking_v1_ranking_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rank.ProtoReflect.Descriptor instead.
func (*Rank) Descriptor() ([]byte, []int) {
	return file_api_ranking_v1_ranking_proto_rawDescGZIP(), []int{12}
}

func (x *Rank) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Rank) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Rank) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Rank) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Rank) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Rank) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Rank) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Rank) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_api_ranking_v1_ranking_proto protoreflect.FileDescriptor

var file_api_ranking_v1_ranking_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0b,
	0x54, 0x6f, 0x70, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x09, 0x54, 0x6f, 0x70, 0x4e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x2e, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x22, 0x46, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52,
	0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x44, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e,
	0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x56,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x07, 0x72,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x44, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x07,
	0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x61, 0x6e, 0x6b, 0x52, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x26, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69,
	0x6e, 0x67, 0x22, 0x3c, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0x44, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x08, 0x72, 0x61,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xd2, 0x01, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xe6, 0x04, 0x0a, 0x07,
	0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x3e, 0x0a, 0x04, 0x54, 0x6f, 0x70, 0x4e, 0x12,
	0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x6f, 0x70, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f,
	0x70, 0x4e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x6d, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x72, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6f, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x2a, 0x0c, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x63, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x62, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12,
	0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x79, 0x2f,
	0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ranking_v1_ranking_proto_rawDescOnce sync.Once
	file_api_ranking_v1_ranking_proto_rawDescData = file_api_ranking_v1_ranking_proto_rawDesc
)

func file_api_ranking_v1_ranking_proto_rawDescGZIP() []byte {
	file_api_ranking_v1_ranking_proto_rawDescOnce.Do(func() {
		file_api_ranking_v1_ranking_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ranking_v1_ranking_proto_rawDescData)
	})
	return file_api_ranking_v1_ranking_proto_rawDescData
}

var file_api_ranking_v1_ranking_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_ranking_v1_ranking_proto_goTypes = []any{
	(*TopNRequest)(nil),          // 0: api.ranking.v1.TopNRequest
	(*TopNReply)(nil),            // 1: api.ranking.v1.TopNReply
	(*CreateRankingRequest)(nil), // 2: api.ranking.v1.CreateRankingRequest
	(*CreateRankingReply)(nil),   // 3: api.ranking.v1.CreateRankingReply
	(*UpdateRankingRequest)(nil), // 4: api.ranking.v1.UpdateRankingRequest
	(*UpdateRankingReply)(nil),   // 5: api.ranking.v1.UpdateRankingReply
	(*DeleteRankingRequest)(nil), // 6: api.ranking.v1.DeleteRankingRequest
	(*DeleteRankingReply)(nil),   // 7: api.ranking.v1.DeleteRankingReply
	(*GetRankingRequest)(nil),    // 8: api.ranking.v1.GetRankingRequest
	(*GetRankingReply)(nil),      // 9: api.ranking.v1.GetRankingReply
	(*ListRankingRequest)(nil),   // 10: api.ranking.v1.ListRankingRequest
	(*ListRankingReply)(nil),     // 11: api.ranking.v1.ListRankingReply
	(*Rank)(nil),                 // 12: api.ranking.v1.Rank
}
var file_api_ranking_v1_ranking_proto_depIdxs = []int32{
	12, // 0: api.ranking.v1.TopNReply.ranking:type_name -> api.ranking.v1.Rank
	12, // 1: api.ranking.v1.CreateRankingRequest.ranking:type_name -> api.ranking.v1.Rank
	12, // 2: api.ranking.v1.CreateRankingReply.ranking:type_name -> api.ranking.v1.Rank
	12, // 3: api.ranking.v1.UpdateRankingRequest.ranking:type_name -> api.ranking.v1.Rank
	12, // 4: api.ranking.v1.UpdateRankingReply.ranking:type_name -> api.ranking.v1.Rank
	12, // 5: api.ranking.v1.GetRankingReply.ranking:type_name -> api.ranking.v1.Rank
	12, // 6: api.ranking.v1.ListRankingReply.rankings:type_name -> api.ranking.v1.Rank
	0,  // 7: api.ranking.v1.Ranking.TopN:input_type -> api.ranking.v1.TopNRequest
	2,  // 8: api.ranking.v1.Ranking.CreateRanking:input_type -> api.ranking.v1.CreateRankingRequest
	4,  // 9: api.ranking.v1.Ranking.UpdateRanking:input_type -> api.ranking.v1.UpdateRankingRequest
	6,  // 10: api.ranking.v1.Ranking.DeleteRanking:input_type -> api.ranking.v1.DeleteRankingRequest
	8,  // 11: api.ranking.v1.Ranking.GetRanking:input_type -> api.ranking.v1.GetRankingRequest
	10, // 12: api.ranking.v1.Ranking.ListRanking:input_type -> api.ranking.v1.ListRankingRequest
	1,  // 13: api.ranking.v1.Ranking.TopN:output_type -> api.ranking.v1.TopNReply
	3,  // 14: api.ranking.v1.Ranking.CreateRanking:output_type -> api.ranking.v1.CreateRankingReply
	5,  // 15: api.ranking.v1.Ranking.UpdateRanking:output_type -> api.ranking.v1.UpdateRankingReply
	7,  // 16: api.ranking.v1.Ranking.DeleteRanking:output_type -> api.ranking.v1.DeleteRankingReply
	9,  // 17: api.ranking.v1.Ranking.GetRanking:output_type -> api.ranking.v1.GetRankingReply
	11, // 18: api.ranking.v1.Ranking.ListRanking:output_type -> api.ranking.v1.ListRankingReply
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_ranking_v1_ranking_proto_init() }
func file_api_ranking_v1_ranking_proto_init() {
	if File_api_ranking_v1_ranking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ranking_v1_ranking_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TopNRequest); i {
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
		file_api_ranking_v1_ranking_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TopNReply); i {
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
		file_api_ranking_v1_ranking_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRankingRequest); i {
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
		file_api_ranking_v1_ranking_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRankingReply); i {
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
		file_api_ranking_v1_ranking_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRankingRequest); i {
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
		file_api_ranking_v1_ranking_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRankingReply); i {
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
		file_api_ranking_v1_ranking_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRankingRequest); i {
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
		file_api_ranking_v1_ranking_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRankingReply); i {
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
		file_api_ranking_v1_ranking_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetRankingRequest); i {
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
		file_api_ranking_v1_ranking_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetRankingReply); i {
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
		file_api_ranking_v1_ranking_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ListRankingRequest); i {
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
		file_api_ranking_v1_ranking_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ListRankingReply); i {
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
		file_api_ranking_v1_ranking_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*Rank); i {
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
			RawDescriptor: file_api_ranking_v1_ranking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ranking_v1_ranking_proto_goTypes,
		DependencyIndexes: file_api_ranking_v1_ranking_proto_depIdxs,
		MessageInfos:      file_api_ranking_v1_ranking_proto_msgTypes,
	}.Build()
	File_api_ranking_v1_ranking_proto = out.File
	file_api_ranking_v1_ranking_proto_rawDesc = nil
	file_api_ranking_v1_ranking_proto_goTypes = nil
	file_api_ranking_v1_ranking_proto_depIdxs = nil
}
