package middleware

import (
	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/repository"
	"morning-box-hackfest-be/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware : to verify all authorized operations
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationToken := c.GetHeader("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
		if idToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id token not available"})
			c.Abort()
			return
		}

		// get user
		firestoreClient := config.GetFirestoreClient()
		userRepo := repository.NewUserRepository(firestoreClient)
		userService := service.NewUserService(userRepo)

		user, err := userService.GetUser(idToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
