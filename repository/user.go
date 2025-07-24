package repository

import (
	"case-management/appcore/appcore_handler"
	"case-management/model"
	"context"
	"errors"
	"log"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const tableUserMetrix = "user_metrixes"
const tableUser = "users"

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

func (r *authRepo) GetAllUsers(c *gin.Context, limit, offset int, filter model.UserFilter) ([]*model.User, error) {
	var users []*model.User

	query := r.DB.Debug().WithContext(c).Model(&model.User{}).
		Preload("Role").Preload("Center").
		Joins("LEFT JOIN roles ON roles.id = users.role_id").
		Joins("LEFT JOIN centers ON centers.id = users.center_id")

	if filter.IsActive != nil {
		query = query.Where("users.is_active = ?", *filter.IsActive)
	}
	if filter.Role != "" {
		query = query.Where("roles.name = ?", filter.Role)
	}
	if filter.Team != "" {
		query = query.Where("users.team = ?", filter.Team)
	}
	if filter.Center != "" {
		query = query.Where("TRIM(centers.name) = ?", strings.TrimSpace(filter.Center))

	}
	if filter.Sort != "" {
		query = query.Order(filter.Sort)
	}

	if err := query.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *authRepo) CountUsersWithFilter(c *gin.Context, filter model.UserFilter) (int, error) {
	var count int64
	query := r.DB.WithContext(c).Model(&model.User{}).
		Joins("LEFT JOIN roles ON roles.id = users.role_id").
		Joins("LEFT JOIN centers ON centers.id = users.center_id")

	if filter.IsActive != nil {
		query = query.Where("users.is_active = ?", *filter.IsActive)
	}
	if filter.Role != "" {
		query = query.Where("roles.name = ?", filter.Role)
	}
	if filter.Team != "" {
		query = query.Where("users.team = ?", filter.Team)
	}
	if filter.Center != "" {

		query = query.Where("centers.name ILIKE ?", strings.TrimSpace(filter.Center))
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *authRepo) CountUsers(c *gin.Context) (int, error) {
	var count int64
	if err := r.DB.WithContext(c).Model(&model.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *authRepo) GetUserByID(c *gin.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.DB.WithContext(c).
		Preload("Role").
		Preload("Center").
		Where("users.id = ?", id).
		First(&user).Error; err != nil {
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

func (r *authRepo) UpdateUser(c *gin.Context, userID uint, input model.UserFilter) error {
	updateData := map[string]interface{}{}

	if input.Name != "" {
		updateData["name"] = input.Name
	}
	if input.IsActive != nil {
		updateData["is_active"] = *input.IsActive
	}
	if input.RoleID != 0 {
		updateData["role_id"] = input.RoleID
	}
	if input.Team != "" {
		updateData["team"] = input.Team
	}
	if input.CenterID != 0 {
		updateData["center_id"] = input.CenterID
	}

	if len(updateData) == 0 {
		return errors.New("no valid fields to update")
	}

	if err := r.DB.WithContext(c).Model(&model.User{}).Where("id = ?", userID).Updates(updateData).Error; err != nil {
		return err
	}

	return nil
}

func (r *authRepo) GetUserMetrix(ctx context.Context, role string) (*model.UserMetrix, error) {
	var userMetrix model.UserMetrix

	if err := r.DB.WithContext(ctx).Table(tableUserMetrix).Where("role = ?", role).Find(&userMetrix).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			details := map[string]string{
				"db": "ไม่พบผู้ใช้ในระบบ",
			}
			appErr := appcore_handler.NewAppError(
				appcore_handler.ErrNotFound.Code,
				appcore_handler.ErrNotFound.Message,
				appcore_handler.ErrNotFound.HTTPStatus,
				details,
			)
			return nil, appErr
		}

		return nil, appcore_handler.ErrInternalServer
	}

	return &userMetrix, nil
}

func (r *authRepo) GetUser(ctx context.Context, username string) (*model.User, error) {
	var user model.User

	log.Println("user:", username)
	if err := r.DB.WithContext(ctx).Table(tableUser).Where("username = ?", username).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			details := map[string]string{
				"db": "ไม่พบผู้ใช้ในระบบ",
			}
			appErr := appcore_handler.NewAppError(
				appcore_handler.ErrNotFound.Code,
				appcore_handler.ErrNotFound.Message,
				appcore_handler.ErrNotFound.HTTPStatus,
				details,
			)
			return nil, appErr
		}

		return nil, appcore_handler.ErrInternalServer
	}

	if (user == model.User{}) {
		details := map[string]string{
			"db": "ไม่พบผู้ใช้ในระบบ",
		}
		appErr := appcore_handler.NewAppError(
			appcore_handler.ErrNotFound.Code,
			appcore_handler.ErrNotFound.Message,
			appcore_handler.ErrNotFound.HTTPStatus,
			details,
		)
		return nil, appErr
	}

	return &user, nil
}

func (r *authRepo) BulkInsertUsers(c context.Context, users []model.User) error {
	tx := r.DB.WithContext(c).Create(&users)
	return tx.Error
}
