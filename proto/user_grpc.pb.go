// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
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
	UserService_UserSignup_FullMethodName         = "/user.UserService/UserSignup"
	UserService_VerifyingEmail_FullMethodName     = "/user.UserService/VerifyingEmail"
	UserService_Login_FullMethodName              = "/user.UserService/Login"
	UserService_AdminLogin_FullMethodName         = "/user.UserService/AdminLogin"
	UserService_AddCategory_FullMethodName        = "/user.UserService/AddCategory"
	UserService_AddSkill_FullMethodName           = "/user.UserService/AddSkill"
	UserService_UpdateBio_FullMethodName          = "/user.UserService/UpdateBio"
	UserService_FreelacerAddSkill_FullMethodName  = "/user.UserService/FreelacerAddSkill"
	UserService_GetUserProfile_FullMethodName     = "/user.UserService/GetUserProfile"
	UserService_DeleteSkill_FullMethodName        = "/user.UserService/DeleteSkill"
	UserService_GetCategory_FullMethodName        = "/user.UserService/GetCategory"
	UserService_GetSkill_FullMethodName           = "/user.UserService/GetSkill"
	UserService_AdDeleteSkill_FullMethodName      = "/user.UserService/AdDeleteSkill"
	UserService_DeleteCategory_FullMethodName     = "/user.UserService/DeleteCategory"
	UserService_GetCategoryByName_FullMethodName  = "/user.UserService/GetCategoryByName"
	UserService_GetAllUsers_FullMethodName        = "/user.UserService/GetAllUsers"
	UserService_UploadProfilePhoto_FullMethodName = "/user.UserService/UploadProfilePhoto"
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
	FreelacerAddSkill(ctx context.Context, in *FreeAddSkillsReq, opts ...grpc.CallOption) (*CommonRes, error)
	GetUserProfile(ctx context.Context, in *ProfileReq, opts ...grpc.CallOption) (*ProfileRes, error)
	DeleteSkill(ctx context.Context, in *DeleteSkillRes, opts ...grpc.CallOption) (*CommonRes, error)
	GetCategory(ctx context.Context, in *EmtpyReq, opts ...grpc.CallOption) (*GetCategoryRes, error)
	GetSkill(ctx context.Context, in *EmtpyReq, opts ...grpc.CallOption) (*GetSkillsRes, error)
	AdDeleteSkill(ctx context.Context, in *ADeleteSkillReq, opts ...grpc.CallOption) (*EmtpyRes, error)
	DeleteCategory(ctx context.Context, in *DeleteCatReq, opts ...grpc.CallOption) (*EmtpyRes, error)
	GetCategoryByName(ctx context.Context, in *CategoryName, opts ...grpc.CallOption) (*CategoryIdRes, error)
	GetAllUsers(ctx context.Context, in *EmtpyReq, opts ...grpc.CallOption) (*GetAllUserRes, error)
	UploadProfilePhoto(ctx context.Context, in *UpProilePicReq, opts ...grpc.CallOption) (*CommonRes, error)
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

