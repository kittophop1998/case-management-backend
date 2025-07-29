package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusResponse struct {
	Status string `json:"status"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type AccessTokenRequest struct {
	Access_token string `json:"access_token" binding:"required"`
}

type UpdatePermissionRolesRequest struct {
	Permission string   `json:"permission" binding:"required"`
	Roles      []string `json:"roles" binding:"required"`
}

type PermissionWithRolesResponse struct {
	Permission string   `json:"permission"`
	Label      string   `json:"label"`
	Roles      []string `json:"roles"`
}
type FormFilter struct {
	Limit  int                    `json:"limit"`
	Page   int                    `json:"page"`
	Sort   string                 `json:"sort"`
	Filter map[string]interface{} `json:"filter"`
}

type Model struct {
	ID        uuid.UUID      `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" swaggerignore:"true"`
}
