package repository

import (
	"todos/internal/entity"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(NewUser entity.User) (int, error)
	GetUser(InputUsername string) (entity.User, error)
	GetUserById(userId int) (entity.User, error)
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
		Authorization: NewAuthRepo(db),
		ToDoItem:      NewToDoItemRepo(db),
	}
}
