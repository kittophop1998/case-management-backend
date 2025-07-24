package repository

import (
	"case-management/appcore/appcore_config"
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"context"
	"time"

	"github.com/golang-jwt/jwt"
)

func (r *authRepo) SaveAccressLog(ctx context.Context, accessLog model.AccessLogs) error {

	r.DB.Save(&accessLog)

	return nil
}

func (a *authRepo) GenerateToken(ttl time.Duration, metadata *appcore_model.Metadata) (signedToken string, err error) {
	claims := &appcore_model.JwtClaims{
		Metadata: *metadata,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(ttl).Unix(),
			Issuer:    "casemanagement",
		},
	}

	secretKey := []byte(appcore_config.Config.SecretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
