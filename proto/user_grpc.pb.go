// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: user.proto

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
	UserService_UserSignup_FullMethodName        = "/user.UserService/UserSignup"
	UserService_VerifyingEmail_FullMethodName    = "/user.UserService/VerifyingEmail"
	UserService_Login_FullMethodName             = "/user.UserService/Login"
	UserService_AdminLogin_FullMethodName        = "/user.UserService/AdminLogin"
	UserService_AddCategory_FullMethodName       = "/user.UserService/AddCategory"
	UserService_AddSkill_FullMethodName          = "/user.UserService/AddSkill"
	UserService_UpdateBio_FullMethodName         = "/user.UserService/UpdateBio"
	UserService_UpdateProfilePic_FullMethodName  = "/user.UserService/UpdateProfilePic"
	UserService_FreelacerAddSkill_FullMethodName = "/user.UserService/FreelacerAddSkill"
	UserService_GetUserProfile_FullMethodName    = "/user.UserService/GetUserProfile"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserSignup(ctx context.Context, in *SignupReq, opts ...grpc.CallOption) (*SignupRes, error)
	VerifyingEmail(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRes, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	AdminLogin(ctx context.Context, in *AdLoginReq, opts ...grpc.CallOption) (*CommonRes, error)
	AddCategory(ctx context.Context, in *CategoryReq, opts ...grpc.CallOption) (*CommonRes, error)
	AddSkill(ctx context.Context, in *AddSkillReq, opts ...grpc.CallOption) (*CommonRes, error)
	UpdateBio(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*CommonRes, error)
	UpdateProfilePic(ctx context.Context, in *UpdatePP_Req, opts ...grpc.CallOption) (*CommonRes, error)
	FreelacerAddSkill(ctx context.Context, in *FreeAddSkillsReq, opts ...grpc.CallOption) (*CommonRes, error)
	GetUserProfile(ctx context.Context, in *ProfileReq, opts ...grpc.CallOption) (*ProfileRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserSignup(ctx context.Context, in *SignupReq, opts ...grpc.CallOption) (*SignupRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignupRes)
	err := c.cc.Invoke(ctx, UserService_UserSignup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) VerifyingEmail(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyRes)
	err := c.cc.Invoke(ctx, UserService_VerifyingEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AdminLogin(ctx context.Context, in *AdLoginReq, opts ...grpc.CallOption) (*CommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, UserService_AdminLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddCategory(ctx context.Context, in *CategoryReq, opts ...grpc.CallOption) (*CommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, UserService_AddCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddSkill(ctx context.Context, in *AddSkillReq, opts ...grpc.CallOption) (*CommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, UserService_AddSkill_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateBio(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*CommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, UserService_UpdateBio_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateProfilePic(ctx context.Context, in *UpdatePP_Req, opts ...grpc.CallOption) (*CommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, UserService_UpdateProfilePic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FreelacerAddSkill(ctx context.Context, in *FreeAddSkillsReq, opts ...grpc.CallOption) (*CommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, UserService_FreelacerAddSkill_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserProfile(ctx context.Context, in *ProfileReq, opts ...grpc.CallOption) (*ProfileRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfileRes)
	err := c.cc.Invoke(ctx, UserService_GetUserProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	UserSignup(context.Context, *SignupReq) (*SignupRes, error)
	VerifyingEmail(context.Context, *VerifyReq) (*VerifyRes, error)
	Login(context.Context, *LoginReq) (*LoginRes, error)
	AdminLogin(context.Context, *AdLoginReq) (*CommonRes, error)
	AddCategory(context.Context, *CategoryReq) (*CommonRes, error)
	AddSkill(context.Context, *AddSkillReq) (*CommonRes, error)
	UpdateBio(context.Context, *UpdateProfileReq) (*CommonRes, error)
	UpdateProfilePic(context.Context, *UpdatePP_Req) (*CommonRes, error)
	FreelacerAddSkill(context.Context, *FreeAddSkillsReq) (*CommonRes, error)
	GetUserProfile(context.Context, *ProfileReq) (*ProfileRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) UserSignup(context.Context, *SignupReq) (*SignupRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignup not implemented")
}
func (UnimplementedUserServiceServer) VerifyingEmail(context.Context, *VerifyReq) (*VerifyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyingEmail not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginReq) (*LoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) AdminLogin(context.Context, *AdLoginReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedUserServiceServer) AddCategory(context.Context, *CategoryReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedUserServiceServer) AddSkill(context.Context, *AddSkillReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSkill not implemented")
}
func (UnimplementedUserServiceServer) UpdateBio(context.Context, *UpdateProfileReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBio not implemented")
}
func (UnimplementedUserServiceServer) UpdateProfilePic(context.Context, *UpdatePP_Req) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfilePic not implemented")
}
func (UnimplementedUserServiceServer) FreelacerAddSkill(context.Context, *FreeAddSkillsReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreelacerAddSkill not implemented")
}
func (UnimplementedUserServiceServer) GetUserProfile(context.Context, *ProfileReq) (*ProfileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserSignup(ctx, req.(*SignupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_VerifyingEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerifyingEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_VerifyingEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerifyingEmail(ctx, req.(*VerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AdminLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AdminLogin(ctx, req.(*AdLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddCategory(ctx, req.(*CategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSkillReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddSkill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddSkill(ctx, req.(*AddSkillReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateBio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateBio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateBio_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateBio(ctx, req.(*UpdateProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateProfilePic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePP_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateProfilePic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateProfilePic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateProfilePic(ctx, req.(*UpdatePP_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FreelacerAddSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreeAddSkillsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FreelacerAddSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FreelacerAddSkill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FreelacerAddSkill(ctx, req.(*FreeAddSkillsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserProfile(ctx, req.(*ProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserSignup",
			Handler:    _UserService_UserSignup_Handler,
		},
		{
			MethodName: "VerifyingEmail",
			Handler:    _UserService_VerifyingEmail_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "AdminLogin",
			Handler:    _UserService_AdminLogin_Handler,
		},
		{
			MethodName: "AddCategory",
			Handler:    _UserService_AddCategory_Handler,
		},
		{
			MethodName: "AddSkill",
			Handler:    _UserService_AddSkill_Handler,
		},
		{
			MethodName: "UpdateBio",
			Handler:    _UserService_UpdateBio_Handler,
		},
		{
			MethodName: "UpdateProfilePic",
			Handler:    _UserService_UpdateProfilePic_Handler,
		},
		{
			MethodName: "FreelacerAddSkill",
			Handler:    _UserService_FreelacerAddSkill_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _UserService_GetUserProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
