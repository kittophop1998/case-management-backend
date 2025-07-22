package mock

import (
	"case-management/model"
	"time"

	"github.com/gin-gonic/gin"
)

type MockRepository struct{}

func (m *MockRepository) GetAllUsers(c *gin.Context) ([]*model.User, error) {
	return []*model.User{
		{
			Id:             "1234",
			UserName:       "john_doe",
			Email:          "john.doe@example.com",
			RoleId:         "1",
			Team:           "Inbound",
			IsActive:       "ACTIVE",
			CreateDatetime: time.Now(),
			UpdateDatetime: time.Now(),
		},
	}, nil
}

func (m *MockRepository) GetUserByID(c *gin.Context, id string) (*model.User, error) {
	return &model.User{Id: id, UserName: "Mock_User", Email: "mock.user@example.com"}, nil
}

func (m *MockRepository) CreateUser(c *gin.Context, user *model.User) (string, error) {
	return user.Id, nil
}

func (m *MockRepository) DeleteUserByID(c *gin.Context, id string) error {
	return nil
}
