package repository

import (
	"case-management/model"

	"github.com/gin-gonic/gin"
)

func (a *authRepo) CreateAuditLog(c *gin.Context, log model.AuditLog) error {
	return a.DB.WithContext(c.Request.Context()).Create(&log).Error
}
