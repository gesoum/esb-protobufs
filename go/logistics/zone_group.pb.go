// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.9
// source: zone_group.proto

package logistics

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ZoneGroupId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ZoneGroupId) Reset() {
	*x = ZoneGroupId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneGroupId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneGroupId) ProtoMessage() {}

func (x *ZoneGroupId) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneGroupId.ProtoReflect.Descriptor instead.
func (*ZoneGroupId) Descriptor() ([]byte, []int) {
	return file_zone_group_proto_rawDescGZIP(), []int{0}
}

func (x *ZoneGroupId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ZoneId int32 `protobuf:"varint,2,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
}

func (x *DeleteZoneRequest) Reset() {
	*x = DeleteZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteZoneRequest) ProtoMessage() {}

func (x *DeleteZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteZoneRequest.ProtoReflect.Descriptor instead.
func (*DeleteZoneRequest) Descriptor() ([]byte, []int) {
	return file_zone_group_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteZoneRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteZoneRequest) GetZoneId() int32 {
	if x != nil {
		return x.ZoneId
	}
	return 0
}

type ZoneGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsActive  bool    `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt string  `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string  `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Zones     []*Zone `protobuf:"bytes,6,rep,name=zones,proto3" json:"zones,omitempty"`
}

func (x *ZoneGroup) Reset() {
	*x = ZoneGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneGroup) ProtoMessage() {}

func (x *ZoneGroup) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneGroup.ProtoReflect.Descriptor instead.
func (*ZoneGroup) Descriptor() ([]byte, []int) {
	return file_zone_group_proto_rawDescGZIP(), []int{2}
}

func (x *ZoneGroup) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ZoneGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ZoneGroup) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *ZoneGroup) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ZoneGroup) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ZoneGroup) GetZones() []*Zone {
	if x != nil {
		return x.Zones
	}
	return nil
}

type UpdateZoneGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsActive bool    `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	ZoneIds  []int32 `protobuf:"varint,4,rep,packed,name=zone_ids,json=zoneIds,proto3" json:"zone_ids,omitempty"`
}

func (x *UpdateZoneGroup) Reset() {
	*x = UpdateZoneGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateZoneGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateZoneGroup) ProtoMessage() {}

func (x *UpdateZoneGroup) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateZoneGroup.ProtoReflect.Descriptor instead.
func (*UpdateZoneGroup) Descriptor() ([]byte, []int) {
	return file_zone_group_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateZoneGroup) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateZoneGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateZoneGroup) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *UpdateZoneGroup) GetZoneIds() []int32 {
	if x != nil {
		return x.ZoneIds
	}
	return nil
}

type CreateZoneGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsActive bool    `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	ZoneIds  []int32 `protobuf:"varint,3,rep,packed,name=zone_ids,json=zoneIds,proto3" json:"zone_ids,omitempty"`
}

func (x *CreateZoneGroup) Reset() {
	*x = CreateZoneGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateZoneGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateZoneGroup) ProtoMessage() {}

func (x *CreateZoneGroup) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateZoneGroup.ProtoReflect.Descriptor instead.
func (*CreateZoneGroup) Descriptor() ([]byte, []int) {
	return file_zone_group_proto_rawDescGZIP(), []int{4}
}

func (x *CreateZoneGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateZoneGroup) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *CreateZoneGroup) GetZoneIds() []int32 {
	if x != nil {
		return x.ZoneIds
	}
	return nil
}

type ListZoneGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	Sort   string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *ListZoneGroupRequest) Reset() {
	*x = ListZoneGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListZoneGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZoneGroupRequest) ProtoMessage() {}

