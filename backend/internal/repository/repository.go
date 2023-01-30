package repository

import (
	repository_routes "todos/internal/repository/repository_routes"
)

type Authorization interface {
}

type ToDoItem interface {
}

type Repository struct {
	Authorization
	ToDoItem
}

func NewRepository() *Repository {
	return &Repository{
		Authorization: repository_routes.NewAuthRepo(),
		ToDoItem:      repository_routes.NewToDoItemRepo(),
	}
}
