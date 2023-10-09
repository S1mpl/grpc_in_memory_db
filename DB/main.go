package main

import (
	"context"
	"errors"
	"github.com/google/uuid"
	protobuf "github.com/grpc_in_memory/DB/protobuf/compiled"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
)

type server struct {
	protobuf.UnimplementedDatabaseServiceServer
	db map[string]*protobuf.Row
}

func (s *server) GetAll(ctx context.Context, in *protobuf.GetAllRequest) (*protobuf.AllRows, error) {
	var rows []*protobuf.Row

	for key, row := range s.db {
		if strings.HasPrefix(key, in.Table) {
			rows = append(rows, row)
		}
	}

	return &protobuf.AllRows{Rows: rows}, nil
}

func (s *server) GetByKey(ctx context.Context, in *protobuf.GetByKeyRequest) (*protobuf.Row, error) {
	return s.db[in.Key], nil
}

func (s *server) Insert(ctx context.Context, in *protobuf.Row) (*protobuf.Response, error) {
	in.Data["id"].Value = uuid.New().String()
	in.Data["id"].Unique = true

	for key, row := range s.db {
		if strings.HasPrefix(key, in.Key) {
			for k, v := range in.Data {
				if v.Unique && row.Data[k].Value == v.Value {
					return &protobuf.Response{Success: false}, errors.New("unique value exists")
				}
			}
		}
	}

	s.db[in.Key+"."+in.Data["id"].Value] = in
	return &protobuf.Response{Success: true}, nil
}

func (s *server) Update(ctx context.Context, in *protobuf.Row) (*protobuf.Response, error) {
	s.db[in.Key] = in
	return &protobuf.Response{Success: true}, nil
}

func (s *server) DeleteByKey(ctx context.Context, in *protobuf.DeleteByKeyRequest) (*protobuf.Response, error) {
	delete(s.db, in.Key)
	return &protobuf.Response{Success: true}, nil
}

func main() {
	godotenv.Load(".env")
	lis, err := net.Listen("tcp", os.Getenv("PORT"))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	protobuf.RegisterDatabaseServiceServer(grpcServer, &server{db: make(map[string]*protobuf.Row)})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
