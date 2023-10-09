package server

import (
	"context"
	protobuf "github.com/grpc_in_memory/UserService/protobuf/compiled"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *UserService) CreateUser(ctx context.Context, user *protobuf.User) (*protobuf.User, error) {

	user, err := u.userUC.CreateUser(user)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return user, nil

}

func (u *UserService) UpdateUser(ctx context.Context, user *protobuf.User) (*protobuf.User, error) {

	user, err := u.userUC.UpdateUser(user)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return user, nil
}
func (u *UserService) DeleteUser(ctx context.Context, user *protobuf.User) (*protobuf.Response, error) {

	err := u.userUC.DeleteUser(user)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &protobuf.Response{Response: &protobuf.Response_Status{Status: &protobuf.Status{
		Code:    int32(codes.OK),
		Message: "User deleted",
		Error:   false,
	}}}, nil
}
