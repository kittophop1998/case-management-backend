package usecase

import (
	"case-management/model"
	"log/slog"

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
	CreateUser(c *gin.Context, user *model.User) (string, error)
	GetAllUsers(c *gin.Context) ([]*model.User, error)
	GetUserByID(c *gin.Context, id string) (*model.User, error)
	DeleteUserByID(c *gin.Context, id string) error
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
