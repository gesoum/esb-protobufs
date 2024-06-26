// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.9
// source: payment_method.proto

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

type PaymentMethodId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PaymentMethodId) Reset() {
	*x = PaymentMethodId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_method_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethodId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethodId) ProtoMessage() {}

func (x *PaymentMethodId) ProtoReflect() protoreflect.Message {
	mi := &file_payment_method_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethodId.ProtoReflect.Descriptor instead.
func (*PaymentMethodId) Descriptor() ([]byte, []int) {
	return file_payment_method_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentMethodId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PaymentMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TitleRu       string `protobuf:"bytes,2,opt,name=title_ru,json=titleRu,proto3" json:"title_ru,omitempty"`
	TitleInternal string `protobuf:"bytes,3,opt,name=title_internal,json=titleInternal,proto3" json:"title_internal,omitempty"`
	Code          string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	IsActive      bool   `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt     string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *PaymentMethod) Reset() {
	*x = PaymentMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_method_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethod) ProtoMessage() {}

func (x *PaymentMethod) ProtoReflect() protoreflect.Message {
	mi := &file_payment_method_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethod.ProtoReflect.Descriptor instead.
func (*PaymentMethod) Descriptor() ([]byte, []int) {
	return file_payment_method_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentMethod) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentMethod) GetTitleRu() string {
	if x != nil {
		return x.TitleRu
	}
	return ""
}

func (x *PaymentMethod) GetTitleInternal() string {
	if x != nil {
		return x.TitleInternal
	}
	return ""
}

func (x *PaymentMethod) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PaymentMethod) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *PaymentMethod) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PaymentMethod) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ListPaymentMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	Sort   string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *ListPaymentMethodRequest) Reset() {
	*x = ListPaymentMethodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_method_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPaymentMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPaymentMethodRequest) ProtoMessage() {}

func (x *ListPaymentMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_method_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPaymentMethodRequest.ProtoReflect.Descriptor instead.
func (*ListPaymentMethodRequest) Descriptor() ([]byte, []int) {
	return file_payment_method_proto_rawDescGZIP(), []int{2}
}

func (x *ListPaymentMethodRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListPaymentMethodRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListPaymentMethodRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListPaymentMethodRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type ListPaymentMethodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*PaymentMethod `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListPaymentMethodResponse) Reset() {
	*x = ListPaymentMethodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_method_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPaymentMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPaymentMethodResponse) ProtoMessage() {}

func (x *ListPaymentMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_method_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPaymentMethodResponse.ProtoReflect.Descriptor instead.
func (*ListPaymentMethodResponse) Descriptor() ([]byte, []int) {
	return file_payment_method_proto_rawDescGZIP(), []int{3}
}

func (x *ListPaymentMethodResponse) GetResults() []*PaymentMethod {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListPaymentMethodResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type SuggestPaymentMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *SuggestPaymentMethodRequest) Reset() {
	*x = SuggestPaymentMethodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_method_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestPaymentMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestPaymentMethodRequest) ProtoMessage() {}

func (x *SuggestPaymentMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_method_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestPaymentMethodRequest.ProtoReflect.Descriptor instead.
func (*SuggestPaymentMethodRequest) Descriptor() ([]byte, []int) {
	return file_payment_method_proto_rawDescGZIP(), []int{4}
}

func (x *SuggestPaymentMethodRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SuggestPaymentMethodRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type SuggestPaymentMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TitleRu  string `protobuf:"bytes,2,opt,name=title_ru,json=titleRu,proto3" json:"title_ru,omitempty"`
	IsActive bool   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *SuggestPaymentMethod) Reset() {
	*x = SuggestPaymentMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_method_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestPaymentMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestPaymentMethod) ProtoMessage() {}

func (x *SuggestPaymentMethod) ProtoReflect() protoreflect.Message {
	mi := &file_payment_method_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestPaymentMethod.ProtoReflect.Descriptor instead.
func (*SuggestPaymentMethod) Descriptor() ([]byte, []int) {
	return file_payment_method_proto_rawDescGZIP(), []int{5}
}

func (x *SuggestPaymentMethod) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SuggestPaymentMethod) GetTitleRu() string {
	if x != nil {
		return x.TitleRu
	}
	return ""
}

func (x *SuggestPaymentMethod) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type SuggestPaymentMethodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*SuggestPaymentMethod `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32                   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *SuggestPaymentMethodResponse) Reset() {
	*x = SuggestPaymentMethodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_method_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestPaymentMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestPaymentMethodResponse) ProtoMessage() {}

func (x *SuggestPaymentMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_method_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestPaymentMethodResponse.ProtoReflect.Descriptor instead.
func (*SuggestPaymentMethodResponse) Descriptor() ([]byte, []int) {
	return file_payment_method_proto_rawDescGZIP(), []int{6}
}

func (x *SuggestPaymentMethodResponse) GetResults() []*SuggestPaymentMethod {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *SuggestPaymentMethodResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_payment_method_proto protoreflect.FileDescriptor

var file_payment_method_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0xd0, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x72, 0x75, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x75, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x74, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x65, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x4b, 0x0a, 0x1b, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x5e, 0x0a, 0x14,
	0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x72, 0x75,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x75, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x6f, 0x0a, 0x1c,
	0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x9a, 0x05,
	0x0a, 0x14, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x5f, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x18, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a,
	0x01, 0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x60, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x1a, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x1a, 0x18, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12,
	0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x64, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x1a, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x61, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x82, 0x01, 0x0a, 0x07, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x2f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x6f,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_payment_method_proto_rawDescOnce sync.Once
	file_payment_method_proto_rawDescData = file_payment_method_proto_rawDesc
)

func file_payment_method_proto_rawDescGZIP() []byte {
	file_payment_method_proto_rawDescOnce.Do(func() {
		file_payment_method_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_method_proto_rawDescData)
	})
	return file_payment_method_proto_rawDescData
}

var file_payment_method_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_payment_method_proto_goTypes = []interface{}{
	(*PaymentMethodId)(nil),              // 0: logistics.PaymentMethodId
	(*PaymentMethod)(nil),                // 1: logistics.PaymentMethod
	(*ListPaymentMethodRequest)(nil),     // 2: logistics.ListPaymentMethodRequest
	(*ListPaymentMethodResponse)(nil),    // 3: logistics.ListPaymentMethodResponse
	(*SuggestPaymentMethodRequest)(nil),  // 4: logistics.SuggestPaymentMethodRequest
	(*SuggestPaymentMethod)(nil),         // 5: logistics.SuggestPaymentMethod
	(*SuggestPaymentMethodResponse)(nil), // 6: logistics.SuggestPaymentMethodResponse
	(*emptypb.Empty)(nil),                // 7: google.protobuf.Empty
}
var file_payment_method_proto_depIdxs = []int32{
	1, // 0: logistics.ListPaymentMethodResponse.results:type_name -> logistics.PaymentMethod
	5, // 1: logistics.SuggestPaymentMethodResponse.results:type_name -> logistics.SuggestPaymentMethod
	2, // 2: logistics.PaymentMethodService.List:input_type -> logistics.ListPaymentMethodRequest
	1, // 3: logistics.PaymentMethodService.Create:input_type -> logistics.PaymentMethod
	0, // 4: logistics.PaymentMethodService.Get:input_type -> logistics.PaymentMethodId
	1, // 5: logistics.PaymentMethodService.Update:input_type -> logistics.PaymentMethod
	0, // 6: logistics.PaymentMethodService.Delete:input_type -> logistics.PaymentMethodId
	4, // 7: logistics.PaymentMethodService.Suggest:input_type -> logistics.SuggestPaymentMethodRequest
	3, // 8: logistics.PaymentMethodService.List:output_type -> logistics.ListPaymentMethodResponse
	1, // 9: logistics.PaymentMethodService.Create:output_type -> logistics.PaymentMethod
	1, // 10: logistics.PaymentMethodService.Get:output_type -> logistics.PaymentMethod
	1, // 11: logistics.PaymentMethodService.Update:output_type -> logistics.PaymentMethod
	7, // 12: logistics.PaymentMethodService.Delete:output_type -> google.protobuf.Empty
	6, // 13: logistics.PaymentMethodService.Suggest:output_type -> logistics.SuggestPaymentMethodResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_payment_method_proto_init() }
func file_payment_method_proto_init() {
	if File_payment_method_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_method_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethodId); i {
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
		file_payment_method_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethod); i {
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
		file_payment_method_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPaymentMethodRequest); i {
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
		file_payment_method_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPaymentMethodResponse); i {
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
		file_payment_method_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestPaymentMethodRequest); i {
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
		file_payment_method_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestPaymentMethod); i {
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
		file_payment_method_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestPaymentMethodResponse); i {
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
			RawDescriptor: file_payment_method_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_method_proto_goTypes,
		DependencyIndexes: file_payment_method_proto_depIdxs,
		MessageInfos:      file_payment_method_proto_msgTypes,
	}.Build()
	File_payment_method_proto = out.File
	file_payment_method_proto_rawDesc = nil
	file_payment_method_proto_goTypes = nil
	file_payment_method_proto_depIdxs = nil
}
