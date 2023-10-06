package database

import (
	protobuf "github.com/grpc_in_memory/UserService/protobuf/compiled"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"os"
)

func InitMyDB() protobuf.DBServiceClient {
	conn, err := grpc.Dial(os.Getenv("DB_URL"))
	if err != nil {
		log.Err(err)
	}

	client := protobuf.NewDBServiceClient(conn)
	return client
}
