package repository

import (
	"errors"
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
		return nil, errors.New("список дел пуст")
	}

	return items, nil
}

func (r *ToDoItemRepo) AddToDoItem(toDoItemForAdd entity.ToDoItem) (int, error) {

	result := r.db.Where("user_id = ?", toDoItemForAdd.UserId).Create(&toDoItemForAdd)

	if result.RowsAffected == 0 {
		log.Print("No created")
		return -1, errors.New("не удалось создать дело")
	}

	return toDoItemForAdd.Id, nil
}

func (r *ToDoItemRepo) UpdateToDoItem(toDoItemForUpdate entity.ToDoItem) error {

	result := r.db.Model(&toDoItemForUpdate).Where("id=?", toDoItemForUpdate.Id).Updates(&toDoItemForUpdate)
	if result.RowsAffected == 0 {
		log.Print("No update")
		return errors.New("не удалось обновить информацию о деле с id = " + strconv.Itoa(toDoItemForUpdate.Id))
	}

	return nil
}

func (r *ToDoItemRepo) DeleteToDoItem(toDoItemId int) (int, error) {
	var items entity.ToDoItem

	result := r.db.Where("id= ?", toDoItemId).Delete(&items)
	if result.RowsAffected == 0 {
		log.Print("No delete")
		return -1, errors.New("не удалось удалить дело с id = " + strconv.Itoa(toDoItemId))
	}

	return toDoItemId, nil
}
