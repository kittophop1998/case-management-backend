package usecase

import (
	"case-management/model"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (u *UseCase) SaveLog(userId uuid.UUID, method, endpoint string, req, res []byte, status, duration uint, errMsg string) error {
	log := model.ApiLogs{
		UserId:          userId,
		Method:          method,
		Endpoint:        endpoint,
		RequestPayload:  json.RawMessage(req),
		ResponsePayload: json.RawMessage(res),
		StatusCode:      status,
		DurationMs:      duration,
		ErrorMessage:    errMsg,
		CreatedAt:       time.Now(),
	}
	return u.caseManagementRepository.SaveLog(&log)
}

func (u *UseCase) GetLogs(c *gin.Context) ([]model.ApiLogs, error) {
	return u.caseManagementRepository.GetAllLogs(c)
}
