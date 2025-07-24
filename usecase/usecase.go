package usecase

import (
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"context"
	"log/slog"
	"time"

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
	CreateUser(c *gin.Context, user *model.User) (uint, error)
	GetAllUsers(c *gin.Context, limit, offset int, filter model.UserFilter) ([]*model.User, error)
	GetUserByID(c *gin.Context, id string) (*model.User, error)
	DeleteUserByID(c *gin.Context, id string) error
	CountUsers(c *gin.Context) (int, error)
	CountUsersWithFilter(c *gin.Context, filter model.UserFilter) (int, error)
	UpdateUser(c *gin.Context, userID uint, input model.UserFilter) error
	SaveAccressLog(ctx context.Context, accessLog model.AccessLogs) error
	GetUserMetrix(ctx context.Context, role string) (*model.UserMetrix, error)
	GetUser(ctx context.Context, username string) (*model.User, error)
	GenerateToken(ttl time.Duration, metadata *appcore_model.Metadata) (signedToken string, err error)
	BulkInsertUsers(c context.Context, users []model.User) error
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
