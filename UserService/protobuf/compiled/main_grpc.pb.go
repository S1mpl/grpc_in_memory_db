// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: protobuf/raw/main.proto

package compiled

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_GetUserById_FullMethodName = "/UserService/GetUserById"
	UserService_GetUserList_FullMethodName = "/UserService/GetUserList"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserById(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	GetUserList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserList, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserById(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_GetUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, UserService_GetUserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUserById(context.Context, *User) (*User, error)
	GetUserList(context.Context, *emptypb.Empty) (*UserList, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUserById(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServiceServer) GetUserList(context.Context, *emptypb.Empty) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserById(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _UserService_GetUserById_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _UserService_GetUserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/raw/main.proto",
}

const (
	AdminUserService_CreateUser_FullMethodName = "/AdminUserService/CreateUser"
	AdminUserService_UpdateUser_FullMethodName = "/AdminUserService/UpdateUser"
	AdminUserService_DeleteUser_FullMethodName = "/AdminUserService/DeleteUser"
)

// AdminUserServiceClient is the client API for AdminUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminUserServiceClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
}

type adminUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminUserServiceClient(cc grpc.ClientConnInterface) AdminUserServiceClient {
	return &adminUserServiceClient{cc}
}

func (c *adminUserServiceClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, AdminUserService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminUserServiceClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, AdminUserService_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminUserServiceClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, AdminUserService_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminUserServiceServer is the server API for AdminUserService service.
// All implementations must embed UnimplementedAdminUserServiceServer
// for forward compatibility
type AdminUserServiceServer interface {
	CreateUser(context.Context, *User) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
	DeleteUser(context.Context, *User) (*Response, error)
	mustEmbedUnimplementedAdminUserServiceServer()
}

// UnimplementedAdminUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminUserServiceServer struct {
}

func (UnimplementedAdminUserServiceServer) CreateUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAdminUserServiceServer) UpdateUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedAdminUserServiceServer) DeleteUser(context.Context, *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedAdminUserServiceServer) mustEmbedUnimplementedAdminUserServiceServer() {}

// UnsafeAdminUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminUserServiceServer will
// result in compilation errors.
type UnsafeAdminUserServiceServer interface {
	mustEmbedUnimplementedAdminUserServiceServer()
}

func RegisterAdminUserServiceServer(s grpc.ServiceRegistrar, srv AdminUserServiceServer) {
	s.RegisterService(&AdminUserService_ServiceDesc, srv)
}

func _AdminUserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminUserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminUserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminUserServiceServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminUserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminUserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminUserService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminUserServiceServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminUserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminUserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminUserService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminUserServiceServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminUserService_ServiceDesc is the grpc.ServiceDesc for AdminUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AdminUserService",
	HandlerType: (*AdminUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _AdminUserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AdminUserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AdminUserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/raw/main.proto",
}
