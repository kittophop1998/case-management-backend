package repository

import (
	"case-management/model"
	"case-management/utils"
	"context"
	"log/slog"
)

func (a *authRepo) CreateUser(ctx context.Context, user *model.User) (uint, error) {
	a.Logger.Info("Creating user", slog.String("username", user.Username))

	// Hash password
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		a.Logger.Error("Failed to hash password", slog.Any("error", err))
		return 0, err
	}
	user.Password = hashedPwd

	// Save to DB
	if err := a.DB.Create(user).Error; err != nil {
		a.Logger.Error("Failed to create user", slog.Any("error", err))
		return 0, err
	}

	a.Logger.Info("User created successfully", slog.Uint64("user_id", uint64(user.Id)))
	return user.Id, nil
}

func (r *authRepo) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.DB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *authRepo) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	if err := r.DB.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepo) DeleteUserByID(ctx context.Context, id uint) error {
	if err := r.DB.WithContext(ctx).Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
