package handler

import (
	"case-management/appcore/appcore_handler"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) UploadAttachment(c *gin.Context) {
	userIdRaw, exists := c.Get("userId")
	if !exists {
		h.Logger.Error("user_id not found in context")
		c.JSON(http.StatusUnauthorized, appcore_handler.NewResponseError("Unauthorized", "user_id not found"))
		return
	}
	userId, ok := userIdRaw.(uuid.UUID)
	if !ok {
		h.Logger.Error("invalid user_id format")
		c.JSON(http.StatusInternalServerError, appcore_handler.NewResponseError("Invalid user_id format", "user_id error"))
		return
	}

	fileInput, err := c.MultipartForm()
	if err != nil {
		h.Logger.Error("Failed to parse form", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, appcore_handler.NewResponseError(err.Error(), "Failed to parse form"))
		return
	}
	files := fileInput.File["file"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, appcore_handler.NewResponseError("No file provided", "file missing"))
		return
	}

	caseIDStr := c.Param("case_id")
	caseID, err := uuid.Parse(caseIDStr)
	if err != nil {
		h.Logger.Error("Invalid case_id", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, appcore_handler.NewResponseError("Invalid case_id", "Invalid UUID"))
		return
	}

	err = h.UseCase.UploadAttachment(c, files, caseID, userId)
	if err != nil {
		h.Logger.Error("Upload failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, appcore_handler.NewResponseError("Upload failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, appcore_handler.NewResponseCreated("File(s) uploaded successfully"))
}
