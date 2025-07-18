package usecase

import (
	"case-management/model"
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/redis/go-redis/v9"
)

type UseCase struct {
	// mu                      sync.Mutex
	Cache                   *redis.Client
	Logger                  *slog.Logger
	caseMangementRepository CaseManagementRepository
}

type CaseManagementRepository interface {
	CreateUser(c context.Context, user *model.User) (uint, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, id uint) (*model.User, error)
	DeleteUserByID(ctx context.Context, id uint) error
}

func New(caseMangementRepository CaseManagementRepository,
	cache *redis.Client,
	logger *slog.Logger,
	s3 *s3.Client) *UseCase {
	return &UseCase{
		caseMangementRepository: caseMangementRepository,
		Cache:                   cache,
		Logger:                  logger,
	}
}
