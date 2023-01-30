package service

import (
	"todos/internal/repository"
	service_routes "todos/internal/service/service_routes"
)

type Authorization interface {
	PlugFunc()
}

type ToDoItem interface {
	PlugFunc()
}

type Service struct {
	Authorization
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: service_routes.NewAuthService(repos.Authorization),
		ToDoItem:      service_routes.NewToDoItemService(repos.ToDoItem),
	}
}
