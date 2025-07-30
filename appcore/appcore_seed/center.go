package appcore_seed

import (
	"case-management/model"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedCenter(db *gorm.DB) map[string]uuid.UUID {
	centerMaps := make(map[string]uuid.UUID)

	centers := []model.Center{
		{Name: "BKK"},
	}

	for _, center := range centers {
		var r model.Center

		err := db.Where("name = ?", center.Name).FirstOrCreate(&r, center).Error
		if err != nil {
			log.Printf("failed to seed center %s: %v", center.Name, err)
		} else {
			centerMaps[center.Name] = r.ID
		}
	}

	return centerMaps
}
