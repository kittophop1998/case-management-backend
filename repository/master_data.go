package repository

import (
	"case-management/model"

	"github.com/gin-gonic/gin"
)

func (r *authRepo) GetAllLookups(ctx *gin.Context) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	var teams []model.Team
	if err := r.DB.WithContext(ctx).Model(&model.Team{}).Find(&teams).Error; err != nil {
		return nil, err
	}
	result["teams"] = teams

	var queues []model.Queue
	if err := r.DB.WithContext(ctx).Model(&model.Queue{}).Find(&queues).Error; err != nil {
		return nil, err
	}
	result["queues"] = queues

	var roles []model.Role
	if err := r.DB.WithContext(ctx).Preload("Permissions").Find(&roles).Error; err != nil {
		return nil, err
	}
	result["roles"] = roles

	var centers []model.Center
	if err := r.DB.WithContext(ctx).Model(&model.Center{}).Find(&centers).Error; err != nil {
		return nil, err
	}
	result["centers"] = centers

	var permissions []model.Permission
	if err := r.DB.WithContext(ctx).Model(&model.Permission{}).Find(&permissions).Error; err != nil {
		return nil, err
	}
	result["permissions"] = permissions

	var departments []model.Department
	if err := r.DB.WithContext(ctx).Model(&model.Department{}).Find(&departments).Error; err != nil {
		return nil, err
	}
	result["departments"] = departments

	return result, nil
}

func (r *authRepo) GetAllPermissionsWithRoles(ctx *gin.Context, limit, offset int) ([]model.PermissionWithRolesResponse, error) {
	var permissions []model.Permission

	if err := r.DB.WithContext(ctx).
		Preload("Roles").
		Limit(limit).
		Offset(offset).
		Order("name").
		Find(&permissions).Error; err != nil {
		return nil, err
	}

	var result []model.PermissionWithRolesResponse
	for _, p := range permissions {
		var roleNames []string
		for _, role := range p.Roles {
			roleNames = append(roleNames, role.Name)
		}

		if len(p.Roles) == 0 {
			roleNames = []string{}
		}

		result = append(result, model.PermissionWithRolesResponse{
			Permission: p.Key,
			Name:       p.Name,
			Roles:      roleNames,
		})
	}

	return result, nil
}
