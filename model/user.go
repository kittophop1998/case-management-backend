package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Model
	AgentID    uint      `json:"agentId"`
	Username   string    `gorm:"type:varchar(50)" json:"username" example:"john.doe"`
	Email      string    `gorm:"type:varchar(100)" json:"email" example:"user@example.com"`
	Name       string    `gorm:"type:varchar(100)" json:"name"`
	Team       string    `gorm:"type:varchar(50)" json:"team" example:"CEN123456"`
	IsActive   *bool     `json:"isActive"`
	CenterID   uuid.UUID `json:"centerId"`
	Center     Center    `gorm:"foreignKey:CenterID" json:"center"`
	RoleID     uuid.UUID `json:"roleId"`
	Role       Role      `gorm:"foreignKey:RoleID" json:"role"`
	OperatorID uint      `json:"operatorId"`
}

type UserFilter struct {
	Name     string    `gorm:"type:varchar(100)" json:"name"`
	Sort     string    `json:"sort"`
	IsActive *bool     `json:"isActive"`
	Role     string    `json:"role"`
	Team     string    `json:"team"`
	Center   string    `json:"center"`
	RoleID   uuid.UUID `json:"roleID,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	TeamID   uuid.UUID `json:"teamID,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	CenterID uuid.UUID `json:"centerID,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
}

type AccessLogs struct {
	ID            uuid.UUID       `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID        uuid.UUID       `json:"userId" gorm:"type:uuid;default:uuid_generate_v4()"`
	Action        string          `json:"action"`
	IPAddress     string          `json:"ipAddress"`
	UserAgent     string          `json:"userAgent" gorm:"type:text"`
	Details       json.RawMessage `gorm:"type:jsonb" json:"details"`
	CreatedAt     time.Time       `json:"createdAt"`
	Username      string          `gorm:"type:varchar(20)" json:"username"`
	LogonDatetime time.Time       `gorm:"type:timestamp" json:"logonDatetime"`
	LogonResult   string          `gorm:"type:varchar(10)" json:"logonResult"`
}

type UserResponse struct {
	Username string    `json:"username"`
	RoleId   uuid.UUID `json:"roleId"`
	Role     Role      `json:"role"`
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

type ImportStatus struct {
	Progress   int      `json:"progress"`
	Errors     []string `json:"errors"`
	Total      int      `json:"total"`
	Successful int      `json:"successful"`
}

func (User) TableName() string {
	return "users"
}

func (UserMetrix) TableName() string {
	return "user_metrixes"
}

func (AccessLogs) TableName() string {
	return "access_logs"
}
