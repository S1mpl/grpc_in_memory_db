package server

import (
	"context"
	protobuf "github.com/grpc_in_memory/DBService/protobuf/compiled"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (u *UserService) CreateUser(ctx context.Context, user *protobuf.User) (*protobuf.User, error) {

	user, err := u.UserRepository.CreateUser(user)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return user, nil

}

func (u *UserService) UpdateUser(ctx context.Context, user *protobuf.User) (*protobuf.User, error) {

	user, err := u.UserRepository.UpdateUser(user)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return user, nil
}
func (u *UserService) DeleteUser(ctx context.Context, user *protobuf.User) (*protobuf.Response, error) {

	err := u.UserRepository.DeleteUser(user)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &protobuf.Response{Response: &protobuf.Response_Status{Status: &protobuf.Status{
		Code:    int32(codes.OK),
		Message: "User deleted",
		Error:   false,
	}}}, nil
}

func (u *UserService) GetUserById(ctx context.Context, user *protobuf.User) (*protobuf.User, error) {

	user, err := u.UserRepository.GetUserById(user)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return user, nil
}

func (u *UserService) GetUserList(context.Context, *emptypb.Empty) (*protobuf.UserList, error) {

	userList, err := u.UserRepository.GetUserList()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return userList, nil
}
