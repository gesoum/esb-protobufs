// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.9
// source: delivery_method.proto

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

type DeliveryMethodId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeliveryMethodId) Reset() {
	*x = DeliveryMethodId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_method_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryMethodId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryMethodId) ProtoMessage() {}

func (x *DeliveryMethodId) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_method_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryMethodId.ProtoReflect.Descriptor instead.
func (*DeliveryMethodId) Descriptor() ([]byte, []int) {
	return file_delivery_method_proto_rawDescGZIP(), []int{0}
}

func (x *DeliveryMethodId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeliveryMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsActive  bool   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *DeliveryMethod) Reset() {
	*x = DeliveryMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_method_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryMethod) ProtoMessage() {}

func (x *DeliveryMethod) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_method_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryMethod.ProtoReflect.Descriptor instead.
func (*DeliveryMethod) Descriptor() ([]byte, []int) {
	return file_delivery_method_proto_rawDescGZIP(), []int{1}
}

func (x *DeliveryMethod) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeliveryMethod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeliveryMethod) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *DeliveryMethod) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DeliveryMethod) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ListDeliveryMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	Sort   string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *ListDeliveryMethodRequest) Reset() {
	*x = ListDeliveryMethodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_method_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeliveryMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeliveryMethodRequest) ProtoMessage() {}

