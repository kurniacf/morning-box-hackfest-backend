package routes

import (
	"github.com/gin-gonic/gin"

	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/handler"
	"morning-box-hackfest-be/repository"
	"morning-box-hackfest-be/service"
)

func AddAuthRoutes(r *gin.Engine) {
	firestoreClient := config.GetFirestoreClient()
	authRepo := repository.NewAuthRepository(firestoreClient)
	userRepo := repository.NewUserRepository(firestoreClient)
	authService := service.NewAuthService(authRepo)
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(authService, userService)

	r.POST("/signin", authHandler.Signin)
	r.POST("/signup", authHandler.Signup)
}
