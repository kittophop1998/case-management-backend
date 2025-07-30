package appcore_store

import (
	"case-management/appcore/appcore_config"
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBStore *gorm.DB

func InitPostgresDBStore(logger *slog.Logger) {
	logger.Info("Init DB Store")
	logger.Info("Connecting to database")

	// dsn := appcore_config.Config.ProdPostgresURL
	dsn := appcore_config.Config.PostgresConnString

	// ✅ Parse the config using pgx/v5
	cfg, err := pgx.ParseConfig(dsn)
	if err != nil {
		panic("failed to parse DSN: " + err.Error())
	}

	// ✅ Convert to stdlib-compatible DB
	sqlDB := stdlib.OpenDB(*cfg)

	// ✅ Use GORM with pgx connection
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	DBStore = db

	// DB tuning
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.PingContext(context.Background()); err != nil {
		panic("database ping failed: " + err.Error())
	}

	logger.Info("Connecting to database success")
}
