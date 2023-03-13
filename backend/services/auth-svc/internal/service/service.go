package service

import (
	"context"
	"time"
	"todos/internal/entity"
	"todos/services/auth-svc/internal/pb"
	"todos/services/auth-svc/internal/repository"
)

/************************/
/*  Прототипы сервисов  */
/************************/
type Authorization interface {
	// Работа с пользователем
	CreateUser(newUser entity.Student) (int, error)
	GetUser(inputUsername string) (entity.Student, error)
	GetUserById(userId int) (entity.User, error)

	// Работа с токенами
	CreateToken(userId int, ttl time.Duration, privateKey string, isRefreshToken bool, requestInfo *entity.RequestAdditionalInfo) (string, error)
	ValidateToken(token string, publicKey string, isRefreshToken bool, requestInfo *entity.RequestAdditionalInfo) (interface{}, error)
	InvalidateRefreshToken(refreshToken string, publicKey string) error

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

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization),
	}
}
