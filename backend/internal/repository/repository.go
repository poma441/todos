package repository

type Authorization interface {
}

type ToDoItem interface {
}

type Repository struct {
	Authorization
	ToDoItem
}

func NewRepository() *Repository {
	return &Repository{
		Authorization: NewAuthRepo(),
		ToDoItem:      NewToDoItemRepo(),
	}
}
