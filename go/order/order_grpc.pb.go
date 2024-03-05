// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: proto/order.proto

package order

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OutboxService_GetList_FullMethodName     = "/orders.OutboxService/GetList"
	OutboxService_GetByID_FullMethodName     = "/orders.OutboxService/GetByID"
	OutboxService_SetBodyByID_FullMethodName = "/orders.OutboxService/SetBodyByID"
	OutboxService_DeleteByID_FullMethodName  = "/orders.OutboxService/DeleteByID"
)

// OutboxServiceClient is the client API for OutboxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OutboxServiceClient interface {
	GetList(ctx context.Context, in *GetListRequestForOrder, opts ...grpc.CallOption) (*GetListOutBoxResponse, error)
	GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDOutBoxResponse, error)
	SetBodyByID(ctx context.Context, in *SetByIDOutboxBodyRequest, opts ...grpc.CallOption) (*SetByIDOutboxBodyResponse, error)
	DeleteByID(ctx context.Context, in *DeleteByIDOutboxRequest, opts ...grpc.CallOption) (*DeleteByIDOutboxResponse, error)
}

type outboxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOutboxServiceClient(cc grpc.ClientConnInterface) OutboxServiceClient {
	return &outboxServiceClient{cc}
}

func (c *outboxServiceClient) GetList(ctx context.Context, in *GetListRequestForOrder, opts ...grpc.CallOption) (*GetListOutBoxResponse, error) {
	out := new(GetListOutBoxResponse)
	err := c.cc.Invoke(ctx, OutboxService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboxServiceClient) GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDOutBoxResponse, error) {
	out := new(GetByIDOutBoxResponse)
	err := c.cc.Invoke(ctx, OutboxService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboxServiceClient) SetBodyByID(ctx context.Context, in *SetByIDOutboxBodyRequest, opts ...grpc.CallOption) (*SetByIDOutboxBodyResponse, error) {
	out := new(SetByIDOutboxBodyResponse)
	err := c.cc.Invoke(ctx, OutboxService_SetBodyByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboxServiceClient) DeleteByID(ctx context.Context, in *DeleteByIDOutboxRequest, opts ...grpc.CallOption) (*DeleteByIDOutboxResponse, error) {
	out := new(DeleteByIDOutboxResponse)
	err := c.cc.Invoke(ctx, OutboxService_DeleteByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OutboxServiceServer is the server API for OutboxService service.
// All implementations should embed UnimplementedOutboxServiceServer
// for forward compatibility
type OutboxServiceServer interface {
	GetList(context.Context, *GetListRequestForOrder) (*GetListOutBoxResponse, error)
	GetByID(context.Context, *GetByIDRequest) (*GetByIDOutBoxResponse, error)
	SetBodyByID(context.Context, *SetByIDOutboxBodyRequest) (*SetByIDOutboxBodyResponse, error)
	DeleteByID(context.Context, *DeleteByIDOutboxRequest) (*DeleteByIDOutboxResponse, error)
}

// UnimplementedOutboxServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOutboxServiceServer struct {
}

func (UnimplementedOutboxServiceServer) GetList(context.Context, *GetListRequestForOrder) (*GetListOutBoxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedOutboxServiceServer) GetByID(context.Context, *GetByIDRequest) (*GetByIDOutBoxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedOutboxServiceServer) SetBodyByID(context.Context, *SetByIDOutboxBodyRequest) (*SetByIDOutboxBodyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBodyByID not implemented")
}
func (UnimplementedOutboxServiceServer) DeleteByID(context.Context, *DeleteByIDOutboxRequest) (*DeleteByIDOutboxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByID not implemented")
}

// UnsafeOutboxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OutboxServiceServer will
// result in compilation errors.
type UnsafeOutboxServiceServer interface {
	mustEmbedUnimplementedOutboxServiceServer()
}

func RegisterOutboxServiceServer(s grpc.ServiceRegistrar, srv OutboxServiceServer) {
	s.RegisterService(&OutboxService_ServiceDesc, srv)
}

func _OutboxService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequestForOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OutboxService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServiceServer).GetList(ctx, req.(*GetListRequestForOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboxService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OutboxService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServiceServer).GetByID(ctx, req.(*GetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboxService_SetBodyByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetByIDOutboxBodyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServiceServer).SetBodyByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OutboxService_SetBodyByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServiceServer).SetBodyByID(ctx, req.(*SetByIDOutboxBodyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboxService_DeleteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteByIDOutboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServiceServer).DeleteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OutboxService_DeleteByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServiceServer).DeleteByID(ctx, req.(*DeleteByIDOutboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OutboxService_ServiceDesc is the grpc.ServiceDesc for OutboxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OutboxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.OutboxService",
	HandlerType: (*OutboxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _OutboxService_GetList_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _OutboxService_GetByID_Handler,
		},
		{
			MethodName: "SetBodyByID",
			Handler:    _OutboxService_SetBodyByID_Handler,
		},
		{
			MethodName: "DeleteByID",
			Handler:    _OutboxService_DeleteByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/order.proto",
}

const (
	OrderHistoryService_GetList_FullMethodName = "/orders.OrderHistoryService/GetList"
	OrderHistoryService_GetByID_FullMethodName = "/orders.OrderHistoryService/GetByID"
)

// OrderHistoryServiceClient is the client API for OrderHistoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderHistoryServiceClient interface {
	GetList(ctx context.Context, in *GetListRequestForOrderAndPosition, opts ...grpc.CallOption) (*GetListOrderHistoryResponse, error)
	GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDOrderHistoryResponse, error)
}

type orderHistoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderHistoryServiceClient(cc grpc.ClientConnInterface) OrderHistoryServiceClient {
	return &orderHistoryServiceClient{cc}
}

func (c *orderHistoryServiceClient) GetList(ctx context.Context, in *GetListRequestForOrderAndPosition, opts ...grpc.CallOption) (*GetListOrderHistoryResponse, error) {
	out := new(GetListOrderHistoryResponse)
	err := c.cc.Invoke(ctx, OrderHistoryService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderHistoryServiceClient) GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDOrderHistoryResponse, error) {
	out := new(GetByIDOrderHistoryResponse)
	err := c.cc.Invoke(ctx, OrderHistoryService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderHistoryServiceServer is the server API for OrderHistoryService service.
// All implementations should embed UnimplementedOrderHistoryServiceServer
// for forward compatibility
type OrderHistoryServiceServer interface {
	GetList(context.Context, *GetListRequestForOrderAndPosition) (*GetListOrderHistoryResponse, error)
	GetByID(context.Context, *GetByIDRequest) (*GetByIDOrderHistoryResponse, error)
}

// UnimplementedOrderHistoryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderHistoryServiceServer struct {
}

func (UnimplementedOrderHistoryServiceServer) GetList(context.Context, *GetListRequestForOrderAndPosition) (*GetListOrderHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedOrderHistoryServiceServer) GetByID(context.Context, *GetByIDRequest) (*GetByIDOrderHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}

// UnsafeOrderHistoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderHistoryServiceServer will
// result in compilation errors.
type UnsafeOrderHistoryServiceServer interface {
	mustEmbedUnimplementedOrderHistoryServiceServer()
}

func RegisterOrderHistoryServiceServer(s grpc.ServiceRegistrar, srv OrderHistoryServiceServer) {
	s.RegisterService(&OrderHistoryService_ServiceDesc, srv)
}

func _OrderHistoryService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequestForOrderAndPosition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderHistoryServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderHistoryService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderHistoryServiceServer).GetList(ctx, req.(*GetListRequestForOrderAndPosition))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderHistoryService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderHistoryServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderHistoryService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderHistoryServiceServer).GetByID(ctx, req.(*GetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderHistoryService_ServiceDesc is the grpc.ServiceDesc for OrderHistoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderHistoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.OrderHistoryService",
	HandlerType: (*OrderHistoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _OrderHistoryService_GetList_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _OrderHistoryService_GetByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/order.proto",
}

const (
	PositionHistoryService_GetList_FullMethodName = "/orders.PositionHistoryService/GetList"
	PositionHistoryService_GetByID_FullMethodName = "/orders.PositionHistoryService/GetByID"
)

// PositionHistoryServiceClient is the client API for PositionHistoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PositionHistoryServiceClient interface {
	GetList(ctx context.Context, in *GetListRequestForPosition, opts ...grpc.CallOption) (*GetListPositionHistoryResponse, error)
	GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDPositionHistoryResponse, error)
}

type positionHistoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionHistoryServiceClient(cc grpc.ClientConnInterface) PositionHistoryServiceClient {
	return &positionHistoryServiceClient{cc}
}

func (c *positionHistoryServiceClient) GetList(ctx context.Context, in *GetListRequestForPosition, opts ...grpc.CallOption) (*GetListPositionHistoryResponse, error) {
	out := new(GetListPositionHistoryResponse)
	err := c.cc.Invoke(ctx, PositionHistoryService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionHistoryServiceClient) GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDPositionHistoryResponse, error) {
	out := new(GetByIDPositionHistoryResponse)
	err := c.cc.Invoke(ctx, PositionHistoryService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionHistoryServiceServer is the server API for PositionHistoryService service.
// All implementations should embed UnimplementedPositionHistoryServiceServer
// for forward compatibility
type PositionHistoryServiceServer interface {
	GetList(context.Context, *GetListRequestForPosition) (*GetListPositionHistoryResponse, error)
	GetByID(context.Context, *GetByIDRequest) (*GetByIDPositionHistoryResponse, error)
}

// UnimplementedPositionHistoryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPositionHistoryServiceServer struct {
}

func (UnimplementedPositionHistoryServiceServer) GetList(context.Context, *GetListRequestForPosition) (*GetListPositionHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedPositionHistoryServiceServer) GetByID(context.Context, *GetByIDRequest) (*GetByIDPositionHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}

// UnsafePositionHistoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionHistoryServiceServer will
// result in compilation errors.
type UnsafePositionHistoryServiceServer interface {
	mustEmbedUnimplementedPositionHistoryServiceServer()
}

func RegisterPositionHistoryServiceServer(s grpc.ServiceRegistrar, srv PositionHistoryServiceServer) {
	s.RegisterService(&PositionHistoryService_ServiceDesc, srv)
}

func _PositionHistoryService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequestForPosition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionHistoryServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PositionHistoryService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionHistoryServiceServer).GetList(ctx, req.(*GetListRequestForPosition))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionHistoryService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionHistoryServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PositionHistoryService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionHistoryServiceServer).GetByID(ctx, req.(*GetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PositionHistoryService_ServiceDesc is the grpc.ServiceDesc for PositionHistoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PositionHistoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.PositionHistoryService",
	HandlerType: (*PositionHistoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _PositionHistoryService_GetList_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _PositionHistoryService_GetByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/order.proto",
}

const (
	OrderService_GetList_FullMethodName         = "/orders.OrderService/GetList"
	OrderService_GetByID_FullMethodName         = "/orders.OrderService/GetByID"
	OrderService_GetListForAdmin_FullMethodName = "/orders.OrderService/GetListForAdmin"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	GetList(ctx context.Context, in *GetListRequestForPosition, opts ...grpc.CallOption) (*GetListOrderResponse, error)
	GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDOrderResponse, error)
	GetListForAdmin(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListForAdminResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) GetList(ctx context.Context, in *GetListRequestForPosition, opts ...grpc.CallOption) (*GetListOrderResponse, error) {
	out := new(GetListOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDOrderResponse, error) {
	out := new(GetByIDOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetListForAdmin(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListForAdminResponse, error) {
	out := new(GetListForAdminResponse)
	err := c.cc.Invoke(ctx, OrderService_GetListForAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations should embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	GetList(context.Context, *GetListRequestForPosition) (*GetListOrderResponse, error)
	GetByID(context.Context, *GetByIDRequest) (*GetByIDOrderResponse, error)
	GetListForAdmin(context.Context, *GetListRequest) (*GetListForAdminResponse, error)
}

// UnimplementedOrderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) GetList(context.Context, *GetListRequestForPosition) (*GetListOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedOrderServiceServer) GetByID(context.Context, *GetByIDRequest) (*GetByIDOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedOrderServiceServer) GetListForAdmin(context.Context, *GetListRequest) (*GetListForAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListForAdmin not implemented")
}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequestForPosition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetList(ctx, req.(*GetListRequestForPosition))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetByID(ctx, req.(*GetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetListForAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetListForAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetListForAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetListForAdmin(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _OrderService_GetList_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _OrderService_GetByID_Handler,
		},
		{
			MethodName: "GetListForAdmin",
			Handler:    _OrderService_GetListForAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/order.proto",
}
