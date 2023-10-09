package mydb

import (
	"context"
	protobuf "github.com/grpc_in_memory/DBService/protobuf/compiled"
)

type UserRepository struct {
	db protobuf.DatabaseServiceClient
}

func NewUserRepository(db protobuf.DatabaseServiceClient) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *protobuf.User) (*protobuf.User, error) {
	data := make(map[string]*protobuf.Value)
	data = protobuf.ProtoUserToValueMyDB(user)
	_, err := r.db.Insert(context.Background(), &protobuf.Row{
		Key:  "users",
		Data: data,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserRepository) UpdateUser(user *protobuf.User) (*protobuf.User, error) {
	data := make(map[string]*protobuf.Value)
	data = protobuf.ProtoUserToValueMyDB(user)
	_, err := r.db.Update(context.Background(), &protobuf.Row{
		Key:  "users",
		Data: data,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserRepository) DeleteUser(user *protobuf.User) error {
	_, err := r.db.DeleteByKey(context.Background(), &protobuf.DeleteByKeyRequest{
		Key: user.GetId(),
	})
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) GetUserById(user *protobuf.User) (*protobuf.User, error) {
	row, err := r.db.GetByKey(context.Background(), &protobuf.GetByKeyRequest{Key: user.GetId()})
	if err != nil {
		return nil, err
	}
	user = protobuf.MyDBValueToProtoUser(row.Data)
	return user, nil
}
func (r *UserRepository) GetUserList() (*protobuf.UserList, error) {
	rows, err := r.db.GetAll(context.Background(), &protobuf.GetAllRequest{
		Table: "users",
	})
	if err != nil {
		return nil, err
	}
	userList := protobuf.MyDBValueToProtoUserList(rows)
	return userList, nil
}
func (r *UserRepository) CheckUserExist(user *protobuf.User) (*protobuf.User, error) {
	row, err := r.db.GetByKey(context.Background(), &protobuf.GetByKeyRequest{Key: user.GetId()})
	user = protobuf.MyDBValueToProtoUser(row.Data)
	if err != nil {
		return nil, err
	}
	return user, nil
}
