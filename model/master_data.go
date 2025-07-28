package model

import "github.com/google/uuid"

type Permission struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string    `gorm:"uniqueIndex;not null"`
}

type Role struct {
	ID          uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string       `gorm:"uniqueIndex;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

type RolePermission struct {
	RoleID       uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	PermissionID uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
}

type Center struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string    `gorm:"type:varchar(100)" json:"name"`
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

func (Center) TableName() string {
	return "centers"
}
