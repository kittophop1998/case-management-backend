package handler

import (
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_router"
)

func (h *Handler) ModuleAPI(r *appcore_router.Router) {
	api := r.Engine.Group("/api/v1")

	// Auth routes
	authRoutes := api.Group("/auth")
	authRoutes.Use(h.APILogger())

	{
		authRoutes.POST("/login", h.Login)
		authRoutes.GET("/profile", appcore_handler.MiddlewareCheckAccessToken(), h.Profile)
		authRoutes.POST("/logout", appcore_handler.MiddlewareCheckAccessToken(), h.Logout)
	}

	// API Logs Routes
	apiLogsRoutes := api.Group("api_logs")
	{
		apiLogsRoutes.GET("", h.GetAPILogs)
	}

	// Look Up Routes
	userLookUpsRoutes := api.Group("/lookups")
	userLookUpsRoutes.Use(h.APILogger())

	{
		userLookUpsRoutes.GET("", h.GetAllLookups)
	}

	// Permission Routes
	permissionsWithRolesRoutes := api.Group("/permissions")
	permissionsWithRolesRoutes.Use(h.APILogger())
	{
		permissionsWithRolesRoutes.GET("", h.GetPermissionsWithRoles)
		permissionsWithRolesRoutes.PATCH("update", h.UpdatePermissionRoles)
	}

	// User routes
	userRoutes := api.Group("/users")
	userRoutes.Use(
		h.APILogger(),
		appcore_handler.MiddlewareCheckAccessToken(),
	)
	{
		userRoutes.POST("", h.CreateUser)
		userRoutes.GET("", h.GetAllUsers)
		userRoutes.GET("/:id", h.GetUserByID)
		userRoutes.PUT("/:id", h.UpdateUserByID)
		userRoutes.DELETE("/:id", h.DeleteUserByID)
		// userRoutes.POST("/import", h.ImportCSV)
	}

	// Attachment routes
	attachmentsRoutes := api.Group("/attachment")
	attachmentsRoutes.Use(h.APILogger())
	attachmentsRoutes.Use(appcore_handler.MiddlewareCheckAccessToken())
	{
		attachmentsRoutes.POST("/:case_id", h.UploadAttachment)
		attachmentsRoutes.GET("/file/*objectName", h.GetFile)
	}

	caseManagementRoutes := api.Group("/cases")
	caseManagementRoutes.Use(
		h.APILogger(),
		appcore_handler.MiddlewareCheckAccessToken(),
	)
	{
		caseManagementRoutes.POST("", h.CreateCase)
		caseManagementRoutes.GET("", h.GetAllCases)
		caseManagementRoutes.POST("/note_type", h.CreateNoteType)
		caseManagementRoutes.GET("/:id", h.GetCaseByID)
		caseManagementRoutes.POST("/add-initial-description", h.AddInitialDescription)
		caseManagementRoutes.GET("/note_type/:id", h.GetNoteTypeById)
		caseManagementRoutes.POST("/customer/note", h.CreateCustomerNote)
	}

	customerRoutes := api.Group("/customers")
	customerRoutes.Use(
		h.APILogger(),
		appcore_handler.MiddlewareCheckAccessToken(),
	)
	{
		customerRoutes.GET("/search", h.CustomerSearch)
		// customerRoutes.POST("/note", h.CreateCustomer)
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
