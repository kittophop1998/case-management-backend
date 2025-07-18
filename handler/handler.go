package handler

import (
	"case-management/usecase"
	"log/slog"
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
