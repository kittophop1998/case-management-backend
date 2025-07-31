package appcore_seed

import (
	"case-management/model"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedTeam(db *gorm.DB) map[string]uuid.UUID {
	teamMaps := make(map[string]uuid.UUID)

	teams := []model.Team{
		{Name: "Inbound"},
		{Name: "CHD"},
		{Name: "CHB"},
		{Name: "Convince"},
		{Name: "Tele"},
		{Name: "JSJ"},
		{Name: "EDP"},
	}

	for _, team := range teams {
		var r model.Team

		err := db.Where("name = ?", team.Name).FirstOrCreate(&r, team).Error
		if err != nil {
			log.Printf("failed to seed team %s: %v", team.Name, err)
		} else {
			teamMaps[team.Name] = r.ID
		}
	}

	return teamMaps
}
