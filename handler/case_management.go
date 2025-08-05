package handler

import (
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCase(c *gin.Context) {
	var caseInput model.Cases

	if err := c.ShouldBindJSON(&caseInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	id, err := h.UseCase.CreateCase(c, &caseInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create case", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Case created successfully", "case_id": id})
}

func (h *Handler) GetAllCases(c *gin.Context) {
	limit, err := getLimit(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, appcore_handler.NewResponseError(err.Error(), errorInvalidRequest))
		return
	}

	page, err := getPage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, appcore_handler.NewResponseError(err.Error(), errorInvalidRequest))
		return
	}

	sort := c.DefaultQuery("sort", "created_at desc")
	keyword := c.Query("keyword")
	statusIDStr := c.Query("statusId")
	priorityIDStr := c.Query("priorityId")
	slaDateFromStr := c.Query("slaDateFrom")
	slaDateToStr := c.Query("slaDateTo")

	layout := "2006-01-02"
	var filter model.CaseFilter

	filter.Keyword = keyword
	filter.Sort = sort

	if statusIDStr != "" {
		if v, err := strconv.ParseUint(statusIDStr, 10, 32); err == nil {
			u := uint(v)
			filter.StatusID = &u
		}
	}
	if priorityIDStr != "" {
		if v, err := strconv.ParseUint(priorityIDStr, 10, 32); err == nil {
			u := uint(v)
			filter.PriorityID = &u
		}
	}

	if slaDateFromStr != "" {
		if t, err := time.Parse(layout, slaDateFromStr); err == nil {
			filter.SLADateFrom = &t
		}
	}
	if slaDateToStr != "" {
		if t, err := time.Parse(layout, slaDateToStr); err == nil {
			filter.SLADateTo = &t
		}
	}

	cases, total, err := h.UseCase.GetAllCases(c, page, limit, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appcore_model.NewPaginatedResponse(cases, page, limit, total))

}

func (h *Handler) CreateNoteType(c *gin.Context) {
	var note model.NoteTypes
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.UseCase.CreateNoteType(c, note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create note type"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetCaseByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing case ID"})
		return
	}

	caseData, err := h.UseCase.GetCaseByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": caseData})
}
