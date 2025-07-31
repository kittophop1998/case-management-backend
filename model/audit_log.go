package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

const (
	EventCreated     string = "Created"
	EventUpdated     string = "Updated"
	EventSoftDeleted string = "Soft-Deleted"
	EventHardDeleted string = "Hard-Deleted"
)

type AuditLog struct {
	Model
	EventType   string         `json:"event_type"`
	ChangeTable string         `json:"change_table"`
	ChangeId    string         `json:"change_id"`
	Diff        datatypes.JSON `json:"diff" gorm:"diff"`
	UserID      uuid.UUID      `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
}

func (AuditLog) TableName() string {
	return "case_management_audit_log_info"
}
