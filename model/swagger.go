package model

import "github.com/google/uuid"

type CreateUserResponse struct {
	Data struct {
		ID uuid.UUID `json:"id"`
	} `json:"data"`
}

type UserRequest struct {
	AgentID    string    `json:"agentId" example:"12337"`
	UserName   string    `json:"userName" example:"Janet Adebayo"`
	Email      string    `json:"email" example:"Janet@exam.com"`
	Team       string    `json:"team" example:"Inbound"`
	IsActive   bool      `json:"isActive" example:"true"`
	OperatorID string    `json:"operatorId" example:"1233"`
	CenterID   uuid.UUID `json:"centerId" example:"b94eee08-8324-4d4f-b166-d82775553a7e"`
	RoleID     uuid.UUID `json:"roleId" example:"538cd6c5-4cb3-4463-b7d5-ac6645815476"`
}

type DeleteUserResponse struct {
	Message string `json:"message" example:"User deleted successfully"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Invalid ID"`
}
