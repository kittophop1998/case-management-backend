package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Model
	AgentID      uint       `json:"agentId"`
	Username     string     `gorm:"type:varchar(50)" json:"username" example:"john.doe"`
	DomainName   string     `gorm:"type:varchar(50)" json:"domainName" example:"user"`
	Email        string     `gorm:"type:varchar(100)" json:"email" example:"user@example.com"`
	Name         string     `gorm:"type:varchar(100)" json:"name"`
	TeamID       uuid.UUID  `json:"teamId"`
	Team         Team       `gorm:"foreignKey:TeamID" json:"team"`
	IsActive     *bool      `json:"isActive" gorm:"default:true"`
	CenterID     uuid.UUID  `json:"centerId"`
	Center       Center     `gorm:"foreignKey:CenterID" json:"center"`
	RoleID       uuid.UUID  `json:"roleId"`
	Role         Role       `gorm:"foreignKey:RoleID" json:"role"`
	QueueID      uuid.UUID  `json:"queueId"`
	Queue        Queue      `gorm:"foreignKey:QueueID" json:"queue"`
	OperatorID   uint       `json:"operatorId"`
	Department   Department `gorm:"foreignKey:DepartmentID" json:"department"`
	DepartmentID uuid.UUID  `json:"departmentId"`
}

type UserFilter struct {
	Keyword  string    `json:"keyword"`
	Name     string    `gorm:"type:varchar(100)" json:"name"`
	Sort     string    `json:"sort"`
	IsActive *bool     `json:"isActive"`
	Role     string    `json:"role"`
	Team     Team      `json:"team"`
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

// Struct For Create User
type CreateUserRequest struct {
	AgentID      uint      `json:"agentId" validate:"required" example:"12337"`
	Username     string    `json:"username" validate:"required" example:"Janet Adebayo"`
	Email        string    `json:"email" validate:"required" example:"Janet@exam.com"`
	TeamID       uuid.UUID `json:"teamId" validate:"required" example:"b94eee08-8324-4d4f-b166-d82775553a7e"`
	OperatorID   uint      `json:"operatorId" validate:"required" example:"1233"`
	CenterID     uuid.UUID `json:"centerId" validate:"required" example:"b94eee08-8324-4d4f-b166-d82775553a7e"`
	RoleID       uuid.UUID `json:"roleId" validate:"required" example:"538cd6c5-4cb3-4463-b7d5-ac6645815476"`
	QueueID      uuid.UUID `json:"queueId" validate:"required" example:"b94eee08-8324-4d4f-b166-d82775553a7e"`
	DepartmentID uuid.UUID `json:"departmentId" validate:"required" example:"b94eee08-8324-4d4f-b166-d82775553a7e"`
	IsActive     bool      `json:"isActive" validate:"required" example:"true"`
}

type CreateUserResponse struct {
	Data struct {
		ID uuid.UUID `json:"id"`
	} `json:"data"`
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
