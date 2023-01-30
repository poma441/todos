package service

import (
	"todos/internal/entity"
	"todos/internal/repository"
)

type Authorization interface {
	PlugFunc()
}

type ToDoItem interface {
	GetToDoItemsList(userId int) ([]entity.ToDoItem, error)
	AddToDoItem(userId int) (entity.ToDoItem, error)
	UpdateToDoItem(toDoItemId int) (entity.ToDoItem, error)
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
