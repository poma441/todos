package service

import (
	"time"
	"todos/internal/entity"
	"todos/internal/repository"
)

type Authorization interface {
	// Работа с пользователем
	CreateUser(newUser entity.User) (int, error)
	GetUser(inputUsername string) (entity.User, error)
	GetUserById(userId int) (entity.User, error)

	// Работа с токенами
	CreateToken(userId int, ttl time.Duration, privateKey string, isRefreshToken bool, requestInfo *entity.RequestAdditionalInfo) (string, error)
	ValidateToken(token string, publicKey string, isRefreshToken bool, requestInfo *entity.RequestAdditionalInfo) (interface{}, error)
	InvalidateRefreshToken(refreshToken string, publicKey string) error
}

type ToDoItem interface {
	GetToDoItemsList(userId int) ([]entity.ToDoItem, error)
	AddToDoItem(toDoItem entity.ToDoItem, toDoItemId int) (int, error)
	UpdateToDoItem(toDoItem entity.ToDoItem, toDoItemId int) error
	DeleteToDoItem(toDoItemId int) (int, error)
}

type Service struct {
	Authorization
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ToDoItem:      NewToDoItemService(repos.ToDoItem),
	}
}
