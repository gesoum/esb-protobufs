// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.18.1
// source: transport_company.proto

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

type TransportCompanyId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TransportCompanyId) Reset() {
	*x = TransportCompanyId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransportCompanyId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportCompanyId) ProtoMessage() {}

func (x *TransportCompanyId) ProtoReflect() protoreflect.Message {
	mi := &file_transport_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportCompanyId.ProtoReflect.Descriptor instead.
func (*TransportCompanyId) Descriptor() ([]byte, []int) {
	return file_transport_company_proto_rawDescGZIP(), []int{0}
}

func (x *TransportCompanyId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TransportCompany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TrackUrl                     string `protobuf:"bytes,3,opt,name=track_url,json=trackUrl,proto3" json:"track_url,omitempty"`
	HasIntegration               bool   `protobuf:"varint,4,opt,name=has_integration,json=hasIntegration,proto3" json:"has_integration,omitempty"`
	IsDeliveryUseOnlyWorkingDays bool   `protobuf:"varint,5,opt,name=is_delivery_use_only_working_days,json=isDeliveryUseOnlyWorkingDays,proto3" json:"is_delivery_use_only_working_days,omitempty"`
	IsActive                     bool   `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt                    string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt                    string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *TransportCompany) Reset() {
	*x = TransportCompany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_company_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransportCompany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportCompany) ProtoMessage() {}

func (x *TransportCompany) ProtoReflect() protoreflect.Message {
	mi := &file_transport_company_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportCompany.ProtoReflect.Descriptor instead.
func (*TransportCompany) Descriptor() ([]byte, []int) {
	return file_transport_company_proto_rawDescGZIP(), []int{1}
}

func (x *TransportCompany) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TransportCompany) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TransportCompany) GetTrackUrl() string {
	if x != nil {
		return x.TrackUrl
	}
	return ""
}

func (x *TransportCompany) GetHasIntegration() bool {
	if x != nil {
		return x.HasIntegration
	}
	return false
}

func (x *TransportCompany) GetIsDeliveryUseOnlyWorkingDays() bool {
	if x != nil {
		return x.IsDeliveryUseOnlyWorkingDays
	}
	return false
}

func (x *TransportCompany) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *TransportCompany) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *TransportCompany) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ListTransportCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	Sort   string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *ListTransportCompanyRequest) Reset() {
	*x = ListTransportCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_company_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransportCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransportCompanyRequest) ProtoMessage() {}

func (x *ListTransportCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_company_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransportCompanyRequest.ProtoReflect.Descriptor instead.
func (*ListTransportCompanyRequest) Descriptor() ([]byte, []int) {
	return file_transport_company_proto_rawDescGZIP(), []int{2}
}

func (x *ListTransportCompanyRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTransportCompanyRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListTransportCompanyRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListTransportCompanyRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type ListTransportCompanyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*TransportCompany `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListTransportCompanyResponse) Reset() {
	*x = ListTransportCompanyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_company_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransportCompanyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransportCompanyResponse) ProtoMessage() {}

func (x *ListTransportCompanyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_company_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransportCompanyResponse.ProtoReflect.Descriptor instead.
func (*ListTransportCompanyResponse) Descriptor() ([]byte, []int) {
	return file_transport_company_proto_rawDescGZIP(), []int{3}
}

func (x *ListTransportCompanyResponse) GetResults() []*TransportCompany {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListTransportCompanyResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_transport_company_proto protoreflect.FileDescriptor

var file_transport_company_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x24, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa0, 0x02, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x68,
	0x61, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x21, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x1c, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x4f, 0x6e,
	0x6c, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x79, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x77, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x22, 0x6b, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xc2,
	0x04, 0x0a, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7a, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x26, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x68, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x1b, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x3a, 0x01, 0x2a,
	0x12, 0x69, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x1a, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6d, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x1a, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22,
	0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x32, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x67, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x2a, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_company_proto_rawDescOnce sync.Once
	file_transport_company_proto_rawDescData = file_transport_company_proto_rawDesc
)

func file_transport_company_proto_rawDescGZIP() []byte {
	file_transport_company_proto_rawDescOnce.Do(func() {
		file_transport_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_company_proto_rawDescData)
	})
	return file_transport_company_proto_rawDescData
}

var file_transport_company_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_transport_company_proto_goTypes = []interface{}{
	(*TransportCompanyId)(nil),           // 0: logistics.TransportCompanyId
	(*TransportCompany)(nil),             // 1: logistics.TransportCompany
	(*ListTransportCompanyRequest)(nil),  // 2: logistics.ListTransportCompanyRequest
	(*ListTransportCompanyResponse)(nil), // 3: logistics.ListTransportCompanyResponse
	(*emptypb.Empty)(nil),                // 4: google.protobuf.Empty
}
var file_transport_company_proto_depIdxs = []int32{
	1, // 0: logistics.ListTransportCompanyResponse.results:type_name -> logistics.TransportCompany
	2, // 1: logistics.TransportCompanyService.List:input_type -> logistics.ListTransportCompanyRequest
	1, // 2: logistics.TransportCompanyService.Create:input_type -> logistics.TransportCompany
	0, // 3: logistics.TransportCompanyService.Get:input_type -> logistics.TransportCompanyId
	1, // 4: logistics.TransportCompanyService.Update:input_type -> logistics.TransportCompany
	0, // 5: logistics.TransportCompanyService.Delete:input_type -> logistics.TransportCompanyId
	3, // 6: logistics.TransportCompanyService.List:output_type -> logistics.ListTransportCompanyResponse
	1, // 7: logistics.TransportCompanyService.Create:output_type -> logistics.TransportCompany
	1, // 8: logistics.TransportCompanyService.Get:output_type -> logistics.TransportCompany
	1, // 9: logistics.TransportCompanyService.Update:output_type -> logistics.TransportCompany
	4, // 10: logistics.TransportCompanyService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transport_company_proto_init() }
func file_transport_company_proto_init() {
	if File_transport_company_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_company_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransportCompanyId); i {
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
		file_transport_company_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransportCompany); i {
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
		file_transport_company_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTransportCompanyRequest); i {
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
		file_transport_company_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTransportCompanyResponse); i {
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
			RawDescriptor: file_transport_company_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transport_company_proto_goTypes,
		DependencyIndexes: file_transport_company_proto_depIdxs,
		MessageInfos:      file_transport_company_proto_msgTypes,
	}.Build()
	File_transport_company_proto = out.File
	file_transport_company_proto_rawDesc = nil
	file_transport_company_proto_goTypes = nil
	file_transport_company_proto_depIdxs = nil
}
