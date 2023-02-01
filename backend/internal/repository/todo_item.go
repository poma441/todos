package repository

import (
	"fmt"
	"log"
	"todos/internal/entity"

	"gorm.io/gorm"
)

type ToDoItemRepo struct {
	db *gorm.DB
}

func NewToDoItemRepo(db *gorm.DB) *ToDoItemRepo {
	return &ToDoItemRepo{db: db}
}

func (r *ToDoItemRepo) AddToDoItem(userId int) (entity.ToDoItem, error) {
	var items entity.ToDoItem
	fmt.Println(items)
	result := r.db.Where("user_id = ?", userId).Create(&items)
	fmt.Println(&result)
	if result.RowsAffected == 0 {
		log.Fatal("No created")
	}
	fmt.Println(items)
	return items, nil
}

func (r *ToDoItemRepo) GetToDoItemsList(userId int) ([]entity.ToDoItem, error) {
	var items []entity.ToDoItem

	result := r.db.Where("user_id = ?", userId).Find(&items)

	if result.RowsAffected == 0 {
		log.Fatal("No select")
	}

	return items, nil
}

func (r *ToDoItemRepo) UpdateToDoItem(toDoItemId int) (entity.ToDoItem, error) {
	var items entity.ToDoItem

	result := r.db.Model(&items).Where("id=?", toDoItemId).Updates(items)
	if result.RowsAffected == 0 {
		log.Fatal("No update")
	}

	return items, nil
}

func (r *ToDoItemRepo) DeleteToDoItem(toDoItemId int) (int, error) {
	var items entity.ToDoItem

	result := r.db.Where("id= ?", toDoItemId).Delete(&items)
	if result.RowsAffected == 0 {
		log.Fatal("No delete")
	}

	return toDoItemId, nil
}
