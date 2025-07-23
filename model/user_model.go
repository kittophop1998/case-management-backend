package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(50)" json:"userName" example:"john.doe"`
	Email    string `gorm:"type:varchar(100)" json:"email" example:"user@example.com"`
	RoleId   string `gorm:"type:varchar(50)" json:"roleId" example:"ROLE_ADMIN"`
	Team     string `gorm:"type:varchar(50)" json:"team" example:"CEN123456"`
	CenterId string `gorm:"type:varchar(50)" json:"centerId" example:"Team A"`
	IsActive string `gorm:"type:varchar(50)" json:"isActive" example:"true"`
}

func (User) TableName() string {
	return "users"
}
