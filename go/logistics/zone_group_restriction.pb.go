// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.9
// source: zone_group_restriction.proto

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

type ZoneGroupRestrictionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ZoneGroupRestrictionId) Reset() {
	*x = ZoneGroupRestrictionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_restriction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneGroupRestrictionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneGroupRestrictionId) ProtoMessage() {}

func (x *ZoneGroupRestrictionId) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_restriction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneGroupRestrictionId.ProtoReflect.Descriptor instead.
func (*ZoneGroupRestrictionId) Descriptor() ([]byte, []int) {
	return file_zone_group_restriction_proto_rawDescGZIP(), []int{0}
}

func (x *ZoneGroupRestrictionId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateZoneGroupRestriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneGroupId    int32  `protobuf:"varint,1,opt,name=zone_group_id,json=zoneGroupId,proto3" json:"zone_group_id,omitempty"`
	CountryIsoCode string `protobuf:"bytes,2,opt,name=country_iso_code,json=countryIsoCode,proto3" json:"country_iso_code,omitempty"`
	MessageId      int32  `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	PopupId        int32  `protobuf:"varint,4,opt,name=popup_id,json=popupId,proto3" json:"popup_id,omitempty"`
	IsActive       bool   `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *CreateZoneGroupRestriction) Reset() {
	*x = CreateZoneGroupRestriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_restriction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateZoneGroupRestriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateZoneGroupRestriction) ProtoMessage() {}

func (x *CreateZoneGroupRestriction) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_restriction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateZoneGroupRestriction.ProtoReflect.Descriptor instead.
func (*CreateZoneGroupRestriction) Descriptor() ([]byte, []int) {
	return file_zone_group_restriction_proto_rawDescGZIP(), []int{1}
}

func (x *CreateZoneGroupRestriction) GetZoneGroupId() int32 {
	if x != nil {
		return x.ZoneGroupId
	}
	return 0
}

func (x *CreateZoneGroupRestriction) GetCountryIsoCode() string {
	if x != nil {
		return x.CountryIsoCode
	}
	return ""
}

func (x *CreateZoneGroupRestriction) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *CreateZoneGroupRestriction) GetPopupId() int32 {
	if x != nil {
		return x.PopupId
	}
	return 0
}

func (x *CreateZoneGroupRestriction) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type UpdateZoneGroupRestriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ZoneGroupId    int32  `protobuf:"varint,2,opt,name=zone_group_id,json=zoneGroupId,proto3" json:"zone_group_id,omitempty"`
	CountryIsoCode string `protobuf:"bytes,3,opt,name=country_iso_code,json=countryIsoCode,proto3" json:"country_iso_code,omitempty"`
	MessageId      int32  `protobuf:"varint,4,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	PopupId        int32  `protobuf:"varint,5,opt,name=popup_id,json=popupId,proto3" json:"popup_id,omitempty"`
	IsActive       bool   `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *UpdateZoneGroupRestriction) Reset() {
	*x = UpdateZoneGroupRestriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_restriction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateZoneGroupRestriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateZoneGroupRestriction) ProtoMessage() {}

func (x *UpdateZoneGroupRestriction) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_restriction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateZoneGroupRestriction.ProtoReflect.Descriptor instead.
func (*UpdateZoneGroupRestriction) Descriptor() ([]byte, []int) {
	return file_zone_group_restriction_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateZoneGroupRestriction) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateZoneGroupRestriction) GetZoneGroupId() int32 {
	if x != nil {
		return x.ZoneGroupId
	}
	return 0
}

func (x *UpdateZoneGroupRestriction) GetCountryIsoCode() string {
	if x != nil {
		return x.CountryIsoCode
	}
	return ""
}

func (x *UpdateZoneGroupRestriction) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *UpdateZoneGroupRestriction) GetPopupId() int32 {
	if x != nil {
		return x.PopupId
	}
	return 0
}

func (x *UpdateZoneGroupRestriction) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type ZoneGroupRestriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ZoneGroup *ZoneGroup `protobuf:"bytes,2,opt,name=zone_group,json=zoneGroup,proto3" json:"zone_group,omitempty"`
	Country   *Country   `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Message   *Message   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Popup     *Popup     `protobuf:"bytes,5,opt,name=popup,proto3" json:"popup,omitempty"`
	IsActive  bool       `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt string     `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string     `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ZoneGroupRestriction) Reset() {
	*x = ZoneGroupRestriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_restriction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneGroupRestriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneGroupRestriction) ProtoMessage() {}

