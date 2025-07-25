package model

import "github.com/google/uuid"

type Permission struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string
}

type Role struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

type RolePermission struct {
	RoleID       uuid.UUID `gorm:"type:uuid;not null"`
	PermissionID uuid.UUID `gorm:"type:uuid;not null"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}

func (Permission) TableName() string {
	return "permissions"
}

func (Role) TableName() string {
	return "roles"
}
