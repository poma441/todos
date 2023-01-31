package service

import (
	"errors"
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
	// Заглушка
	var list []entity.ToDoItem = make([]entity.ToDoItem, 3)

	item1 := entity.ToDoItem{
		Id:     0,
		Title:  "Get Test 1",
		Status: true,
		UserId: 0,
	}

	item2 := entity.ToDoItem{
		Id:     0,
		Title:  "Get Test 2",
		Status: true,
		UserId: 0,
	}

	item3 := entity.ToDoItem{
		Id:     0,
		Title:  "Get Test 3",
		Status: true,
		UserId: 0,
	}

	list[0] = item1
	list[1] = item2
	list[2] = item3

	return list, errors.New("")
}

func (s *ToDoItemService) AddToDoItem(userId int) (entity.ToDoItem, error) {
	// Заглушка
	var list entity.ToDoItem
	list.Id = 0
	list.Title = "Add Test"
	list.Status = true
	list.UserId = 0

	return list, errors.New("")
}

func (s *ToDoItemService) UpdateToDoItem(toDoItemId int) (entity.ToDoItem, error) {
	// Заглушка
	var list entity.ToDoItem
	list.Id = 0
	list.Title = "Update Test"
	list.Status = true
	list.UserId = 0

	return list, errors.New("")
}

func (s *ToDoItemService) DeleteToDoItem(toDoItemId int) (int, error) {
	// Заглушка
	return 0, errors.New("")
}
