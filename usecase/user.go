package usecase

import (
	"case-management/model"

	"github.com/gin-gonic/gin"
)

func (u *UseCase) CreateUser(c *gin.Context, user *model.User) (uint, error) {
	id, err := u.caseMangementRepository.CreateUser(c, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UseCase) GetAllUsers(c *gin.Context) ([]*model.User, error) {
	return u.caseMangementRepository.GetAllUsers(c)
}

func (u *UseCase) GetUserByID(c *gin.Context, id uint) (*model.User, error) {
	return u.caseMangementRepository.GetUserByID(c, id)
}

func (u *UseCase) DeleteUserByID(c *gin.Context, id uint) error {
	err := u.caseMangementRepository.DeleteUserByID(c, id)
	if err != nil {
		return err
	}
	return nil
}
