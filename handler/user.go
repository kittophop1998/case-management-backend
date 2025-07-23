package handler

import (
	"case-management/appcore/appcore_handler"
	"case-management/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided JSON body
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User data"
// @Success 201 {object} model.CreateUserResponse
// @Router /users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.UseCase.CreateUser(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusCreated, gin.H{"message": "User created", "userId": id})
	c.JSON(http.StatusCreated, appcore_handler.NewResponseCreated(id))
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} model.User
// @Router  /users [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.UseCase.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appcore_handler.NewResponseObject(users))
}

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Retrieve user information by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  model.User
// @Router       /users/{id} [get]
func (h *Handler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := h.UseCase.GetUserByID(c, idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, appcore_handler.NewResponseObject(user))
}

// DeleteUserByID godoc
// @Summary Delete user by ID
// @Description Delete a user by its ID
// @Tags Users
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {object} model.DeleteUserResponse
// @Router /users/{id} [delete]
func (h *Handler) DeleteUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.UseCase.DeleteUserByID(c, idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	c.JSON(http.StatusOK, appcore_handler.NewResponseObject(
		model.StatusResponse{
			Status: "success",
		},
	))
}
