package middleware

import (
	"context"
	"morning-box-hackfest-be/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware : to verify all authorized operations
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		firebaseAuth := config.GetAuthClient()

		authorizationToken := c.GetHeader("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
		if idToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id token not available"})
			c.Abort()
			return
		}
		//verify token
		token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("token", token.UID)
		c.Next()
	}
}
