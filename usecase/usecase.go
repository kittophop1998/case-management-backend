package usecase

import (
	"case-management/model"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type UseCase struct {
	// mu                      sync.Mutex
	Cache                   *redis.Client
	Logger                  *slog.Logger
	caseMangementRepository CaseManagementRepository
}

type CaseManagementRepository interface {
	CreateUser(c *gin.Context, user *model.User) (uint, error)
	GetAllUsers(c *gin.Context) ([]*model.User, error)
	GetUserByID(c *gin.Context, id uint) (*model.User, error)
	DeleteUserByID(c *gin.Context, id uint) error
}

func New(caseMangementRepository CaseManagementRepository,
	cache *redis.Client,
	logger *slog.Logger) *UseCase {
	return &UseCase{
		caseMangementRepository: caseMangementRepository,
		Cache:                   cache,
		Logger:                  logger,
	}
}
