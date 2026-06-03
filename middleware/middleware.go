package middleware

import (
	"Eccomerce-GO--recap-project-/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        jwt_token, err := c.Cookie("jwt")
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        claims, err := utils.ValidateToken(jwt_token, "access")

        if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"error": "Unauthorized"})
        }

        c.Set("userID", claims.UserID)
        c.Next()
    }
}
