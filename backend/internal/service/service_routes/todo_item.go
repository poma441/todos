package service

import "todos/internal/repository"

type ToDoItemService struct {
	repo repository.ToDoItem
}

func NewToDoItemService(repo repository.ToDoItem) *ToDoItemService {
	return &ToDoItemService{
		repo: repo,
	}
}

func (s *ToDoItemService) PlugFunc() {

}
