package usecase

import (
	"case-management/appcore/appcore_config"
	"case-management/utils"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"case-management/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go"
)

func (u *UseCase) UploadAttachment(c *gin.Context, files []*multipart.FileHeader, caseID, userID uuid.UUID) error {
	for _, file := range files {
		// ดึงนามสกุลไฟล์
		extension := filepath.Ext(file.Filename)

		// สุ่มชื่อไฟล์ใหม่
		randomName, err := utils.RandStringRunes(10)
		if err != nil {
			return err
		}

		newFileName := fmt.Sprintf("/case/%s/%s%s", caseID.String(), randomName, extension)
		fileContent, err := file.Open()
		if err != nil {
			return err
		}
		defer fileContent.Close()

		// ข้อมูลการอัปโหลด
		bucket := appcore_config.Config.MinioBucketName
		fileSize := file.Size

		_, err = u.Storage.PutObject(bucket, newFileName, fileContent, fileSize, minio.PutObjectOptions{})
		if err != nil {
			return err
		}

		attachment := model.Attachment{
			CaseId:           caseID,
			FileName:         file.Filename,
			FilePath:         newFileName,
			FileType:         file.Header.Get("Content-Type"),
			FileSizeBytes:    uint64(fileSize),
			UploadedByUserId: userID,
			UploadedAt:       time.Now(),
		}

		id, err := u.caseManagementRepository.UploadAttachment(c, caseID, attachment)
		if err != nil {
			return err
		}

		if err := u.CreateAuditLog(c, userID, model.EventCreated, id.String(), "case_attachment", attachment); err != nil {
			return err
		}
	}
	return nil
}

func (u *UseCase) GetFile(c *gin.Context, objectName string) (*minio.Object, string, error) {
	bucketName := appcore_config.Config.MinioBucketName
	contentType := ""
	object, err := u.Storage.GetObject(bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return object, contentType, err
	}
	return object, contentType, nil
}
