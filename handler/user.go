package handler

import (
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"case-management/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		c.JSON(http.StatusBadRequest, appcore_handler.NewResponseError(
			err.Error(),
			errorSystem,
		))
		return
	}

	id, err := h.UseCase.CreateUser(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appcore_handler.NewResponseError(
			err.Error(),
			"error",
		))
		return
	}

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
	limit, err := getLimit(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, appcore_handler.NewResponseError(err.Error(), errorInvalidRequest))
		return
	}

	page, err := getPage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, appcore_handler.NewResponseError(err.Error(), errorInvalidRequest))
		return
	}

	sort := c.DefaultQuery("sort", "created_at desc")
	role := c.Query("role")
	team := c.Query("team")
	center := c.Query("center")

	isActiveStr := c.Query("is_active")
	var isActive *bool = nil

	if isActiveStr != "" {
		val := isActiveStr == "true"
		isActive = &val
	}

	filter := model.UserFilter{
		IsActive: isActive,
		Sort:     sort,
		Role:     role,
		Team:     team,
		Center:   center,
	}

	users, total, err := h.UseCase.GetAllUsers(c, page, limit, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appcore_handler.NewResponseError(err.Error(), errorSystem))
		return
	}

	c.JSON(http.StatusOK, appcore_model.NewPaginatedResponse(users, page, limit, total))
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
		c.JSON(http.StatusBadRequest, appcore_handler.NewResponseError(
			err.Error(),
			errorInvalidRequest,
		))
		return
	}

	user, err := h.UseCase.GetUserByID(c, idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, appcore_handler.NewResponseError(
			err.Error(),
			"user not found",
		))
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
		c.JSON(http.StatusBadRequest, appcore_handler.NewResponseError(
			err.Error(),
			errorInvalidRequest,
		))
		return
	}

	err = h.UseCase.DeleteUserByID(c, idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appcore_handler.NewResponseError(
			err.Error(),
			errorSystem,
		))
		return
	}

	// c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	c.JSON(http.StatusOK, appcore_handler.NewResponseObject(
		model.StatusResponse{
			Status: "success",
		},
	))
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var input model.UserFilter
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, appcore_handler.NewResponseError(
			err.Error(),
			errorInvalidRequest,
		))
		return
	}

	idParam := c.Param("id")
	userID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, appcore_handler.NewResponseError("invalid user ID", "invalid_request"))
		return
	}

	err = h.UseCase.UpdateUser(c, uint(userID), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appcore_handler.NewResponseError(err.Error(), "error"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

// func (h *Handler) ImportCSV(c *gin.Context) {
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot get file"})
// 		return
// 	}

// 	src, err := file.Open()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot open file"})
// 		return
// 	}
// 	defer src.Close()

// 	err = h.UseCase.ImportUsersFromCSV(c, src)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "import success"})
// }

func (h *Handler) ImportCSV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot get file"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot open file"})
		return
	}
	defer src.Close()

	taskID := c.Query("taskID")
	if taskID == "" {
		taskID = uuid.NewString()
	}

	cCopy := c.Copy()
	go func() {
		err := h.UseCase.ImportUsersFromCSVWithProgress(cCopy, src, taskID)
		if err != nil {
			log.Printf("Import error: %v", err)
			utils.SetProgress(taskID, 100)
		}
	}()

	c.JSON(http.StatusAccepted, gin.H{
		"message": "import started",
		"taskID":  taskID,
	})

}
