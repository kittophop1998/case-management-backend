package appcore_handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIErrorResponse struct {
	Errors ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	StatusCode int         `json:"statusCode"`
	Error      string      `json:"error"`
	Message    string      `json:"message"`
	Timestamp  string      `json:"timestamp"`
	Path       string      `json:"path"`
	Details    interface{} `json:"details,omitempty"`
}

type AppError struct {
	Code       string
	Message    string
	HTTPStatus int
	Details    interface{}
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code, message string, status int, details interface{}) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: status,
		Details:    details,
	}
}

// 🔁 Common errors
var (
	ErrBadRequest         = NewAppError("BAD_REQUEST", "ข้อมูลที่ส่งมาไม่ถูกต้องตามที่ระบบรองรับ", http.StatusBadRequest, nil)
	ErrRequiredParam      = NewAppError("REQUIRED_PARAMETER", "ไม่สามารถดึงข้อมูลได้ เนื่องจากระบบขาดข้อมูลบางส่วน", http.StatusBadRequest, nil)
	ErrFilterRequired     = NewAppError("FILTER_REQUIRED", "โปรดเลือกอย่างน้อย 1 เงื่อนไขเพื่อค้นหา", http.StatusBadRequest, nil)
	ErrNotFound           = NewAppError("NOT_FOUND", "ไม่พบข้อมูลที่ค้นหา", http.StatusNotFound, nil)
	ErrInternalServer     = NewAppError("INTERNAL_SERVER_ERROR", "เกิดข้อผิดพลาดในระบบ กรุณาลองใหม่ภายหลัง", http.StatusInternalServerError, nil)
	ErrServiceUnavailable = NewAppError("SERVICE_UNAVAILABLE", "The service is temporarily unavailable or in maintenance", http.StatusServiceUnavailable, nil)
	ErrGatewayTimeout     = NewAppError("NO_RESPONSE", "No response from an upstream service", http.StatusGatewayTimeout, nil)
)

func HandleError(c *gin.Context, err error) {
	// ตรวจสอบว่าเป็น AppError หรือไม่
	var appErr *AppError
	if errors.As(err, &appErr) {
		renderAppError(c, appErr)
		return
	}

	// ตรวจสอบว่าเป็น gorm.ErrRecordNotFound หรือไม่
	if errors.Is(err, gorm.ErrRecordNotFound) {
		renderAppError(c, ErrNotFound)
		return
	}

	// fallback: internal server error
	renderAppError(c, ErrInternalServer)
}

// ✅ แยกการสร้าง response
func renderAppError(c *gin.Context, appErr *AppError) {
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
