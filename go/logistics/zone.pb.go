// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.9
// source: zone.proto

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

type ZoneId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ZoneId) Reset() {
	*x = ZoneId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneId) ProtoMessage() {}

func (x *ZoneId) ProtoReflect() protoreflect.Message {
	mi := &file_zone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneId.ProtoReflect.Descriptor instead.
func (*ZoneId) Descriptor() ([]byte, []int) {
	return file_zone_proto_rawDescGZIP(), []int{0}
}

func (x *ZoneId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Zone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsActive      bool             `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt     string           `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string           `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ZoneGroups    []*ZoneGroupInfo `protobuf:"bytes,6,rep,name=zone_groups,json=zoneGroups,proto3" json:"zone_groups,omitempty"`
	GeoIdQuantity int32            `protobuf:"varint,7,opt,name=geo_id_quantity,json=geoIdQuantity,proto3" json:"geo_id_quantity,omitempty"`
	GeoIds        []string         `protobuf:"bytes,8,rep,name=geo_ids,json=geoIds,proto3" json:"geo_ids,omitempty"`
}

func (x *Zone) Reset() {
	*x = Zone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zone) ProtoMessage() {}

func (x *Zone) ProtoReflect() protoreflect.Message {
	mi := &file_zone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Zone.ProtoReflect.Descriptor instead.
func (*Zone) Descriptor() ([]byte, []int) {
	return file_zone_proto_rawDescGZIP(), []int{1}
}

func (x *Zone) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Zone) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Zone) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Zone) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Zone) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Zone) GetZoneGroups() []*ZoneGroupInfo {
	if x != nil {
		return x.ZoneGroups
	}
	return nil
}

func (x *Zone) GetGeoIdQuantity() int32 {
	if x != nil {
		return x.GeoIdQuantity
	}
	return 0
}

func (x *Zone) GetGeoIds() []string {
	if x != nil {
		return x.GeoIds
	}
	return nil
}

type ZoneGroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsActive  bool   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ZoneGroupInfo) Reset() {
	*x = ZoneGroupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneGroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneGroupInfo) ProtoMessage() {}

func (x *ZoneGroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_zone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneGroupInfo.ProtoReflect.Descriptor instead.
func (*ZoneGroupInfo) Descriptor() ([]byte, []int) {
	return file_zone_proto_rawDescGZIP(), []int{2}
}

func (x *ZoneGroupInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ZoneGroupInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ZoneGroupInfo) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *ZoneGroupInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ZoneGroupInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ListZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	Sort   string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *ListZoneRequest) Reset() {
	*x = ListZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZoneRequest) ProtoMessage() {}

func (x *ListZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zone_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListZoneRequest.ProtoReflect.Descriptor instead.
func (*ListZoneRequest) Descriptor() ([]byte, []int) {
	return file_zone_proto_rawDescGZIP(), []int{3}
}

func (x *ListZoneRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListZoneRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListZoneRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListZoneRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type ListZoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Zone `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListZoneResponse) Reset() {
	*x = ListZoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListZoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZoneResponse) ProtoMessage() {}

func (x *ListZoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zone_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListZoneResponse.ProtoReflect.Descriptor instead.
func (*ListZoneResponse) Descriptor() ([]byte, []int) {
	return file_zone_proto_rawDescGZIP(), []int{4}
}

func (x *ListZoneResponse) GetResults() []*Zone {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListZoneResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type SuggestZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *SuggestZoneRequest) Reset() {
	*x = SuggestZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestZoneRequest) ProtoMessage() {}

func (x *SuggestZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zone_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestZoneRequest.ProtoReflect.Descriptor instead.
func (*SuggestZoneRequest) Descriptor() ([]byte, []int) {
	return file_zone_proto_rawDescGZIP(), []int{5}
}

func (x *SuggestZoneRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SuggestZoneRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type SuggestZone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsActive bool   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *SuggestZone) Reset() {
	*x = SuggestZone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestZone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestZone) ProtoMessage() {}

func (x *SuggestZone) ProtoReflect() protoreflect.Message {
	mi := &file_zone_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestZone.ProtoReflect.Descriptor instead.
func (*SuggestZone) Descriptor() ([]byte, []int) {
	return file_zone_proto_rawDescGZIP(), []int{6}
}

func (x *SuggestZone) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SuggestZone) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SuggestZone) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type SuggestZoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*SuggestZone `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *SuggestZoneResponse) Reset() {
	*x = SuggestZoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestZoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestZoneResponse) ProtoMessage() {}

func (x *SuggestZoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zone_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestZoneResponse.ProtoReflect.Descriptor instead.
func (*SuggestZoneResponse) Descriptor() ([]byte, []int) {
	return file_zone_proto_rawDescGZIP(), []int{7}
}

func (x *SuggestZoneResponse) GetResults() []*SuggestZone {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *SuggestZoneResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_zone_proto protoreflect.FileDescriptor

var file_zone_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x18, 0x0a, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x81, 0x02, 0x0a,
	0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x7a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x67, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x67, 0x65, 0x6f, 0x49, 0x64, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x65, 0x6f, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6f, 0x49, 0x64, 0x73,
	0x22, 0x8e, 0x01, 0x0a, 0x0d, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x6b, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x53,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x42, 0x0a, 0x12, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x5a, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x4e, 0x0a, 0x0b, 0x53, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x5d, 0x0a, 0x13, 0x53, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xf1, 0x03, 0x0a, 0x0b, 0x5a, 0x6f, 0x6e, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x5a,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12,
	0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x43, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x1a, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f,
	0x6e, 0x65, 0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7a,
	0x6f, 0x6e, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x48, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a,
	0x6f, 0x6e, 0x65, 0x1a, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x32,
	0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x4e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a,
	0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x66, 0x0a, 0x07, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f,
	0x6e, 0x65, 0x2f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x6f,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_zone_proto_rawDescOnce sync.Once
	file_zone_proto_rawDescData = file_zone_proto_rawDesc
)

func file_zone_proto_rawDescGZIP() []byte {
	file_zone_proto_rawDescOnce.Do(func() {
		file_zone_proto_rawDescData = protoimpl.X.CompressGZIP(file_zone_proto_rawDescData)
	})
	return file_zone_proto_rawDescData
}

var file_zone_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_zone_proto_goTypes = []interface{}{
	(*ZoneId)(nil),              // 0: logistics.ZoneId
	(*Zone)(nil),                // 1: logistics.Zone
	(*ZoneGroupInfo)(nil),       // 2: logistics.ZoneGroupInfo
	(*ListZoneRequest)(nil),     // 3: logistics.ListZoneRequest
	(*ListZoneResponse)(nil),    // 4: logistics.ListZoneResponse
	(*SuggestZoneRequest)(nil),  // 5: logistics.SuggestZoneRequest
	(*SuggestZone)(nil),         // 6: logistics.SuggestZone
	(*SuggestZoneResponse)(nil), // 7: logistics.SuggestZoneResponse
	(*emptypb.Empty)(nil),       // 8: google.protobuf.Empty
}
var file_zone_proto_depIdxs = []int32{
	2, // 0: logistics.Zone.zone_groups:type_name -> logistics.ZoneGroupInfo
	1, // 1: logistics.ListZoneResponse.results:type_name -> logistics.Zone
	6, // 2: logistics.SuggestZoneResponse.results:type_name -> logistics.SuggestZone
	3, // 3: logistics.ZoneService.List:input_type -> logistics.ListZoneRequest
	1, // 4: logistics.ZoneService.Create:input_type -> logistics.Zone
	0, // 5: logistics.ZoneService.Get:input_type -> logistics.ZoneId
	1, // 6: logistics.ZoneService.Update:input_type -> logistics.Zone
	0, // 7: logistics.ZoneService.Delete:input_type -> logistics.ZoneId
	5, // 8: logistics.ZoneService.Suggest:input_type -> logistics.SuggestZoneRequest
	4, // 9: logistics.ZoneService.List:output_type -> logistics.ListZoneResponse
	1, // 10: logistics.ZoneService.Create:output_type -> logistics.Zone
	1, // 11: logistics.ZoneService.Get:output_type -> logistics.Zone
	1, // 12: logistics.ZoneService.Update:output_type -> logistics.Zone
	8, // 13: logistics.ZoneService.Delete:output_type -> google.protobuf.Empty
	7, // 14: logistics.ZoneService.Suggest:output_type -> logistics.SuggestZoneResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_zone_proto_init() }
func file_zone_proto_init() {
	if File_zone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneId); i {
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
		file_zone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zone); i {
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
		file_zone_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneGroupInfo); i {
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
		file_zone_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListZoneRequest); i {
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
		file_zone_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListZoneResponse); i {
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
		file_zone_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestZoneRequest); i {
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
		file_zone_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestZone); i {
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
		file_zone_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestZoneResponse); i {
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
			RawDescriptor: file_zone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zone_proto_goTypes,
		DependencyIndexes: file_zone_proto_depIdxs,
		MessageInfos:      file_zone_proto_msgTypes,
	}.Build()
	File_zone_proto = out.File
	file_zone_proto_rawDesc = nil
	file_zone_proto_goTypes = nil
	file_zone_proto_depIdxs = nil
}
