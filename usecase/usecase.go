package usecase

import (
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"case-management/services/mailer"
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type UseCase struct {
	// mu                      sync.Mutex
	Cache                    *redis.Client
	Logger                   *slog.Logger
	caseManagementRepository CaseManagementRepository
	Storage                  *minio.Client
	Mail                     mailer.Email
}

type CaseManagementRepository interface {
	// User
	CreateUser(c *gin.Context, user *model.CreateUserRequest) (uuid.UUID, error)
	GetAllUsers(c *gin.Context, limit, offset int, filter model.UserFilter) ([]*model.User, error)
	GetUserByID(c *gin.Context, id uuid.UUID) (*model.User, error)
	DeleteUserByID(c *gin.Context, id string) error
	CountUsers(c *gin.Context) (int, error)
	CountUsersWithFilter(c *gin.Context, filter model.UserFilter) (int, error)
	UpdateUser(c *gin.Context, userID uuid.UUID, input model.UserFilter) error
	SaveAccessLog(ctx context.Context, accessLog model.AccessLogs) error
	GetUserByUserName(c *gin.Context, username string) (*model.User, error)
	GenerateToken(ttl time.Duration, metadata *appcore_model.Metadata) (signedToken string, err error)
	BulkInsertUsers(c context.Context, users []model.User) error
	StoreToken(c *gin.Context, accessToken string) error
	ValidateToken(signedToken string) (claims *appcore_model.JwtClaims, err error)
	DeleteToken(c *gin.Context, accessToken string) error
	GetAllLookups(ctx *gin.Context) (map[string]interface{}, error)
	GetAllPermissionsWithRoles(ctx *gin.Context, limit, offset int) ([]model.PermissionWithRolesResponse, error)
	UpdatePermissionRoles(ctx *gin.Context, req model.UpdatePermissionRolesRequest) error
	CountPermissions(ctx *gin.Context) (int, error)

	// Case Management
	CreateCase(ctx *gin.Context, c *model.Cases) (uuid.UUID, error)
	GetAllCases(c *gin.Context, limit, offset int, filter model.CaseFilter) ([]*model.Cases, error)
	CountCasesWithFilter(c *gin.Context, filter model.CaseFilter) (int, error)
	CreateNoteType(c *gin.Context, note model.NoteTypes) (*model.NoteTypes, error)
	GetCaseByID(c *gin.Context, id uuid.UUID) (*model.Cases, error)
	AddInitialDescription(c *gin.Context, caseID uuid.UUID, newDescription string) error
	GetNoteTypeById(c *gin.Context, noteTypeID uuid.UUID) (*model.NoteTypes, error)
	CreateCustomerNote(c *gin.Context, note *model.CustomerNote) error

	// Attachment
	UploadAttachment(c *gin.Context, caseID uuid.UUID, file model.Attachment) (uuid.UUID, error)

	// Audit Log
	CreateAuditLog(c *gin.Context, log model.AuditLog) error

	// API Log
	SaveLog(log *model.ApiLogs) error
	GetAllLogs(c *gin.Context) ([]model.ApiLogs, error)
}

func New(
	caseManagementRepository CaseManagementRepository,
	cache *redis.Client,
	logger *slog.Logger,
	storage *minio.Client,
	mail mailer.Email,
) *UseCase {
	return &UseCase{
		caseManagementRepository: caseManagementRepository,
		Cache:                    cache,
		Logger:                   logger,
		Storage:                  storage,
		Mail:                     mail,
	}
}
