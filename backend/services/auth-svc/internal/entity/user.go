package entity

type User struct {
	Id       int    `json:"-"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
