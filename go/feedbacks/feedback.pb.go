// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.18.1
// source: feedback.proto

package feedbacks

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

type FeedbackPriority int32

const (
	FeedbackPriority_low    FeedbackPriority = 0
	FeedbackPriority_normal FeedbackPriority = 1
	FeedbackPriority_high   FeedbackPriority = 2
)

// Enum value maps for FeedbackPriority.
var (
	FeedbackPriority_name = map[int32]string{
		0: "low",
		1: "normal",
		2: "high",
	}
	FeedbackPriority_value = map[string]int32{
		"low":    0,
		"normal": 1,
		"high":   2,
	}
)

func (x FeedbackPriority) Enum() *FeedbackPriority {
	p := new(FeedbackPriority)
	*p = x
	return p
}

func (x FeedbackPriority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedbackPriority) Descriptor() protoreflect.EnumDescriptor {
	return file_feedback_proto_enumTypes[0].Descriptor()
}

func (FeedbackPriority) Type() protoreflect.EnumType {
	return &file_feedback_proto_enumTypes[0]
}

func (x FeedbackPriority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedbackPriority.Descriptor instead.
func (FeedbackPriority) EnumDescriptor() ([]byte, []int) {
	return file_feedback_proto_rawDescGZIP(), []int{0}
}

type FeedbackCreationType int32

const (
	FeedbackCreationType_H FeedbackCreationType = 0
	FeedbackCreationType_A FeedbackCreationType = 1
)

// Enum value maps for FeedbackCreationType.
var (
	FeedbackCreationType_name = map[int32]string{
		0: "H",
		1: "A",
	}
	FeedbackCreationType_value = map[string]int32{
		"H": 0,
		"A": 1,
	}
)

func (x FeedbackCreationType) Enum() *FeedbackCreationType {
	p := new(FeedbackCreationType)
	*p = x
	return p
}

func (x FeedbackCreationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedbackCreationType) Descriptor() protoreflect.EnumDescriptor {
	return file_feedback_proto_enumTypes[1].Descriptor()
}

func (FeedbackCreationType) Type() protoreflect.EnumType {
	return &file_feedback_proto_enumTypes[1]
}

func (x FeedbackCreationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedbackCreationType.Descriptor instead.
func (FeedbackCreationType) EnumDescriptor() ([]byte, []int) {
	return file_feedback_proto_rawDescGZIP(), []int{1}
}

type ListFeedbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	Sort   string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	// filters
	SourceIds       []int32                `protobuf:"varint,5,rep,packed,name=source_ids,json=sourceIds,proto3" json:"source_ids,omitempty"`
	CreatedBy       int32                  `protobuf:"varint,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	StatusIds       []int32                `protobuf:"varint,7,rep,packed,name=status_ids,json=statusIds,proto3" json:"status_ids,omitempty"`
	SolutionIds     []int32                `protobuf:"varint,8,rep,packed,name=solution_ids,json=solutionIds,proto3" json:"solution_ids,omitempty"`
	UserId          int32                  `protobuf:"varint,9,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DatetimeFrom    string                 `protobuf:"bytes,10,opt,name=datetime_from,json=datetimeFrom,proto3" json:"datetime_from,omitempty"`
	DatetimeTo      string                 `protobuf:"bytes,11,opt,name=datetime_to,json=datetimeTo,proto3" json:"datetime_to,omitempty"`
	ClientStatusIds []int32                `protobuf:"varint,12,rep,packed,name=client_status_ids,json=clientStatusIds,proto3" json:"client_status_ids,omitempty"`
	CreationType    []FeedbackCreationType `protobuf:"varint,13,rep,packed,name=creation_type,json=creationType,proto3,enum=feedbacks.FeedbackCreationType" json:"creation_type,omitempty"`
	IsDouble        bool                   `protobuf:"varint,14,opt,name=is_double,json=isDouble,proto3" json:"is_double,omitempty"`
}

func (x *ListFeedbackRequest) Reset() {
	*x = ListFeedbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeedbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeedbackRequest) ProtoMessage() {}

func (x *ListFeedbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeedbackRequest.ProtoReflect.Descriptor instead.
func (*ListFeedbackRequest) Descriptor() ([]byte, []int) {
	return file_feedback_proto_rawDescGZIP(), []int{0}
}

func (x *ListFeedbackRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListFeedbackRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListFeedbackRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListFeedbackRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *ListFeedbackRequest) GetSourceIds() []int32 {
	if x != nil {
		return x.SourceIds
	}
	return nil
}

func (x *ListFeedbackRequest) GetCreatedBy() int32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *ListFeedbackRequest) GetStatusIds() []int32 {
	if x != nil {
		return x.StatusIds
	}
	return nil
}

func (x *ListFeedbackRequest) GetSolutionIds() []int32 {
	if x != nil {
		return x.SolutionIds
	}
	return nil
}

func (x *ListFeedbackRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListFeedbackRequest) GetDatetimeFrom() string {
	if x != nil {
		return x.DatetimeFrom
	}
	return ""
}

func (x *ListFeedbackRequest) GetDatetimeTo() string {
	if x != nil {
		return x.DatetimeTo
	}
	return ""
}

func (x *ListFeedbackRequest) GetClientStatusIds() []int32 {
	if x != nil {
		return x.ClientStatusIds
	}
	return nil
}

func (x *ListFeedbackRequest) GetCreationType() []FeedbackCreationType {
	if x != nil {
		return x.CreationType
	}
	return nil
}

func (x *ListFeedbackRequest) GetIsDouble() bool {
	if x != nil {
		return x.IsDouble
	}
	return false
}

type Feedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Source       string           `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	User         string           `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Client       string           `protobuf:"bytes,5,opt,name=client,proto3" json:"client,omitempty"`
	ClientStatus string           `protobuf:"bytes,6,opt,name=client_status,json=clientStatus,proto3" json:"client_status,omitempty"`
	Phone        string           `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Email        string           `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	Status       string           `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Solution     string           `protobuf:"bytes,10,opt,name=solution,proto3" json:"solution,omitempty"`
	IsManual     bool             `protobuf:"varint,11,opt,name=is_manual,json=isManual,proto3" json:"is_manual,omitempty"`
	IsDouble     bool             `protobuf:"varint,12,opt,name=is_double,json=isDouble,proto3" json:"is_double,omitempty"`
	Priority     FeedbackPriority `protobuf:"varint,13,opt,name=priority,proto3,enum=feedbacks.FeedbackPriority" json:"priority,omitempty"`
	CreatedAt    string           `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy    string           `protobuf:"bytes,15,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
}

func (x *Feedback) Reset() {
	*x = Feedback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feedback) ProtoMessage() {}

func (x *Feedback) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feedback.ProtoReflect.Descriptor instead.
func (*Feedback) Descriptor() ([]byte, []int) {
	return file_feedback_proto_rawDescGZIP(), []int{1}
}

func (x *Feedback) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Feedback) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Feedback) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Feedback) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *Feedback) GetClientStatus() string {
	if x != nil {
		return x.ClientStatus
	}
	return ""
}

func (x *Feedback) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Feedback) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Feedback) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Feedback) GetSolution() string {
	if x != nil {
		return x.Solution
	}
	return ""
}

func (x *Feedback) GetIsManual() bool {
	if x != nil {
		return x.IsManual
	}
	return false
}

func (x *Feedback) GetIsDouble() bool {
	if x != nil {
		return x.IsDouble
	}
	return false
}

func (x *Feedback) GetPriority() FeedbackPriority {
	if x != nil {
		return x.Priority
	}
	return FeedbackPriority_low
}

func (x *Feedback) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Feedback) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type ListFeedbackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Feedback `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListFeedbackResponse) Reset() {
	*x = ListFeedbackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeedbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeedbackResponse) ProtoMessage() {}

