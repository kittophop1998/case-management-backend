package main

import (
	"case-management/appcore/appcore_cache"
	"case-management/appcore/appcore_logger"
	"case-management/appcore/appcore_store"
	"case-management/handler"
	"case-management/model"
	"case-management/repository"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	appcore_store.InitPostgresDBStore(logger)
	db := appcore_store.DBStore
	if db == nil {
		logger.Error("DB is nil, cannot continue")
		return
	}

	// Auto-migrate
	if err := db.AutoMigrate(&model.User{}); err != nil {
		logger.Error("Auto-migrate failed", slog.Any("error", err))
		return
	}

	// สร้าง repository
	authRepo := repository.New(appcore_store.DBStore, appcore_logger.Logger, appcore_cache.Cache)

	// สร้าง handler
	userHandler := handler.NewUserHandler(authRepo)

	// Setup Gin และ Route
	router := gin.Default()
	handler.SetupRoutes(router, userHandler)

	// Run
	router.Run(":8000")
}
