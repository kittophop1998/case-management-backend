package handler

import (
	"case-management/appcore/appcore_config"
	"case-management/appcore/appcore_internal/appcore_model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (h *Handler) MiddlewareCheckRefreshToken() gin.HandlerFunc {
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
