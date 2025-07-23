package model

type StatusResponse struct {
	Status string `json:"status"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type AccessTokenRequest struct {
	Access_token string `json:"access_token" binding:"required"`
}
