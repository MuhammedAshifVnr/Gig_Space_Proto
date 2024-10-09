// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: payment.proto

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
	PaymentService_CreateSubscription_FullMethodName  = "/payment.PaymentService/CreateSubscription"
	PaymentService_UpdatePaymentStatus_FullMethodName = "/payment.PaymentService/UpdatePaymentStatus"
	PaymentService_RenewSubscription_FullMethodName   = "/payment.PaymentService/RenewSubscription"
	PaymentService_GetSubDetails_FullMethodName       = "/payment.PaymentService/GetSubDetails"
	PaymentService_CreatePlan_FullMethodName          = "/payment.PaymentService/CreatePlan"
	PaymentService_GetPlan_FullMethodName             = "/payment.PaymentService/GetPlan"
	PaymentService_DeletePlan_FullMethodName          = "/payment.PaymentService/DeletePlan"
	PaymentService_HandleWebhook_FullMethodName       = "/payment.PaymentService/HandleWebhook"
	PaymentService_CreatePaymentOrder_FullMethodName  = "/payment.PaymentService/CreatePaymentOrder"
	PaymentService_CreateWallet_FullMethodName        = "/payment.PaymentService/CreateWallet"
	PaymentService_GetWallet_FullMethodName           = "/payment.PaymentService/GetWallet"
	PaymentService_CreateBankAccount_FullMethodName   = "/payment.PaymentService/CreateBankAccount"
)

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	CreateSubscription(ctx context.Context, in *CreateSubscriptionRequest, opts ...grpc.CallOption) (*CreateSubscriptionResponse, error)
	UpdatePaymentStatus(ctx context.Context, in *UpdatePaymentReq, opts ...grpc.CallOption) (*PaymentCommonRes, error)
	RenewSubscription(ctx context.Context, in *RenewSubReq, opts ...grpc.CallOption) (*PaymentCommonRes, error)
	GetSubDetails(ctx context.Context, in *GetSubReq, opts ...grpc.CallOption) (*GetSubRes, error)
	CreatePlan(ctx context.Context, in *CreatePlanReq, opts ...grpc.CallOption) (*PaymentCommonRes, error)
	GetPlan(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*GetPlanRes, error)
	DeletePlan(ctx context.Context, in *DeletePlanReq, opts ...grpc.CallOption) (*PaymentCommonRes, error)
	HandleWebhook(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookResponse, error)
	CreatePaymentOrder(ctx context.Context, in *CreatePaymentOrderReq, opts ...grpc.CallOption) (*PaymentCommonRes, error)
	CreateWallet(ctx context.Context, in *CreateWalletReq, opts ...grpc.CallOption) (*PaymentCommonRes, error)
	GetWallet(ctx context.Context, in *GetwalletReq, opts ...grpc.CallOption) (*WalletRes, error)
	CreateBankAccount(ctx context.Context, in *CreaBankReq, opts ...grpc.CallOption) (*PaymentCommonRes, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) CreateSubscription(ctx context.Context, in *CreateSubscriptionRequest, opts ...grpc.CallOption) (*CreateSubscriptionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSubscriptionResponse)
	err := c.cc.Invoke(ctx, PaymentService_CreateSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) UpdatePaymentStatus(ctx context.Context, in *UpdatePaymentReq, opts ...grpc.CallOption) (*PaymentCommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymentCommonRes)
	err := c.cc.Invoke(ctx, PaymentService_UpdatePaymentStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) RenewSubscription(ctx context.Context, in *RenewSubReq, opts ...grpc.CallOption) (*PaymentCommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymentCommonRes)
	err := c.cc.Invoke(ctx, PaymentService_RenewSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetSubDetails(ctx context.Context, in *GetSubReq, opts ...grpc.CallOption) (*GetSubRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSubRes)
	err := c.cc.Invoke(ctx, PaymentService_GetSubDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CreatePlan(ctx context.Context, in *CreatePlanReq, opts ...grpc.CallOption) (*PaymentCommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymentCommonRes)
	err := c.cc.Invoke(ctx, PaymentService_CreatePlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPlan(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*GetPlanRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPlanRes)
	err := c.cc.Invoke(ctx, PaymentService_GetPlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) DeletePlan(ctx context.Context, in *DeletePlanReq, opts ...grpc.CallOption) (*PaymentCommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymentCommonRes)
	err := c.cc.Invoke(ctx, PaymentService_DeletePlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) HandleWebhook(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookResponse)
	err := c.cc.Invoke(ctx, PaymentService_HandleWebhook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CreatePaymentOrder(ctx context.Context, in *CreatePaymentOrderReq, opts ...grpc.CallOption) (*PaymentCommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymentCommonRes)
	err := c.cc.Invoke(ctx, PaymentService_CreatePaymentOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CreateWallet(ctx context.Context, in *CreateWalletReq, opts ...grpc.CallOption) (*PaymentCommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymentCommonRes)
	err := c.cc.Invoke(ctx, PaymentService_CreateWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetWallet(ctx context.Context, in *GetwalletReq, opts ...grpc.CallOption) (*WalletRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WalletRes)
	err := c.cc.Invoke(ctx, PaymentService_GetWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CreateBankAccount(ctx context.Context, in *CreaBankReq, opts ...grpc.CallOption) (*PaymentCommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymentCommonRes)
	err := c.cc.Invoke(ctx, PaymentService_CreateBankAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility.
type PaymentServiceServer interface {
	CreateSubscription(context.Context, *CreateSubscriptionRequest) (*CreateSubscriptionResponse, error)
	UpdatePaymentStatus(context.Context, *UpdatePaymentReq) (*PaymentCommonRes, error)
	RenewSubscription(context.Context, *RenewSubReq) (*PaymentCommonRes, error)
	GetSubDetails(context.Context, *GetSubReq) (*GetSubRes, error)
	CreatePlan(context.Context, *CreatePlanReq) (*PaymentCommonRes, error)
	GetPlan(context.Context, *EmptyReq) (*GetPlanRes, error)
	DeletePlan(context.Context, *DeletePlanReq) (*PaymentCommonRes, error)
	HandleWebhook(context.Context, *WebhookRequest) (*WebhookResponse, error)
	CreatePaymentOrder(context.Context, *CreatePaymentOrderReq) (*PaymentCommonRes, error)
	CreateWallet(context.Context, *CreateWalletReq) (*PaymentCommonRes, error)
	GetWallet(context.Context, *GetwalletReq) (*WalletRes, error)
	CreateBankAccount(context.Context, *CreaBankReq) (*PaymentCommonRes, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentServiceServer struct{}

func (UnimplementedPaymentServiceServer) CreateSubscription(context.Context, *CreateSubscriptionRequest) (*CreateSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscription not implemented")
}
func (UnimplementedPaymentServiceServer) UpdatePaymentStatus(context.Context, *UpdatePaymentReq) (*PaymentCommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePaymentStatus not implemented")
}
func (UnimplementedPaymentServiceServer) RenewSubscription(context.Context, *RenewSubReq) (*PaymentCommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewSubscription not implemented")
}
func (UnimplementedPaymentServiceServer) GetSubDetails(context.Context, *GetSubReq) (*GetSubRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubDetails not implemented")
}
func (UnimplementedPaymentServiceServer) CreatePlan(context.Context, *CreatePlanReq) (*PaymentCommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlan not implemented")
}
func (UnimplementedPaymentServiceServer) GetPlan(context.Context, *EmptyReq) (*GetPlanRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlan not implemented")
}
func (UnimplementedPaymentServiceServer) DeletePlan(context.Context, *DeletePlanReq) (*PaymentCommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlan not implemented")
}
func (UnimplementedPaymentServiceServer) HandleWebhook(context.Context, *WebhookRequest) (*WebhookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleWebhook not implemented")
}
func (UnimplementedPaymentServiceServer) CreatePaymentOrder(context.Context, *CreatePaymentOrderReq) (*PaymentCommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePaymentOrder not implemented")
}
func (UnimplementedPaymentServiceServer) CreateWallet(context.Context, *CreateWalletReq) (*PaymentCommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedPaymentServiceServer) GetWallet(context.Context, *GetwalletReq) (*WalletRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWallet not implemented")
}
func (UnimplementedPaymentServiceServer) CreateBankAccount(context.Context, *CreaBankReq) (*PaymentCommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBankAccount not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}
func (UnimplementedPaymentServiceServer) testEmbeddedByValue()                        {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	// If the following call pancis, it indicates UnimplementedPaymentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_CreateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_CreateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreateSubscription(ctx, req.(*CreateSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_UpdatePaymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).UpdatePaymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_UpdatePaymentStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).UpdatePaymentStatus(ctx, req.(*UpdatePaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_RenewSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewSubReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).RenewSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_RenewSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).RenewSubscription(ctx, req.(*RenewSubReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetSubDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetSubDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetSubDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetSubDetails(ctx, req.(*GetSubReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_CreatePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreatePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_CreatePlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreatePlan(ctx, req.(*CreatePlanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPlan(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_DeletePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).DeletePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_DeletePlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).DeletePlan(ctx, req.(*DeletePlanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_HandleWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).HandleWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_HandleWebhook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).HandleWebhook(ctx, req.(*WebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_CreatePaymentOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreatePaymentOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_CreatePaymentOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreatePaymentOrder(ctx, req.(*CreatePaymentOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_CreateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreateWallet(ctx, req.(*CreateWalletReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetwalletReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetWallet(ctx, req.(*GetwalletReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_CreateBankAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreaBankReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreateBankAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_CreateBankAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreateBankAccount(ctx, req.(*CreaBankReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubscription",
			Handler:    _PaymentService_CreateSubscription_Handler,
		},
		{
			MethodName: "UpdatePaymentStatus",
			Handler:    _PaymentService_UpdatePaymentStatus_Handler,
		},
		{
			MethodName: "RenewSubscription",
			Handler:    _PaymentService_RenewSubscription_Handler,
		},
		{
			MethodName: "GetSubDetails",
			Handler:    _PaymentService_GetSubDetails_Handler,
		},
		{
			MethodName: "CreatePlan",
			Handler:    _PaymentService_CreatePlan_Handler,
		},
		{
			MethodName: "GetPlan",
			Handler:    _PaymentService_GetPlan_Handler,
		},
		{
			MethodName: "DeletePlan",
			Handler:    _PaymentService_DeletePlan_Handler,
		},
		{
			MethodName: "HandleWebhook",
			Handler:    _PaymentService_HandleWebhook_Handler,
		},
		{
			MethodName: "CreatePaymentOrder",
			Handler:    _PaymentService_CreatePaymentOrder_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _PaymentService_CreateWallet_Handler,
		},
		{
			MethodName: "GetWallet",
			Handler:    _PaymentService_GetWallet_Handler,
		},
		{
			MethodName: "CreateBankAccount",
			Handler:    _PaymentService_CreateBankAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment.proto",
}
