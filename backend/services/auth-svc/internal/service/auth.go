package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"
	"todos/internal/entity"
	"todos/services/auth-svc/internal/pb"
	"todos/services/auth-svc/internal/repository"
	jwthelper "todos/services/auth-svc/pkg/jwt_helper"
	hasher "todos/services/auth-svc/pkg/password_hasher"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{
		Success: true,
		Error:   "",
		Uuid:    "",
		TokensInfo: &pb.TokensInfo{
			AccessToken:       "",
			RefreshToken:      "",
			RefreshCookiePath: "",
			RefreshCookieHost: "",
			LogoutCookiePath:  "",
			LogoutCookieHost:  "",
		},
	}, nil
}

func (s *AuthService) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		Success: true,
		Error:   "",
		Uuid:    "",
		TokensInfo: &pb.TokensInfo{
			AccessToken:       "",
			RefreshToken:      "",
			RefreshCookiePath: "",
			RefreshCookieHost: "",
			LogoutCookiePath:  "",
			LogoutCookieHost:  "",
		},
	}, nil
}

func (s *AuthService) Validate(ctx context.Context, r *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return &pb.ValidateResponse{
		Success: true,
		Error:   "",
		Uuid:    "",
	}, nil
}

func (s *AuthService) RefreshTokens(ctx context.Context, r *pb.RefreshTokensRequest) (*pb.RefreshTokensResponse, error) {
	return &pb.RefreshTokensResponse{
		Success: true,
		Error:   "",
		TokensInfo: &pb.TokensInfo{
			AccessToken:       "",
			RefreshToken:      "",
			RefreshCookiePath: "",
			RefreshCookieHost: "",
			LogoutCookiePath:  "",
			LogoutCookieHost:  "",
		},
	}, nil
}

func (s *AuthService) InvalidateTokens(ctx context.Context, r *pb.InvalidateTokensRequest) (*pb.InvalidateTokensResponse, error) {
	return &pb.InvalidateTokensResponse{
		TokensInfo: &pb.TokensInfo{
			AccessToken:       "",
			RefreshToken:      "",
			RefreshCookiePath: "",
			RefreshCookieHost: "",
			LogoutCookiePath:  "",
			LogoutCookieHost:  "",
		},
	}, nil
}

/*
*	Создание пользователя в БД
 */
func (s *AuthService) CreateUser(newUser entity.Student) (int, error) {
	return s.repo.CreateUser(newUser)
}

/*
*	Получение пользователя из БД по логину
 */
func (s *AuthService) GetUser(inputUsername string) (entity.Student, error) {
	return s.repo.GetUser(inputUsername)
}

/*
*	Получение пользователя из БД по id
 */
func (s *AuthService) GetUserById(userId int) (entity.User, error) {
	return s.repo.GetUserById(userId)
}

/*
*	Функция создания токенов, как access, так и refresh
 */
func (s *AuthService) CreateToken(userId int, ttl time.Duration, privateKey string, isRefreshToken bool, requestInfo *entity.RequestAdditionalInfo) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()
	claims := make(jwt.MapClaims)

	if isRefreshToken {
		randomHash, err := hasher.GenerateRandomHash256()
		if err != nil {
			return "", err
		}

		claims["sub"] = userId
		claims["jti"] = randomHash
		claims["exp"] = now.Add(ttl).Unix()

		// Запись информации о refresh токене в кэш redis
		if err := s.repo.SaveRefreshToken(userId, randomHash, ttl, requestInfo); err != nil {
			return "", err
		}

	} else {
		claims["sub"] = userId
		claims["exp"] = now.Add(ttl).Unix()
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) ValidateToken(token string, publicKey string, isRefreshToken bool, requestInfo *entity.RequestAdditionalInfo) (interface{}, error) {
	tokenParsed, err := jwthelper.RsaSignedTokenParse(token, publicKey)
	if err != nil {
		return nil, err
	}

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok || !tokenParsed.Valid {
		return nil, fmt.Errorf("validate: токен не валиден")
	}

	if isRefreshToken {
		requestInfo, err := s.repo.CheckRefreshToken(int(claims["sub"].(float64)), claims["jti"].(string))
		if err != nil {
			return nil, err
		}

		// Заглушка для requestInfo, возможно, в дальнейшем будет использование для выявления подозрительной активности
		fmt.Print(requestInfo)
	}

	return claims["sub"], nil
}

func (s *AuthService) InvalidateRefreshToken(refreshToken string, publicKey string) error {
	tokenParsed, err := jwthelper.RsaSignedTokenParse(refreshToken, publicKey)
	if err != nil {
		return err
	}

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok || !tokenParsed.Valid {
		return fmt.Errorf("validate: токен не валиден")
	}

	// Удаление из кэша redis
	userId := claims["sub"].(float64)
	refreshTokenHash := claims["jti"].(string)

	if err := s.repo.DeleteRefreshToken(int(userId), refreshTokenHash); err != nil {
		return err
	}

	return nil
}
