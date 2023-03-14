package models

type User struct {
	Id       int    `json:"-" gorm:"primaryKey"`
	Uuid     string `json:"-" gorm:"unique"`
	Role     string `json:"role"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
}
