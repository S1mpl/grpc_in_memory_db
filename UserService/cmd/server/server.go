package main

import (
	"github.com/grpc_in_memory/UserService/config"
	protobuf "github.com/grpc_in_memory/UserService/protobuf/compiled"
	"github.com/grpc_in_memory/UserService/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	config.Init()

	log.Log().Msgf("listening on port %s", os.Getenv("PORT"))
	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Error().Msgf("failed to listen: %v", err)
	}

	userService := server.NewUserService()

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(userService.ValidateInterceptor, userService.ValidateInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	protobuf.RegisterUserServiceServer(grpcServer, userService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Error().Msg(err.Error())
	}
}
