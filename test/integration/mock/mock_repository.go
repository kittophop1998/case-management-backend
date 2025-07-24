package mock

import (
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type MockRepository struct{}

func (m *MockRepository) GetAllUsers(c *gin.Context, limit, offset int, filter model.UserFilter) ([]*model.User, error) {
	return nil, nil
}

func (m *MockRepository) GetUserByID(c *gin.Context, id string) (*model.User, error) {
	return nil, nil
}

func (m *MockRepository) CreateUser(c *gin.Context, user *model.User) (uint, error) {
	return user.ID, nil
}

func (m *MockRepository) DeleteUserByID(c *gin.Context, id string) error {
	return nil
}

func (m *MockRepository) CountUsers(*gin.Context) (int, error) {
	return 0, nil
}

func (m *MockRepository) CountUsersWithFilter(c *gin.Context, filter model.UserFilter) (int, error) {
	return 0, nil
}

func (m *MockRepository) UpdateUser(c *gin.Context, userID uint, input model.UserFilter) error {
	return nil
}

func (m *MockRepository) GetUser(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}

func (m *MockRepository) GetUserMetrix(ctx context.Context, role string) (*model.UserMetrix, error) {
	return nil, nil
}

func (m *MockRepository) SaveAccressLog(ctx context.Context, accessLog model.AccessLogs) error {
	return nil
}

func (m *MockRepository) GenerateToken(ttl time.Duration, metadata *appcore_model.Metadata) (signedToken string, err error) {
	return "", nil
}

func (m *MockRepository) BulkInsertUsers(c context.Context, users []model.User) error {
	return nil
}
