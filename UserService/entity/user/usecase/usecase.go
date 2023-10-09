package usecase

import (
	"github.com/grpc_in_memory/UserService/entity/user/repository/mydb"
	protobuf "github.com/grpc_in_memory/UserService/protobuf/compiled"
)

type UserUseCase struct {
	userRepo *mydb.UserRepository
}

func NewUserUseCase(userRepo *mydb.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (u *UserUseCase) CreateUser(user *protobuf.User) (*protobuf.User, error) {

	user, err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) UpdateUser(user *protobuf.User) (*protobuf.User, error) {
	user, err := u.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (u *UserUseCase) DeleteUser(user *protobuf.User) error {

	err := u.userRepo.DeleteUser(user)
	if err != nil {
		return err
	}

	return nil
}
func (u *UserUseCase) GetUserById(user *protobuf.User) (*protobuf.User, error) {
	user, err := u.userRepo.GetUserById(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (u *UserUseCase) GetUserList() (*protobuf.UserList, error) {

	userList, err := u.userRepo.GetUserList()
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (u *UserUseCase) CheckUserExist(user *protobuf.User) (*protobuf.User, error) {
	user, err := u.userRepo.CheckUserExist(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
