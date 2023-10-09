package user

import protobuf "github.com/grpc_in_memory/UserService/protobuf/compiled"

type UserUseCase interface {
	CreateUser(user *protobuf.User) (*protobuf.User, error)
	UpdateUser(user *protobuf.User) (*protobuf.User, error)
	DeleteUser(user *protobuf.User) (*protobuf.Response, error)
	GetUserById(user *protobuf.User) (*protobuf.User, error)
	GetUserList() (*protobuf.UserList, error)
	CheckUserExist(user *protobuf.User) (*protobuf.User, error)
}
