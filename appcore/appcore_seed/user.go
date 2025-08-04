package appcore_seed

import (
	"case-management/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB, roleMap, teamMap, centerMap, departmentMap map[string]uuid.UUID) {
	isActive := true
	users := []model.User{
		{
			Username:     "admin",
			TeamID:       teamMap["Inbound"],
			CenterID:     centerMap["BKK"],
			RoleID:       roleMap["Admin"],
			AgentID:      1,
			IsActive:     &isActive,
			Name:         "admin",
			Email:        "admin@admin.com",
			OperatorID:   1,
			DepartmentID: departmentMap["Marketing"],
		},
	}

	for _, user := range users {
		var existingUser model.User
		if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// User does not exist, create it
				if err := db.Create(&user).Error; err != nil {
					panic("Failed to seed user: " + err.Error())
				}
			} else {
				panic("Failed to check existing user: " + err.Error())
			}
		}
	}
}
