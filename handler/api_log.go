package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type bodyCaptureWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyCaptureWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (h *Handler) APILogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Clone body for reading
		var requestBody []byte
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBody = bodyBytes
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Re-assign body
		}

		// Capture response
		writer := &bodyCaptureWriter{ResponseWriter: c.Writer, body: bytes.NewBufferString("")}
		c.Writer = writer

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()

		var responsePayload json.RawMessage
		_ = json.Unmarshal(writer.body.Bytes(), &responsePayload)

		var userID uuid.UUID
		if uid, exists := c.Get("user_id"); exists {
			userID, _ = uid.(uuid.UUID)
		}

		_ = h.UseCase.SaveLog(userID, c.Request.Method, c.FullPath(), requestBody, writer.body.Bytes(), uint(status), uint(duration.Milliseconds()), "")
	}
}

func (h *Handler) GetAPILogs(c *gin.Context) {
	logs, err := h.UseCase.GetLogs(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch logs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": logs})
}
