package handler

import (
	"case-management/appcore/appcore_handler"
	"case-management/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {

	var userData model.LoginRequest
	reqID := c.GetHeader("X-Request-ID")

	if reqID == "" {
		appcore_handler.HandleError(c, appcore_handler.ErrFilterRequired)
		return
	}

	if err := c.ShouldBindJSON(&userData); err != nil {
		appcore_handler.HandleError(c, appcore_handler.ErrBadRequest)
		return
	}

	if userData.UserData == "" {
		appcore_handler.HandleError(c, appcore_handler.ErrRequiredParam)
		return
	}

	resp, err := h.UseCase.Login(c.Request.Context(), userData)
	if err != nil {

		username := ""
		if resp != nil {
			if resp.User.Username != "" {
				username = resp.User.Username
			}
		}

		accessLog := model.AccessLogs{
			Username:      username,
			LogonDatetime: time.Now(),
			LogonResult:   "failed",
		}

		// err := h.authUsecase.SaveAccessLog(c.Request.Context(), accessLog)
		if err := h.UseCase.SaveAccessLog(c.Request.Context(), accessLog); err != nil {
			appcore_handler.HandleError(c, err)
			return
		}

		appcore_handler.HandleError(c, err)
		return
	}

	accessLog := model.AccessLogs{
		Username:      resp.User.Username,
		LogonDatetime: time.Now(),
		LogonResult:   "success",
	}

	if err := h.UseCase.SaveAccessLog(c.Request.Context(), accessLog); err != nil {
		appcore_handler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
