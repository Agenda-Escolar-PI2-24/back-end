package pkg

import (
	"agenda-escolar/internal/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var userService services.UserService

// AuthenticationMiddleware checks if the user has a valid JWT token
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})
			return
		}

		// The token should be prefixed with "Bearer "
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
			return
		}

		tokenString = tokenParts[1]

		claims, err := VerifyToken(tokenString)
		username, ok := claims["username"].(string)
		exists, _ := userService.Exists(username)
		id, _ := claims["user_id"].(float64)
		if err != nil || !exists || id == 0 || !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
			return
		}

		c.Set("username", claims["username"])
		c.Set("user_id", int(id))
		c.Next()
	}
}
