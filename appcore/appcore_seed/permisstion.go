package appcore_seed

import (
	"case-management/model"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedPermission(db *gorm.DB) map[string]uuid.UUID {
	permissionMap := make(map[string]uuid.UUID)

	permissions := []model.Permission{
		{Name: "case.create"},
		{Name: "case.view"},
		{Name: "case.edit"},
		{Name: "case.delete"},
		{Name: "user.manage"},
	}

	for _, permission := range permissions {
		err := db.Where("name = ?", permission.Name).FirstOrCreate(&permission).Error
		if err != nil {
			// Log the error if seeding fails
			log.Printf("failed to seed permission %s: %v", permission.Name, err)
		} else {
			permissionMap[permission.Name] = permission.ID
		}
	}

	return permissionMap
}
