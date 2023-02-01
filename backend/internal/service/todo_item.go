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
	return s.repo.GetToDoItemsList(userId)
}

func (s *ToDoItemService) AddToDoItem(toDoItemForAdd entity.ToDoItem) (int, error) {
	return s.repo.AddToDoItem(toDoItemForAdd)
}

func (s *ToDoItemService) UpdateToDoItem(toDoItemForUpdate entity.ToDoItem) error {
	return s.repo.UpdateToDoItem(toDoItemForUpdate)
}

func (s *ToDoItemService) DeleteToDoItem(toDoItemId int) (int, error) {
	return s.repo.DeleteToDoItem(toDoItemId)
}
