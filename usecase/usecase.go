package usecase

import (
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type UseCase struct {
	// mu                      sync.Mutex
	Cache                    *redis.Client
	Logger                   *slog.Logger
	caseManagementRepository CaseManagementRepository
}

type CaseManagementRepository interface {
	CreateUser(c *gin.Context, user *model.User) (uuid.UUID, error)
	GetAllUsers(c *gin.Context, limit, offset int, filter model.UserFilter) ([]*model.User, error)
	GetUserByID(c *gin.Context, id uuid.UUID) (*model.User, error)
	DeleteUserByID(c *gin.Context, id string) error
	CountUsers(c *gin.Context) (int, error)
	CountUsersWithFilter(c *gin.Context, filter model.UserFilter) (int, error)
	UpdateUser(c *gin.Context, userID uuid.UUID, input model.UserFilter) error
	SaveAccessLog(ctx context.Context, accessLog model.AccessLogs) error
	GetUserByUserName(c *gin.Context, username string) (*model.User, error)
	GenerateToken(ttl time.Duration, metadata *appcore_model.Metadata) (signedToken string, err error)
	BulkInsertUsers(c context.Context, users []model.User) error
	StoreToken(c *gin.Context, accessToken string) error
	ValidateToken(signedToken string) (claims *appcore_model.JwtClaims, err error)
	DeleteToken(c *gin.Context, accessToken string) error
	GetAllLookups(ctx *gin.Context) (map[string]interface{}, error)
	GetAllPermissionsWithRoles(ctx *gin.Context) ([]model.PermissionWithRolesResponse, error)
	UpdatePermissionRoles(ctx *gin.Context, req model.UpdatePermissionRolesRequest) error
}

func New(caseManagementRepository CaseManagementRepository,
	cache *redis.Client,
	logger *slog.Logger) *UseCase {
	return &UseCase{
		caseManagementRepository: caseManagementRepository,
		Cache:                    cache,
		Logger:                   logger,
	}
}
