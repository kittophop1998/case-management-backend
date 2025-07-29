package handler

import (
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_router"
)

func (h *Handler) ModuleAPI(r *appcore_router.Router) {
	api := r.Engine.Group("/api/v1")

	// Auth routes
	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("/login", h.Login)
		authRoutes.GET("/profile", appcore_handler.MiddlewareCheckAccessToken(), h.Profile)
	}

	healthcheckRoutes := api.Group("/health")
	{
		healthcheckRoutes.GET("", h.HealthCheck)
	}

	userLookUpsRoutes := api.Group("/lookups")
	{
		userLookUpsRoutes.GET("", h.GetAllLookups)
	}

	// User routes
	userRoutes := api.Group("/users")
	userRoutes.Use(appcore_handler.MiddlewareCheckAccessToken())
	{
		userRoutes.POST("", h.CreateUser)
		userRoutes.GET("", h.GetAllUsers)
		userRoutes.GET("/:id", h.GetUserByID)
		userRoutes.PUT("/:id", h.UpdateUser)
		userRoutes.DELETE("/:id", h.DeleteUserByID)
		// userRoutes.POST("/import", h.ImportCSV)
	}

	// Refresh token api
	// refreshTokenAPI := api.Group("/")
	// refreshTokenAPI.Use(normalRateLimiter, h.MiddlewareCheckRefreshToken())
	// {
	// 	refreshTokenAPI.POST("/refresh", h.RefreshAccessToken)
	// }

	// // 300 request / min
	// normalRateLimiter := rate_limiter.NewRateLimiter(appcore_cache.Cache, &rate_limiter.RateLimit{
	// 	Rate:  time.Minute,
	// 	Limit: 100,
	// })

	// // 5 request / min
	// otpRateLimiter := rate_limiter.NewRateLimiter(appcore_cache.Cache, &rate_limiter.RateLimit{
	// 	Rate:  time.Minute,
	// 	Limit: 5,
	// })

	// rateLimiter := rate_limiter.NewRateLimiter(appcore_cache.Cache, &rate_limiter.RateLimit{
	// 	Rate:  time.Minute,
	// 	Limit: 5,
	// })
}
