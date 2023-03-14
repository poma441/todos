package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"
	"todos/services/auth-svc/config"
	"todos/services/auth-svc/internal/pb"
	"todos/services/auth-svc/internal/repository"
	"todos/services/auth-svc/models"
	jwthelper "todos/services/auth-svc/pkg/jwt_helper"
	hasher "todos/services/auth-svc/pkg/password_hasher"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type AuthService struct {
	cfg  *config.Config
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization, cfg *config.Config) *AuthService {
	return &AuthService{
		cfg:  cfg,
		repo: repo,
	}
}

func (s *AuthService) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	requestInfo := &models.RequestAdditionalInfo{
		UserAgent: r.RequestInfo.GetUserAgent(),
		SrcIP:     r.RequestInfo.GetSrcIpAddr(),
	}

	user := models.User{
		Email:    r.Email,
		Password: hasher.HashPass(r.Password),
		Role:     r.Role,
		Phone:    r.Phone,
		FullName: r.Fullname,
	}

	// Создание пользователя
	user.Uuid = uuid.New().String()

	_, err := s.repo.CreateUser(user)
	if err != nil {
		return &pb.RegisterResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Генерация access токена
	accessToken, err := s.createToken(user.Uuid, s.cfg.Token.Access.TTL, s.cfg.Token.Keys.PrivateKey, false, requestInfo)
	if err != nil {
		return &pb.RegisterResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Генерация refresh токена
	refreshToken, err := s.createToken(user.Uuid, s.cfg.Token.Refresh.TTL, s.cfg.Token.Keys.PrivateKey, true, requestInfo)
	if err != nil {
		return &pb.RegisterResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Success: true,
		Error:   "",
		Uuid:    user.Uuid,
		TokensInfo: &pb.TokensInfo{
			AccessToken:       accessToken,
			RefreshToken:      refreshToken,
			RefreshCookieHost: s.cfg.Server.Host,
			LogoutCookieHost:  s.cfg.Server.Host,
			RefreshCookiePath: s.cfg.Token.Refresh.RefreshCookiePath,
			LogoutCookiePath:  s.cfg.Token.Refresh.LogoutCookiePath,
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
*	Получение пользователя из БД по логину
 */
func (s *AuthService) getUser(inputUsername string) (models.User, error) {
	return s.repo.GetUser(inputUsername)
}

/*
*	Получение пользователя из БД по id
 */
func (s *AuthService) getUserById(userId int) (models.User, error) {
	return s.repo.GetUserById(userId)
}

/*
*	Функция создания токенов, как access, так и refresh
 */
func (s *AuthService) createToken(uuid string, ttl time.Duration, privateKey string, isRefreshToken bool, requestInfo *models.RequestAdditionalInfo) (string, error) {
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

		claims["sub"] = uuid
		claims["jti"] = randomHash
		claims["exp"] = now.Add(ttl).Unix()

		// Запись информации о refresh токене в кэш redis
		if err := s.repo.SaveRefreshToken(uuid, randomHash, ttl, requestInfo); err != nil {
			return "", err
		}

	} else {
		claims["sub"] = uuid
		claims["exp"] = now.Add(ttl).Unix()
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) validateToken(token string, publicKey string, isRefreshToken bool, requestInfo *models.RequestAdditionalInfo) (interface{}, error) {
	tokenParsed, err := jwthelper.RsaSignedTokenParse(token, publicKey)
	if err != nil {
		return nil, err
	}

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok || !tokenParsed.Valid {
		return nil, fmt.Errorf("validate: токен не валиден")
	}

	if isRefreshToken {
		requestInfo, err := s.repo.CheckRefreshToken(claims["sub"].(string), claims["jti"].(string))
		if err != nil {
			return nil, err
		}

		// Заглушка для requestInfo, возможно, в дальнейшем будет использование для выявления подозрительной активности
		fmt.Print(requestInfo)
	}

	return claims["sub"], nil
}

func (s *AuthService) invalidateRefreshToken(refreshToken string, publicKey string) error {
	tokenParsed, err := jwthelper.RsaSignedTokenParse(refreshToken, publicKey)
	if err != nil {
		return err
	}

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok || !tokenParsed.Valid {
		return fmt.Errorf("validate: токен не валиден")
	}

	// Удаление из кэша redis
	uuid := claims["sub"].(string)
	refreshTokenHash := claims["jti"].(string)

	if err := s.repo.DeleteRefreshToken(uuid, refreshTokenHash); err != nil {
		return err
	}

	return nil
}
