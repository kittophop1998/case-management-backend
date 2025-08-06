package handler

import (
	"case-management/appcore/appcore_config"
	"case-management/appcore/appcore_handler"
	"case-management/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Login godoc
// @Summary Login a user
// @Description Login a user with the provided credentials
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} model.LoginResponse
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req model.LoginRequest

	// Validate required header
	if reqID := c.GetHeader("X-Request-ID"); reqID == "" {
		appcore_handler.HandleError(c, appcore_config.ErrFilterRequired)
		return
	}

	// Validate body
	if err := c.ShouldBindJSON(&req); err != nil {
		appcore_handler.HandleError(c, appcore_config.ErrBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		appcore_handler.HandleError(c, appcore_config.ErrRequiredParam)
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

// Logout godoc
// @Summary Logout a user
// @Description Logout a user and clear session
// @Tags Auth
// @Success 200 {object} map[string]string
// @Accept json
// @Produce json
// @Router /auth/logout [post]
func (h *Handler) Logout(c *gin.Context) {
	if err := h.UseCase.Logout(c); err != nil {
		appcore_handler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// Profile godoc
// @Summary Get user profile
// @Description Get the profile of the logged-in user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} model.UserResponse
// @Router /auth/profile [get]
func (h *Handler) Profile(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		appcore_handler.HandleError(c, appcore_config.ErrBadRequest)
		return
	}

	uid, ok := userId.(uuid.UUID)
	if !ok {
		appcore_handler.HandleError(c, appcore_config.ErrBadRequest)
		return
	}

	resp, err := h.UseCase.GetUserByID(c, uid)
	if err != nil {
		appcore_handler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
