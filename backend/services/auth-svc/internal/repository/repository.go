package repository

import (
	"time"
	"todos/services/auth-svc/models"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(newUser models.User) (int, error)
	GetUser(inputUsername string) (models.User, error)
	GetUserById(userId int) (models.User, error)

	// Работа с кэшем refresh токенов
	SaveRefreshToken(uuid string, refreshTokenHash string, ttl time.Duration, requestInfo *models.RequestAdditionalInfo) error
	CheckRefreshToken(uuid string, refreshTokenHash string) (*models.RequestAdditionalInfo, error)
	DeleteRefreshToken(uuid string, refreshTokenHash string) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB, redisConn *redis.Client) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db, redisConn),
	}
}
