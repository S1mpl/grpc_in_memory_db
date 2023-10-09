package server

import (
	"github.com/bufbuild/protovalidate-go"
	userRepo "github.com/grpc_in_memory/UserService/entity/user/repository/mydb"
	userUC "github.com/grpc_in_memory/UserService/entity/user/usecase"
	"github.com/grpc_in_memory/UserService/pkg/database"
	protobuf "github.com/grpc_in_memory/UserService/protobuf/compiled"
	"github.com/rs/zerolog/log"
	"sync"
)

type UserService struct {
	protobuf.UnimplementedUserServiceServer
	protobuf.UnimplementedAdminUserServiceServer
	mu        sync.Mutex
	Validator *protovalidate.Validator
	DB        protobuf.DBServiceClient
	userUC    *userUC.UserUseCase
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

	userRepository := userRepo.NewUserRepository(db)

	userService := &UserService{
		Validator: validator,
		DB:        db,
		userUC:    userUC.NewUserUseCase(userRepository),
	}

	return userService
}
