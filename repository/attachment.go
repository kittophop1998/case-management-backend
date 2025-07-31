package repository

import (
	"case-management/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a authRepo) UploadAttachment(c *gin.Context, caseID uuid.UUID, file model.Attachment) (uuid.UUID, error) {
	tx := a.DB.WithContext(c.Request.Context()).Create(&file)
	return file.ID, tx.Error
}
