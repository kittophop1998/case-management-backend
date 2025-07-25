package model

import (
	"encoding/json"
	"time"

	"encore.dev/types/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type User struct {
	Model
	UserName string `gorm:"type:varchar(50)" json:"userName" example:"john.doe"`
	Email    string `gorm:"type:varchar(100)" json:"email" example:"user@example.com"`
	Team     string `gorm:"type:varchar(50)" json:"team" example:"CEN123456"`
	IsActive *bool  `json:"isActive"`
	CenterID uint   `json:"-"` // ไม่แสดงใน response
	Center   Center `gorm:"foreignKey:CenterID" json:"center"`
	RoleID   uint   `json:"-"` // ไม่แสดงใน response
	Role     Role   `gorm:"foreignKey:RoleID" json:"role"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
}

type UserFilter struct {
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Sort     string `json:"sort"`
	IsActive *bool  `json:"isActive"`
	Role     string `json:"role"`
	Team     string `json:"team"`
	Center   string `json:"center"`
	RoleID   uint   `json:"roleId"`
	TeamID   uint   `json:"teamId"`
	CenterID uint   `json:"centerId"`
}

type Role struct {
	Model
	Name string `gorm:"type:varchar(100)" json:"name"`
}

type Center struct {
	Model
	Name string `gorm:"type:varchar(100)" json:"name"`
}

type AccessLogs struct {
	ID            uuid.UUID       `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID        uuid.UUID       `json:"userId" gorm:"type:uuid;default:uuid_generate_v4()"`
	Action        string          `json:"action"`
	IPAddress     string          `json:"ipAddress"`
	UserAgent     string          `json:"userAgent" gorm:"type:text"`
	Details       json.RawMessage `gorm:"type:jsonb" json:"details"`
	CreatedAt     time.Time       `json:"createdAt"`
	Username      string          `gorm:"primaryKey;type:varchar(20)" json:"username"`
	LogonDatetime time.Time       `gorm:"type:timestamp" json:"logonDatetime"`
	LogonResult   string          `gorm:"type:varchar(10)" json:"logonResult"`
}

type UserResponse struct {
	Username   string             `json:"username"`
	UserMetrix UserMetrixResponse `json:"userMetrix"`
}

type UserMetrixResponse struct {
	Role        string `json:"role"`
	Create      bool   `json:"create"`
	Update      bool   `json:"update"`
	Delete      bool   `json:"delete"`
	CreateEvent bool   `json:"createEvent"`
	UpdateEvent bool   `json:"updateEvent"`
	DeleteEvent bool   `json:"deleteEvent"`
}

type UserMetrix struct {
	Role           string    `gorm:"primaryKey;type:varchar(20)" json:"role"`
	Create         bool      `json:"create"`
	Update         bool      `json:"update"`
	Delete         bool      `json:"delete"`
	CreateEvent    bool      `json:"createEvent"`
	UpdateEvent    bool      `json:"updateEvent"`
	DeleteEvent    bool      `json:"deleteEvent"`
	CreateDatetime time.Time `gorm:"type:timestamp" json:"createDatetime"`
	CreateBy       string    `gorm:"type:varchar(20)" json:"createBy"`
	UpdateDatetime time.Time `gorm:"type:timestamp" json:"updateDatetime"`
	UpdateBy       string    `gorm:"type:varchar(20)" json:"updateBy"`
	DeleteDatetime time.Time `gorm:"type:timestamp" json:"deleteDatetime"`
	DeleteBy       string    `gorm:"type:varchar(20)" json:"deleteBy"`
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

func (UserMetrix) TableName() string {
	return "user_metrixes"
}
