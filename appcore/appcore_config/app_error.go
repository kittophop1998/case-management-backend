package appcore_config

import (
	"net/http"
)

type AppError struct {
	Code       string
	Message    Message
	HTTPStatus int
	Details    interface{}
}

type Message struct {
	Th string `json:"th"`
	En string `json:"en"`
}

func (e *AppError) Error() string {
	return e.Message.Th
}

func NewAppError(code string, message Message, status int, details interface{}) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: status,
		Details:    details,
	}
}

// 🔁 Common errors
var (
	ErrBadRequest         = newAppError("BAD_REQUEST", "ข้อมูลที่ส่งมาไม่ถูกต้องตามที่ระบบรองรับ", "The data provided is invalid or unsupported by the system", http.StatusBadRequest)
	ErrRequiredParam      = newAppError("REQUIRED_PARAMETER", "ไม่สามารถดึงข้อมูลได้ เนื่องจากระบบขาดข้อมูลบางส่วน", "Missing required parameters to retrieve data", http.StatusBadRequest)
	ErrFilterRequired     = newAppError("FILTER_REQUIRED", "โปรดเลือกอย่างน้อย 1 เงื่อนไขเพื่อค้นหา", "Please select at least one filter condition for search", http.StatusBadRequest)
	ErrNotFound           = newAppError("NOT_FOUND", "ไม่พบข้อมูลที่ค้นหา", "Requested data not found", http.StatusNotFound)
	ErrInternalServer     = newAppError("INTERNAL_SERVER_ERROR", "เกิดข้อผิดพลาดในระบบ กรุณาลองใหม่ภายหลัง", "An internal server error occurred. Please try again later", http.StatusInternalServerError)
	ErrServiceUnavailable = newAppError("SERVICE_UNAVAILABLE", "ไม่สามารถใช้บริการได้ชั่วคราว หรืออยู่ระหว่างการบำรุงรักษา", "The service is temporarily unavailable or under maintenance", http.StatusServiceUnavailable)
	ErrGatewayTimeout     = newAppError("NO_RESPONSE", "ไม่ได้รับการตอบสนองจากบริการภายนอก", "No response from an upstream service", http.StatusGatewayTimeout)
)

func newAppError(code string, th string, en string, status int) *AppError {
	return NewAppError(code, Message{Th: th, En: en}, status, nil)
}

func (e *AppError) WithDetails(details interface{}) *AppError {
	return &AppError{
		Code:       e.Code,
		Message:    e.Message,
		HTTPStatus: e.HTTPStatus,
		Details:    details,
	}
}

func (e *AppError) WithMessage(message Message) *AppError {
	return &AppError{
		Code:       e.Code,
		Message:    message,
		HTTPStatus: e.HTTPStatus,
		Details:    e.Details,
	}
}
