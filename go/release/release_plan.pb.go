// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.9
// source: release_plan.proto

package release

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

type ArticleReleaseType int32

const (
	ArticleReleaseType_site    ArticleReleaseType = 0
	ArticleReleaseType_mercaux ArticleReleaseType = 1
)

// Enum value maps for ArticleReleaseType.
var (
	ArticleReleaseType_name = map[int32]string{
		0: "site",
		1: "mercaux",
	}
	ArticleReleaseType_value = map[string]int32{
		"site":    0,
		"mercaux": 1,
	}
)

func (x ArticleReleaseType) Enum() *ArticleReleaseType {
	p := new(ArticleReleaseType)
	*p = x
	return p
}

func (x ArticleReleaseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArticleReleaseType) Descriptor() protoreflect.EnumDescriptor {
	return file_release_plan_proto_enumTypes[0].Descriptor()
}

func (ArticleReleaseType) Type() protoreflect.EnumType {
	return &file_release_plan_proto_enumTypes[0]
}

func (x ArticleReleaseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArticleReleaseType.Descriptor instead.
func (ArticleReleaseType) EnumDescriptor() ([]byte, []int) {
	return file_release_plan_proto_rawDescGZIP(), []int{0}
}

type ParamsCreateArticle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article     string             `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	TriggerTime string             `protobuf:"bytes,2,opt,name=trigger_time,json=triggerTime,proto3" json:"trigger_time,omitempty"`
	Type        ArticleReleaseType `protobuf:"varint,3,opt,name=type,proto3,enum=release.ArticleReleaseType" json:"type,omitempty"`
}

func (x *ParamsCreateArticle) Reset() {
	*x = ParamsCreateArticle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamsCreateArticle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamsCreateArticle) ProtoMessage() {}

func (x *ParamsCreateArticle) ProtoReflect() protoreflect.Message {
	mi := &file_release_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamsCreateArticle.ProtoReflect.Descriptor instead.
func (*ParamsCreateArticle) Descriptor() ([]byte, []int) {
	return file_release_plan_proto_rawDescGZIP(), []int{0}
}

func (x *ParamsCreateArticle) GetArticle() string {
	if x != nil {
		return x.Article
	}
	return ""
}

func (x *ParamsCreateArticle) GetTriggerTime() string {
	if x != nil {
		return x.TriggerTime
	}
	return ""
}

func (x *ParamsCreateArticle) GetType() ArticleReleaseType {
	if x != nil {
		return x.Type
	}
	return ArticleReleaseType_site
}

type ArticleReleaseActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool  `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Id int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ArticleReleaseActionResponse) Reset() {
	*x = ArticleReleaseActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_plan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleReleaseActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleReleaseActionResponse) ProtoMessage() {}

func (x *ArticleReleaseActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_release_plan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleReleaseActionResponse.ProtoReflect.Descriptor instead.
func (*ArticleReleaseActionResponse) Descriptor() ([]byte, []int) {
	return file_release_plan_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleReleaseActionResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ArticleReleaseActionResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ParamsGetArticleRelease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit     int32              `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset    int32              `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Type      ArticleReleaseType `protobuf:"varint,3,opt,name=type,proto3,enum=release.ArticleReleaseType" json:"type,omitempty"`
	IsDeleted bool               `protobuf:"varint,4,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	IsActive  bool               `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *ParamsGetArticleRelease) Reset() {
	*x = ParamsGetArticleRelease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_plan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamsGetArticleRelease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamsGetArticleRelease) ProtoMessage() {}

func (x *ParamsGetArticleRelease) ProtoReflect() protoreflect.Message {
	mi := &file_release_plan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamsGetArticleRelease.ProtoReflect.Descriptor instead.
func (*ParamsGetArticleRelease) Descriptor() ([]byte, []int) {
	return file_release_plan_proto_rawDescGZIP(), []int{2}
}

func (x *ParamsGetArticleRelease) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ParamsGetArticleRelease) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ParamsGetArticleRelease) GetType() ArticleReleaseType {
	if x != nil {
		return x.Type
	}
	return ArticleReleaseType_site
}

func (x *ParamsGetArticleRelease) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *ParamsGetArticleRelease) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type ArticleReleaseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok      bool              `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Total   int32             `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Article []*ArticleRelease `protobuf:"bytes,3,rep,name=article,proto3" json:"article,omitempty"`
}

func (x *ArticleReleaseList) Reset() {
	*x = ArticleReleaseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_plan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleReleaseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleReleaseList) ProtoMessage() {}

func (x *ArticleReleaseList) ProtoReflect() protoreflect.Message {
	mi := &file_release_plan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleReleaseList.ProtoReflect.Descriptor instead.
func (*ArticleReleaseList) Descriptor() ([]byte, []int) {
	return file_release_plan_proto_rawDescGZIP(), []int{3}
}

func (x *ArticleReleaseList) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ArticleReleaseList) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ArticleReleaseList) GetArticle() []*ArticleRelease {
	if x != nil {
		return x.Article
	}
	return nil
}

type ArticleRelease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Article     string             `protobuf:"bytes,2,opt,name=article,proto3" json:"article,omitempty"`
	TriggerTime string             `protobuf:"bytes,3,opt,name=trigger_time,json=triggerTime,proto3" json:"trigger_time,omitempty"`
	Type        ArticleReleaseType `protobuf:"varint,4,opt,name=type,proto3,enum=release.ArticleReleaseType" json:"type,omitempty"`
	IsDeleted   bool               `protobuf:"varint,5,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	IsActive    bool               `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *ArticleRelease) Reset() {
	*x = ArticleRelease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_plan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleRelease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleRelease) ProtoMessage() {}

