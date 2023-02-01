package repository

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"todos/internal/entity"

	"gorm.io/gorm"
)

type ToDoItemRepo struct {
	db *gorm.DB
}

func NewToDoItemRepo(db *gorm.DB) *ToDoItemRepo {
	return &ToDoItemRepo{db: db}
}

func (r *ToDoItemRepo) GetToDoItemsList(userId int) ([]entity.ToDoItem, error) {
	var items []entity.ToDoItem

	result := r.db.Where("user_id = ?", userId).Find(&items)

	if result.RowsAffected == 0 {
		log.Print("No select")
		return nil, errors.New("Список дел пуст")
	}

	return items, nil
}

func (r *ToDoItemRepo) AddToDoItem(toDoItemForAdd entity.ToDoItem) (int, error) {
	var items entity.ToDoItem
	fmt.Println(items)
	result := r.db.Where("user_id = ?", toDoItemForAdd.UserId).Create(&items)
	fmt.Println(&result)
	if result.RowsAffected == 0 {
		log.Print("No created")
		return -1, errors.New("Не удалось создать дело")
	}
	fmt.Println(items)
	return items.Id, nil
}

func (r *ToDoItemRepo) UpdateToDoItem(toDoItemForUpdate entity.ToDoItem) error {
	var items entity.ToDoItem

	result := r.db.Model(&items).Where("id=?", toDoItemForUpdate.Id).Updates(items)
	if result.RowsAffected == 0 {
		log.Print("No update")
		return errors.New("Не удалось обновить информацию о деле с id = " + strconv.Itoa(toDoItemForUpdate.Id))
	}

	return nil
}

func (r *ToDoItemRepo) DeleteToDoItem(toDoItemId int) (int, error) {
	var items entity.ToDoItem

	result := r.db.Where("id= ?", toDoItemId).Delete(&items)
	if result.RowsAffected == 0 {
		log.Print("No delete")
		return -1, errors.New("Не удалось удалить дело с id = " + strconv.Itoa(toDoItemId))
	}

	return toDoItemId, nil
}
