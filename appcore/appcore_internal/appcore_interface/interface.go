package inf

import (
	"case-management/appcore/appcore_router"

	"gorm.io/gorm"
)

type Handler struct {
	Store   *gorm.DB
	Version string
}

type Module interface {
	ModuleAPI(r *appcore_router.Router)
	// GrpcServer() *grpc.Server
}
