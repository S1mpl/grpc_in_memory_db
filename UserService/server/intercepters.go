package server

import (
	"context"
	"encoding/base64"
	protobuf "github.com/grpc_in_memory/UserService/protobuf/compiled"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"strings"
)

func (u *UserService) BasicAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	authorization := md["authorization"]
	if len(authorization) < 1 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	token := strings.TrimPrefix(authorization[0], "Basic ")
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "incorrect authorization token")
	}
	userAndPassword := strings.SplitN(string(data), ":", 2)
	if len(userAndPassword) < 2 {
		return nil, status.Errorf(codes.Unauthenticated, "incorrect authorization token")
	}

	user, err := u.userUC.CheckUserExist(&protobuf.User{
		Username: &userAndPassword[0],
		Password: &userAndPassword[1],
	})

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "incorrect login or password")
	}

	if strings.HasPrefix(info.FullMethod, "/proto.Admin/") {
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "incorrect login or password")
		}
		if !*user.Admin {
			return nil, status.Error(codes.PermissionDenied, "You are not admin")
		}
	}
	return handler(ctx, req)
}

func (u *UserService) ValidateInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	err := u.Validator.Validate(req.(proto.Message))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Message")
	}

	return handler(ctx, req)
}
