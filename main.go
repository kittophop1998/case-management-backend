// @title           Case Management API
// @version         1.0
// @description     This is a sample server for managing cases.
// @termsOfService  http://swagger.io/terms/

// @contact.name   SYE
// @contact.url    https://aeon.co.th
// @contact.email  sye@aeon.co.th

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8000
// @BasePath        /api/v1

package main

import (
	"case-management/appcore"
	"case-management/appcore/appcore_cache"
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_logger"
	appcore_migation "case-management/appcore/appcore_migration"
	appcoreseed "case-management/appcore/appcore_seed"
	"case-management/appcore/appcore_store"
	"case-management/handler"
	"case-management/repository"
	"case-management/usecase"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	loc, _ := time.LoadLocation("UTC")
	// handle err
	time.Local = loc // -> this is setting the global timezone

	serviceName := "caseManagement"
	version := "v1.0.0"

	// output current time zone
	timeZone, offset := time.Now().Zone()
	slog.Info("service time zone", "timeZone", timeZone, "offset", offset)
	slog.Info("currentTime", "time", time.Now())

	// Initialize Migration
	if err := appcore_migation.Migrate(); err != nil {
		slog.Error("migration failed", "error", err)
		os.Exit(1)
	}

	// Seeder
	appcoreseed.SeedAll(appcore_store.DBStore)

	caseManagementRepo := repository.New(appcore_store.DBStore, appcore_logger.Logger, appcore_cache.Cache)
	caseManagementUseCase := usecase.New(caseManagementRepo, appcore_cache.Cache, appcore_logger.Logger)
	caseManagementHandler := handler.NewHandler(caseManagementUseCase, appcore_logger.Logger)

	// สร้าง handler
	appCoreHandler := appcore_handler.NewHandler(serviceName, version, caseManagementHandler)

	service := appcore.NewService(serviceName, version, appCoreHandler)
	service.Run()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	sig := <-c
	appcore_logger.Logger.Info("shutting down the server", "received signal", sig)
	appcore_logger.Logger.Info("shutting down gracefully, press Ctrl+C again to force")
	service.Stop()
}
