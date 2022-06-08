// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/rpc/user/user.proto

package rpcv3

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	CreateUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.User, error)
	GetUsers(ctx context.Context, in *v31.QueryOptions, opts ...grpc.CallOption) (*v3.UserList, error)
	GetUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.User, error)
	GetUserInfo(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.UserInfo, error)
	UpdateUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.User, error)
	DeleteUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	DownloadCliConfig(ctx context.Context, in *CliConfigRequest, opts ...grpc.CallOption) (*v31.HttpBody, error)
	UserListApiKeys(ctx context.Context, in *ApiKeyRequest, opts ...grpc.CallOption) (*ApiKeyResponseList, error)
	UserDeleteApiKeys(ctx context.Context, in *ApiKeyRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	UserForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) CreateUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.User, error) {
	out := new(v3.User)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUsers(ctx context.Context, in *v31.QueryOptions, opts ...grpc.CallOption) (*v3.UserList, error) {
	out := new(v3.UserList)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.User/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.User, error) {
	out := new(v3.User)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.User/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserInfo(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.UserInfo, error) {
	out := new(v3.UserInfo)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.User/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*v3.User, error) {
	out := new(v3.User)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUser(ctx context.Context, in *v3.User, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.User/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DownloadCliConfig(ctx context.Context, in *CliConfigRequest, opts ...grpc.CallOption) (*v31.HttpBody, error) {
	out := new(v31.HttpBody)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.User/DownloadCliConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserListApiKeys(ctx context.Context, in *ApiKeyRequest, opts ...grpc.CallOption) (*ApiKeyResponseList, error) {
	out := new(ApiKeyResponseList)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.User/UserListApiKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserDeleteApiKeys(ctx context.Context, in *ApiKeyRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.User/UserDeleteApiKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error) {
	out := new(ForgotPasswordResponse)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.User/UserForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations should embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	CreateUser(context.Context, *v3.User) (*v3.User, error)
	GetUsers(context.Context, *v31.QueryOptions) (*v3.UserList, error)
	GetUser(context.Context, *v3.User) (*v3.User, error)
	GetUserInfo(context.Context, *v3.User) (*v3.UserInfo, error)
	UpdateUser(context.Context, *v3.User) (*v3.User, error)
	DeleteUser(context.Context, *v3.User) (*DeleteUserResponse, error)
	DownloadCliConfig(context.Context, *CliConfigRequest) (*v31.HttpBody, error)
	UserListApiKeys(context.Context, *ApiKeyRequest) (*ApiKeyResponseList, error)
	UserDeleteApiKeys(context.Context, *ApiKeyRequest) (*DeleteUserResponse, error)
	UserForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error)
}

// UnimplementedUserServer should be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) CreateUser(context.Context, *v3.User) (*v3.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServer) GetUsers(context.Context, *v31.QueryOptions) (*v3.UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserServer) GetUser(context.Context, *v3.User) (*v3.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) GetUserInfo(context.Context, *v3.User) (*v3.UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServer) UpdateUser(context.Context, *v3.User) (*v3.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServer) DeleteUser(context.Context, *v3.User) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServer) DownloadCliConfig(context.Context, *CliConfigRequest) (*v31.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadCliConfig not implemented")
}
func (UnimplementedUserServer) UserListApiKeys(context.Context, *ApiKeyRequest) (*ApiKeyResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserListApiKeys not implemented")
}
func (UnimplementedUserServer) UserDeleteApiKeys(context.Context, *ApiKeyRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDeleteApiKeys not implemented")
}
func (UnimplementedUserServer) UserForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserForgotPassword not implemented")
}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*v3.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v31.QueryOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.User/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUsers(ctx, req.(*v31.QueryOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.User/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*v3.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.User/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserInfo(ctx, req.(*v3.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*v3.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.User/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUser(ctx, req.(*v3.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DownloadCliConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CliConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DownloadCliConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.User/DownloadCliConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DownloadCliConfig(ctx, req.(*CliConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserListApiKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserListApiKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.User/UserListApiKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserListApiKeys(ctx, req.(*ApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserDeleteApiKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserDeleteApiKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.User/UserDeleteApiKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserDeleteApiKeys(ctx, req.(*ApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.User/UserForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserForgotPassword(ctx, req.(*ForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rafay.dev.rpc.v3.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _User_GetUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _User_GetUserInfo_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _User_DeleteUser_Handler,
		},
		{
			MethodName: "DownloadCliConfig",
			Handler:    _User_DownloadCliConfig_Handler,
		},
		{
			MethodName: "UserListApiKeys",
			Handler:    _User_UserListApiKeys_Handler,
		},
		{
			MethodName: "UserDeleteApiKeys",
			Handler:    _User_UserDeleteApiKeys_Handler,
		},
		{
			MethodName: "UserForgotPassword",
			Handler:    _User_UserForgotPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/user/user.proto",
}
