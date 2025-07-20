package appcore_handler

import (
	"case-management/appcore/appcore_cache"
	inf "case-management/appcore/appcore_internal/appcore_interface"
	"case-management/appcore/appcore_store"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

type ApiHandler struct {
	ServiceName string
	Version     string
	Module      inf.Module

	DB    *gorm.DB
	Cache *redis.Client
}

func NewHandler(serviceName, version string, module inf.Module) *ApiHandler {
	return &ApiHandler{
		ServiceName: serviceName,
		Version:     version,
		Module:      module,

		DB:    appcore_store.DBStore,
		Cache: appcore_cache.Cache,
	}
}

func (h *ApiHandler) HealthCheck(c *gin.Context) {
	ctx := context.Background()

	isError := false
	errorMessage := ""

	defer func() {
		if isError {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"service": h.ServiceName,
				"version": h.Version,
				"error":   errorMessage,
			})
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"service": h.ServiceName,
				"message": "pong",
				"version": h.Version,
			})
		}
	}()

	if h.DB != nil {
		sql, err := h.DB.DB()
		if err != nil {
			isError = true
			errorMessage = "cannot ping to database service"
			return
		}
		err = sql.Ping()
		if err != nil {
			isError = true
			errorMessage = "cannot ping to database service"
			return
		}
	}

	if h.Cache != nil {
		status := h.Cache.Ping(ctx)
		if status.Err() != nil {
			isError = true
			errorMessage = "cannot ping to cache service"
			return
		}
	}

}
