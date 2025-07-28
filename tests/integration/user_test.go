package integration

import (
	"case-management/handler"
	"case-management/test/integration/mock"
	"case-management/usecase"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func Setup() *gin.Engine {
	// Initialize the mock repository and Redis client
	mockRepo := &mock.MockRepository{}
	mockRedis := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	logger := slog.Default()

	u := usecase.New(mockRepo, mockRedis, logger)
	h := handler.NewHandler(u, logger)

	// Gin router
	router := gin.Default()
	router.GET("/users", h.GetAllUsers)
	return router
}

func TestGetAllUsers_ReturnSuccess(t *testing.T) {
	router := Setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}
