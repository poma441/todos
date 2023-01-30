package entity

type ToDoItem struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
	UserId int    `json:"userid"`
}
