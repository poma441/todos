package entity

type User struct {
	Id       int    `json:"-"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
