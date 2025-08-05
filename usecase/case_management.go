package usecase

import (
	"case-management/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (u *UseCase) CreateCase(c *gin.Context, caseManage *model.Cases) (uuid.UUID, error) {
	id, err := u.caseManagementRepository.CreateCase(c, caseManage)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (u *UseCase) GetAllCases(c *gin.Context, page, limit int, filter model.CaseFilter) ([]*model.Cases, int, error) {
	offset := (page - 1) * limit

	cases, err := u.caseManagementRepository.GetAllCases(c, limit, offset, filter)
	if err != nil {
		return nil, 0, err
	}

	total, err := u.caseManagementRepository.CountCasesWithFilter(c, filter)
	if err != nil {
		return nil, 0, err
	}

	return cases, total, nil
}

func (u *UseCase) CreateNoteType(c *gin.Context, note model.NoteTypes) (*model.NoteTypes, error) {
	return u.caseManagementRepository.CreateNoteType(c, note)
}

func (u *UseCase) GetCaseByID(ctx *gin.Context, id uuid.UUID) (*model.Cases, error) {
	return u.caseManagementRepository.GetCaseByID(ctx, id)
}
