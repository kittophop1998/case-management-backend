package appcore_handler

import (
	"case-management/appcore/appcore_config"
	"case-management/appcore/appcore_internal/appcore_model"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	errorInvalidTokenMsg = "invalid token"
)

func MiddlewareCheckAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractBearerToken(c)
		if err != nil {
			HandleError(c, ErrBadRequest)
			return
		}

		claims, err := parseToken(token)
		if err != nil {
			HandleError(c, ErrBadRequest)
			return
		}

		// Optional: Check Redis blacklist
		// if isRevoked(token) {
		//     c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token has been revoked"})
		//     return
		// }

		c.Set("userId", claims.UserId)
		c.Set("username", claims.Username)
		c.Next()
	}
}

func parseToken(tokenString string) (*appcore_model.JwtClaims, error) {
	secretKey := appcore_config.Config.SecretKey
	token, err := jwt.ParseWithClaims(tokenString, &appcore_model.JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*appcore_model.JwtClaims)
	if !ok || !token.Valid {
		return nil, errors.New("couldn't parse claims or token is invalid")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

func extractBearerToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("authorization header format must be Bearer {token}")
	}

	return parts[1], nil
}

func MiddlewareCheckRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errorInvalidTokenMsg})
			return
		}

		refreshToken := strings.TrimPrefix(authHeader, "Bearer ")
		if refreshToken == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errorInvalidTokenMsg})
			return
		}

		claims, err := parseToken(refreshToken)
		if err != nil {
			status := http.StatusBadRequest
			if err.Error() == "token expired" {
				status = http.StatusUnauthorized
			}
			c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
			return
		}

		c.Set("refresh_token_claims", claims)
		c.Next()
	}
}
