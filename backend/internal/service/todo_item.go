package service

import (
	"todos/internal/entity"
	"todos/internal/repository"
)

type ToDoItemService struct {
	repo repository.ToDoItem
}

func NewToDoItemService(repo repository.ToDoItem) *ToDoItemService {
	return &ToDoItemService{
		repo: repo,
	}
}

func (s *ToDoItemService) GetToDoItemsList(userId int) ([]entity.ToDoItem, error) {
}

func (s *ToDoItemService) AddToDoItem(userId int) (entity.ToDoItem, error) {
}

func (s *ToDoItemService) UpdateToDoItem(toDoItemId int) (entity.ToDoItem, error) {
}

func (s *ToDoItemService) DeleteToDoItem(toDoItemId int) (int, error) {
}
