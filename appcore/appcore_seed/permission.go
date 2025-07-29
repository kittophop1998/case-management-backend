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
		{Key: "case.create", Name: "Case Create"},
		{Key: "case.view", Name: "Case View"},
		{Key: "case.edit", Name: "Case Edit"},
		{Key: "case.delete", Name: "Case Delete"},
		{Key: "user.manage", Name: "User Management"},
	}

	for _, p := range permissions {
		var permission model.Permission
		err := db.Where("key = ?", p.Key).FirstOrCreate(&permission, p).Error
		if err != nil {
			log.Printf("failed to seed permission %s: %v", p.Name, err)
		} else {
			permissionMap[p.Key] = permission.ID
		}
	}

	return permissionMap
}
