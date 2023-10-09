package server

import (
	"context"
	protobuf "github.com/grpc_in_memory/UserService/protobuf/compiled"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (u *UserService) GetUserById(ctx context.Context, user *protobuf.User) (*protobuf.User, error) {

	user, err := u.userUC.GetUserById(user)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return user, nil
}

func (u *UserService) GetUserList(context.Context, *emptypb.Empty) (*protobuf.UserList, error) {

	userList, err := u.userUC.GetUserList()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return userList, nil
}
