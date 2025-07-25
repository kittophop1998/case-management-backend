package usecase

import (
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/ldap.v2"
)

func (u *UseCase) Login(ctx *gin.Context, req model.LoginRequest) (*model.LoginResponse, error) {
	// Optional LDAP authentication
	if err := u.authenticateWithLDAP(req.Username, req.Password); err != nil {
		return nil, err
	}

	user, err := u.caseManagementRepository.GetUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	userMetrix, err := u.caseManagementRepository.GetUserMetrix(ctx, user.Role.Name)
	if err != nil {
		return nil, err
	}

	// Convert userMetrix to UserMetrixResponse
	userMetrixResp := &model.UserMetrixResponse{
		Role:        userMetrix.Role,
		Create:      userMetrix.Create,
		Update:      userMetrix.Update,
		Delete:      userMetrix.Delete,
		CreateEvent: userMetrix.CreateEvent,
		UpdateEvent: userMetrix.UpdateEvent,
		DeleteEvent: userMetrix.DeleteEvent,
	}

	if err := u.SaveAccessLog(ctx, req.Username, true); err != nil {
		return nil, err
	}

	// Generate JWT token and set in cookie
	token, err := u.GenerateToken(24*time.Hour, &appcore_model.Metadata{
		UserID:   user.ID,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	if err := u.setAccessTokenCookie(ctx, token); err != nil {
		return nil, err
	}

	// Map and return response

	resp := u.buildLoginResponse(user, userMetrixResp)
	return &resp, nil
}

// LDAP authentication
func (u *UseCase) authenticateWithLDAP(username, password string) error {
	conn, err := ldap.Dial("tcp", "192.168.129.239:389")
	if err != nil {
		return appcore_handler.ErrInternalServer
	}
	defer conn.Close()

	if err := conn.Bind("HEADOFFICE\\"+username, password); err != nil {
		return appcore_handler.NewAppError(
			appcore_handler.ErrNotFound.Code,
			appcore_handler.ErrNotFound.Message,
			appcore_handler.ErrNotFound.HTTPStatus,
			map[string]string{"ldap": "ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง"},
		)
	}
	return nil
}

// Save Log Access
func (u *UseCase) SaveAccessLog(ctx context.Context, username string, success bool) error {
	logResult := "success"
	if !success {
		logResult = "failed"
	}

	return u.caseManagementRepository.SaveAccessLog(ctx, model.AccessLogs{
		Username:      username,
		LogonDatetime: time.Now(),
		LogonResult:   logResult,
	})
}

// Set access token in cookie
func (u *UseCase) setAccessTokenCookie(c *gin.Context, token string) error {
	isSecure := gin.Mode() == gin.ReleaseMode

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		MaxAge:   86400, // 1 วัน
		HttpOnly: true,
		Secure:   isSecure,
		SameSite: http.SameSiteLaxMode,
		Domain:   "localhost", // หรือเว้นว่างถ้าไม่ได้ข้ามโดเมน
	})
	return nil
}

// Token generation (pass-through)
func (u *UseCase) GenerateToken(ttl time.Duration, metadata *appcore_model.Metadata) (string, error) {
	return u.caseManagementRepository.GenerateToken(ttl, metadata)
}

// Build response from user + metrix
func (u *UseCase) buildLoginResponse(user *model.User, metrix *model.UserMetrixResponse) model.LoginResponse {
	return model.LoginResponse{
		User: model.UserResponse{
			Username: user.UserName,
			UserMetrix: model.UserMetrixResponse{
				Role:        metrix.Role,
				Create:      metrix.Create,
				Update:      metrix.Update,
				Delete:      metrix.Delete,
				CreateEvent: metrix.CreateEvent,
				UpdateEvent: metrix.UpdateEvent,
				DeleteEvent: metrix.DeleteEvent,
			},
		},
	}
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
