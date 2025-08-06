package appcore_handler

import (
	"case-management/appcore/appcore_config"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIErrorResponse struct {
	Errors ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	StatusCode int                    `json:"statusCode"`
	Error      string                 `json:"error"`
	Message    appcore_config.Message `json:"message"`
	Timestamp  string                 `json:"timestamp"`
	Path       string                 `json:"path"`
	Details    interface{}            `json:"details,omitempty"`
}

func HandleError(c *gin.Context, err error) {
	// ตรวจสอบว่าเป็น AppError หรือไม่
	var appErr *appcore_config.AppError
	if errors.As(err, &appErr) {
		renderAppError(c, appErr)
		return
	}

	// ตรวจสอบว่าเป็น gorm.ErrRecordNotFound หรือไม่
	if errors.Is(err, gorm.ErrRecordNotFound) {
		renderAppError(c, appcore_config.ErrNotFound)
		return
	}

	// fallback: internal server error
	renderAppError(c, appcore_config.ErrInternalServer)
}

// ✅ แยกการสร้าง response
func renderAppError(c *gin.Context, appErr *appcore_config.AppError) {
	c.AbortWithStatusJSON(appErr.HTTPStatus, APIErrorResponse{
		Errors: ErrorDetail{
			StatusCode: appErr.HTTPStatus,
			Error:      appErr.Code,
			Message:    appErr.Message,
			Timestamp:  time.Now().UTC().Format(time.RFC3339Nano),
			Path:       c.FullPath(),
			Details:    appErr.Details,
		},
	})
}
