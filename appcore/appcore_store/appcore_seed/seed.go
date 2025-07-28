package appcoreseed

import (
	"log"

	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	// Seed roles
	roleMap := SeedRole(db)
	// SeedRole(db)

	// Seed centers
	centerMap := SeedCenter(db)

	// Seed permissions
	SeedPermission(db)

	// Seed users
	SeedUser(db, roleMap, centerMap)

	// Seed role_permissions
	if err := SeedRolePermission(db); err != nil {
		log.Fatalf("Failed to seed role_permissions: %v", err)
	}
}
