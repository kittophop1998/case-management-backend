package appcore_model

import (
	"github.com/google/uuid"

	"github.com/golang-jwt/jwt"
)

type Metadata struct {
	TokenID       string    `json:"token_id"`
	UserID        uuid.UUID `json:"user_id"`
	Email         string    `json:"email"`
	PartnerID     uint      `json:"partner_id"`
	ConsentStatus uint      `json:"consent_status"`
	Username      string    `json:"username"`
	Owner         string    `json:"owner"`
	Permissions   []string  `json:"permissions"`
	Roles         []string  `json:"roles"`
}

type JwtClaims struct {
	jwt.StandardClaims
	Metadata
}
