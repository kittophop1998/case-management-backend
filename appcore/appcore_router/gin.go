package appcore_router

import (
	"case-management/appcore/appcore_logger"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	requestid "github.com/sumit-tembe/gin-requestid"
)

type Router struct {
	*gin.Engine
}

func New() *Router {
	r := gin.New()
	//config := cors.DefaultConfig()
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Time-Zone"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	//config.AllowAllOrigins = true
	//config.AllowCredentials = true
	// config.AddAllowHeaders("authorization")
	r.Use(cors.New(config))

	// if appcore_config.Config.IsObserve {
	// 	r.Use(otelgin.Middleware("Epson"))
	// }

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(requestid.RequestID(nil))

	/*r.Use(limit.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP() // limit rate by client ip
	}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
		return rate.NewLimiter(20, 1), 1 * time.Hour
	}, func(c *gin.Context) {
		c.AbortWithStatus(429) // handle exceed rate limit request
	}))*/
	return &Router{r}
}

func (r *Router) ListenAndServe(ip, port string) *http.Server {
	appcore_logger.Logger.Info("Listening on API", "address", ip+":"+port)
	s := &http.Server{
		Addr:    ip + ":" + port,
		Handler: r,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 10 * time.Second,
		// IdleTimeout:    1 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appcore_logger.Logger.Error("listen: %s", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()
	return s
}