func (c *userServiceClient) DeleteSkill(ctx context.Context, in *DeleteSkillRes, opts ...grpc.CallOption) (*CommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, UserService_DeleteSkill_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCategory(ctx context.Context, in *EmtpyReq, opts ...grpc.CallOption) (*GetCategoryRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoryRes)
	err := c.cc.Invoke(ctx, UserService_GetCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetSkill(ctx context.Context, in *EmtpyReq, opts ...grpc.CallOption) (*GetSkillsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSkillsRes)
	err := c.cc.Invoke(ctx, UserService_GetSkill_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AdDeleteSkill(ctx context.Context, in *ADeleteSkillReq, opts ...grpc.CallOption) (*EmtpyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmtpyRes)
	err := c.cc.Invoke(ctx, UserService_AdDeleteSkill_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteCategory(ctx context.Context, in *DeleteCatReq, opts ...grpc.CallOption) (*EmtpyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmtpyRes)
	err := c.cc.Invoke(ctx, UserService_DeleteCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCategoryByName(ctx context.Context, in *CategoryName, opts ...grpc.CallOption) (*CategoryIdRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryIdRes)
	err := c.cc.Invoke(ctx, UserService_GetCategoryByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *EmtpyReq, opts ...grpc.CallOption) (*GetAllUserRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllUserRes)
	err := c.cc.Invoke(ctx, UserService_GetAllUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UploadProfilePhoto(ctx context.Context, in *UpProilePicReq, opts ...grpc.CallOption) (*CommonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonRes)
	err := c.cc.Invoke(ctx, UserService_UploadProfilePhoto_FullMethodName, in, out, cOpts...)
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
	FreelacerAddSkill(context.Context, *FreeAddSkillsReq) (*CommonRes, error)
	GetUserProfile(context.Context, *ProfileReq) (*ProfileRes, error)
	DeleteSkill(context.Context, *DeleteSkillRes) (*CommonRes, error)
	GetCategory(context.Context, *EmtpyReq) (*GetCategoryRes, error)
	GetSkill(context.Context, *EmtpyReq) (*GetSkillsRes, error)
	AdDeleteSkill(context.Context, *ADeleteSkillReq) (*EmtpyRes, error)
	DeleteCategory(context.Context, *DeleteCatReq) (*EmtpyRes, error)
	GetCategoryByName(context.Context, *CategoryName) (*CategoryIdRes, error)
	GetAllUsers(context.Context, *EmtpyReq) (*GetAllUserRes, error)
	UploadProfilePhoto(context.Context, *UpProilePicReq) (*CommonRes, error)
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
func (UnimplementedUserServiceServer) FreelacerAddSkill(context.Context, *FreeAddSkillsReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreelacerAddSkill not implemented")
}
func (UnimplementedUserServiceServer) GetUserProfile(context.Context, *ProfileReq) (*ProfileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedUserServiceServer) DeleteSkill(context.Context, *DeleteSkillRes) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSkill not implemented")
}
func (UnimplementedUserServiceServer) GetCategory(context.Context, *EmtpyReq) (*GetCategoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedUserServiceServer) GetSkill(context.Context, *EmtpyReq) (*GetSkillsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkill not implemented")
}
func (UnimplementedUserServiceServer) AdDeleteSkill(context.Context, *ADeleteSkillReq) (*EmtpyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdDeleteSkill not implemented")
}
func (UnimplementedUserServiceServer) DeleteCategory(context.Context, *DeleteCatReq) (*EmtpyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedUserServiceServer) GetCategoryByName(context.Context, *CategoryName) (*CategoryIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryByName not implemented")
}
func (UnimplementedUserServiceServer) GetAllUsers(context.Context, *EmtpyReq) (*GetAllUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUserServiceServer) UploadProfilePhoto(context.Context, *UpProilePicReq) (*CommonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadProfilePhoto not implemented")
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

func _UserService_DeleteSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSkillRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteSkill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteSkill(ctx, req.(*DeleteSkillRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmtpyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCategory(ctx, req.(*EmtpyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmtpyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetSkill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetSkill(ctx, req.(*EmtpyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AdDeleteSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ADeleteSkillReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AdDeleteSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AdDeleteSkill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AdDeleteSkill(ctx, req.(*ADeleteSkillReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteCategory(ctx, req.(*DeleteCatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCategoryByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCategoryByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetCategoryByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCategoryByName(ctx, req.(*CategoryName))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmtpyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsers(ctx, req.(*EmtpyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UploadProfilePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpProilePicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UploadProfilePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UploadProfilePhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UploadProfilePhoto(ctx, req.(*UpProilePicReq))
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
			MethodName: "FreelacerAddSkill",
			Handler:    _UserService_FreelacerAddSkill_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _UserService_GetUserProfile_Handler,
		},
		{
			MethodName: "DeleteSkill",
			Handler:    _UserService_DeleteSkill_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _UserService_GetCategory_Handler,
		},
		{
			MethodName: "GetSkill",
			Handler:    _UserService_GetSkill_Handler,
		},
		{
			MethodName: "AdDeleteSkill",
			Handler:    _UserService_AdDeleteSkill_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _UserService_DeleteCategory_Handler,
		},
		{
			MethodName: "GetCategoryByName",
			Handler:    _UserService_GetCategoryByName_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _UserService_GetAllUsers_Handler,
		},
		{
			MethodName: "UploadProfilePhoto",
			Handler:    _UserService_UploadProfilePhoto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
