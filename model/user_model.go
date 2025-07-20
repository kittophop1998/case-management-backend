package model

type User struct {
	Id        uint   `json:"id"`
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string `json:"-"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

func (User) TableName() string {
	return "users"
}
