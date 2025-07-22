package model

import "time"

// ผู้ใช้งานระบบ
type User struct {
	Id             string    `gorm:"primaryKey;type:varchar(50)" json:"id"`
	UserName       string    `gorm:"type:varchar(50)" json:"userName"`
	Email          string    `gorm:"type:varchar(100)" json:"email"`
	RoleId         string    `gorm:"type:varchar(50)" json:"roleId"`
	Team           string    `gorm:"type:varchar(50)" json:"team"`
	CenterId       string    `gorm:"type:varchar(50)" json:"centerId"`
	IsActive       string    `gorm:"type:varchar(50)" json:"isActive"`
	CreateDatetime time.Time `gorm:"type:timestamp" json:"createDatetime"`
	UpdateDatetime time.Time `gorm:"type:timestamp" json:"updateDatetime"`
}

func (User) TableName() string {
	return "users"
}
