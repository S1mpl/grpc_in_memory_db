package mydb

import (
	"context"
	protobuf "github.com/grpc_in_memory/UserService/protobuf/compiled"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserRepository struct {
	db protobuf.DBServiceClient
}

func NewUserRepository(db protobuf.DBServiceClient) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *protobuf.User) (*protobuf.User, error) {
	user, err := r.db.CreateUser(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserRepository) UpdateUser(user *protobuf.User) (*protobuf.User, error) {
	user, err := r.db.UpdateUser(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserRepository) DeleteUser(user *protobuf.User) error {
	_, err := r.db.DeleteUser(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) GetUserById(user *protobuf.User) (*protobuf.User, error) {
	user, err := r.db.GetUserById(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserRepository) GetUserList() (*protobuf.UserList, error) {
	userList, err := r.db.GetUserList(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return userList, nil
}
func (r *UserRepository) CheckUserExist(user *protobuf.User) (*protobuf.User, error) {
	user, err := r.db.CheckUserExist(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