func (x *ArticleRelease) ProtoReflect() protoreflect.Message {
	mi := &file_release_plan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleRelease.ProtoReflect.Descriptor instead.
func (*ArticleRelease) Descriptor() ([]byte, []int) {
	return file_release_plan_proto_rawDescGZIP(), []int{4}
}

func (x *ArticleRelease) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ArticleRelease) GetArticle() string {
	if x != nil {
		return x.Article
	}
	return ""
}

func (x *ArticleRelease) GetTriggerTime() string {
	if x != nil {
		return x.TriggerTime
	}
	return ""
}

func (x *ArticleRelease) GetType() ArticleReleaseType {
	if x != nil {
		return x.Type
	}
	return ArticleReleaseType_site
}

func (x *ArticleRelease) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *ArticleRelease) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type ArticlesReleaseId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ArticlesReleaseId) Reset() {
	*x = ArticlesReleaseId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_plan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticlesReleaseId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticlesReleaseId) ProtoMessage() {}

func (x *ArticlesReleaseId) ProtoReflect() protoreflect.Message {
	mi := &file_release_plan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticlesReleaseId.ProtoReflect.Descriptor instead.
func (*ArticlesReleaseId) Descriptor() ([]byte, []int) {
	return file_release_plan_proto_rawDescGZIP(), []int{5}
}

func (x *ArticlesReleaseId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_release_plan_proto protoreflect.FileDescriptor

var file_release_plan_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x22, 0x83, 0x01,
	0x0a, 0x13, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x3e, 0x0a, 0x1c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xb4, 0x01, 0x0a, 0x17, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x2f, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x6d, 0x0a, 0x12, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x31, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0xca, 0x01, 0x0a, 0x0e, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x2b, 0x0a, 0x12, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x6d,
	0x65, 0x72, 0x63, 0x61, 0x75, 0x78, 0x10, 0x01, 0x32, 0x8c, 0x03, 0x0a, 0x0f, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x1a, 0x25, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x1a, 0x1b, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x1a, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x1a, 0x17, 0x2e, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x1a,
	0x25, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x1a, 0x25,
	0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x6f, 0x2f, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_release_plan_proto_rawDescOnce sync.Once
	file_release_plan_proto_rawDescData = file_release_plan_proto_rawDesc
)

func file_release_plan_proto_rawDescGZIP() []byte {
	file_release_plan_proto_rawDescOnce.Do(func() {
		file_release_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_release_plan_proto_rawDescData)
	})
	return file_release_plan_proto_rawDescData
}

var file_release_plan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_release_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_release_plan_proto_goTypes = []interface{}{
	(ArticleReleaseType)(0),              // 0: release.ArticleReleaseType
	(*ParamsCreateArticle)(nil),          // 1: release.ParamsCreateArticle
	(*ArticleReleaseActionResponse)(nil), // 2: release.ArticleReleaseActionResponse
	(*ParamsGetArticleRelease)(nil),      // 3: release.ParamsGetArticleRelease
	(*ArticleReleaseList)(nil),           // 4: release.ArticleReleaseList
	(*ArticleRelease)(nil),               // 5: release.ArticleRelease
	(*ArticlesReleaseId)(nil),            // 6: release.ArticlesReleaseId
}
var file_release_plan_proto_depIdxs = []int32{
	0, // 0: release.ParamsCreateArticle.type:type_name -> release.ArticleReleaseType
	0, // 1: release.ParamsGetArticleRelease.type:type_name -> release.ArticleReleaseType
	5, // 2: release.ArticleReleaseList.article:type_name -> release.ArticleRelease
	0, // 3: release.ArticleRelease.type:type_name -> release.ArticleReleaseType
	1, // 4: release.ArticlesRelease.Create:input_type -> release.ParamsCreateArticle
	3, // 5: release.ArticlesRelease.Get:input_type -> release.ParamsGetArticleRelease
	6, // 6: release.ArticlesRelease.GetById:input_type -> release.ArticlesReleaseId
	1, // 7: release.ArticlesRelease.Update:input_type -> release.ParamsCreateArticle
	6, // 8: release.ArticlesRelease.Delete:input_type -> release.ArticlesReleaseId
	2, // 9: release.ArticlesRelease.Create:output_type -> release.ArticleReleaseActionResponse
	4, // 10: release.ArticlesRelease.Get:output_type -> release.ArticleReleaseList
	5, // 11: release.ArticlesRelease.GetById:output_type -> release.ArticleRelease
	2, // 12: release.ArticlesRelease.Update:output_type -> release.ArticleReleaseActionResponse
	2, // 13: release.ArticlesRelease.Delete:output_type -> release.ArticleReleaseActionResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_release_plan_proto_init() }
func file_release_plan_proto_init() {
	if File_release_plan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_release_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamsCreateArticle); i {
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
		file_release_plan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleReleaseActionResponse); i {
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
		file_release_plan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamsGetArticleRelease); i {
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
		file_release_plan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleReleaseList); i {
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
		file_release_plan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleRelease); i {
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
		file_release_plan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticlesReleaseId); i {
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
			RawDescriptor: file_release_plan_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_release_plan_proto_goTypes,
		DependencyIndexes: file_release_plan_proto_depIdxs,
		EnumInfos:         file_release_plan_proto_enumTypes,
		MessageInfos:      file_release_plan_proto_msgTypes,
	}.Build()
	File_release_plan_proto = out.File
	file_release_plan_proto_rawDesc = nil
	file_release_plan_proto_goTypes = nil
	file_release_plan_proto_depIdxs = nil
}
