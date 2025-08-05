package model

type DeleteUserResponse struct {
	Message string `json:"message" example:"User deleted successfully"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Invalid ID"`
}