func (x *ZoneGroupRestriction) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_restriction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneGroupRestriction.ProtoReflect.Descriptor instead.
func (*ZoneGroupRestriction) Descriptor() ([]byte, []int) {
	return file_zone_group_restriction_proto_rawDescGZIP(), []int{3}
}

func (x *ZoneGroupRestriction) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ZoneGroupRestriction) GetZoneGroup() *ZoneGroup {
	if x != nil {
		return x.ZoneGroup
	}
	return nil
}

func (x *ZoneGroupRestriction) GetCountry() *Country {
	if x != nil {
		return x.Country
	}
	return nil
}

func (x *ZoneGroupRestriction) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *ZoneGroupRestriction) GetPopup() *Popup {
	if x != nil {
		return x.Popup
	}
	return nil
}

func (x *ZoneGroupRestriction) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *ZoneGroupRestriction) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ZoneGroupRestriction) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ListZoneGroupRestrictionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit                int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search               string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	Sort                 string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	IsCountryRestriction bool   `protobuf:"varint,5,opt,name=is_country_restriction,json=isCountryRestriction,proto3" json:"is_country_restriction,omitempty"`
}

func (x *ListZoneGroupRestrictionRequest) Reset() {
	*x = ListZoneGroupRestrictionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_restriction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListZoneGroupRestrictionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZoneGroupRestrictionRequest) ProtoMessage() {}

func (x *ListZoneGroupRestrictionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_restriction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListZoneGroupRestrictionRequest.ProtoReflect.Descriptor instead.
func (*ListZoneGroupRestrictionRequest) Descriptor() ([]byte, []int) {
	return file_zone_group_restriction_proto_rawDescGZIP(), []int{4}
}

func (x *ListZoneGroupRestrictionRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListZoneGroupRestrictionRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListZoneGroupRestrictionRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListZoneGroupRestrictionRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *ListZoneGroupRestrictionRequest) GetIsCountryRestriction() bool {
	if x != nil {
		return x.IsCountryRestriction
	}
	return false
}

type ListZoneGroupRestrictionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ZoneGroupRestriction `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32                   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListZoneGroupRestrictionResponse) Reset() {
	*x = ListZoneGroupRestrictionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_restriction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListZoneGroupRestrictionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZoneGroupRestrictionResponse) ProtoMessage() {}

func (x *ListZoneGroupRestrictionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_restriction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListZoneGroupRestrictionResponse.ProtoReflect.Descriptor instead.
func (*ListZoneGroupRestrictionResponse) Descriptor() ([]byte, []int) {
	return file_zone_group_restriction_proto_rawDescGZIP(), []int{5}
}

func (x *ListZoneGroupRestrictionResponse) GetResults() []*ZoneGroupRestriction {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListZoneGroupRestrictionResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryCode   string `protobuf:"bytes,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	CountryNameRu string `protobuf:"bytes,2,opt,name=country_name_ru,json=countryNameRu,proto3" json:"country_name_ru,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zone_group_restriction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_zone_group_restriction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_zone_group_restriction_proto_rawDescGZIP(), []int{6}
}

func (x *Country) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *Country) GetCountryNameRu() string {
	if x != nil {
		return x.CountryNameRu
	}
	return ""
}

var File_zone_group_restriction_proto protoreflect.FileDescriptor

var file_zone_group_restriction_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x70, 0x6f, 0x70, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x28, 0x0a, 0x16, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc1, 0x01, 0x0a,
	0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x7a,
	0x6f, 0x6e, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x7a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12,
	0x28, 0x0a, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x73, 0x6f, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x49, 0x73, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x70, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x6f, 0x70, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x22, 0xd1, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x0d, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x7a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69,
	0x73, 0x6f, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x73, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x6f, 0x70, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x70, 0x6f, 0x70, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x22, 0xba, 0x02, 0x0a, 0x14, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a,
	0x0a, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f,
	0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x7a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x2c, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x2c, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26,
	0x0a, 0x05, 0x70, 0x6f, 0x70, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x6f, 0x70, 0x75, 0x70, 0x52,
	0x05, 0x70, 0x6f, 0x70, 0x75, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xb1, 0x01, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12,
	0x34, 0x0a, 0x16, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x14, 0x69, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x73, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x54, 0x0a, 0x07, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x75,
	0x32, 0xed, 0x04, 0x0a, 0x1b, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x81, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x75, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x25,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01,
	0x2a, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x75, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x23, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x32, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7a, 0x6f,
	0x6e, 0x65, 0x12, 0x70, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x1f, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x21,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1f, 0x2a, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zone_group_restriction_proto_rawDescOnce sync.Once
	file_zone_group_restriction_proto_rawDescData = file_zone_group_restriction_proto_rawDesc
)

