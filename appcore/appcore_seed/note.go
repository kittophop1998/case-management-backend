package appcore_seed

import (
	"case-management/model"
	"log"

	"gorm.io/gorm"
)

func SeedNote(db *gorm.DB) error {
	notes := []model.NoteTypes{
		{
			Name:        "Internal Note",
			Description: "This is an internal note type",
		},
	}

	for _, note := range notes {
		var r model.NoteTypes

		err := db.Where("name = ?", note.Name).Assign().FirstOrCreate(&r, note).Error
		if err != nil {
			log.Printf("failed to seed note %s: %v", note.Name, err)
		}
	}

	return nil
}
