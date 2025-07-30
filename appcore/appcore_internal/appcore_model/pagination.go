package appcore_model

import (
	"context"
	"math"

	"gorm.io/gorm"
)

type Pagination[T any] struct {
	Limit      int    `json:"limit,omitempty;query:limit"`
	Page       int    `json:"page,omitempty;query:page"`
	Sort       string `json:"sort,omitempty;query:sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	Rows       []T    `json:"rows"`
}

func (p *Pagination[T]) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination[T]) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination[T]) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination[T]) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}

type AdvanceSQL struct {
	SelectQuery    string
	ConditionQuery string
	JoinQuery      string
	Scopes         func(db *gorm.DB) *gorm.DB
}

func Paginate[T any](ctx context.Context, modelPointer *[]T, pagination *Pagination[T], db *gorm.DB, s AdvanceSQL) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	if s.Scopes != nil {
		db.WithContext(ctx).Model(modelPointer).Scopes(s.Scopes).Joins(s.JoinQuery).Select(s.SelectQuery).Where(s.ConditionQuery).Count(&totalRows)
	} else {
		db.WithContext(ctx).Model(modelPointer).Joins(s.JoinQuery).Select(s.SelectQuery).Where(s.ConditionQuery).Count(&totalRows)
	}

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		if s.Scopes != nil {
			return db.WithContext(ctx).Offset(pagination.GetOffset()).Scopes(s.Scopes).Joins(s.JoinQuery).Limit(pagination.GetLimit()).Order(pagination.GetSort())
		} else {
			return db.WithContext(ctx).Offset(pagination.GetOffset()).Joins(s.JoinQuery).Limit(pagination.GetLimit()).Order(pagination.GetSort())
		}
	}
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Total      int         `json:"total"`
	TotalPages int         `json:"totalPages"`
}

func NewPaginatedResponse(data interface{}, page, limit, total int) PaginatedResponse {
	totalPages := (total + limit - 1) / limit // ปัดขึ้น
	return PaginatedResponse{
		Data:       data,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}
}
