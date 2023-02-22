package repository

import (
	"errors"
	"todos/internal/entity"

	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db: db}
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
	result := r.db.First(&user, "username=?", inputUsername)
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
