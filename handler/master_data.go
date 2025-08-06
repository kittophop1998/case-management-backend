package handler

import (
	"case-management/appcore/appcore_config"
	"case-management/appcore/appcore_handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllLookups godoc
// @Summary Get all lookup values
// @Description Get all teams, roles, centers, and permissions
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} appcore_handler.ResponseObject
// @Router /lookups [get]
func (h *Handler) GetAllLookups(c *gin.Context) {
	data, err := h.UseCase.GetAllLookups(c)
	if err != nil {
		appcore_handler.HandleError(c, appcore_config.NewAppError(
			"LOOKUP_ERROR",
			appcore_config.Message{Th: "ไม่สามารถดึงข้อมูล lookup ได้", En: "Failed to fetch lookup data"},
			http.StatusInternalServerError,
			err,
		))
		return
	}

	c.JSON(http.StatusOK, appcore_handler.NewResponseObject(data))
}
