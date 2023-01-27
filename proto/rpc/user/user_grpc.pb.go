// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/rpc/user/user.proto

package userv3

import (
	context "context"
	v31 "github.com/paralus/paralus/proto/types/commonpb/v3"
	v3 "github.com/paralus/paralus/proto/types/userpb/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	AuditLogWebhook(ctx context.Context, in *UserLoginAuditRequest, opts ...grpc.CallOption) (*UserLoginAuditResponse, error)
	CreateUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.User, error)
	GetUsers(ctx context.Context, in *v31.QueryOptions, opts ...grpc.CallOption) (*v3.UserList, error)
	GetUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.User, error)
	GetUserInfo(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.UserInfo, error)
	UpdateUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.User, error)
	UpdateUserForceReset(ctx context.Context, in *UpdateForceResetRequest, opts ...grpc.CallOption) (*UpdateForceResetResponse, error)
	DeleteUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*UserDeleteApiKeysResponse, error)
	DownloadCliConfig(ctx context.Context, in *CliConfigRequest, opts ...grpc.CallOption) (*v31.HttpBody, error)
	UserListApiKeys(ctx context.Context, in *ApiKeyRequest, opts ...grpc.CallOption) (*UserListApiKeysResponse, error)
	UserDeleteApiKeys(ctx context.Context, in *ApiKeyRequest, opts ...grpc.CallOption) (*UserDeleteApiKeysResponse, error)
	UserForgotPassword(ctx context.Context, in *UserForgotPasswordRequest, opts ...grpc.CallOption) (*UserForgotPasswordResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AuditLogWebhook(ctx context.Context, in *UserLoginAuditRequest, opts ...grpc.CallOption) (*UserLoginAuditResponse, error) {
	out := new(UserLoginAuditResponse)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.UserService/AuditLogWebhook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.User, error) {
	out := new(v3.User)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *v31.QueryOptions, opts ...grpc.CallOption) (*v3.UserList, error) {
	out := new(v3.UserList)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.UserService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.User, error) {
	out := new(v3.User)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.UserInfo, error) {
	out := new(v3.UserInfo)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.UserService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.User, error) {
	out := new(v3.User)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserForceReset(ctx context.Context, in *UpdateForceResetRequest, opts ...grpc.CallOption) (*UpdateForceResetResponse, error) {
	out := new(UpdateForceResetResponse)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.UserService/UpdateUserForceReset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*UserDeleteApiKeysResponse, error) {
	out := new(UserDeleteApiKeysResponse)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DownloadCliConfig(ctx context.Context, in *CliConfigRequest, opts ...grpc.CallOption) (*v31.HttpBody, error) {
	out := new(v31.HttpBody)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.UserService/DownloadCliConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserListApiKeys(ctx context.Context, in *ApiKeyRequest, opts ...grpc.CallOption) (*UserListApiKeysResponse, error) {
	out := new(UserListApiKeysResponse)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.UserService/UserListApiKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserDeleteApiKeys(ctx context.Context, in *ApiKeyRequest, opts ...grpc.CallOption) (*UserDeleteApiKeysResponse, error) {
	out := new(UserDeleteApiKeysResponse)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.UserService/UserDeleteApiKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserForgotPassword(ctx context.Context, in *UserForgotPasswordRequest, opts ...grpc.CallOption) (*UserForgotPasswordResponse, error) {
	out := new(UserForgotPasswordResponse)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.user.v3.UserService/UserForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	AuditLogWebhook(context.Context, *UserLoginAuditRequest) (*UserLoginAuditResponse, error)
	CreateUser(context.Context, *v3.User) (*v3.User, error)
	GetUsers(context.Context, *v31.QueryOptions) (*v3.UserList, error)
	GetUser(context.Context, *v3.User) (*v3.User, error)
	GetUserInfo(context.Context, *v3.User) (*v3.UserInfo, error)
	UpdateUser(context.Context, *v3.User) (*v3.User, error)
	UpdateUserForceReset(context.Context, *UpdateForceResetRequest) (*UpdateForceResetResponse, error)
	DeleteUser(context.Context, *v3.User) (*UserDeleteApiKeysResponse, error)
	DownloadCliConfig(context.Context, *CliConfigRequest) (*v31.HttpBody, error)
	UserListApiKeys(context.Context, *ApiKeyRequest) (*UserListApiKeysResponse, error)
	UserDeleteApiKeys(context.Context, *ApiKeyRequest) (*UserDeleteApiKeysResponse, error)
	UserForgotPassword(context.Context, *UserForgotPasswordRequest) (*UserForgotPasswordResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) AuditLogWebhook(context.Context, *UserLoginAuditRequest) (*UserLoginAuditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuditLogWebhook not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *v3.User) (*v3.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUsers(context.Context, *v31.QueryOptions) (*v3.UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *v3.User) (*v3.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfo(context.Context, *v3.User) (*v3.UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *v3.User) (*v3.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserForceReset(context.Context, *UpdateForceResetRequest) (*UpdateForceResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserForceReset not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *v3.User) (*UserDeleteApiKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) DownloadCliConfig(context.Context, *CliConfigRequest) (*v31.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadCliConfig not implemented")
}
func (UnimplementedUserServiceServer) UserListApiKeys(context.Context, *ApiKeyRequest) (*UserListApiKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserListApiKeys not implemented")
}
func (UnimplementedUserServiceServer) UserDeleteApiKeys(context.Context, *ApiKeyRequest) (*UserDeleteApiKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDeleteApiKeys not implemented")
}
func (UnimplementedUserServiceServer) UserForgotPassword(context.Context, *UserForgotPasswordRequest) (*UserForgotPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserForgotPassword not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_AuditLogWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginAuditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AuditLogWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.UserService/AuditLogWebhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AuditLogWebhook(ctx, req.(*UserLoginAuditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*v3.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v31.QueryOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.UserService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsers(ctx, req.(*v31.QueryOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*v3.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.UserService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfo(ctx, req.(*v3.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*v3.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserForceReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateForceResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserForceReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.UserService/UpdateUserForceReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserForceReset(ctx, req.(*UpdateForceResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*v3.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DownloadCliConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CliConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DownloadCliConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.UserService/DownloadCliConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DownloadCliConfig(ctx, req.(*CliConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserListApiKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserListApiKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.UserService/UserListApiKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserListApiKeys(ctx, req.(*ApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserDeleteApiKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserDeleteApiKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.UserService/UserDeleteApiKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserDeleteApiKeys(ctx, req.(*ApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.user.v3.UserService/UserForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserForgotPassword(ctx, req.(*UserForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "paralus.dev.rpc.user.v3.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuditLogWebhook",
			Handler:    _UserService_AuditLogWebhook_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserService_GetUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserService_GetUserInfo_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateUserForceReset",
			Handler:    _UserService_UpdateUserForceReset_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "DownloadCliConfig",
			Handler:    _UserService_DownloadCliConfig_Handler,
		},
		{
			MethodName: "UserListApiKeys",
			Handler:    _UserService_UserListApiKeys_Handler,
		},
		{
			MethodName: "UserDeleteApiKeys",
			Handler:    _UserService_UserDeleteApiKeys_Handler,
		},
		{
			MethodName: "UserForgotPassword",
			Handler:    _UserService_UserForgotPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/user/user.proto",
}
