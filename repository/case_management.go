package repository

import (
	"case-management/model"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *authRepo) CreateCase(ctx *gin.Context, c *model.Cases) (uuid.UUID, error) {
	a.Logger.Info("Creating case", slog.String("title", c.Title))

	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	if err := a.DB.Create(c).Error; err != nil {
		a.Logger.Error("Failed to create case", slog.Any("error", err))
		return uuid.Nil, err
	}

	a.Logger.Info("Case created successfully", slog.Any("case_id", c.ID))
	return c.ID, nil
}

func (r *authRepo) GetAllCases(c *gin.Context, limit, offset int, filter model.CaseFilter) ([]*model.Cases, error) {
	var cases []*model.Cases

	query := r.DB.WithContext(c).Model(&model.Cases{})

	if filter.Keyword != "" {
		kw := "%" + strings.TrimSpace(filter.Keyword) + "%"
		query = query.Where(
			r.DB.Where("title ILIKE ?", kw).
				Or("customer_id ILIKE ?", kw).
				Or("created_by ILIKE ?", kw).
				Or("CAST(sla_date AS TEXT) ILIKE ?", kw).
				Or("CAST(created_at AS TEXT) ILIKE ?", kw),
		)
	}

	if filter.StatusID != nil {
		query = query.Where("status_id = ?", *filter.StatusID)
	}

	if filter.PriorityID != nil {
		query = query.Where("priority_id = ?", *filter.PriorityID)
	}

	if filter.SLADateFrom != nil {
		query = query.Where("sla_date >= ?", *filter.SLADateFrom)
	}
	if filter.SLADateTo != nil {
		query = query.Where("sla_date <= ?", *filter.SLADateTo)
	}

	if err := query.Limit(limit).Offset(offset).Order("created_at desc").Find(&cases).Error; err != nil {
		return nil, err
	}

	return cases, nil
}

func (r *authRepo) CountCasesWithFilter(c *gin.Context, filter model.CaseFilter) (int, error) {
	var count int64
	query := r.DB.WithContext(c).Model(&model.Cases{})

	if filter.Keyword != "" {
		kw := "%" + strings.TrimSpace(filter.Keyword) + "%"
		query = query.Where(
			r.DB.Where("title ILIKE ?", kw).
				Or("customer_id ILIKE ?", kw).
				Or("created_by ILIKE ?", kw).
				Or("CAST(sla_date AS TEXT) ILIKE ?", kw).
				Or("CAST(created_at AS TEXT) ILIKE ?", kw),
		)
	}

	if filter.StatusID != nil {
		query = query.Where("status_id = ?", *filter.StatusID)
	}

	if filter.PriorityID != nil {
		query = query.Where("priority_id = ?", *filter.PriorityID)
	}

	if filter.SLADateFrom != nil {
		query = query.Where("sla_date >= ?", *filter.SLADateFrom)
	}

	if filter.SLADateTo != nil {
		query = query.Where("sla_date <= ?", *filter.SLADateTo)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *authRepo) CreateNoteType(c *gin.Context, note model.NoteTypes) (*model.NoteTypes, error) {
	if err := r.DB.WithContext(c).Create(&note).Error; err != nil {
		return nil, err
	}
	return &note, nil
}

func (r *authRepo) GetCaseByID(c *gin.Context, id string) (*model.CaseWithInitialDescriptions, error) {
	var selectedCase model.Cases
	if err := r.DB.WithContext(c).First(&selectedCase, "id = ?", id).Error; err != nil {
		return nil, err
	}

	var descriptions []string
	if err := r.DB.WithContext(c).
		Model(&model.Cases{}).
		Pluck("initial_description", &descriptions).Error; err != nil {
		return nil, err
	}

	return &model.CaseWithInitialDescriptions{
		Cases:               selectedCase,
		InitialDescriptions: descriptions,
	}, nil
}
