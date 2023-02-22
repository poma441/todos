package entity

type Student struct {
	Id       int    `json:"-" gorm:"primaryKey"`
	Role     string `json:"role"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
}

type User struct {
	Id       int    `json:"-"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
