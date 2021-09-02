// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// FPTBookingClient is the client API for FPTBooking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FPTBookingClient interface {
	CustomerBooking(ctx context.Context, in *BookingRequest, opts ...grpc.CallOption) (*Booking, error)
	ViewBooking(ctx context.Context, in *ViewBookingRequest, opts ...grpc.CallOption) (*ViewBookingResponse, error)
	CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*Empty, error)
	ViewBookingByID(ctx context.Context, in *ViewBookingByIDRequest, opts ...grpc.CallOption) (*ViewBookingByIDResponse, error)
}

type fPTBookingClient struct {
	cc grpc.ClientConnInterface
}

func NewFPTBookingClient(cc grpc.ClientConnInterface) FPTBookingClient {
	return &fPTBookingClient{cc}
}

func (c *fPTBookingClient) CustomerBooking(ctx context.Context, in *BookingRequest, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/training.FPTBooking/CustomerBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fPTBookingClient) ViewBooking(ctx context.Context, in *ViewBookingRequest, opts ...grpc.CallOption) (*ViewBookingResponse, error) {
	out := new(ViewBookingResponse)
	err := c.cc.Invoke(ctx, "/training.FPTBooking/ViewBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fPTBookingClient) CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/training.FPTBooking/CancelBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fPTBookingClient) ViewBookingByID(ctx context.Context, in *ViewBookingByIDRequest, opts ...grpc.CallOption) (*ViewBookingByIDResponse, error) {
	out := new(ViewBookingByIDResponse)
	err := c.cc.Invoke(ctx, "/training.FPTBooking/ViewBookingByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FPTBookingServer is the server API for FPTBooking service.
// All implementations must embed UnimplementedFPTBookingServer
// for forward compatibility
type FPTBookingServer interface {
	CustomerBooking(context.Context, *BookingRequest) (*Booking, error)
	ViewBooking(context.Context, *ViewBookingRequest) (*ViewBookingResponse, error)
	CancelBooking(context.Context, *CancelBookingRequest) (*Empty, error)
	ViewBookingByID(context.Context, *ViewBookingByIDRequest) (*ViewBookingByIDResponse, error)
	mustEmbedUnimplementedFPTBookingServer()
}

// UnimplementedFPTBookingServer must be embedded to have forward compatible implementations.
type UnimplementedFPTBookingServer struct {
}

func (UnimplementedFPTBookingServer) CustomerBooking(context.Context, *BookingRequest) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerBooking not implemented")
}
func (UnimplementedFPTBookingServer) ViewBooking(context.Context, *ViewBookingRequest) (*ViewBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewBooking not implemented")
}
func (UnimplementedFPTBookingServer) CancelBooking(context.Context, *CancelBookingRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}
func (UnimplementedFPTBookingServer) ViewBookingByID(context.Context, *ViewBookingByIDRequest) (*ViewBookingByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewBookingByID not implemented")
}
func (UnimplementedFPTBookingServer) mustEmbedUnimplementedFPTBookingServer() {}

// UnsafeFPTBookingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FPTBookingServer will
// result in compilation errors.
type UnsafeFPTBookingServer interface {
	mustEmbedUnimplementedFPTBookingServer()
}

func RegisterFPTBookingServer(s grpc.ServiceRegistrar, srv FPTBookingServer) {
	s.RegisterService(&FPTBooking_ServiceDesc, srv)
}

func _FPTBooking_CustomerBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTBookingServer).CustomerBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTBooking/CustomerBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTBookingServer).CustomerBooking(ctx, req.(*BookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FPTBooking_ViewBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTBookingServer).ViewBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTBooking/ViewBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTBookingServer).ViewBooking(ctx, req.(*ViewBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FPTBooking_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTBookingServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTBooking/CancelBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTBookingServer).CancelBooking(ctx, req.(*CancelBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FPTBooking_ViewBookingByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewBookingByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FPTBookingServer).ViewBookingByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/training.FPTBooking/ViewBookingByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FPTBookingServer).ViewBookingByID(ctx, req.(*ViewBookingByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FPTBooking_ServiceDesc is the grpc.ServiceDesc for FPTBooking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FPTBooking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "training.FPTBooking",
	HandlerType: (*FPTBookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CustomerBooking",
			Handler:    _FPTBooking_CustomerBooking_Handler,
		},
		{
			MethodName: "ViewBooking",
			Handler:    _FPTBooking_ViewBooking_Handler,
		},
		{
			MethodName: "CancelBooking",
			Handler:    _FPTBooking_CancelBooking_Handler,
		},
		{
			MethodName: "ViewBookingByID",
			Handler:    _FPTBooking_ViewBookingByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}