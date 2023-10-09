package server

import (
	"github.com/grpc_in_memory/DBService/entity/user/repository/mydb"
	userRepo "github.com/grpc_in_memory/DBService/entity/user/repository/mydb"
	"github.com/grpc_in_memory/DBService/pkg/database"
	protobuf "github.com/grpc_in_memory/DBService/protobuf/compiled"
	"github.com/joho/godotenv"
	"sync"
)

type UserService struct {
	protobuf.UnimplementedUserServiceServer
	mu             sync.Mutex
	DB             protobuf.DatabaseServiceClient
	UserRepository *mydb.UserRepository
}

func NewUserService() *UserService {
	godotenv.Load(".env")

	db := database.InitMyDB()

	userRepository := userRepo.NewUserRepository(db)

	userService := &UserService{
		DB:             db,
		UserRepository: userRepository,
	}

	return userService
}
