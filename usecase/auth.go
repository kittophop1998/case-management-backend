package usecase

import (
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/ldap.v2"
)

func (u *UseCase) Login(ctx *gin.Context, req model.LoginRequest) (*model.UserResponse, error) {
	// Check if it's a hardcoded admin login
	if isAdminLogin(req) {
		return u.loginAsAdmin(ctx, req.Username)
	}

	// Authenticate via LDAP
	if err := u.authenticateWithLDAP(req.Username, req.Password); err != nil {
		return nil, err
	}

	// Fetch user from DB
	user, err := u.caseManagementRepository.GetUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	// Generate token and set cookie
	token, err := u.GenerateToken(24*time.Hour, &appcore_model.Metadata{
		UserID:   user.ID,
		Username: user.UserName,
	})
	if err != nil {
		return nil, err
	}
	if err := u.setAccessTokenCookie(ctx, token); err != nil {
		return nil, err
	}

	// Map to response
	resp := &model.UserResponse{
		Username: user.UserName,
		RoleId:   user.Role.ID,
		Role:     user.Role,
	}
	return resp, nil
}

// loginAsAdmin handles hardcoded admin login
func (u *UseCase) loginAsAdmin(ctx *gin.Context, username string) (*model.UserResponse, error) {
	token, err := u.GenerateToken(24*time.Hour, &appcore_model.Metadata{
		UserID:   1,
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	if err := u.setAccessTokenCookie(ctx, token); err != nil {
		return nil, err
	}

	adminRoleID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	return &model.UserResponse{
		Username: username,
		RoleId:   adminRoleID,
		Role: model.Role{
			ID:   adminRoleID,
			Name: "Admin",
		},
	}, nil
}

// isAdminLogin checks if login credentials match hardcoded admin
func isAdminLogin(req model.LoginRequest) bool {
	return req.Username == "admin" || req.Password == "admin"
}

// authenticateWithLDAP authenticates via external LDAP
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

// SaveAccessLog logs login attempts
func (u *UseCase) SaveAccessLog(ctx context.Context, username string, success bool) error {
	status := "success"
	if !success {
		status = "failed"
	}

	return u.caseManagementRepository.SaveAccessLog(ctx, model.AccessLogs{
		Username:      username,
		LogonDatetime: time.Now(),
		LogonResult:   status,
	})
}

// setAccessTokenCookie sets JWT token in HTTP cookie
func (u *UseCase) setAccessTokenCookie(c *gin.Context, token string) error {
	isSecure := gin.Mode() == gin.ReleaseMode

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
		Secure:   isSecure,
		SameSite: http.SameSiteLaxMode,
		Domain:   "localhost",
	})
	return nil
}

// Token functions (delegate to repository)
func (u *UseCase) GenerateToken(ttl time.Duration, metadata *appcore_model.Metadata) (string, error) {
	return u.caseManagementRepository.GenerateToken(ttl, metadata)
}
func (u *UseCase) StoreToken(c *gin.Context, token string) error {
	return u.caseManagementRepository.StoreToken(c, token)
}
func (u *UseCase) ValidateToken(signedToken string) (*appcore_model.JwtClaims, error) {
	return u.caseManagementRepository.ValidateToken(signedToken)
}
func (u *UseCase) DeleteToken(c *gin.Context, token string) error {
	return u.caseManagementRepository.DeleteToken(c, token)
}
