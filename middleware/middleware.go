package middleware

import (
	"net/http"
	"strings"

	"Eccomerce-GO--recap-project-/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token from Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    "unauthorized",
				"message": "Authorization header required",
			})
			return
		}

		// Extract Bearer token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    "unauthorized",
				"message": "Invalid authorization header format. Use: Bearer <token>",
			})
			return
		}

		token := parts[1]
		claims, err := utils.ValidateToken(token, "access")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    "unauthorized",
				"message": "Invalid or expired token",
			})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
