package handler

import (
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing case ID"})
		return
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	caseData, err := h.UseCase.GetCaseByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": caseData})
}

func (h *Handler) AddInitialDescription(c *gin.Context) {
	var req model.AddInitialDescriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	caseUUID, err := uuid.Parse(req.CaseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid case ID"})
		return
	}

	if err := h.UseCase.AddInitialDescription(c, caseUUID, req.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "initial description added"})
}

func (h *Handler) GetNoteTypeById(c *gin.Context) {
	noteTypeIDStr := c.Param("id")
	noteTypeID, err := uuid.Parse(noteTypeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, appcore_handler.NewResponseError("Invalid UUID", "INVALID_ID"))
		return
	}

	noteType, err := h.UseCase.GetNoteTypeById(c, noteTypeID)
	if err != nil {
		c.JSON(http.StatusNotFound, appcore_handler.NewResponseError("NoteType not found", "NOT_FOUND"))
		return
	}

	c.JSON(http.StatusOK, appcore_handler.NewResponseObject(noteType))
}

func (h *Handler) CreateCustomerNote(c *gin.Context) {
	var req model.CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ดึง customerID แบบ UUID โดยตรง
	customerIDValue, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "customer_id not found"})
		return
	}

	customerID, ok := customerIDValue.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid customer_id type in context"})
		return
	}

	noteTypeUUID, err := uuid.Parse(req.NoteTypeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid noteTypeId"})
		return
	}

	// ส่ง customerID เข้า usecase ด้วย
	if err := h.UseCase.CreateCustomerNote(c, customerID, noteTypeUUID, req.Note); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "note created successfully"})
}
