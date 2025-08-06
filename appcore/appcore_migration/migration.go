package appcore_migration

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
		&model.Team{},
		&model.Cases{},
		&model.Center{},
		&model.Department{},
		&model.NoteTypes{},
		&model.CaseTypes{},
		&model.CaseStatus{},
		&model.CasePriorities{},
		&model.CaseNotes{},
		&model.Attachment{},
		&model.AuditLog{},
		&model.ApiLogs{},
		&model.CustomerNote{},
	); err != nil {
		return err
	}

	return nil // Migration successful
}
