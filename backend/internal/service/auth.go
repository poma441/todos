package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"
	"todos/internal/entity"
	"todos/internal/repository"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) HashPass(inputPass string) string {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(inputPass), bcrypt.DefaultCost)
	if err != nil {
		return ""

	}

	return string(hashPass)
}

func (s *AuthService) CreateUser(newUser entity.User) (int, error) {
	return s.repo.CreateUser(newUser)
}

func (s *AuthService) GetUser(inputUsername string) (entity.User, error) {

	return s.repo.GetUser(inputUsername)
}

func (s *AuthService) GetUserById(userId int) (entity.User, error) {

	return s.repo.GetUserById(userId)
}

func (s *AuthService) ComparePass(userHashPass string, inputPass string) error {

	err := bcrypt.CompareHashAndPassword([]byte(userHashPass), []byte(inputPass))
	if err != nil {
		return errors.New("неверное имя пользователя или пароль")
	}
	return err
}

func (s *AuthService) CreateToken(ttl time.Duration, userId int, privateKey string) (string, error) {

	decodePrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodePrivateKey)
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)

	claims["sub"] = userId
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (s *AuthService) ValidateToken(token string, publicKey string) (interface{}, error) {
	decodePublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodePublicKey)
	if err != nil {
		return "", fmt.Errorf("validate: ошибка: %w", err)
	}

	tokenParsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("непредвиденный метод: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok || !tokenParsed.Valid {
		return nil, fmt.Errorf("validate: токен не валиден")
	}

	return claims["sub"], nil

}
