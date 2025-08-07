package handler

import (
	"case-management/appcore/appcore_config"
	"case-management/appcore/appcore_handler"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CustomerSearch(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		appcore_handler.HandleError(c, appcore_config.ErrBadRequest.WithDetails(appcore_config.Message{
			Th: "กรุณาระบุ ID ของลูกค้า",
			En: "Please specify the customer ID",
		}))
		return
	}

	c.JSON(200, gin.H{
		"message": "Customer search successful",
	})
}

func (h *Handler) CustomerDashBoard(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		appcore_handler.HandleError(c, appcore_config.ErrRequiredParam.WithDetails(appcore_config.Message{
			Th: "กรุณาระบุ ID ของลูกค้า",
			En: "Please specify the customer ID",
		}))
		return
	}

	c.JSON(200, gin.H{
		"message": "Customer dashboard retrieved successfully",
	})
}
