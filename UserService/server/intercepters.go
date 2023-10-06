package server

import (
	"context"
	"google.golang.org/grpc"
)

func (u *UserService) BasicAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Проверяем, доступен ли метод для администраторов только
	/*if strings.HasPrefix(info.FullMethod, "/proto.AdminService/") {
		// Извлекаем имя пользователя и пароль из заголовка Authorization
		// ...
		user, err := s.db.GetUser(username)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "Неверные учетные данные")
		}
		if !user.IsAdmin {
			return nil, status.Error(codes.PermissionDenied, "Требуются админские права")
		}
	}*/
	return handler(ctx, req)
}
