package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"todos/internal/entity"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type AuthRepo struct {
	db        *gorm.DB
	redisConn *redis.Client
}

func NewAuthRepo(db *gorm.DB, redisConn *redis.Client) *AuthRepo {
	return &AuthRepo{
		db:        db,
		redisConn: redisConn,
	}
}

func (r *AuthRepo) CreateUser(newUser entity.User) (int, error) {

	check := r.db.Where("email=?", newUser.Email).Find(&newUser)
	if check.RowsAffected != 0 {
		return -1, errors.New("пользователем с таким именем существует")
	}
	result := r.db.Create(&newUser)
	if result.RowsAffected == 0 {
		return -1, errors.New("не удалось создать пользователя")
	}

	return newUser.Id, nil
}

func (r *AuthRepo) GetUser(inputUsername string) (entity.User, error) {
	var user entity.User
	result := r.db.First(&user, "email=?", inputUsername)
	if result.RowsAffected == 0 {
		return user, errors.New("неверное имя пользователя или пароль")
	}

	return user, nil
}

func (r *AuthRepo) GetUserById(userId int) (entity.User, error) {

	var currentUser entity.User
	result := r.db.First(&currentUser, "id=?", userId)

	if result.RowsAffected == 0 {
		return currentUser, errors.New("id не найден")
	}

	return currentUser, nil
}

/*
*	Функция сохранения refresh токена в кэше redis
 */
func (r *AuthRepo) SaveRefreshToken(userId int, refreshTokenHash string, ttl time.Duration, requestInfo *entity.RequestAdditionalInfo) error {
	requestInfoJson, err := json.Marshal(requestInfo)
	if err != nil {
		return err
	}

	err = r.redisConn.Set(fmt.Sprintf("user:%d:%s", userId, refreshTokenHash), requestInfoJson, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

/*
*	Функция проверки наличия рефреш токена в кэше
 */
func (r *AuthRepo) CheckRefreshToken(userId int, refreshTokenHash string) (*entity.RequestAdditionalInfo, error) {
	requestUnmarshalledInfo := new(entity.RequestAdditionalInfo)
	key := fmt.Sprintf("user:%d:%s", userId, refreshTokenHash)

	// Проверка существования данного ключа
	if existance := r.redisConn.Exists(key).Val(); existance == 0 {
		return nil, errors.New("refresh токен не найден в кэше")
	}

	// Достаем информацию о запросе пользователя для ее дальнейшего анализа в usecase
	requestInfoJson, err := r.redisConn.Get(key).Result()
	if err == redis.Nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(requestInfoJson), requestUnmarshalledInfo); err != nil {
		return nil, err
	}

	return requestUnmarshalledInfo, nil
}

/*
*	Функция удаления refresh token'a из кэша redis
 */
func (r *AuthRepo) DeleteRefreshToken(userId int, refreshTokenHash string) error {
	key := fmt.Sprintf("user:%d:%s", userId, refreshTokenHash)

	// Проверка существования данного ключа
	if existance := r.redisConn.Exists(key).Val(); existance == 0 {
		return errors.New("refresh токен не найден в кэше")
	}

	// Удаление данного ключа
	if success := r.redisConn.Del(key).Val(); success == 0 {
		return errors.New("не удалось удалить refresh token")
	}

	return nil
}
