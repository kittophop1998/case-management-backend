package main

import (
	"case-management/appcore"
	"case-management/appcore/appcore_cache"
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_logger"
	"case-management/appcore/appcore_store"
	"case-management/handler"
	"case-management/model"
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

	err := appcore_store.DBStore.AutoMigrate(&model.User{})

	if err != nil {
		panic(err.Error())
	}

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
