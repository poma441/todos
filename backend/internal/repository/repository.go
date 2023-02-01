package repository

import (
	"todos/internal/entity"

	"gorm.io/gorm"
)

type Authorization interface {
}

type ToDoItem interface {
	GetToDoItemsList(userId int) ([]entity.ToDoItem, error)
	AddToDoItem(userId int) (entity.ToDoItem, error)
	UpdateToDoItem(toDoItemId int) (entity.ToDoItem, error)
	DeleteToDoItem(toDoItemId int) (int, error)
}

type Repository struct {
	Authorization
	ToDoItem
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(),
		ToDoItem:      NewToDoItemRepo(db),
	}
}
