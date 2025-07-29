package appcore_migation

import (
	"case-management/appcore/appcore_store"
	"case-management/model"
)

func Migrate() error {
	// Ensure the uuid-ossp extension is enabled for UUID generation
	appcore_store.DBStore.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	// Migrate the model
	if err := appcore_store.DBStore.AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.AccessLogs{},
		&model.Permission{},
		&model.RolePermission{},
	); err != nil {
		return err
	}

	return nil // Migration successful
}
