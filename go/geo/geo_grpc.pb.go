// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: proto/geo.proto

package geo

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

// GeoClient is the client API for Geo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeoClient interface {
	// Suggest по странам. Возвращет страны из БД
	SuggestCountry(ctx context.Context, in *SuggestCountryParams, opts ...grpc.CallOption) (*SuggestCountryResponse, error)
	// Детальная информация по стране из БД
	CountryDetails(ctx context.Context, in *CountryDetailsParams, opts ...grpc.CallOption) (*Country, error)
	// Suggest по городам. Если ничего не найдено - возвращает список городов по умолчанию (см. DefaultCityList)
	// Первый возвращаемый город - город, определённый по IP (если удалось)
	SuggestCity(ctx context.Context, in *SuggestCityParams, opts ...grpc.CallOption) (*SuggestCityResponse, error)
	// Детали города (используется нормализация дадаты)
	CityDetails(ctx context.Context, in *CityDetailsParams, opts ...grpc.CallOption) (*City, error)
	// Детали города по GeoID
	CityDetailsByGeoID(ctx context.Context, in *CityDetailsByGeoIDParams, opts ...grpc.CallOption) (*City, error)
	// Метод возвращает список городов по-умолчанию, в зависимости от количества заказов
	DefaultCityList(ctx context.Context, in *DefaultCityListParams, opts ...grpc.CallOption) (*SuggestCityResponse, error)
	// Определение города по IP (MaxMind & Dadata + cache)
	CityDetailsByIP(ctx context.Context, in *CityDetailsByIPParams, opts ...grpc.CallOption) (*City, error)
	// Suggest по адресам
	SuggestAddress(ctx context.Context, in *SuggestAddressParams, opts ...grpc.CallOption) (*SuggestAddressResponse, error)
	// Нормализация адреса
	AddressDetails(ctx context.Context, in *AddressDetailsParams, opts ...grpc.CallOption) (*Address, error)
	// Детали адреса по GeoID
	AddressDetailsByGeoID(ctx context.Context, in *AddressDetailsByGeoIDParams, opts ...grpc.CallOption) (*Address, error)
}

type geoClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoClient(cc grpc.ClientConnInterface) GeoClient {
	return &geoClient{cc}
}

