package usecase

import (
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func (u *UseCase) Login(ctx context.Context, userData model.LoginRequest) (*model.LoginResponse, error) {

	parts := strings.Split(userData.UserData, ".")
	if len(parts) != 3 {

		return nil, appcore_handler.ErrBadRequest
	}

	payload := parts[1]

	// base64url decode
	payloadBytes, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return nil, appcore_handler.ErrInternalServer

	}

	var data map[string]string
	if err := json.Unmarshal(payloadBytes, &data); err != nil {
		return nil, appcore_handler.ErrInternalServer

	}

	fmt.Printf("Decoded payload: %+v\n", data)
	username := data["username"]
	// password := data["password"]

	// // create ldap connection
	// conn, err := ldap.Dial("tcp", "192.168.129.239:389")
	// if err != nil {

	// 	// err := models.ErrorResponse{
	// 	// 	// ErrorCode:    "CDP001",
	// 	// 	// ErrorMessage: "LDAP connection failed",
	// 	// }
	// 	return nil, appcore_handler.ErrInternalServer
	// 	// c.JSON(http.StatusInternalServerError, err)
	// 	// return

	// }

	// userBide := "HEADOFFICE\\" + username

	// // bind ldap user
	// if err := conn.Bind(userBide, password); err != nil {
	// 	details := map[string]string{
	// 		"ldap": "ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง",
	// 	}
	// 	appErr := appcore_handler.NewAppError(
	// 		appcore_handler.ErrNotFound.Code,
	// 		appcore_handler.ErrNotFound.Message,
	// 		appcore_handler.ErrNotFound.HTTPStatus,
	// 		details,
	// 	)
	// 	// HandleError(c, appcore_handler.ErrNotFound)
	// 	return nil, appErr
	// 	// return nil, appcore_handler.ErrInternalServer
	// 	// HandleError(c, appcore_handler.ErrNotFound)
	// 	// return
	// }

	userResp, err := u.caseManagementRepository.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}

	userMetrixResp, err := u.caseManagementRepository.GetUserMetrix(ctx, userResp.Role.Name)
	if err != nil {
		return nil, err
	}

	accessLog := model.AccessLogs{
		Username:      username,
		LogonDatetime: time.Now(),
		LogonResult:   "success",
	}

	isSave := u.SaveAccressLog(ctx, accessLog) //แก้ด้วย

	if isSave != nil {
		return nil, err
	}

	userMetrix := model.UserMetrixResponse{
		Role:        userMetrixResp.Role,
		Create:      userMetrixResp.Create,
		Update:      userMetrixResp.Update,
		Delete:      userMetrixResp.Delete,
		CreateEvent: userMetrixResp.CreateEvent,
		UpdateEvent: userMetrixResp.UpdateEvent,
		DeleteEvent: userMetrixResp.DeleteEvent,
	}
	user := model.UserResponse{
		Username:   username,
		UserMetrix: userMetrix,
	}

	login := model.LoginResponse{
		User: user,
	}

	return &login, nil
}

func (u *UseCase) SaveAccressLog(ctx context.Context, accessLog model.AccessLogs) error {

	err := u.caseManagementRepository.SaveAccressLog(ctx, accessLog) //แก้ด้วย

	if err != nil {
		return appcore_handler.ErrInternalServer
	}

	return nil
}

func (u *UseCase) GenerateToken(ttl time.Duration, metadata *appcore_model.Metadata) (signedToken string, err error) {
	return u.caseManagementRepository.GenerateToken(ttl, metadata)
}
