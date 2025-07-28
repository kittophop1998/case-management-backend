package appcoreseed

import (
	"case-management/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB, roleMap, centerMap map[string]uuid.UUID) {
	users := []model.User{
		{
			UserName: "admin",
			Team:     "BKK",
			CenterID: centerMap["BKK"],
			RoleID:   roleMap["Admin"],
		},
	}

	for _, user := range users {
		var existingUser model.User
		if err := db.Where("user_name = ?", user.UserName).First(&existingUser).Error; err != nil {
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
