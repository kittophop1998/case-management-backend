package appcore_store

import (
	"case-management/appcore/appcore_config"
	"context"
	"log"
	"log/slog"
	"strconv"
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

	var dsn string
	if appcore_config.Config.Mode == "development" {
		dsn = appcore_config.Config.PostgresRailwayURL
	} else {
		dsn = "host=" + appcore_config.Config.PostgresHost +
			" port=" + strconv.Itoa(appcore_config.Config.PostgresPort) +
			" user=" + appcore_config.Config.PostgresUser +
			" password=" + appcore_config.Config.PostgresPassword +
			" dbname=" + appcore_config.Config.PostgresDBName +
			" sslmode=" + appcore_config.Config.PostgresSSLMode
	}

	// ✅ Parse the config using pgx/v5
	log.Printf("Postgres DSN: %s", dsn)
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
