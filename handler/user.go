package handler

import (
	"case-management/appcore/appcore_config"
	"case-management/appcore/appcore_handler"
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided JSON body
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.CreateUserRequest true "User data"
// @Success 201 {object} model.CreateUserResponse
// @Router /users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var user model.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		appcore_handler.HandleError(c, appcore_config.ErrBadRequest)
		return
	}

	id, err := h.UseCase.CreateUser(c, &user)
	if err != nil {
		appcore_handler.HandleError(c, appcore_config.ErrInternalServer.WithDetails(err.Error()))
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
		appcore_handler.HandleError(c, appcore_config.ErrBadRequest)
		return
	}

	page, err := getPage(c)
	if err != nil {
		appcore_handler.HandleError(c, appcore_config.ErrBadRequest)
		return
	}

	sort := c.DefaultQuery("sort", "is_active desc")
	keyword := c.Query("keyword")
	roleIDStr := c.Query("roleId")
	teamIDStr := c.Query("teamId")
	centerIDStr := c.Query("centerId")

	isActiveStr := c.Query("is_active")
	var isActive *bool = nil
	if isActiveStr != "" {
		val := isActiveStr == "true"
		isActive = &val
	}

	var roleID, teamID, centerID uuid.UUID
	if roleIDStr != "" {
		if id, err := uuid.Parse(roleIDStr); err == nil {
			roleID = id
		}
	}
	if teamIDStr != "" {
		if id, err := uuid.Parse(teamIDStr); err == nil {
			teamID = id
		}
	}
	if centerIDStr != "" {
		if id, err := uuid.Parse(centerIDStr); err == nil {
			centerID = id
		}
	}

	filter := model.UserFilter{
		Keyword:  keyword,
		Sort:     sort,
		IsActive: isActive,
		RoleID:   roleID,
		TeamID:   teamID,
		CenterID: centerID,
	}

	users, total, err := h.UseCase.GetAllUsers(c, page, limit, filter)
	if err != nil {
		appcore_handler.HandleError(c, err)
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
// @Param        id   path  string  true  "User ID"
// @Success      200  {object}  model.User
// @Router       /users/{id} [get]
func (h *Handler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	uid, err := uuid.Parse(idParam)
	if err != nil {
		appcore_handler.HandleError(c, appcore_config.ErrBadRequest)
		return
	}

	user, err := h.UseCase.GetUserByID(c, uid)
	if err != nil {
		appcore_handler.HandleError(c, appcore_config.ErrNotFound)
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
	_, err := uuid.Parse(idParam)
	if err != nil {
		appcore_handler.HandleError(c, appcore_config.ErrBadRequest.WithMessage(appcore_config.Message{
			Th: "รหัสผู้ใช้ไม่ถูกต้อง",
			En: "Invalid user ID",
		}))
		return
	}

	err = h.UseCase.DeleteUserByID(c, idParam)
	if err != nil {
		appcore_handler.HandleError(c, appcore_config.ErrInternalServer.WithMessage(appcore_config.Message{
			Th: "ไม่สามารถลบผู้ใช้ได้",
			En: "Failed to delete user",
		}))
		return
	}

	// c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	c.JSON(http.StatusOK, appcore_handler.NewResponseObject(
		model.StatusResponse{
			Status: "success",
		},
	))
}

// UpdateUser godoc
// @Summary Update user by ID
// @Description Update user information by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body model.CreateUserRequest true "User data"
// @Success 200 {object} model.CreateUserResponse
// @Router /users/{id} [put]
func (h *Handler) UpdateUserByID(c *gin.Context) {
	var input model.UserFilter
	if err := c.ShouldBindJSON(&input); err != nil {
		appcore_handler.HandleError(c, appcore_config.ErrBadRequest)
		return
	}

	idParam := c.Param("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		appcore_handler.HandleError(c, appcore_config.ErrBadRequest.WithMessage(appcore_config.Message{
			Th: "รหัสผู้ใช้ไม่ถูกต้อง",
			En: "Invalid user ID",
		}))
		return
	}

	err = h.UseCase.UpdateUser(c, userID, input)
	if err != nil {
		appcore_handler.HandleError(c, appcore_config.ErrInternalServer.WithMessage(appcore_config.Message{
			Th: "ไม่สามารถอัปเดตผู้ใช้ได้",
			En: "Failed to update user",
		}))
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

// ImportCSV godoc
// @Summary Import users from CSV file
// @Description Import user data from a CSV file asynchronously and track progress via task ID
// @Tags Users
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV file to upload"
// @Param taskID query string false "Optional task ID to track import progress"
// @Success 202 {object} model.MessageResponse
// @Router /users/import [post]
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

// 	taskID := c.Query("taskID")
// 	if taskID == "" {
// 		taskID = uuid.NewString()
// 	}

// 	cCopy := c.Copy()
// 	go func() {
// 		err := h.UseCase.ImportUsersFromCSVWithProgress(cCopy, src, taskID)
// 		if err != nil {
// 			log.Printf("Import error: %v", err)
// 			utils.SetProgress(taskID, 100)
// 		}
// 	}()

// 	c.JSON(http.StatusAccepted, gin.H{
// 		"message": "import started",
// 		"taskID":  taskID,
// 	})

// }

// func (h *Handler) GetImportStatus(c *gin.Context) {
// 	taskID := c.Query("taskID")
// 	log.Printf("GetImportStatus called with taskID=%s", taskID)
// 	if taskID == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "missing taskID"})
// 		return
// 	}

// 	status := utils.GetImportStatus(taskID)
// 	c.JSON(http.StatusOK, status)
// }
