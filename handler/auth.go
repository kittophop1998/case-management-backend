package handler

import (
	"case-management/appcore/appcore_handler"
	"case-management/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Login godoc
// @Summary      Login
// @Description  User login endpoint
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        X-Request-ID  header  string  true  "Request ID"
// @Param        body  body  model.LoginRequest  true  "Login request"
// @Success      200  {object}  model.LoginResponse
func (h *Handler) Login(c *gin.Context) {
	var req model.LoginRequest

	// Validate required header
	if reqID := c.GetHeader("X-Request-ID"); reqID == "" {
		appcore_handler.HandleError(c, appcore_handler.ErrFilterRequired)
		return
	}

	// Validate body
	if err := c.ShouldBindJSON(&req); err != nil {
		appcore_handler.HandleError(c, appcore_handler.ErrBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		appcore_handler.HandleError(c, appcore_handler.ErrRequiredParam)
		return
	}

	// Main login use case
	resp, err := h.UseCase.Login(c, req)
	success := err == nil

	// Log access (even on failure)
	_ = h.UseCase.SaveAccessLog(c.Request.Context(), req.Username, success)

	if err != nil {
		appcore_handler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Logout(c *gin.Context) {
	if err := h.UseCase.Logout(c); err != nil {
		appcore_handler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (h *Handler) Profile(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		appcore_handler.HandleError(c, appcore_handler.ErrBadRequest)
		return
	}

	uid, ok := userId.(uuid.UUID)
	if !ok {
		appcore_handler.HandleError(c, appcore_handler.ErrBadRequest)
		return
	}

	resp, err := h.UseCase.GetUserByID(c, uid)
	if err != nil {
		appcore_handler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
