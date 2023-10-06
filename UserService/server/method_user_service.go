package server

import (
	"context"
	protobuf "github.com/grpc_in_memory/UserService/protobuf/compiled"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (u *UserService) GetUserById(context.Context, *protobuf.User) (*protobuf.User, error) {

	return nil, nil
}

func (u *UserService) GetUserList(context.Context, *emptypb.Empty) (*protobuf.UserList, error) {

	return nil, nil
}

func (u *UserService) CreateUser(context.Context, *protobuf.User) (*protobuf.Response, error) {

	return nil, nil
}

func (u *UserService) UpdateUser(context.Context, *protobuf.User) (*protobuf.Response, error) {

	return nil, nil
}
func (u *UserService) DeleteUser(context.Context, *protobuf.User) (*protobuf.Response, error) {

	return nil, nil
}
