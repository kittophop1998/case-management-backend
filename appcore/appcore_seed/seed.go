package appcore_seed

import (
	"log"

	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {

	// Seed roles
	roleMap := SeedRole(db)

	// Seed centers
	centerMap := SeedCenter(db)

	// Seed teams
	teamMap := SeedTeam(db)

	// Seed Department
	departmentMap := SeedDepartment(db)

	// Seed permissions
	SeedPermission(db)

	// Seed users
	SeedUser(db, roleMap, teamMap, centerMap, departmentMap)

	// Seed role_permissions
	if err := SeedRolePermission(db); err != nil {
		log.Fatalf("Failed to seed role_permissions: %v", err)
	}
}
