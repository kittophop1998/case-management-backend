package appcore_seed

import (
	"case-management/model"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedDepartment(db *gorm.DB) map[string]uuid.UUID {
	departmentMaps := make(map[string]uuid.UUID)

	departments := []model.Department{
		{Name: "Marketing"},
	}

	for _, department := range departments {
		var r model.Department

		err := db.Where("name = ?", department.Name).FirstOrCreate(&r, department).Error
		if err != nil {
			log.Printf("failed to seed department %s: %v", department.Name, err)
		} else {
			departmentMaps[department.Name] = r.ID
		}
	}

	return departmentMaps
}
