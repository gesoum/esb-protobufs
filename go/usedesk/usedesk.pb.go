// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.9
// source: proto/usedesk.proto

package usedesk

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

type SaveParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId          int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientId        int32  `protobuf:"varint,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	TicketId        int32  `protobuf:"varint,4,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	Rating          int32  `protobuf:"varint,5,opt,name=rating,proto3" json:"rating,omitempty"`
	CompanyId       int32  `protobuf:"varint,6,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	TicketCommentId int32  `protobuf:"varint,7,opt,name=ticket_comment_id,json=ticketCommentId,proto3" json:"ticket_comment_id,omitempty"`
	ChannelId       int32  `protobuf:"varint,8,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Comment         string `protobuf:"bytes,9,opt,name=comment,proto3" json:"comment,omitempty"`
	CreatedAt       string `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SaveParams) Reset() {
	*x = SaveParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_usedesk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveParams) ProtoMessage() {}

func (x *SaveParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_usedesk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveParams.ProtoReflect.Descriptor instead.
func (*SaveParams) Descriptor() ([]byte, []int) {
	return file_proto_usedesk_proto_rawDescGZIP(), []int{0}
}

func (x *SaveParams) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SaveParams) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SaveParams) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *SaveParams) GetTicketId() int32 {
	if x != nil {
		return x.TicketId
	}
	return 0
}

func (x *SaveParams) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *SaveParams) GetCompanyId() int32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *SaveParams) GetTicketCommentId() int32 {
	if x != nil {
		return x.TicketCommentId
	}
	return 0
}

func (x *SaveParams) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *SaveParams) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *SaveParams) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SaveParams) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SaveResponse) Reset() {
	*x = SaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_usedesk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResponse) ProtoMessage() {}

func (x *SaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_usedesk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResponse.ProtoReflect.Descriptor instead.
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return file_proto_usedesk_proto_rawDescGZIP(), []int{1}
}

func (x *SaveResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_proto_usedesk_proto protoreflect.FileDescriptor

var file_proto_usedesk_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x64, 0x65, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73, 0x65, 0x64, 0x65, 0x73, 0x6b, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x02, 0x0a,
	0x0a, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x1e, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0x47, 0x0a, 0x03, 0x43, 0x53, 0x49, 0x12,
	0x40, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x64, 0x65, 0x73,
	0x6b, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x15, 0x2e, 0x75,
	0x73, 0x65, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x0c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x06, 0x3a, 0x01, 0x2a, 0x22, 0x01,
	0x2f, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x64, 0x65, 0x73, 0x6b, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_usedesk_proto_rawDescOnce sync.Once
	file_proto_usedesk_proto_rawDescData = file_proto_usedesk_proto_rawDesc
)

func file_proto_usedesk_proto_rawDescGZIP() []byte {
	file_proto_usedesk_proto_rawDescOnce.Do(func() {
		file_proto_usedesk_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_usedesk_proto_rawDescData)
	})
	return file_proto_usedesk_proto_rawDescData
}

var file_proto_usedesk_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_usedesk_proto_goTypes = []interface{}{
	(*SaveParams)(nil),   // 0: usedesk.SaveParams
	(*SaveResponse)(nil), // 1: usedesk.SaveResponse
}
var file_proto_usedesk_proto_depIdxs = []int32{
	0, // 0: usedesk.CSI.Save:input_type -> usedesk.SaveParams
	1, // 1: usedesk.CSI.Save:output_type -> usedesk.SaveResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_usedesk_proto_init() }
func file_proto_usedesk_proto_init() {
	if File_proto_usedesk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_usedesk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveParams); i {
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
		file_proto_usedesk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveResponse); i {
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
			RawDescriptor: file_proto_usedesk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_usedesk_proto_goTypes,
		DependencyIndexes: file_proto_usedesk_proto_depIdxs,
		MessageInfos:      file_proto_usedesk_proto_msgTypes,
	}.Build()
	File_proto_usedesk_proto = out.File
	file_proto_usedesk_proto_rawDesc = nil
	file_proto_usedesk_proto_goTypes = nil
	file_proto_usedesk_proto_depIdxs = nil
}
