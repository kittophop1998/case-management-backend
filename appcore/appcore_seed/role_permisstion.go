package appcore_seed

import (
	"case-management/model"

	"gorm.io/gorm"
)

func SeedRolePermission(db *gorm.DB) error {
	var admin model.Role
	if err := db.Preload("Permissions").Where("name = ?", "Admin").First(&admin).Error; err != nil {
		return err
	}

	var perms []model.Permission
	if err := db.Where("name IN ?", []string{"case.create", "case.view", "case.edit", "case.delete", "user.manage"}).Find(&perms).Error; err != nil {
		return err
	}

	return db.Model(&admin).Association("Permissions").Replace(perms)
}
