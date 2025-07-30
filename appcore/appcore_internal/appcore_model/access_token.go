package appcore_model

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Metadata struct {
	TokenId       string    `json:"tokenId"`
	UserId        uuid.UUID `json:"userId"`
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
