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

func MiddlewareCheckAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("access_token")
		if err != nil || token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errorInvalidTokenMsg})
			return
		}

		claims, err := parseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// TODO: ตรวจสอบ token ใน Redis เพื่อเช็คว่า token นี้ยัง valid อยู่หรือถูก revoked หรือไม่
		// ตัวอย่าง:
		// exists, err := appcore_cache.Cache.Exists(c, token).Result()
		// if err != nil || exists != 1 {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token revoked or invalid"})
		// 	return
		// }

		c.Set("userId", claims.UserId)
		c.Set("username", claims.Username)
		c.Next()
	}
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