func (x *ListZoneGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListZoneGroupRequest.ProtoReflect.Descriptor instead.
func (*ListZoneGroupRequest) Descriptor() ([]byte, []int) {
	return file_zone_group_proto_rawDescGZIP(), []int{5}
}

func (x *ListZoneGroupRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListZoneGroupRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListZoneGroupRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListZoneGroupRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type ListZoneGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ZoneGroup `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListZoneGroupResponse) Reset() {
	*x = ListZoneGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListZoneGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZoneGroupResponse) ProtoMessage() {}

func (x *ListZoneGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListZoneGroupResponse.ProtoReflect.Descriptor instead.
func (*ListZoneGroupResponse) Descriptor() ([]byte, []int) {
	return file_zone_group_proto_rawDescGZIP(), []int{6}
}

func (x *ListZoneGroupResponse) GetResults() []*ZoneGroup {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListZoneGroupResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type SuggestZoneGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *SuggestZoneGroupRequest) Reset() {
	*x = SuggestZoneGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestZoneGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestZoneGroupRequest) ProtoMessage() {}

func (x *SuggestZoneGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestZoneGroupRequest.ProtoReflect.Descriptor instead.
func (*SuggestZoneGroupRequest) Descriptor() ([]byte, []int) {
	return file_zone_group_proto_rawDescGZIP(), []int{7}
}

func (x *SuggestZoneGroupRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SuggestZoneGroupRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type SuggestZoneGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsActive bool   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *SuggestZoneGroup) Reset() {
	*x = SuggestZoneGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestZoneGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestZoneGroup) ProtoMessage() {}

func (x *SuggestZoneGroup) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestZoneGroup.ProtoReflect.Descriptor instead.
func (*SuggestZoneGroup) Descriptor() ([]byte, []int) {
	return file_zone_group_proto_rawDescGZIP(), []int{8}
}

func (x *SuggestZoneGroup) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SuggestZoneGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SuggestZoneGroup) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type SuggestZoneGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*SuggestZoneGroup `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *SuggestZoneGroupResponse) Reset() {
	*x = SuggestZoneGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestZoneGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestZoneGroupResponse) ProtoMessage() {}

func (x *SuggestZoneGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestZoneGroupResponse.ProtoReflect.Descriptor instead.
func (*SuggestZoneGroupResponse) Descriptor() ([]byte, []int) {
	return file_zone_group_proto_rawDescGZIP(), []int{9}
}

func (x *SuggestZoneGroupResponse) GetResults() []*SuggestZoneGroup {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *SuggestZoneGroupResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_zone_group_proto protoreflect.FileDescriptor

var file_zone_group_proto_rawDesc = []byte{
	0x0a, 0x10, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x0b, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5a, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49,
	0x64, 0x22, 0xb1, 0x01, 0x0a, 0x09, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x25,
	0x0a, 0x05, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x05,
	0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x22, 0x6d, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5a,
	0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x6f, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x7a, 0x6f, 0x6e,
	0x65, 0x49, 0x64, 0x73, 0x22, 0x5d, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5a, 0x6f,
	0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x6f, 0x6e, 0x65,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x7a, 0x6f, 0x6e, 0x65,
	0x49, 0x64, 0x73, 0x22, 0x70, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x5d, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x22, 0x47, 0x0a, 0x17, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x5a,
	0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x53, 0x0a,
	0x10, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x22, 0x67, 0x0a, 0x18, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x5a, 0x6f, 0x6e,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xd1, 0x05, 0x0a, 0x10,
	0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x65, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f, 0x6e,
	0x65, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x59, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x14, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2d, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x54, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x1a, 0x14, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f,
	0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12,
	0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2d, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1a, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x14,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x32,
	0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2d, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x59, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a,
	0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x76, 0x0a, 0x07, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12,
	0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2d, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x12, 0x72, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x2a, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x7b, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42,
	0x0e, 0x5a, 0x0c, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zone_group_proto_rawDescOnce sync.Once
	file_zone_group_proto_rawDescData = file_zone_group_proto_rawDesc
)

func file_zone_group_proto_rawDescGZIP() []byte {
	file_zone_group_proto_rawDescOnce.Do(func() {
		file_zone_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_zone_group_proto_rawDescData)
	})
	return file_zone_group_proto_rawDescData
}

var file_zone_group_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_zone_group_proto_goTypes = []interface{}{
	(*ZoneGroupId)(nil),              // 0: logistics.ZoneGroupId
	(*DeleteZoneRequest)(nil),        // 1: logistics.DeleteZoneRequest
	(*ZoneGroup)(nil),                // 2: logistics.ZoneGroup
	(*UpdateZoneGroup)(nil),          // 3: logistics.UpdateZoneGroup
	(*CreateZoneGroup)(nil),          // 4: logistics.CreateZoneGroup
	(*ListZoneGroupRequest)(nil),     // 5: logistics.ListZoneGroupRequest
	(*ListZoneGroupResponse)(nil),    // 6: logistics.ListZoneGroupResponse
	(*SuggestZoneGroupRequest)(nil),  // 7: logistics.SuggestZoneGroupRequest
	(*SuggestZoneGroup)(nil),         // 8: logistics.SuggestZoneGroup
	(*SuggestZoneGroupResponse)(nil), // 9: logistics.SuggestZoneGroupResponse
	(*Zone)(nil),                     // 10: logistics.Zone
	(*emptypb.Empty)(nil),            // 11: google.protobuf.Empty
}
var file_zone_group_proto_depIdxs = []int32{
	10, // 0: logistics.ZoneGroup.zones:type_name -> logistics.Zone
	2,  // 1: logistics.ListZoneGroupResponse.results:type_name -> logistics.ZoneGroup
	8,  // 2: logistics.SuggestZoneGroupResponse.results:type_name -> logistics.SuggestZoneGroup
	5,  // 3: logistics.ZoneGroupService.List:input_type -> logistics.ListZoneGroupRequest
	4,  // 4: logistics.ZoneGroupService.Create:input_type -> logistics.CreateZoneGroup
	0,  // 5: logistics.ZoneGroupService.Get:input_type -> logistics.ZoneGroupId
	3,  // 6: logistics.ZoneGroupService.Update:input_type -> logistics.UpdateZoneGroup
	0,  // 7: logistics.ZoneGroupService.Delete:input_type -> logistics.ZoneGroupId
	7,  // 8: logistics.ZoneGroupService.Suggest:input_type -> logistics.SuggestZoneGroupRequest
	1,  // 9: logistics.ZoneGroupService.DeleteZone:input_type -> logistics.DeleteZoneRequest
	6,  // 10: logistics.ZoneGroupService.List:output_type -> logistics.ListZoneGroupResponse
	2,  // 11: logistics.ZoneGroupService.Create:output_type -> logistics.ZoneGroup
	2,  // 12: logistics.ZoneGroupService.Get:output_type -> logistics.ZoneGroup
	2,  // 13: logistics.ZoneGroupService.Update:output_type -> logistics.ZoneGroup
	11, // 14: logistics.ZoneGroupService.Delete:output_type -> google.protobuf.Empty
	9,  // 15: logistics.ZoneGroupService.Suggest:output_type -> logistics.SuggestZoneGroupResponse
	11, // 16: logistics.ZoneGroupService.DeleteZone:output_type -> google.protobuf.Empty
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_zone_group_proto_init() }
func file_zone_group_proto_init() {
	if File_zone_group_proto != nil {
		return
	}
	file_zone_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_zone_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneGroupId); i {
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
		file_zone_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteZoneRequest); i {
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
		file_zone_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneGroup); i {
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
		file_zone_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateZoneGroup); i {
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
		file_zone_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateZoneGroup); i {
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
		file_zone_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListZoneGroupRequest); i {
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
		file_zone_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListZoneGroupResponse); i {
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
		file_zone_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestZoneGroupRequest); i {
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
		file_zone_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestZoneGroup); i {
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
		file_zone_group_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestZoneGroupResponse); i {
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
			RawDescriptor: file_zone_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zone_group_proto_goTypes,
		DependencyIndexes: file_zone_group_proto_depIdxs,
		MessageInfos:      file_zone_group_proto_msgTypes,
	}.Build()
	File_zone_group_proto = out.File
	file_zone_group_proto_rawDesc = nil
	file_zone_group_proto_goTypes = nil
	file_zone_group_proto_depIdxs = nil
}
