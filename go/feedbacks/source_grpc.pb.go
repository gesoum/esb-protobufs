// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: source.proto

package feedbacks

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FeedbacksSourcesServiceClient is the client API for FeedbacksSourcesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedbacksSourcesServiceClient interface {
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListSourceResponse, error)
}

type feedbacksSourcesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedbacksSourcesServiceClient(cc grpc.ClientConnInterface) FeedbacksSourcesServiceClient {
	return &feedbacksSourcesServiceClient{cc}
}

func (c *feedbacksSourcesServiceClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListSourceResponse, error) {
	out := new(ListSourceResponse)
	err := c.cc.Invoke(ctx, "/feedbacks.FeedbacksSourcesService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedbacksSourcesServiceServer is the server API for FeedbacksSourcesService service.
// All implementations should embed UnimplementedFeedbacksSourcesServiceServer
// for forward compatibility
type FeedbacksSourcesServiceServer interface {
	List(context.Context, *emptypb.Empty) (*ListSourceResponse, error)
}

// UnimplementedFeedbacksSourcesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFeedbacksSourcesServiceServer struct {
}

func (UnimplementedFeedbacksSourcesServiceServer) List(context.Context, *emptypb.Empty) (*ListSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeFeedbacksSourcesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedbacksSourcesServiceServer will
// result in compilation errors.
type UnsafeFeedbacksSourcesServiceServer interface {
	mustEmbedUnimplementedFeedbacksSourcesServiceServer()
}

func RegisterFeedbacksSourcesServiceServer(s grpc.ServiceRegistrar, srv FeedbacksSourcesServiceServer) {
	s.RegisterService(&FeedbacksSourcesService_ServiceDesc, srv)
}

func _FeedbacksSourcesService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbacksSourcesServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedbacks.FeedbacksSourcesService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbacksSourcesServiceServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedbacksSourcesService_ServiceDesc is the grpc.ServiceDesc for FeedbacksSourcesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedbacksSourcesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feedbacks.FeedbacksSourcesService",
	HandlerType: (*FeedbacksSourcesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _FeedbacksSourcesService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "source.proto",
}
