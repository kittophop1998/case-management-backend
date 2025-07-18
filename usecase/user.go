package usecase

import (
	"case-management/model"
	"context"
)

func (u *UseCase) CreateUser(ctx context.Context, user *model.User, userId uint) (uint, error) {
	id, err := u.caseMangementRepository.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UseCase) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return u.caseMangementRepository.GetAllUsers(ctx)
}

func (u *UseCase) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	return u.caseMangementRepository.GetUserByID(ctx, id)
}

func (u *UseCase) DeleteUserByID(ctx context.Context, id uint, deletedBy uint) error {
	err := u.caseMangementRepository.DeleteUserByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
