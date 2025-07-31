package repository

import (
	"case-management/usecase"
	"log/slog"

	"github.com/minio/minio-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type authRepo struct {
	DB      *gorm.DB
	Cache   *redis.Client
	Logger  *slog.Logger
	Storage *minio.Client
}

func New(db *gorm.DB, logger *slog.Logger, cache *redis.Client, storage *minio.Client) usecase.CaseManagementRepository {
	return &authRepo{
		DB:      db,
		Cache:   cache,
		Logger:  logger,
		Storage: storage,
	}
}
