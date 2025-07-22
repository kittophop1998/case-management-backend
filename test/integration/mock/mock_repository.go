package mock

import (
	"case-management/model"

	"github.com/gin-gonic/gin"
)

type MockRepository struct{}

func (m *MockRepository) GetAllUsers(c *gin.Context) ([]*model.User, error) {
	return []*model.User{
		{Id: 1, Username: "John Doe", Email: "john.doe@example.com"},
		{Id: 2, Username: "Jane Smith", Email: "jane.smith@example.com"},
	}, nil
}

func (m *MockRepository) GetUserByID(c *gin.Context, id uint) (*model.User, error) {
	return &model.User{Id: id, Username: "Mock User", Email: "mock.user@example.com"}, nil
}

func (m *MockRepository) CreateUser(c *gin.Context, user *model.User) (uint, error) {
	return user.Id, nil
}

func (m *MockRepository) DeleteUserByID(c *gin.Context, id uint) error {
	return nil
}
