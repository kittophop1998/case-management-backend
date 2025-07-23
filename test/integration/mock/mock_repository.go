package mock

import (
	"case-management/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MockRepository struct{}

// func (m *MockRepository) GetAllUsers(*gin.Context, int, int) ([]*model.User, error) {
// 	return []*model.User{
// 		{
// 			Model: gorm.Model{
// 				ID:        1,
// 				CreatedAt: time.Now(),
// 				UpdatedAt: time.Now(),
// 			},
// 			UserName: "john_doe",
// 			Email:    "john.doe@example.com",
// 			Team:     "Inbound",
// 			IsActive: "ACTIVE",
// 		},
// 	}, nil
// }

func (m *MockRepository) GetAllUsers(c *gin.Context, limit, offset int, filter model.UserFilter) ([]*model.User, error) {
	val := true
	return []*model.User{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			UserName: "john_doe",
			Email:    "john.doe@example.com",
			Team:     "Inbound",
			IsActive: &val,
		},
	}, nil
}

func (m *MockRepository) GetUserByID(c *gin.Context, id string) (*model.User, error) {
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	val := true
	return &model.User{
		Model:    gorm.Model{ID: uint(uid)},
		UserName: "Mock_User",
		Email:    "mock.user@example.com",
		Team:     "Inbound",
		IsActive: &val,
	}, nil
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
