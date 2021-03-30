// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: proto/logistics/zone.proto

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

type DeleteZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneId int32 `protobuf:"varint,1,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
}

func (x *DeleteZoneRequest) Reset() {
	*x = DeleteZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_zone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteZoneRequest) ProtoMessage() {}

func (x *DeleteZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_zone_proto_msgTypes[0]
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
	return file_proto_logistics_zone_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteZoneRequest) GetZoneId() int32 {
	if x != nil {
		return x.ZoneId
	}
	return 0
}

type UpdateZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneId int32  `protobuf:"varint,1,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateZoneRequest) Reset() {
	*x = UpdateZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_zone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateZoneRequest) ProtoMessage() {}

func (x *UpdateZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_zone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateZoneRequest.ProtoReflect.Descriptor instead.
func (*UpdateZoneRequest) Descriptor() ([]byte, []int) {
	return file_proto_logistics_zone_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateZoneRequest) GetZoneId() int32 {
	if x != nil {
		return x.ZoneId
	}
	return 0
}

func (x *UpdateZoneRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateZoneRequest) Reset() {
	*x = CreateZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_zone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateZoneRequest) ProtoMessage() {}

func (x *CreateZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_zone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateZoneRequest.ProtoReflect.Descriptor instead.
func (*CreateZoneRequest) Descriptor() ([]byte, []int) {
	return file_proto_logistics_zone_proto_rawDescGZIP(), []int{2}
}

func (x *CreateZoneRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateZoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateZoneResponse) Reset() {
	*x = CreateZoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_zone_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateZoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateZoneResponse) ProtoMessage() {}

func (x *CreateZoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_zone_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateZoneResponse.ProtoReflect.Descriptor instead.
func (*CreateZoneResponse) Descriptor() ([]byte, []int) {
	return file_proto_logistics_zone_proto_rawDescGZIP(), []int{3}
}

func (x *CreateZoneResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListZonesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListZonesRequest) Reset() {
	*x = ListZonesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_zone_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListZonesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZonesRequest) ProtoMessage() {}

func (x *ListZonesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_zone_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListZonesRequest.ProtoReflect.Descriptor instead.
func (*ListZonesRequest) Descriptor() ([]byte, []int) {
	return file_proto_logistics_zone_proto_rawDescGZIP(), []int{4}
}

func (x *ListZonesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListZonesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListZonesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Zone `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListZonesResponse) Reset() {
	*x = ListZonesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_zone_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListZonesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZonesResponse) ProtoMessage() {}

func (x *ListZonesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_zone_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListZonesResponse.ProtoReflect.Descriptor instead.
func (*ListZonesResponse) Descriptor() ([]byte, []int) {
	return file_proto_logistics_zone_proto_rawDescGZIP(), []int{5}
}

func (x *ListZonesResponse) GetResults() []*Zone {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListZonesResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetZoneRequest) Reset() {
	*x = GetZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_zone_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZoneRequest) ProtoMessage() {}

func (x *GetZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_zone_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetZoneRequest.ProtoReflect.Descriptor instead.
func (*GetZoneRequest) Descriptor() ([]byte, []int) {
	return file_proto_logistics_zone_proto_rawDescGZIP(), []int{6}
}

func (x *GetZoneRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Zone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsActive bool   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Created  string `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated  string `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *Zone) Reset() {
	*x = Zone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_zone_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zone) ProtoMessage() {}

func (x *Zone) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_zone_proto_msgTypes[7]
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
	return file_proto_logistics_zone_proto_rawDescGZIP(), []int{7}
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

func (x *Zone) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *Zone) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

var File_proto_logistics_zone_proto protoreflect.FileDescriptor

var file_proto_logistics_zone_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x64,
	0x22, 0x40, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x40, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0x54, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7b, 0x0a, 0x04, 0x5a,
	0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x32, 0xbb, 0x03, 0x0a, 0x05, 0x5a, 0x6f, 0x6e,
	0x65, 0x73, 0x12, 0x5b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5a,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12,
	0x49, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f,
	0x6e, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f,
	0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x54, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73,
	0x12, 0x57, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x32, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x7b, 0x7a, 0x6f,
	0x6e, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x5b, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x2a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x7b, 0x7a, 0x6f,
	0x6e, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_logistics_zone_proto_rawDescOnce sync.Once
	file_proto_logistics_zone_proto_rawDescData = file_proto_logistics_zone_proto_rawDesc
)

func file_proto_logistics_zone_proto_rawDescGZIP() []byte {
	file_proto_logistics_zone_proto_rawDescOnce.Do(func() {
		file_proto_logistics_zone_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_logistics_zone_proto_rawDescData)
	})
	return file_proto_logistics_zone_proto_rawDescData
}

var file_proto_logistics_zone_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_logistics_zone_proto_goTypes = []interface{}{
	(*DeleteZoneRequest)(nil),  // 0: logistics.DeleteZoneRequest
	(*UpdateZoneRequest)(nil),  // 1: logistics.UpdateZoneRequest
	(*CreateZoneRequest)(nil),  // 2: logistics.CreateZoneRequest
	(*CreateZoneResponse)(nil), // 3: logistics.CreateZoneResponse
	(*ListZonesRequest)(nil),   // 4: logistics.ListZonesRequest
	(*ListZonesResponse)(nil),  // 5: logistics.ListZonesResponse
	(*GetZoneRequest)(nil),     // 6: logistics.GetZoneRequest
	(*Zone)(nil),               // 7: logistics.Zone
	(*emptypb.Empty)(nil),      // 8: google.protobuf.Empty
}
var file_proto_logistics_zone_proto_depIdxs = []int32{
	7, // 0: logistics.ListZonesResponse.results:type_name -> logistics.Zone
	2, // 1: logistics.Zones.Create:input_type -> logistics.CreateZoneRequest
	6, // 2: logistics.Zones.Get:input_type -> logistics.GetZoneRequest
	4, // 3: logistics.Zones.List:input_type -> logistics.ListZonesRequest
	1, // 4: logistics.Zones.Update:input_type -> logistics.UpdateZoneRequest
	0, // 5: logistics.Zones.Delete:input_type -> logistics.DeleteZoneRequest
	3, // 6: logistics.Zones.Create:output_type -> logistics.CreateZoneResponse
	7, // 7: logistics.Zones.Get:output_type -> logistics.Zone
	5, // 8: logistics.Zones.List:output_type -> logistics.ListZonesResponse
	7, // 9: logistics.Zones.Update:output_type -> logistics.Zone
	8, // 10: logistics.Zones.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_logistics_zone_proto_init() }
func file_proto_logistics_zone_proto_init() {
	if File_proto_logistics_zone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_logistics_zone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_logistics_zone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateZoneRequest); i {
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
		file_proto_logistics_zone_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateZoneRequest); i {
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
		file_proto_logistics_zone_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateZoneResponse); i {
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
		file_proto_logistics_zone_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListZonesRequest); i {
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
		file_proto_logistics_zone_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListZonesResponse); i {
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
		file_proto_logistics_zone_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetZoneRequest); i {
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
		file_proto_logistics_zone_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_logistics_zone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_logistics_zone_proto_goTypes,
		DependencyIndexes: file_proto_logistics_zone_proto_depIdxs,
		MessageInfos:      file_proto_logistics_zone_proto_msgTypes,
	}.Build()
	File_proto_logistics_zone_proto = out.File
	file_proto_logistics_zone_proto_rawDesc = nil
	file_proto_logistics_zone_proto_goTypes = nil
	file_proto_logistics_zone_proto_depIdxs = nil
}