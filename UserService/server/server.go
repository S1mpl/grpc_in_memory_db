package server

import (
	"github.com/bufbuild/protovalidate-go"
	"github.com/grpc_in_memory/UserService/pkg/database"
	protobuf "github.com/grpc_in_memory/UserService/protobuf/compiled"
	"github.com/rs/zerolog/log"
	"sync"
)

type UserService struct {
	protobuf.UnimplementedUserServiceServer
	mu        sync.Mutex
	Validator *protovalidate.Validator
	DB        protobuf.DBServiceClient
}

func NewUserService() *UserService {

	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&protobuf.User{},
			&protobuf.UserList{},
		))

	if err != nil {
		log.Error().Msg(err.Error())
	}

	db := database.InitMyDB()

	userService := &UserService{
		Validator: validator,
		DB:        db,
	}

	return userService
}
