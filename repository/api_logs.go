package repository

import (
	"case-management/model"

	"github.com/gin-gonic/gin"
)

func (r *authRepo) SaveLog(log *model.ApiLogs) error {
	return r.DB.Create(log).Error
}

func (r *authRepo) GetAllLogs(c *gin.Context) ([]model.ApiLogs, error) {
	var logs []model.ApiLogs
	err := r.DB.WithContext(c.Request.Context()).Order("created_at desc").Find(&logs).Error
	return logs, err
}
