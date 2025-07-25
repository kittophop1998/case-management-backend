package usecase

import (
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/ldap.v2"
)

func (u *UseCase) Login(ctx context.Context, userLogin model.LoginRequest) (*model.LoginResponse, error) {
	username := userLogin.Username
	password := userLogin.Password

	// --- Optional LDAP Authentication (currently disabled) ---
	if err := u.authenticateWithLDAP(username, password); err != nil {
		return nil, err
	}

	userResp, err := u.caseManagementRepository.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}

	userMetrixResp, err := u.caseManagementRepository.GetUserMetrix(ctx, userResp.Role.Name)
	if err != nil {
		return nil, err
	}

	if err := u.logAccessSuccess(ctx, username); err != nil {
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
	userResponse := u.mapToLoginResponse(username, userMetrix)
	return &userResponse, nil
}

// Optional LDAP authentication logic
func (u *UseCase) authenticateWithLDAP(username, password string) error {
	conn, err := ldap.Dial("tcp", "192.168.129.239:389")
	if err != nil {
		return appcore_handler.ErrInternalServer
	}
	defer conn.Close()

	userBind := "HEADOFFICE\\" + username
	if err := conn.Bind(userBind, password); err != nil {
		detail := map[string]string{"ldap": "ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง"}
		return appcore_handler.NewAppError(
			appcore_handler.ErrNotFound.Code,
			appcore_handler.ErrNotFound.Message,
			appcore_handler.ErrNotFound.HTTPStatus,
			detail,
		)
	}
	return nil
}

func (u *UseCase) logAccessSuccess(ctx context.Context, username string) error {
	accessLog := model.AccessLogs{
		Username:      username,
		LogonDatetime: time.Now(),
		LogonResult:   "success",
	}
	return u.SaveAccessLog(ctx, accessLog)
}

func (u *UseCase) SaveAccessLog(ctx context.Context, accessLog model.AccessLogs) error {
	if err := u.caseManagementRepository.SaveAccessLog(ctx, accessLog); err != nil {
		return appcore_handler.ErrInternalServer
	}
	return nil
}

func (u *UseCase) mapToLoginResponse(username string, metrix model.UserMetrixResponse) model.LoginResponse {
	userMetrix := model.UserMetrixResponse{
		Role:        metrix.Role,
		Create:      metrix.Create,
		Update:      metrix.Update,
		Delete:      metrix.Delete,
		CreateEvent: metrix.CreateEvent,
		UpdateEvent: metrix.UpdateEvent,
		DeleteEvent: metrix.DeleteEvent,
	}

	user := model.UserResponse{
		Username:   username,
		UserMetrix: userMetrix,
	}

	return model.LoginResponse{User: user}
}

func (u *UseCase) GenerateToken(ttl time.Duration, metadata *appcore_model.Metadata) (string, error) {
	return u.caseManagementRepository.GenerateToken(ttl, metadata)
}

func (u *UseCase) StoreToken(c *gin.Context, accessToken string) error {
	return u.caseManagementRepository.StoreToken(c, accessToken)
}

func (u *UseCase) ValidateToken(signedToken string) (*appcore_model.JwtClaims, error) {
	return u.caseManagementRepository.ValidateToken(signedToken)
}

func (u *UseCase) DeleteToken(c *gin.Context, accessToken string) error {
	return u.caseManagementRepository.DeleteToken(c, accessToken)
}
