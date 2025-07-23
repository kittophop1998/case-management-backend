package model

import "time"

type User struct {
	Id             string    `gorm:"primaryKey;type:varchar(50)" json:"id" example:"1"`
	UserName       string    `gorm:"type:varchar(50)" json:"userName" example:"john.doe"`
	Email          string    `gorm:"type:varchar(100)" json:"email" example:"user@example.com"`
	RoleId         string    `gorm:"type:varchar(50)" json:"roleId" example:"ROLE_ADMIN"`
	Team           string    `gorm:"type:varchar(50)" json:"team" example:"CEN123456"`
	CenterId       string    `gorm:"type:varchar(50)" json:"centerId" example:"Team A"`
	IsActive       string    `gorm:"type:varchar(50)" json:"isActive" example:"true"`
	CreateDatetime time.Time `gorm:"type:timestamp" json:"createDatetime" example:"2025-07-17T09:00:00Z"`
	UpdateDatetime time.Time `gorm:"type:timestamp" json:"updateDatetime" example:"2025-07-17T09:00:00Z"`
}

func (User) TableName() string {
	return "users"
}
