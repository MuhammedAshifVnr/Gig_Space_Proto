// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: gig.proto

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
	GigService_CreateGig_FullMethodName             = "/gig.GigService/CreateGig"
	GigService_GetGigsByFreelancerID_FullMethodName = "/gig.GigService/GetGigsByFreelancerID"
	GigService_UpdateGigByID_FullMethodName         = "/gig.GigService/UpdateGigByID"
	GigService_DeleteGigByID_FullMethodName         = "/gig.GigService/DeleteGigByID"
	GigService_CreateOrder_FullMethodName           = "/gig.GigService/CreateOrder"
	GigService_GetOrders_FullMethodName             = "/gig.GigService/GetOrders"
	GigService_RequestQuote_FullMethodName          = "/gig.GigService/RequestQuote"
	GigService_GetAllQuotes_FullMethodName          = "/gig.GigService/GetAllQuotes"
)

// GigServiceClient is the client API for GigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GigServiceClient interface {
	CreateGig(ctx context.Context, in *CreateGigReq, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetGigsByFreelancerID(ctx context.Context, in *GetGigsByFreelancerIDRequest, opts ...grpc.CallOption) (*GetGigsByFreelancerIDResponse, error)
	UpdateGigByID(ctx context.Context, in *UpdateGigRequest, opts ...grpc.CallOption) (*CommonGigRes, error)
	DeleteGigByID(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*CommonGigRes, error)
	CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CommonGigRes, error)
	GetOrders(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderRes, error)
	RequestQuote(ctx context.Context, in *QuoteReq, opts ...grpc.CallOption) (*CommonGigRes, error)
	GetAllQuotes(ctx context.Context, in *GetAllQuoteReq, opts ...grpc.CallOption) (*GetAllQuoteRes, error)
}

type gigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGigServiceClient(cc grpc.ClientConnInterface) GigServiceClient {
	return &gigServiceClient{cc}
}

func (c *gigServiceClient) CreateGig(ctx context.Context, in *CreateGigReq, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, GigService_CreateGig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gigServiceClient) GetGigsByFreelancerID(ctx context.Context, in *GetGigsByFreelancerIDRequest, opts ...grpc.CallOption) (*GetGigsByFreelancerIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGigsByFreelancerIDResponse)
	err := c.cc.Invoke(ctx, GigService_GetGigsByFreelancerID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gigServiceClient) UpdateGigByID(ctx context.Context, in *UpdateGigRequest, opts ...grpc.CallOption) (*CommonGigRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonGigRes)
	err := c.cc.Invoke(ctx, GigService_UpdateGigByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gigServiceClient) DeleteGigByID(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*CommonGigRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonGigRes)
	err := c.cc.Invoke(ctx, GigService_DeleteGigByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gigServiceClient) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CommonGigRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonGigRes)
	err := c.cc.Invoke(ctx, GigService_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gigServiceClient) GetOrders(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderRes)
	err := c.cc.Invoke(ctx, GigService_GetOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gigServiceClient) RequestQuote(ctx context.Context, in *QuoteReq, opts ...grpc.CallOption) (*CommonGigRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonGigRes)
	err := c.cc.Invoke(ctx, GigService_RequestQuote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gigServiceClient) GetAllQuotes(ctx context.Context, in *GetAllQuoteReq, opts ...grpc.CallOption) (*GetAllQuoteRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllQuoteRes)
	err := c.cc.Invoke(ctx, GigService_GetAllQuotes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GigServiceServer is the server API for GigService service.
// All implementations must embed UnimplementedGigServiceServer
// for forward compatibility.
type GigServiceServer interface {
	CreateGig(context.Context, *CreateGigReq) (*EmptyResponse, error)
	GetGigsByFreelancerID(context.Context, *GetGigsByFreelancerIDRequest) (*GetGigsByFreelancerIDResponse, error)
	UpdateGigByID(context.Context, *UpdateGigRequest) (*CommonGigRes, error)
	DeleteGigByID(context.Context, *DeleteReq) (*CommonGigRes, error)
	CreateOrder(context.Context, *CreateOrderReq) (*CommonGigRes, error)
	GetOrders(context.Context, *GetOrderReq) (*GetOrderRes, error)
	RequestQuote(context.Context, *QuoteReq) (*CommonGigRes, error)
	GetAllQuotes(context.Context, *GetAllQuoteReq) (*GetAllQuoteRes, error)
	mustEmbedUnimplementedGigServiceServer()
}

// UnimplementedGigServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGigServiceServer struct{}

func (UnimplementedGigServiceServer) CreateGig(context.Context, *CreateGigReq) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGig not implemented")
}
func (UnimplementedGigServiceServer) GetGigsByFreelancerID(context.Context, *GetGigsByFreelancerIDRequest) (*GetGigsByFreelancerIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGigsByFreelancerID not implemented")
}
func (UnimplementedGigServiceServer) UpdateGigByID(context.Context, *UpdateGigRequest) (*CommonGigRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGigByID not implemented")
}
func (UnimplementedGigServiceServer) DeleteGigByID(context.Context, *DeleteReq) (*CommonGigRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGigByID not implemented")
}
func (UnimplementedGigServiceServer) CreateOrder(context.Context, *CreateOrderReq) (*CommonGigRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedGigServiceServer) GetOrders(context.Context, *GetOrderReq) (*GetOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedGigServiceServer) RequestQuote(context.Context, *QuoteReq) (*CommonGigRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestQuote not implemented")
}
func (UnimplementedGigServiceServer) GetAllQuotes(context.Context, *GetAllQuoteReq) (*GetAllQuoteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllQuotes not implemented")
}
func (UnimplementedGigServiceServer) mustEmbedUnimplementedGigServiceServer() {}
func (UnimplementedGigServiceServer) testEmbeddedByValue()                    {}

// UnsafeGigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GigServiceServer will
// result in compilation errors.
type UnsafeGigServiceServer interface {
	mustEmbedUnimplementedGigServiceServer()
}

func RegisterGigServiceServer(s grpc.ServiceRegistrar, srv GigServiceServer) {
	// If the following call pancis, it indicates UnimplementedGigServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GigService_ServiceDesc, srv)
}

func _GigService_CreateGig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GigServiceServer).CreateGig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GigService_CreateGig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GigServiceServer).CreateGig(ctx, req.(*CreateGigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GigService_GetGigsByFreelancerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGigsByFreelancerIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GigServiceServer).GetGigsByFreelancerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GigService_GetGigsByFreelancerID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GigServiceServer).GetGigsByFreelancerID(ctx, req.(*GetGigsByFreelancerIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GigService_UpdateGigByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GigServiceServer).UpdateGigByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GigService_UpdateGigByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GigServiceServer).UpdateGigByID(ctx, req.(*UpdateGigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GigService_DeleteGigByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GigServiceServer).DeleteGigByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GigService_DeleteGigByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GigServiceServer).DeleteGigByID(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GigService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GigServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GigService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GigServiceServer).CreateOrder(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GigService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GigServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GigService_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GigServiceServer).GetOrders(ctx, req.(*GetOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GigService_RequestQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuoteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GigServiceServer).RequestQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GigService_RequestQuote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GigServiceServer).RequestQuote(ctx, req.(*QuoteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GigService_GetAllQuotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllQuoteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GigServiceServer).GetAllQuotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GigService_GetAllQuotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GigServiceServer).GetAllQuotes(ctx, req.(*GetAllQuoteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GigService_ServiceDesc is the grpc.ServiceDesc for GigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gig.GigService",
	HandlerType: (*GigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGig",
			Handler:    _GigService_CreateGig_Handler,
		},
		{
			MethodName: "GetGigsByFreelancerID",
			Handler:    _GigService_GetGigsByFreelancerID_Handler,
		},
		{
			MethodName: "UpdateGigByID",
			Handler:    _GigService_UpdateGigByID_Handler,
		},
		{
			MethodName: "DeleteGigByID",
			Handler:    _GigService_DeleteGigByID_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _GigService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _GigService_GetOrders_Handler,
		},
		{
			MethodName: "RequestQuote",
			Handler:    _GigService_RequestQuote_Handler,
		},
		{
			MethodName: "GetAllQuotes",
			Handler:    _GigService_GetAllQuotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gig.proto",
}
