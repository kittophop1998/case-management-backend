package repository

import (
	"case-management/model"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func (a *authRepo) CreateUser(c *gin.Context, user *model.User) (uint, error) {
	a.Logger.Info("Creating user", slog.String("username", user.UserName))

	// Save to DB
	if err := a.DB.Create(user).Error; err != nil {
		a.Logger.Error("Failed to create user", slog.Any("error", err))
		return 0, err
	}

	a.Logger.Info("User created successfully", slog.Any("user_id", user.ID))
	return user.ID, nil
}

func (r *authRepo) GetAllUsers(c *gin.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.DB.WithContext(c).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *authRepo) GetUserByID(c *gin.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.DB.WithContext(c).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepo) DeleteUserByID(c *gin.Context, id string) error {
	if err := r.DB.WithContext(c).Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
