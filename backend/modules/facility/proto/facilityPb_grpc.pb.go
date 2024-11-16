// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.2
// source: modules/facility/proto/facilityPb.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FacilityService_CheckSlotAvailability_FullMethodName  = "/facility.FacilityService/CheckSlotAvailability"
	FacilityService_GetFacilityPrice_FullMethodName       = "/facility.FacilityService/GetFacilityPrice"
	FacilityService_UpdateSlotBookingCount_FullMethodName = "/facility.FacilityService/UpdateSlotBookingCount"
)

// FacilityServiceClient is the client API for FacilityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Facility service definition
type FacilityServiceClient interface {
	CheckSlotAvailability(ctx context.Context, in *CheckSlotRequest, opts ...grpc.CallOption) (*SlotAvailabilityResponse, error)
	GetFacilityPrice(ctx context.Context, in *FacilityPriceRequest, opts ...grpc.CallOption) (*FacilityPriceResponse, error)
	UpdateSlotBookingCount(ctx context.Context, in *UpdateSlotRequest, opts ...grpc.CallOption) (*UpdateSlotResponse, error)
}

type facilityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFacilityServiceClient(cc grpc.ClientConnInterface) FacilityServiceClient {
	return &facilityServiceClient{cc}
}

func (c *facilityServiceClient) CheckSlotAvailability(ctx context.Context, in *CheckSlotRequest, opts ...grpc.CallOption) (*SlotAvailabilityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SlotAvailabilityResponse)
	err := c.cc.Invoke(ctx, FacilityService_CheckSlotAvailability_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *facilityServiceClient) GetFacilityPrice(ctx context.Context, in *FacilityPriceRequest, opts ...grpc.CallOption) (*FacilityPriceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FacilityPriceResponse)
	err := c.cc.Invoke(ctx, FacilityService_GetFacilityPrice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *facilityServiceClient) UpdateSlotBookingCount(ctx context.Context, in *UpdateSlotRequest, opts ...grpc.CallOption) (*UpdateSlotResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSlotResponse)
	err := c.cc.Invoke(ctx, FacilityService_UpdateSlotBookingCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FacilityServiceServer is the server API for FacilityService service.
// All implementations must embed UnimplementedFacilityServiceServer
// for forward compatibility.
//
// Facility service definition
type FacilityServiceServer interface {
	CheckSlotAvailability(context.Context, *CheckSlotRequest) (*SlotAvailabilityResponse, error)
	GetFacilityPrice(context.Context, *FacilityPriceRequest) (*FacilityPriceResponse, error)
	UpdateSlotBookingCount(context.Context, *UpdateSlotRequest) (*UpdateSlotResponse, error)
	mustEmbedUnimplementedFacilityServiceServer()
}

// UnimplementedFacilityServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFacilityServiceServer struct{}

func (UnimplementedFacilityServiceServer) CheckSlotAvailability(context.Context, *CheckSlotRequest) (*SlotAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSlotAvailability not implemented")
}
func (UnimplementedFacilityServiceServer) GetFacilityPrice(context.Context, *FacilityPriceRequest) (*FacilityPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFacilityPrice not implemented")
}
func (UnimplementedFacilityServiceServer) UpdateSlotBookingCount(context.Context, *UpdateSlotRequest) (*UpdateSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSlotBookingCount not implemented")
}
func (UnimplementedFacilityServiceServer) mustEmbedUnimplementedFacilityServiceServer() {}
func (UnimplementedFacilityServiceServer) testEmbeddedByValue()                         {}

// UnsafeFacilityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FacilityServiceServer will
// result in compilation errors.
type UnsafeFacilityServiceServer interface {
	mustEmbedUnimplementedFacilityServiceServer()
}

func RegisterFacilityServiceServer(s grpc.ServiceRegistrar, srv FacilityServiceServer) {
	// If the following call pancis, it indicates UnimplementedFacilityServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FacilityService_ServiceDesc, srv)
}

func _FacilityService_CheckSlotAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacilityServiceServer).CheckSlotAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FacilityService_CheckSlotAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacilityServiceServer).CheckSlotAvailability(ctx, req.(*CheckSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FacilityService_GetFacilityPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FacilityPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacilityServiceServer).GetFacilityPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FacilityService_GetFacilityPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacilityServiceServer).GetFacilityPrice(ctx, req.(*FacilityPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FacilityService_UpdateSlotBookingCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacilityServiceServer).UpdateSlotBookingCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FacilityService_UpdateSlotBookingCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacilityServiceServer).UpdateSlotBookingCount(ctx, req.(*UpdateSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FacilityService_ServiceDesc is the grpc.ServiceDesc for FacilityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FacilityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "facility.FacilityService",
	HandlerType: (*FacilityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckSlotAvailability",
			Handler:    _FacilityService_CheckSlotAvailability_Handler,
		},
		{
			MethodName: "GetFacilityPrice",
			Handler:    _FacilityService_GetFacilityPrice_Handler,
		},
		{
			MethodName: "UpdateSlotBookingCount",
			Handler:    _FacilityService_UpdateSlotBookingCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/facility/proto/facilityPb.proto",
}
