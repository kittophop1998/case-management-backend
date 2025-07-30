package mock

import (
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type MockRepository struct{}

func (m *MockRepository) GetAllUsers(c *gin.Context, limit, offset int, filter model.UserFilter) ([]*model.User, error) {
	isActive := true
	mockUsers := []*model.User{
		{
			Model: model.Model{
				ID: uuid.New(),
			},
			Username: "John Doe",
			Email:    "john.doe@example.com",
			Name:     "Johnathan Doe",
			Team: model.Team{
				ID:   uuid.New(),
				Name: "Inbound",
			},
			IsActive: &isActive,
			Center: model.Center{
				ID:   uuid.New(),
				Name: "Center 1",
			},
			Role: model.Role{
				ID:   uuid.New(),
				Name: "Admin",
			},
		},
	}
	return mockUsers, nil
}

func (m *MockRepository) GetUserByID(c *gin.Context, id uuid.UUID) (*model.User, error) {
	mockUser := &model.User{
		Model: model.Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: "John Doe",
		Email:    "john.doe@example.com",
		Name:     "John",
		Team: model.Team{
			ID:   uuid.New(),
			Name: "Inbound",
		},
		IsActive: func(b bool) *bool { return &b }(true),
		CenterID: uuid.New(),
		RoleID:   uuid.New(),
		Center:   model.Center{Name: "Center A"},
		Role:     model.Role{Name: "Admin"},
	}
	return mockUser, nil
}

func (m *MockRepository) CreateUser(c *gin.Context, user *model.User) (uuid.UUID, error) {
	if user.ID == uuid.Nil {
		newID, err := uuid.NewUUID()
		if err != nil {
			return uuid.Nil, err
		}
		user.ID = newID
	}

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

func (m *MockRepository) UpdateUser(c *gin.Context, userID uuid.UUID, input model.UserFilter) error {
	return nil
}

func (m *MockRepository) GetAllLookups(ctx *gin.Context) (map[string]interface{}, error) {
	mockData := map[string]interface{}{
		"teams": []model.Team{
			{ID: uuid.New(), Name: "Inbound"},
			{ID: uuid.New(), Name: "Outbound"},
		},
		"roles": []model.Role{
			{ID: uuid.New(), Name: "Admin"},
			{ID: uuid.New(), Name: "User"},
		},
		"centers": []model.Center{
			{ID: uuid.New(), Name: "BKK"},
		},
		"permissions": []model.Permission{
			{ID: uuid.New(), Name: "case.create"},
			{ID: uuid.New(), Name: "case.view"},
		},
	}
	return mockData, nil
}

func (m *MockRepository) GetUserByUserName(c *gin.Context, username string) (*model.User, error) {
	return nil, nil
}

func (m *MockRepository) GetUserMetrix(ctx context.Context, role string) (*model.UserMetrix, error) {
	return nil, nil
}

func (m *MockRepository) SaveAccessLog(ctx context.Context, accessLog model.AccessLogs) error {
	return nil
}

func (m *MockRepository) GenerateToken(ttl time.Duration, metadata *appcore_model.Metadata) (signedToken string, err error) {
	return "", nil
}

func (m *MockRepository) BulkInsertUsers(c context.Context, users []model.User) error {
	return nil
}

func (m *MockRepository) StoreToken(c *gin.Context, accessToken string) error {
	return nil
}

func (m *MockRepository) ValidateToken(signedToken string) (claims *appcore_model.JwtClaims, err error) {
	return nil, nil
}

func (m *MockRepository) DeleteToken(c *gin.Context, accessToken string) error {
	return nil
}

func (m *MockRepository) GetAllPermissionsWithRoles(ctx *gin.Context) ([]model.PermissionWithRolesResponse, error) {
	return nil, nil
}

func (m *MockRepository) UpdatePermissionRoles(ctx *gin.Context, req model.UpdatePermissionRolesRequest) error {
	return nil
}
