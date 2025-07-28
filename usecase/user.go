package usecase

import (
	"case-management/model"
	"case-management/utils"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	ColUserName = 0
	ColEmail    = 1
	ColTeam     = 2
	ColIsActive = 3
	ColCenterID = 4
	ColRoleID   = 5
	ColName     = 6
)

func (u *UseCase) CreateUser(c *gin.Context, user *model.User) (uuid.UUID, error) {
	id, err := u.caseManagementRepository.CreateUser(c, user)
	if err != nil {
		return uuid.Nil, err
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

func (u *UseCase) GetUserByID(c *gin.Context, id uuid.UUID) (*model.User, error) {
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

// func (u *UseCase) ImportUsersFromCSV(c context.Context, file io.Reader) error {
// 	reader := csv.NewReader(file)
// 	reader.TrimLeadingSpace = true

// 	// อ่าน header
// 	_, err := reader.Read()
// 	if err != nil {
// 		return fmt.Errorf("cannot read csv header: %v", err)
// 	}

// 	var users []model.User
// 	for {
// 		record, err := reader.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return fmt.Errorf("error reading csv: %v", err)
// 		}

// 		isActive := record[ColIsActive] == "true"
// 		user := model.User{
// 			UserName: record[ColUserName],
// 			Email:    record[ColEmail],
// 			Team:     record[ColTeam],
// 			IsActive: &isActive,
// 			CenterID: parseUint(record[ColCenterID]),
// 			RoleID:   parseUint(record[ColRoleID]),
// 			Name:     record[ColName],
// 		}
// 		users = append(users, user)
// 	}

// 	return u.caseManagementRepository.BulkInsertUsers(c, users)
// }

func (u *UseCase) ImportUsersFromCSVWithProgress(c context.Context, file io.Reader, taskID string) error {
	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	// อ่าน header
	_, err := reader.Read()
	if err != nil {
		return fmt.Errorf("cannot read csv header: %v", err)
	}

	var successCount int
	var total int
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error reading csv line: %v", err)
			continue
		}
		total++
		isActive := record[ColIsActive] == "true"

		user := model.User{
			UserName: record[ColUserName],
			Email:    record[ColEmail],
			Team:     record[ColTeam],
			IsActive: &isActive,
			CenterID: uuid.MustParse(record[ColCenterID]),
			RoleID:   uuid.MustParse(record[ColRoleID]),
			Name:     record[ColName],
		}

		err = u.caseManagementRepository.BulkInsertUsers(c, []model.User{user})
		if err != nil {
			log.Printf("failed to insert user %v: %v", user.Email, err)
			continue
		}

		successCount++
		if total > 0 { // Ensure safe division
			progress := int(float64(successCount) / float64(total) * 100)
			utils.SetProgress(taskID, progress)
		}
	}

	utils.SetProgress(taskID, 100)
	return nil
}
