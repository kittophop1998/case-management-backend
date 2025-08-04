package integration

import (
	"case-management/appcore/appcore_config"
	"case-management/handler"
	"case-management/services/mailer"
	"case-management/tests/integration/mock"
	"fmt"
	"log"
	"strings"

	"case-management/usecase"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func Setup() *gin.Engine {
	// Initialize the mock repository and Redis client
	mockRepo := &mock.MockRepository{}
	mockRedis := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	mockMinio, err := minio.New(
		"localhost:9000",
		"minioadmin",
		"minioadmin",
		false,
	)
	if err != nil {
		log.Fatalf("failed to initialize minio client: %v", err)
	}

	mockeMail := mailer.NewMailTrap(appcore_config.Config.SMTPHost,
		fmt.Sprintf("%s:%s", appcore_config.Config.SMTPHost, appcore_config.Config.SMTPPort),
		appcore_config.Config.SMTPUser,
		appcore_config.Config.SMTPPassword,
	)

	logger := slog.Default()

	u := usecase.New(mockRepo, mockRedis, logger, mockMinio, mockeMail)
	h := handler.NewHandler(u, logger)

	// Gin router
	router := gin.Default()
	router.GET("/users", h.GetAllUsers)
	router.GET("/users/:id", h.GetUserByID)
	router.POST("/users", h.CreateUser)
	router.PUT("/users/:id", h.UpdateUser)
	router.DELETE("/users/:id", h.DeleteUserByID)
	router.GET("/lookups", h.GetAllLookups)

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

func TestGetUserByID_ReturnSuccess(t *testing.T) {
	router := Setup()

	id := uuid.New()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/"+id.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}

func TestCreateUser_ReturnSuccess(t *testing.T) {
	router := Setup()

	payload := `{
		"username": "John Doe",
		"email": "john.doe@example.com",
		"name": "John Doe",
		"team": "CEN123456",
		"isActive": true
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"id"`)
}

func TestDeleteUserByID_ReturnSuccess(t *testing.T) {
	router := Setup()

	id := uuid.New().String()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/"+id, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"status":"success"`)
}

func TestUpdateUser_ReturnSuccess(t *testing.T) {
	router := Setup()

	id := uuid.New()

	updateJSON := `{
		"name": "Updated Name",
		"isActive": false,
		"team": {
		"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
		"name": "Updated Team"
	}
}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/"+id.String(), strings.NewReader(updateJSON))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"message":"user updated successfully"`)
}

func TestGetAllLookups_Success(t *testing.T) {
	router := Setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/lookups", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, `"teams"`)
	assert.Contains(t, body, `"roles"`)
	assert.Contains(t, body, `"centers"`)
	assert.Contains(t, body, `"permissions"`)
}
