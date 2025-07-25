package appcore_handler

import (
	"case-management/appcore/appcore_cache"
	"case-management/appcore/appcore_config"
	"case-management/appcore/appcore_internal/appcore_model"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	errorInvalidTokenMsg string = "invalid token"
)

func validateToken(signedToken string) (*appcore_model.JwtClaims, error) {
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

func MiddlewareCheckAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("authorization")
		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid token",
			})
			return
		}

		accessToken := strings.Split(authorization, "Bearer ")
		if len(accessToken) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid token",
			})
			return
		}

		//validate token
		claims, err := validateToken(accessToken[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		//check token in redis
		exists, err := appcore_cache.Cache.Exists(c, accessToken[1]).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("failed to check access token: %s", err.Error()),
			})
			return
		}
		if exists != 1 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "your token was revoked, please try to login again",
			})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("roles", claims.Roles)
		c.Set("access_token", accessToken[1])
		c.Set("username", claims.Username)

		c.Next()
	}
}

func MiddlewareCheckRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": errorInvalidTokenMsg,
			})
			return
		}

		tokenParts := strings.Split(authorization, "Bearer ")
		if len(tokenParts) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": errorInvalidTokenMsg,
			})
			return
		}

		refreshToken := tokenParts[1]
		if refreshToken == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": errorInvalidTokenMsg,
			})
			return
		}

		token, err := jwt.ParseWithClaims(
			refreshToken,
			&appcore_model.JwtClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(appcore_config.Config.SecretKey), nil
			},
		)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		claims, ok := token.Claims.(*appcore_model.JwtClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "couldn't parse claims or token is invalid",
			})
			return
		}

		if claims.ExpiresAt < time.Now().Unix() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token expired",
			})
			return
		}

		c.Set("refresh_token_claims", claims)
		c.Next()
	}
}
