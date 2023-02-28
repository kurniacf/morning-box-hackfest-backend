package middleware

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware : to verify all authorized operations
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)

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
