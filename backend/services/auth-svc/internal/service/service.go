package service

import (
	"context"
	"time"
	"todos/services/auth-svc/config"
	"todos/services/auth-svc/internal/pb"
	"todos/services/auth-svc/internal/repository"
	"todos/services/auth-svc/models"
)

/************************/
/*  Прототипы сервисов  */
/************************/
type Authorization interface {
	// Работа с пользователем
	getUser(inputUsername string) (models.User, error)
	getUserById(userId int) (models.User, error)

	// Работа с токенами
	createToken(uuid string, ttl time.Duration, privateKey string, isRefreshToken bool, requestInfo *models.RequestAdditionalInfo) (string, error)
	validateToken(token string, publicKey string, isRefreshToken bool, requestInfo *models.RequestAdditionalInfo) (interface{}, error)
	invalidateRefreshToken(refreshToken string, publicKey string) error

	// Процедуры gRPC
	Register(context.Context, *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error)
	Validate(context.Context, *pb.ValidateRequest) (*pb.ValidateResponse, error)
	RefreshTokens(context.Context, *pb.RefreshTokensRequest) (*pb.RefreshTokensResponse, error)
	InvalidateTokens(context.Context, *pb.InvalidateTokensRequest) (*pb.InvalidateTokensResponse, error)
}

/************************/
/*  Хранилище сервисов  */
/************************/
type Service struct {
	Authorization
}

func NewService(r *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization, cfg),
	}
}
