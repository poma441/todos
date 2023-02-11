package service

import (
	"time"
	"todos/internal/entity"
	"todos/internal/repository"
)

type Authorization interface {
	HashPass(inputPass string) string
	GetUser(inputUsername string) (entity.User, error)
	GetUserById(userId int) (entity.User, error)
	CreateUser(newUser entity.User) (int, error)
	ComparePass(userHashPass string, inputPass string) error
	CreateToken(ttl time.Duration, userId int, privateKey string) (string, error)
	ValidateToken(token string, publicKey string) (interface{}, error)
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