func (x *ListDeliveryMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_method_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeliveryMethodRequest.ProtoReflect.Descriptor instead.
func (*ListDeliveryMethodRequest) Descriptor() ([]byte, []int) {
	return file_delivery_method_proto_rawDescGZIP(), []int{2}
}

func (x *ListDeliveryMethodRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListDeliveryMethodRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListDeliveryMethodRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListDeliveryMethodRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type ListDeliveryMethodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*DeliveryMethod `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32             `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListDeliveryMethodResponse) Reset() {
	*x = ListDeliveryMethodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_method_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeliveryMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeliveryMethodResponse) ProtoMessage() {}

func (x *ListDeliveryMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_method_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeliveryMethodResponse.ProtoReflect.Descriptor instead.
func (*ListDeliveryMethodResponse) Descriptor() ([]byte, []int) {
	return file_delivery_method_proto_rawDescGZIP(), []int{3}
}

func (x *ListDeliveryMethodResponse) GetResults() []*DeliveryMethod {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListDeliveryMethodResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type SuggestDeliveryMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *SuggestDeliveryMethodRequest) Reset() {
	*x = SuggestDeliveryMethodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_method_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestDeliveryMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestDeliveryMethodRequest) ProtoMessage() {}

func (x *SuggestDeliveryMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_method_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestDeliveryMethodRequest.ProtoReflect.Descriptor instead.
func (*SuggestDeliveryMethodRequest) Descriptor() ([]byte, []int) {
	return file_delivery_method_proto_rawDescGZIP(), []int{4}
}

func (x *SuggestDeliveryMethodRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SuggestDeliveryMethodRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type SuggestDeliveryMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsActive bool   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *SuggestDeliveryMethod) Reset() {
	*x = SuggestDeliveryMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_method_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestDeliveryMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestDeliveryMethod) ProtoMessage() {}

func (x *SuggestDeliveryMethod) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_method_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestDeliveryMethod.ProtoReflect.Descriptor instead.
func (*SuggestDeliveryMethod) Descriptor() ([]byte, []int) {
	return file_delivery_method_proto_rawDescGZIP(), []int{5}
}

func (x *SuggestDeliveryMethod) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SuggestDeliveryMethod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SuggestDeliveryMethod) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type SuggestDeliveryMethodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*SuggestDeliveryMethod `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32                    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *SuggestDeliveryMethodResponse) Reset() {
	*x = SuggestDeliveryMethodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_method_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestDeliveryMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestDeliveryMethodResponse) ProtoMessage() {}

func (x *SuggestDeliveryMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_method_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestDeliveryMethodResponse.ProtoReflect.Descriptor instead.
func (*SuggestDeliveryMethodResponse) Descriptor() ([]byte, []int) {
	return file_delivery_method_proto_rawDescGZIP(), []int{6}
}

func (x *SuggestDeliveryMethodResponse) GetResults() []*SuggestDeliveryMethod {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *SuggestDeliveryMethodResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_delivery_method_proto protoreflect.FileDescriptor

var file_delivery_method_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x8f, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x75, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x67, 0x0a, 0x1a, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x4c, 0x0a, 0x1c, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x22, 0x58, 0x0a, 0x15, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x71, 0x0a, 0x1d, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xac,
	0x05, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x24, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x62,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x1a, 0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2d, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x63, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x1a, 0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2d, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x67, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x19, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a,
	0x01, 0x2a, 0x32, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x63, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x85, 0x01, 0x0a, 0x07, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2d, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x2f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x42, 0x0e, 0x5a,
	0x0c, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_delivery_method_proto_rawDescOnce sync.Once
	file_delivery_method_proto_rawDescData = file_delivery_method_proto_rawDesc
)

func file_delivery_method_proto_rawDescGZIP() []byte {
	file_delivery_method_proto_rawDescOnce.Do(func() {
		file_delivery_method_proto_rawDescData = protoimpl.X.CompressGZIP(file_delivery_method_proto_rawDescData)
	})
	return file_delivery_method_proto_rawDescData
}

var file_delivery_method_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_delivery_method_proto_goTypes = []interface{}{
	(*DeliveryMethodId)(nil),              // 0: logistics.DeliveryMethodId
	(*DeliveryMethod)(nil),                // 1: logistics.DeliveryMethod
	(*ListDeliveryMethodRequest)(nil),     // 2: logistics.ListDeliveryMethodRequest
	(*ListDeliveryMethodResponse)(nil),    // 3: logistics.ListDeliveryMethodResponse
	(*SuggestDeliveryMethodRequest)(nil),  // 4: logistics.SuggestDeliveryMethodRequest
	(*SuggestDeliveryMethod)(nil),         // 5: logistics.SuggestDeliveryMethod
	(*SuggestDeliveryMethodResponse)(nil), // 6: logistics.SuggestDeliveryMethodResponse
	(*emptypb.Empty)(nil),                 // 7: google.protobuf.Empty
}
var file_delivery_method_proto_depIdxs = []int32{
	1, // 0: logistics.ListDeliveryMethodResponse.results:type_name -> logistics.DeliveryMethod
	5, // 1: logistics.SuggestDeliveryMethodResponse.results:type_name -> logistics.SuggestDeliveryMethod
	2, // 2: logistics.DeliveryMethodService.List:input_type -> logistics.ListDeliveryMethodRequest
	1, // 3: logistics.DeliveryMethodService.Create:input_type -> logistics.DeliveryMethod
	0, // 4: logistics.DeliveryMethodService.Get:input_type -> logistics.DeliveryMethodId
	1, // 5: logistics.DeliveryMethodService.Update:input_type -> logistics.DeliveryMethod
	0, // 6: logistics.DeliveryMethodService.Delete:input_type -> logistics.DeliveryMethodId
	4, // 7: logistics.DeliveryMethodService.Suggest:input_type -> logistics.SuggestDeliveryMethodRequest
	3, // 8: logistics.DeliveryMethodService.List:output_type -> logistics.ListDeliveryMethodResponse
	1, // 9: logistics.DeliveryMethodService.Create:output_type -> logistics.DeliveryMethod
	1, // 10: logistics.DeliveryMethodService.Get:output_type -> logistics.DeliveryMethod
	1, // 11: logistics.DeliveryMethodService.Update:output_type -> logistics.DeliveryMethod
	7, // 12: logistics.DeliveryMethodService.Delete:output_type -> google.protobuf.Empty
	6, // 13: logistics.DeliveryMethodService.Suggest:output_type -> logistics.SuggestDeliveryMethodResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_delivery_method_proto_init() }
func file_delivery_method_proto_init() {
	if File_delivery_method_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_delivery_method_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryMethodId); i {
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
		file_delivery_method_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryMethod); i {
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
		file_delivery_method_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeliveryMethodRequest); i {
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
		file_delivery_method_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeliveryMethodResponse); i {
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
		file_delivery_method_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestDeliveryMethodRequest); i {
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
		file_delivery_method_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestDeliveryMethod); i {
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
		file_delivery_method_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestDeliveryMethodResponse); i {
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
			RawDescriptor: file_delivery_method_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_delivery_method_proto_goTypes,
		DependencyIndexes: file_delivery_method_proto_depIdxs,
		MessageInfos:      file_delivery_method_proto_msgTypes,
	}.Build()
	File_delivery_method_proto = out.File
	file_delivery_method_proto_rawDesc = nil
	file_delivery_method_proto_goTypes = nil
	file_delivery_method_proto_depIdxs = nil
}
