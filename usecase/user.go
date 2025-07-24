package usecase

import (
	"case-management/model"

	"github.com/gin-gonic/gin"
)

func (u *UseCase) CreateUser(c *gin.Context, user *model.User) (uint, error) {
	id, err := u.caseManagementRepository.CreateUser(c, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UseCase) GetAllUsers(c *gin.Context, page, limit int, filter model.UserFilter) ([]*model.User, int, error) {
	offset := (page - 1) * limit

	users, err := u.caseManagementRepository.GetAllUsers(c, limit, offset, filter)
	if err != nil {
		return nil, 0, err
	}

	total, err := u.caseManagementRepository.CountUsersWithFilter(c, filter)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
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

func (u *UseCase) UpdateUser(c *gin.Context, userID uint, input model.UserFilter) error {
	return u.caseManagementRepository.UpdateUser(c, userID, input)
}