func (c *geoClient) SuggestCountry(ctx context.Context, in *SuggestCountryParams, opts ...grpc.CallOption) (*SuggestCountryResponse, error) {
	out := new(SuggestCountryResponse)
	err := c.cc.Invoke(ctx, "/geo.geo/SuggestCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoClient) CountryDetails(ctx context.Context, in *CountryDetailsParams, opts ...grpc.CallOption) (*Country, error) {
	out := new(Country)
	err := c.cc.Invoke(ctx, "/geo.geo/CountryDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoClient) SuggestCity(ctx context.Context, in *SuggestCityParams, opts ...grpc.CallOption) (*SuggestCityResponse, error) {
	out := new(SuggestCityResponse)
	err := c.cc.Invoke(ctx, "/geo.geo/SuggestCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoClient) CityDetails(ctx context.Context, in *CityDetailsParams, opts ...grpc.CallOption) (*City, error) {
	out := new(City)
	err := c.cc.Invoke(ctx, "/geo.geo/CityDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoClient) CityDetailsByGeoID(ctx context.Context, in *CityDetailsByGeoIDParams, opts ...grpc.CallOption) (*City, error) {
	out := new(City)
	err := c.cc.Invoke(ctx, "/geo.geo/CityDetailsByGeoID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoClient) DefaultCityList(ctx context.Context, in *DefaultCityListParams, opts ...grpc.CallOption) (*SuggestCityResponse, error) {
	out := new(SuggestCityResponse)
	err := c.cc.Invoke(ctx, "/geo.geo/DefaultCityList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoClient) CityDetailsByIP(ctx context.Context, in *CityDetailsByIPParams, opts ...grpc.CallOption) (*City, error) {
	out := new(City)
	err := c.cc.Invoke(ctx, "/geo.geo/CityDetailsByIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoClient) SuggestAddress(ctx context.Context, in *SuggestAddressParams, opts ...grpc.CallOption) (*SuggestAddressResponse, error) {
	out := new(SuggestAddressResponse)
	err := c.cc.Invoke(ctx, "/geo.geo/SuggestAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoClient) AddressDetails(ctx context.Context, in *AddressDetailsParams, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/geo.geo/AddressDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoClient) AddressDetailsByGeoID(ctx context.Context, in *AddressDetailsByGeoIDParams, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/geo.geo/AddressDetailsByGeoID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoServer is the server API for Geo service.
// All implementations should embed UnimplementedGeoServer
// for forward compatibility
type GeoServer interface {
	// Suggest по странам. Возвращет страны из БД
	SuggestCountry(context.Context, *SuggestCountryParams) (*SuggestCountryResponse, error)
	// Детальная информация по стране из БД
	CountryDetails(context.Context, *CountryDetailsParams) (*Country, error)
	// Suggest по городам. Если ничего не найдено - возвращает список городов по умолчанию (см. DefaultCityList)
	// Первый возвращаемый город - город, определённый по IP (если удалось)
	SuggestCity(context.Context, *SuggestCityParams) (*SuggestCityResponse, error)
	// Детали города (используется нормализация дадаты)
	CityDetails(context.Context, *CityDetailsParams) (*City, error)
	// Детали города по GeoID
	CityDetailsByGeoID(context.Context, *CityDetailsByGeoIDParams) (*City, error)
	// Метод возвращает список городов по-умолчанию, в зависимости от количества заказов
	DefaultCityList(context.Context, *DefaultCityListParams) (*SuggestCityResponse, error)
	// Определение города по IP (MaxMind & Dadata + cache)
	CityDetailsByIP(context.Context, *CityDetailsByIPParams) (*City, error)
	// Suggest по адресам
	SuggestAddress(context.Context, *SuggestAddressParams) (*SuggestAddressResponse, error)
	// Нормализация адреса
	AddressDetails(context.Context, *AddressDetailsParams) (*Address, error)
	// Детали адреса по GeoID
	AddressDetailsByGeoID(context.Context, *AddressDetailsByGeoIDParams) (*Address, error)
}

// UnimplementedGeoServer should be embedded to have forward compatible implementations.
type UnimplementedGeoServer struct {
}

func (UnimplementedGeoServer) SuggestCountry(context.Context, *SuggestCountryParams) (*SuggestCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestCountry not implemented")
}
func (UnimplementedGeoServer) CountryDetails(context.Context, *CountryDetailsParams) (*Country, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountryDetails not implemented")
}
func (UnimplementedGeoServer) SuggestCity(context.Context, *SuggestCityParams) (*SuggestCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestCity not implemented")
}
func (UnimplementedGeoServer) CityDetails(context.Context, *CityDetailsParams) (*City, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CityDetails not implemented")
}
func (UnimplementedGeoServer) CityDetailsByGeoID(context.Context, *CityDetailsByGeoIDParams) (*City, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CityDetailsByGeoID not implemented")
}
func (UnimplementedGeoServer) DefaultCityList(context.Context, *DefaultCityListParams) (*SuggestCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefaultCityList not implemented")
}
func (UnimplementedGeoServer) CityDetailsByIP(context.Context, *CityDetailsByIPParams) (*City, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CityDetailsByIP not implemented")
}
func (UnimplementedGeoServer) SuggestAddress(context.Context, *SuggestAddressParams) (*SuggestAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestAddress not implemented")
}
func (UnimplementedGeoServer) AddressDetails(context.Context, *AddressDetailsParams) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressDetails not implemented")
}
func (UnimplementedGeoServer) AddressDetailsByGeoID(context.Context, *AddressDetailsByGeoIDParams) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressDetailsByGeoID not implemented")
}

// UnsafeGeoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeoServer will
// result in compilation errors.
type UnsafeGeoServer interface {
	mustEmbedUnimplementedGeoServer()
}

func RegisterGeoServer(s grpc.ServiceRegistrar, srv GeoServer) {
	s.RegisterService(&Geo_ServiceDesc, srv)
}

func _Geo_SuggestCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestCountryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServer).SuggestCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.geo/SuggestCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServer).SuggestCountry(ctx, req.(*SuggestCountryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geo_CountryDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountryDetailsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServer).CountryDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.geo/CountryDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServer).CountryDetails(ctx, req.(*CountryDetailsParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geo_SuggestCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestCityParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServer).SuggestCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.geo/SuggestCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServer).SuggestCity(ctx, req.(*SuggestCityParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geo_CityDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityDetailsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServer).CityDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.geo/CityDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServer).CityDetails(ctx, req.(*CityDetailsParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geo_CityDetailsByGeoID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityDetailsByGeoIDParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServer).CityDetailsByGeoID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.geo/CityDetailsByGeoID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServer).CityDetailsByGeoID(ctx, req.(*CityDetailsByGeoIDParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geo_DefaultCityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultCityListParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServer).DefaultCityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.geo/DefaultCityList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServer).DefaultCityList(ctx, req.(*DefaultCityListParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geo_CityDetailsByIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityDetailsByIPParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServer).CityDetailsByIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.geo/CityDetailsByIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServer).CityDetailsByIP(ctx, req.(*CityDetailsByIPParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geo_SuggestAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestAddressParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServer).SuggestAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.geo/SuggestAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServer).SuggestAddress(ctx, req.(*SuggestAddressParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geo_AddressDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressDetailsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServer).AddressDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.geo/AddressDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServer).AddressDetails(ctx, req.(*AddressDetailsParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geo_AddressDetailsByGeoID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressDetailsByGeoIDParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServer).AddressDetailsByGeoID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.geo/AddressDetailsByGeoID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServer).AddressDetailsByGeoID(ctx, req.(*AddressDetailsByGeoIDParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Geo_ServiceDesc is the grpc.ServiceDesc for Geo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Geo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "geo.geo",
	HandlerType: (*GeoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuggestCountry",
			Handler:    _Geo_SuggestCountry_Handler,
		},
		{
			MethodName: "CountryDetails",
			Handler:    _Geo_CountryDetails_Handler,
		},
		{
			MethodName: "SuggestCity",
			Handler:    _Geo_SuggestCity_Handler,
		},
		{
			MethodName: "CityDetails",
			Handler:    _Geo_CityDetails_Handler,
		},
		{
			MethodName: "CityDetailsByGeoID",
			Handler:    _Geo_CityDetailsByGeoID_Handler,
		},
		{
			MethodName: "DefaultCityList",
			Handler:    _Geo_DefaultCityList_Handler,
		},
		{
			MethodName: "CityDetailsByIP",
			Handler:    _Geo_CityDetailsByIP_Handler,
		},
		{
			MethodName: "SuggestAddress",
			Handler:    _Geo_SuggestAddress_Handler,
		},
		{
			MethodName: "AddressDetails",
			Handler:    _Geo_AddressDetails_Handler,
		},
		{
			MethodName: "AddressDetailsByGeoID",
			Handler:    _Geo_AddressDetailsByGeoID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/geo.proto",
}
