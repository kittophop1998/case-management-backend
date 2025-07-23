package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(50)" json:"user_name" example:"john.doe"`
	Email    string `gorm:"type:varchar(100)" json:"email" example:"user@example.com"`
	Team     string `gorm:"type:varchar(50)" json:"team" example:"CEN123456"`
	IsActive *bool  `json:"is_active"`
	CenterID uint   `json:"center_id"` // foreign key
	Center   Center `gorm:"foreignKey:CenterID" json:"center"`
	RoleID   uint   `json:"role_id"` // foreign key
	Role     Role   `gorm:"foreignKey:RoleID" json:"role"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
}

type UserFilter struct {
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Sort     string
	IsActive *bool `json:"is_active"`
	Role     string
	Team     string
	Center   string
	RoleID   uint `json:"role_id"`
	TeamID   uint `json:"team_id"`
	CenterID uint `json:"center_id"`
}

type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)" json:"name"`
}

type Center struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)" json:"name"`
}

func (User) TableName() string {
	return "users"
}

func (Role) TableName() string {
	return "roles"
}

func (Center) TableName() string {
	return "centers"
}
