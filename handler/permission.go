package handler

import (
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPermissionsWithRoles(c *gin.Context) {
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

	data, total, err := h.UseCase.GetAllPermissionsWithRoles(c, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appcore_handler.NewResponseError(
			"Unable to fetch permission data",
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, appcore_model.NewPaginatedResponse(data, page, limit, total))
}

func (h *Handler) UpdatePermissionRoles(c *gin.Context) {
	var req model.UpdatePermissionRolesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UseCase.UpdatePermissionRoles(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated successfully."})
}
