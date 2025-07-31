package appcore

import (
	"case-management/appcore/appcore_cache"
	"case-management/appcore/appcore_config"
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/appcore/appcore_logger"
	"case-management/appcore/appcore_router"
	"case-management/appcore/appcore_storage"
	"case-management/appcore/appcore_store"
	"flag"

	"context"
	"net/http"
	"time"

	_ "case-management/docs"

	requestID "github.com/sumit-tembe/gin-requestid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type Service struct {
	ServiceName string
	Version     string
	ApiHandler  *appcore_handler.ApiHandler

	restAPIServer *http.Server
	// observeServer func(ctx context.Context) error
}

func init() {
	appcore_model.Port = new(string)
	appcore_model.Port = flag.String("port", "", "your service port")

	appcore_model.IP = new(string)
	appcore_model.IP = flag.String("ip", "", "your service ip")

	flag.Parse()
	if *appcore_model.Port == "" {
		*appcore_model.Port = "8000"
	}

	if *appcore_model.IP == "" {
		*appcore_model.IP = "0.0.0.0"
	}

	appcore_config.InitConfigurations()
	appcore_logger.InitLogger()
	appcore_store.InitPostgresDBStore(appcore_logger.Logger)
	appcore_cache.InitCache(appcore_logger.Logger)
	appcore_storage.InitStorage()

	// appcore_audit_log.InitAuditLog(appcore_store.DBStore)

}

func NewService(serviceName, version string, apiHandler *appcore_handler.ApiHandler) *Service {

	appcore_logger.Logger.Info(">>>>> service : " + serviceName + " " + version + " <<<<<")

	return &Service{
		ServiceName: serviceName,
		Version:     version,
		ApiHandler:  apiHandler,
	}
}

func (s *Service) Run() {
	s.restAPIServer = initGinAPI(s.ApiHandler)
}

func (s *Service) Stop() {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//shutdown api
	appcore_logger.Logger.Info("shutdown API.")
	if err := s.restAPIServer.Shutdown(timeoutCtx); err != nil {
		appcore_logger.Logger.Error(err.Error())
	}

	//shutdown db
	if appcore_store.DBStore != nil {
		appcore_logger.Logger.Info("disconnect DB")
		sqlDB, err := appcore_store.DBStore.DB()
		if err != nil {
			appcore_logger.Logger.Error(err.Error())
		} else {
			err = sqlDB.Close()
			if err != nil {
				appcore_logger.Logger.Error(err.Error())
			}
		}
	}

	//shutdown message broker
	// if appcore_message_broker.MessageBroker != nil {
	// 	appcore_logger.Logger.Info("disconnect MessageBroker")
	// 	appcore_message_broker.MessageBroker.Close()
	// }

	//shutdown cache
	if appcore_cache.Cache != nil {
		appcore_logger.Logger.Info("disconnect Cache")
		err := appcore_cache.Cache.Close()
		if err != nil {
			appcore_logger.Logger.Error(err.Error())
		}
	}

	appcore_logger.Logger.Info("server was shutdown")
}

func initGinAPI(h *appcore_handler.ApiHandler) *http.Server {
	if appcore_config.Config.GinIsReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := appcore_router.New()

	r.Use(gin.LoggerWithConfig(requestID.GetLoggerConfig(nil, nil, nil)))

	v := r.Group("/api")
	v.GET("/ping", h.HealthCheck)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h.Module.ModuleAPI(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"service": h.ServiceName + " service " + h.Version,
			"code":    "PAGE_NOT_FOUND",
			"message": "Page not found"})
	})
	return r.ListenAndServe(*appcore_model.IP, *appcore_model.Port)
}
