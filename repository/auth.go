package repository

import (
	"case-management/appcore/appcore_config"
	"case-management/appcore/appcore_internal/appcore_model"
	"case-management/model"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (r *authRepo) SaveAccessLog(ctx context.Context, accessLog model.AccessLogs) error {
	return r.DB.WithContext(ctx).Create(&accessLog).Error
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

func (a *authRepo) StoreToken(c *gin.Context, accessToken string) error {
	if err := a.Cache.Set(c.Request.Context(), accessToken, nil, 6*time.Hour).Err(); err != nil {
		return fmt.Errorf("failed to store access token: %w", err)
	}
	return nil
}

func (a *authRepo) ValidateToken(signedToken string) (claims *appcore_model.JwtClaims, err error) {
	secretKey := appcore_config.Config.SecretKey
	token, err := jwt.ParseWithClaims(
		signedToken,
		&appcore_model.JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*appcore_model.JwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil
}

func (a *authRepo) DeleteToken(c *gin.Context, accessToken string) error {
	return a.Cache.Del(c.Request.Context(), accessToken).Err()
}
