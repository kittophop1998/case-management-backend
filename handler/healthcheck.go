package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary      Show the status of server
// @Description  Get server health status
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.MessageResponse
// @Router       /health [get]
func (h *Handler) HealthCheck(c *gin.Context) {
	{
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	}
}
