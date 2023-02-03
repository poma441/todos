package repository

import (
	"todos/internal/entity"

	"gorm.io/gorm"
)

type Authorization interface {
}

type ToDoItem interface {
	GetToDoItemsList(int) ([]entity.ToDoItem, error)
	AddToDoItem(toDoItem entity.ToDoItem, toDoItemId int) (int, error)
	UpdateToDoItem(toDoItem entity.ToDoItem, toDoItemId int) error
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
