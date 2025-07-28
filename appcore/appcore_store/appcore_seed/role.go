package appcoreseed

import (
	"case-management/model"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB) map[string]uuid.UUID {
	roleMaps := make(map[string]uuid.UUID)

	roles := []model.Role{
		{Name: "Admin"},
		{Name: "Case Manager"},
		{Name: "Supervisor"},
		{Name: "Read-Only-User"},
	}

	for _, role := range roles {
		if role.Name == "" {
			log.Println("skipping empty role name")
			continue
		}

		err := db.Where("name = ?", role.Name).FirstOrCreate(&role).Error
		if err != nil {
			log.Printf("failed to seed role %s: %v", role.Name, err)
		} else {
			roleMaps[role.Name] = role.ID
		}
	}

	return roleMaps
}
