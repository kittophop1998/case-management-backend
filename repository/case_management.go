package repository

import (
	"time"

	"gorm.io/gorm"
)

type CaseManagement struct {
	gorm.Model
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	CreatedBy string    `json:"created_by"`
	SlaAt     time.Time `json:"sla_at"`
	Subject   string    `json:"subject"`
	Reason    string    `json:"reason"`
}

func (CaseManagement) TableName() string {
	return "case_management"
}
