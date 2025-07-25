package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) error {
	roles := []Role{
		{
			Model: Model{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Admin",
		},
		{
			Model: Model{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "User",
		},
	}

	for _, r := range roles {
		if err := db.FirstOrCreate(&Role{}, Role{Name: r.Name}).Error; err != nil {
			return err
		}
	}
	return nil
}

func SeedCenters(db *gorm.DB) error {
	centers := []Center{
		{
			Model: Model{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Bangkok Center",
		},
		{
			Model: Model{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Chiang Mai Center",
		},
	}

	for _, c := range centers {
		if err := db.FirstOrCreate(&Center{}, Center{Name: c.Name}).Error; err != nil {
			return err
		}
	}
	return nil
}

func SeedUsers(db *gorm.DB) error {
	var adminRole Role
	if err := db.Where("name = ?", "Admin").First(&adminRole).Error; err != nil {
		return err
	}

	var userRole Role
	if err := db.Where("name = ?", "User").First(&userRole).Error; err != nil {
		return err
	}

	var bangkokCenter Center
	if err := db.Where("name = ?", "Bangkok Center").First(&bangkokCenter).Error; err != nil {
		return err
	}

	isActiveTrue := true
	isActiveFalse := false

	users := []User{
		{
			Model: Model{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			UserName: "admin01",
			Email:    "admin01@example.com",
			Team:     "CEN123456",
			IsActive: &isActiveTrue,
			CenterID: bangkokCenter.ID,
			RoleID:   adminRole.ID,
			Name:     "Admin User",
		},
		{
			Model: Model{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			UserName: "user01",
			Email:    "user01@example.com",
			Team:     "CEN123456",
			IsActive: &isActiveFalse,
			CenterID: bangkokCenter.ID,
			RoleID:   userRole.ID,
			Name:     "Normal User",
		},
	}

	for _, u := range users {
		if err := db.Where("user_name = ?", u.UserName).FirstOrCreate(&u).Error; err != nil {
			return err
		}
	}

	return nil
}
