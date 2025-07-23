package mock

import (
	"case-management/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MockRepository struct{}

func (m *MockRepository) GetAllUsers(c *gin.Context) ([]*model.User, error) {
	return []*model.User{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			UserName: "john_doe",
			Email:    "john.doe@example.com",
			RoleId:   "1",
			Team:     "Inbound",
			CenterId: "CEN001",
			IsActive: "ACTIVE",
		},
	}, nil
}

func (m *MockRepository) GetUserByID(c *gin.Context, id string) (*model.User, error) {
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return &model.User{
		Model:    gorm.Model{ID: uint(uid)},
		UserName: "Mock_User",
		Email:    "mock.user@example.com",
		RoleId:   "1",
		Team:     "Inbound",
		CenterId: "Center001",
		IsActive: "ACTIVE",
	}, nil
}

func (m *MockRepository) CreateUser(c *gin.Context, user *model.User) (uint, error) {
	return user.ID, nil
}

func (m *MockRepository) DeleteUserByID(c *gin.Context, id string) error {
	return nil
}