func file_zone_group_restriction_proto_rawDescGZIP() []byte {
	file_zone_group_restriction_proto_rawDescOnce.Do(func() {
		file_zone_group_restriction_proto_rawDescData = protoimpl.X.CompressGZIP(file_zone_group_restriction_proto_rawDescData)
	})
	return file_zone_group_restriction_proto_rawDescData
}

var file_zone_group_restriction_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_zone_group_restriction_proto_goTypes = []interface{}{
	(*ZoneGroupRestrictionId)(nil),           // 0: logistics.ZoneGroupRestrictionId
	(*CreateZoneGroupRestriction)(nil),       // 1: logistics.CreateZoneGroupRestriction
	(*UpdateZoneGroupRestriction)(nil),       // 2: logistics.UpdateZoneGroupRestriction
	(*ZoneGroupRestriction)(nil),             // 3: logistics.ZoneGroupRestriction
	(*ListZoneGroupRestrictionRequest)(nil),  // 4: logistics.ListZoneGroupRestrictionRequest
	(*ListZoneGroupRestrictionResponse)(nil), // 5: logistics.ListZoneGroupRestrictionResponse
	(*Country)(nil),                          // 6: logistics.Country
	(*ZoneGroup)(nil),                        // 7: logistics.ZoneGroup
	(*Message)(nil),                          // 8: logistics.Message
	(*Popup)(nil),                            // 9: logistics.Popup
	(*emptypb.Empty)(nil),                    // 10: google.protobuf.Empty
}
var file_zone_group_restriction_proto_depIdxs = []int32{
	7,  // 0: logistics.ZoneGroupRestriction.zone_group:type_name -> logistics.ZoneGroup
	6,  // 1: logistics.ZoneGroupRestriction.country:type_name -> logistics.Country
	8,  // 2: logistics.ZoneGroupRestriction.message:type_name -> logistics.Message
	9,  // 3: logistics.ZoneGroupRestriction.popup:type_name -> logistics.Popup
	3,  // 4: logistics.ListZoneGroupRestrictionResponse.results:type_name -> logistics.ZoneGroupRestriction
	4,  // 5: logistics.ZoneGroupRestrictionService.List:input_type -> logistics.ListZoneGroupRestrictionRequest
	1,  // 6: logistics.ZoneGroupRestrictionService.Create:input_type -> logistics.CreateZoneGroupRestriction
	2,  // 7: logistics.ZoneGroupRestrictionService.Update:input_type -> logistics.UpdateZoneGroupRestriction
	0,  // 8: logistics.ZoneGroupRestrictionService.Get:input_type -> logistics.ZoneGroupRestrictionId
	0,  // 9: logistics.ZoneGroupRestrictionService.Delete:input_type -> logistics.ZoneGroupRestrictionId
	5,  // 10: logistics.ZoneGroupRestrictionService.List:output_type -> logistics.ListZoneGroupRestrictionResponse
	3,  // 11: logistics.ZoneGroupRestrictionService.Create:output_type -> logistics.ZoneGroupRestriction
	3,  // 12: logistics.ZoneGroupRestrictionService.Update:output_type -> logistics.ZoneGroupRestriction
	3,  // 13: logistics.ZoneGroupRestrictionService.Get:output_type -> logistics.ZoneGroupRestriction
	10, // 14: logistics.ZoneGroupRestrictionService.Delete:output_type -> google.protobuf.Empty
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_zone_group_restriction_proto_init() }
func file_zone_group_restriction_proto_init() {
	if File_zone_group_restriction_proto != nil {
		return
	}
	file_zone_group_proto_init()
	file_message_proto_init()
	file_popup_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_zone_group_restriction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneGroupRestrictionId); i {
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
		file_zone_group_restriction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateZoneGroupRestriction); i {
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
		file_zone_group_restriction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateZoneGroupRestriction); i {
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
		file_zone_group_restriction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneGroupRestriction); i {
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
		file_zone_group_restriction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListZoneGroupRestrictionRequest); i {
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
		file_zone_group_restriction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListZoneGroupRestrictionResponse); i {
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
		file_zone_group_restriction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Country); i {
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
			RawDescriptor: file_zone_group_restriction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zone_group_restriction_proto_goTypes,
		DependencyIndexes: file_zone_group_restriction_proto_depIdxs,
		MessageInfos:      file_zone_group_restriction_proto_msgTypes,
	}.Build()
	File_zone_group_restriction_proto = out.File
	file_zone_group_restriction_proto_rawDesc = nil
	file_zone_group_restriction_proto_goTypes = nil
	file_zone_group_restriction_proto_depIdxs = nil
}
