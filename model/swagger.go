package model

type CreateUserResponse struct {
	Message string `json:"message"`
	UserID  int    `json:"userId"`
}

type DeleteUserResponse struct {
	Message string `json:"message" example:"User deleted successfully"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Invalid ID"`
}
