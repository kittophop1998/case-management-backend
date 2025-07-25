package model

import (
	"time"

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

type FormFilter struct {
	Limit  int                    `json:"limit"`
	Page   int                    `json:"page"`
	Sort   string                 `json:"sort"`
	Filter map[string]interface{} `json:"filter"`
}

type Model struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
