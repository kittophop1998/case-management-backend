package handler

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine, userHandler *UserHandler) {
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetAllUsers)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.DELETE("/users/:id", userHandler.DeleteUserByID)
}
