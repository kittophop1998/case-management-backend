package usecase

import (
	"case-management/model"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

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

func (u *UseCase) ImportUsersFromCSV(c context.Context, file io.Reader) error {
	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	// อ่าน header
	_, err := reader.Read()
	if err != nil {
		return fmt.Errorf("cannot read csv header: %v", err)
	}

	var users []model.User
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading csv: %v", err)
		}

		isActive := record[3] == "true"

		user := model.User{
			UserName: record[0],
			Email:    record[1],
			Team:     record[2],
			IsActive: &isActive,
			CenterID: parseUint(record[4]),
			RoleID:   parseUint(record[5]),
			Name:     record[6],
		}
		users = append(users, user)
	}

	return u.caseManagementRepository.BulkInsertUsers(c, users)
}

func parseUint(s string) uint {
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint(val)
}
