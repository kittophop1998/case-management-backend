package model

import (
	"encoding/json"
	"time"

	"encore.dev/types/uuid"
	"gorm.io/gorm"
)

type Cases struct {
	gorm.Model
	Title               string    `json:"title"`
	CustomerId          string    `json:"customer_id"`
	CreditCardAccountId string    `json:"credit_card_account_id"`
	LoanAccountId       string    `json:"loan_account_id"`
	AssignedToUserId    uuid.UUID `json:"assigned_to_user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	StatusId            uint      `json:"status_id"`
	PriorityId          uint      `json:"priority_id"`
	StartDate           time.Time `json:"start_date" gorm:"type:date"`
	EndDate             time.Time `json:"end_date" gorm:"type:date"`
	InitialDescription  string    `json:"initial_description" gorm:"type:text"`
	Resolution          string    `json:"resolution" gorm:"type:text"`
	CreatedBy           string    `json:"created_by"`
	SLADate             time.Time `json:"sla_date"`
}

type NoteTypes struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type CaseTypes struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description" gorm:"type:text"`
}

type CaseStatus struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	Description    string `json:"description" gorm:"type:text"`
	IsClosedStatus bool   `json:"is_closed_status"`
}

type CasePriorities struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	OrderNumber uint   `json:"order_number"`
}

type CaseNotes struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CaseId      uuid.UUID `json:"case_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserId      uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	NoteTypesId uint      `json:"note_types_id"`
	Content     string    `json:"content" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
}

type Attachment struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CaseId           uuid.UUID `json:"case_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	FileName         string    `json:"file_name"`
	FilePath         string    `json:"file_path"`
	FileType         string    `json:"file_type"`
	FileSizeBytes    uint64    `json:"file_size_bytes" gorm:"type:bigint"`
	UploadedByUserId uuid.UUID `json:"uploaded_by_user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UploadedAt       time.Time `json:"uploaded_at"`
}

type ApiLogs struct {
	ID              uuid.UUID       `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserId          uuid.UUID       `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Endpoint        string          `json:"endpoint"`
	Method          string          `json:"method"`
	RequestPayload  json.RawMessage `gorm:"type:jsonb" json:"request_payload"`
	ResponsePayload json.RawMessage `gorm:"type:jsonb" json:"response_payload"`
	StatusCode      uint            `json:"status_code"`
	DurationMs      uint            `json:"duration_ms"`
	ErrorMessage    string          `json:"error_message" gorm:"type:text"`
	CreatedAt       time.Time       `json:"created_at"`
}

type VerifyQuestionHistory struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CustomerId       string    `json:"customer_id"`
	Question         string    `json:"question" gorm:"type:text"`
	AnswerProvided   string    `json:"answer_provided" gorm:"type:text"`
	IsCorrect        bool      `json:"is_correct"`
	VeryfyByUserId   uuid.UUID `json:"verify_by_user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	VerificationDate time.Time `json:"verification_date"`
	CaseId           uuid.UUID `json:"case_id" gorm:"type:uuid;default:uuid_generate_v4()"`
}

func (Cases) TableName() string {
	return "cases"
}
