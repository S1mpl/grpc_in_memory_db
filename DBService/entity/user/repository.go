package user

import protobuf "github.com/grpc_in_memory/DBService/protobuf/compiled"

type UserRepository interface {
	CreateUser(user *protobuf.User) (*protobuf.User, error)
	UpdateUser(user *protobuf.User) (*protobuf.User, error)
	DeleteUser(user *protobuf.User) error
	GetUserById(user *protobuf.User) (*protobuf.User, error)
	GetUserList() (*protobuf.UserList, error)
	CheckUserExist(user *protobuf.User) (*protobuf.User, error)
}
