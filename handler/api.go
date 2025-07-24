package handler

import (
	"case-management/appcore/appcore_router"
)

const (
	// @todo remove this one (no need to use constant for url path [un-readable])
	roleURL    = "/roles"
	roleID     = "/roles/:id"
	userID     = "/users/:id"
	groupID    = "/groups/:id"
	costID     = "/costs/:id"
	customerID = "/customers/:id"
)

func (h *Handler) ModuleAPI(r *appcore_router.Router) {
	api := r.Engine.Group("/api/v1")

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

	secureAPI := api.Group("/")
	// secureAPI.Use(normalRateLimiter, appcore_handler.MiddlewareCheckAccessToken())
	{
		//user management
		secureAPI.POST("/users", h.CreateUser)
		secureAPI.GET("/users", h.GetAllUsers)
		secureAPI.GET("/users/:id", h.GetUserByID)
		secureAPI.DELETE("/users/:id", h.DeleteUserByID)
		secureAPI.PUT("/users/:id", h.UpdateUser)

	}

	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("/login", h.Login)
	}

	//refresh token api
	// refreshTokenAPI := api.Group("/")
	// refreshTokenAPI.Use(normalRateLimiter, h.MiddlewareCheckRefreshToken())
	// {
	// refreshTokenAPI.POST("/refresh", h.RefreshAccessToken)
	// }

}
