package routes

import (
	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/handler"
	"morning-box-hackfest-be/repository"
	"morning-box-hackfest-be/service"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(r *gin.Engine) {
	firestoreClient := config.GetFirestoreClient()
	userRepo := repository.NewUserRepository(firestoreClient)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userGroup := r.Group("/users")
	{
		userGroup.GET("", userHandler.GetAllUsers)
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.POST("", userHandler.CreateUser)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}
}
