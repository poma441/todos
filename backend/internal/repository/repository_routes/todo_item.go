package repository

type ToDoItemRepo struct {
}

func NewToDoItemRepo() *ToDoItemRepo {
	return &ToDoItemRepo{}
}

func (r *ToDoItemRepo) PlugFunc() {

}
