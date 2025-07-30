package handler

import (
	"case-management/usecase"
	"log/slog"
)

const (
	errorInvalidRequest string = "your request is invalid"
	errorSystem         string = "system error"
	// errorInvalidUsernamePassword string = "invalid username or password"
	// errorInvalidLevelID          string = "invalid level id"
	// errorInvalidLimit            string = "limit number is incorrect"
	// errorInvalidPage             string = "page number is incorrect"
	errorInvalidTokenMsg string = "invalid token"
	// errorInvalidRoleID           string = "invalid role id"
	// errorInvalidPartnerID        string = "invalid partner id"
	// errorInvalidGroupID          string = "invalid group id"
	// errorInvalidCustomerID       string = "invalid customer id"
	// errorInvalidContractID       string = "invalid contract id"
	// errorInvalidModuleID         string = "invalid module id"
	// errorInvalidUserID           string = "invalid user id"
	// errorInvalidMenuID           string = "invalid menu id"
	// errorUserIDNotExist          string = "user id doesn't exist"
	// errorRoleIDNotExist          string = "role id doesn't exist"
	// errorUploadFile              string = "file upload failed"
	// errorPasswordIncorrect       string = "password is incorrect"
)

type Handler struct {
	UseCase *usecase.UseCase
	Logger  *slog.Logger
}

func NewHandler(u *usecase.UseCase, logger *slog.Logger) *Handler {
	return &Handler{
		UseCase: u,
		Logger:  logger,
	}
}
