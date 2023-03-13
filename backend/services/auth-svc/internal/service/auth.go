package service

import (
	"context"
	"todos/services/api-gateway/grpc-svc-clients/auth/pb"
	"todos/services/auth-svc/internal/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func SignUp(ctx context.Context, req *pb.RegisterRequest) {

}
