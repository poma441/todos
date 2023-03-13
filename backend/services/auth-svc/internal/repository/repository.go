package repository

import (
	"time"
	"todos/internal/entity"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(newUser entity.Student) (int, error)
	GetUser(inputUsername string) (entity.Student, error)
	GetUserById(userId int) (entity.User, error)

	// Работа с кэшем refresh токенов
	SaveRefreshToken(userId int, refreshTokenHash string, ttl time.Duration, requestInfo *entity.RequestAdditionalInfo) error
	CheckRefreshToken(userId int, refreshTokenHash string) (*entity.RequestAdditionalInfo, error)
	DeleteRefreshToken(userId int, refreshTokenHash string) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB, redisConn *redis.Client) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db, redisConn),
	}
}
