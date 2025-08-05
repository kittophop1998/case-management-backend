package appcore_seed

import (
	"case-management/model"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedCaseNotes(db *gorm.DB) error {
	var caseIDs []uuid.UUID
	if err := db.Model(&model.Cases{}).Pluck("id", &caseIDs).Error; err != nil {
		return fmt.Errorf("failed to fetch case IDs: %w", err)
	}

	var userIDs []uuid.UUID
	if err := db.Model(&model.User{}).Pluck("id", &userIDs).Error; err != nil {
		return fmt.Errorf("failed to fetch user IDs: %w", err)
	}

	var noteTypeIDs []uuid.UUID
	if err := db.Model(&model.NoteTypes{}).Pluck("id", &noteTypeIDs).Error; err != nil {
		return fmt.Errorf("failed to fetch note type IDs: %w", err)
	}

	if len(caseIDs) == 0 || len(userIDs) == 0 || len(noteTypeIDs) == 0 {
		return fmt.Errorf("not enough data in related tables to seed case notes")
	}

	caseNotes := []model.CaseNotes{
		{
			CaseId:      caseIDs[0],
			UserId:      userIDs[0],
			NoteTypesId: noteTypeIDs[0],
			Content:     "ลูกค้าสอบถามสถานะการชำระเงิน",
		},
		{
			CaseId:      caseIDs[0],
			UserId:      userIDs[1%len(userIDs)],
			NoteTypesId: noteTypeIDs[1%len(noteTypeIDs)],
			Content:     "เจ้าหน้าที่โทรแจ้งยอดล่าสุดให้ลูกค้า",
		},
		{
			CaseId:      caseIDs[1%len(caseIDs)],
			UserId:      userIDs[0],
			NoteTypesId: noteTypeIDs[2%len(noteTypeIDs)],
			Content:     "ลูกค้าแจ้งว่าจะโอนภายในสัปดาห์นี้",
		},
	}

	for _, cn := range caseNotes {
		var existing model.CaseNotes
		err := db.Where("case_id = ? AND user_id = ? AND note_types_id = ? AND content = ?",
			cn.CaseId, cn.UserId, cn.NoteTypesId, cn.Content).
			First(&existing).Error

		if err == gorm.ErrRecordNotFound {
			cn.CreatedAt = time.Now()
			if err := db.Create(&cn).Error; err != nil {
				return fmt.Errorf("failed to seed case note: %w", err)
			}
		} else if err != nil {
			return fmt.Errorf("failed to check existing note: %w", err)
		}
	}

	return nil
}
