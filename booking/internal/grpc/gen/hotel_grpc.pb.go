// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: hotel.proto

package gen

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
	Hotel_GetHotel_FullMethodName = "/gen.Hotel/GetHotel"
)

// HotelClient is the client API for Hotel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HotelClient interface {
	GetHotel(ctx context.Context, in *HotelRequest, opts ...grpc.CallOption) (*HotelResponse, error)
}

type hotelClient struct {
	cc grpc.ClientConnInterface
}

func NewHotelClient(cc grpc.ClientConnInterface) HotelClient {
	return &hotelClient{cc}
}

func (c *hotelClient) GetHotel(ctx context.Context, in *HotelRequest, opts ...grpc.CallOption) (*HotelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HotelResponse)
	err := c.cc.Invoke(ctx, Hotel_GetHotel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelServer is the server API for Hotel service.
// All implementations must embed UnimplementedHotelServer
// for forward compatibility.
type HotelServer interface {
	GetHotel(context.Context, *HotelRequest) (*HotelResponse, error)
	mustEmbedUnimplementedHotelServer()
}

// UnimplementedHotelServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHotelServer struct{}

func (UnimplementedHotelServer) GetHotel(context.Context, *HotelRequest) (*HotelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotel not implemented")
}
func (UnimplementedHotelServer) mustEmbedUnimplementedHotelServer() {}
func (UnimplementedHotelServer) testEmbeddedByValue()               {}

// UnsafeHotelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HotelServer will
// result in compilation errors.
type UnsafeHotelServer interface {
	mustEmbedUnimplementedHotelServer()
}

func RegisterHotelServer(s grpc.ServiceRegistrar, srv HotelServer) {
	// If the following call pancis, it indicates UnimplementedHotelServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Hotel_ServiceDesc, srv)
}

func _Hotel_GetHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HotelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServer).GetHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hotel_GetHotel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServer).GetHotel(ctx, req.(*HotelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hotel_ServiceDesc is the grpc.ServiceDesc for Hotel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hotel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gen.Hotel",
	HandlerType: (*HotelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHotel",
			Handler:    _Hotel_GetHotel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hotel.proto",
}