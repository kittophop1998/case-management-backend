package model

import "gorm.io/gorm"

func SeedRoles(db *gorm.DB) {
	roles := []Role{
		{Name: "Admin"},
		{Name: "Manager"},
		{Name: "Employee"},
		{Name: "Partner"},
		{Name: "Support"},
	}

	for _, role := range roles {
		db.FirstOrCreate(&role, Role{Name: role.Name})
	}
}
