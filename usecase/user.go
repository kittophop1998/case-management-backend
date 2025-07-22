package usecase

import (
	"case-management/model"

	"github.com/gin-gonic/gin"
)

func (u *UseCase) CreateUser(c *gin.Context, user *model.User) (string, error) {
	id, err := u.caseManagementRepository.CreateUser(c, user)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (u *UseCase) GetAllUsers(c *gin.Context) ([]*model.User, error) {
	return u.caseManagementRepository.GetAllUsers(c)
}

func (u *UseCase) GetUserByID(c *gin.Context, id string) (*model.User, error) {
	return u.caseManagementRepository.GetUserByID(c, id)
}

func (u *UseCase) DeleteUserByID(c *gin.Context, id string) error {
	err := u.caseManagementRepository.DeleteUserByID(c, id)
	if err != nil {
		return err
	}
	return nil
}