func (x *ListFeedbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeedbackResponse.ProtoReflect.Descriptor instead.
func (*ListFeedbackResponse) Descriptor() ([]byte, []int) {
	return file_feedback_proto_rawDescGZIP(), []int{2}
}

func (x *ListFeedbackResponse) GetResults() []*Feedback {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListFeedbackResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_feedback_proto protoreflect.FileDescriptor

var file_feedback_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x03, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x49, 0x64, 0x73, 0x12, 0x44, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x22, 0x94, 0x03, 0x0a, 0x08, 0x46, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x22, 0x5b, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x2a, 0x31, 0x0a,
	0x10, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x07, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x6e, 0x6f,
	0x72, 0x6d, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x10, 0x02,
	0x2a, 0x24, 0x0a, 0x14, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x05, 0x0a, 0x01, 0x48, 0x10, 0x00, 0x12,
	0x05, 0x0a, 0x01, 0x41, 0x10, 0x01, 0x32, 0x76, 0x0a, 0x10, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1e, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x42, 0x0e,
	0x5a, 0x0c, 0x67, 0x6f, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feedback_proto_rawDescOnce sync.Once
	file_feedback_proto_rawDescData = file_feedback_proto_rawDesc
)

func file_feedback_proto_rawDescGZIP() []byte {
	file_feedback_proto_rawDescOnce.Do(func() {
		file_feedback_proto_rawDescData = protoimpl.X.CompressGZIP(file_feedback_proto_rawDescData)
	})
	return file_feedback_proto_rawDescData
}

var file_feedback_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_feedback_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_feedback_proto_goTypes = []interface{}{
	(FeedbackPriority)(0),        // 0: feedbacks.FeedbackPriority
	(FeedbackCreationType)(0),    // 1: feedbacks.FeedbackCreationType
	(*ListFeedbackRequest)(nil),  // 2: feedbacks.ListFeedbackRequest
	(*Feedback)(nil),             // 3: feedbacks.Feedback
	(*ListFeedbackResponse)(nil), // 4: feedbacks.ListFeedbackResponse
}
var file_feedback_proto_depIdxs = []int32{
	1, // 0: feedbacks.ListFeedbackRequest.creation_type:type_name -> feedbacks.FeedbackCreationType
	0, // 1: feedbacks.Feedback.priority:type_name -> feedbacks.FeedbackPriority
	3, // 2: feedbacks.ListFeedbackResponse.results:type_name -> feedbacks.Feedback
	2, // 3: feedbacks.FeedbacksService.List:input_type -> feedbacks.ListFeedbackRequest
	4, // 4: feedbacks.FeedbacksService.List:output_type -> feedbacks.ListFeedbackResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_feedback_proto_init() }
func file_feedback_proto_init() {
	if File_feedback_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feedback_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeedbackRequest); i {
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
		file_feedback_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feedback); i {
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
		file_feedback_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeedbackResponse); i {
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
			RawDescriptor: file_feedback_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feedback_proto_goTypes,
		DependencyIndexes: file_feedback_proto_depIdxs,
		EnumInfos:         file_feedback_proto_enumTypes,
		MessageInfos:      file_feedback_proto_msgTypes,
	}.Build()
	File_feedback_proto = out.File
	file_feedback_proto_rawDesc = nil
	file_feedback_proto_goTypes = nil
	file_feedback_proto_depIdxs = nil
}
