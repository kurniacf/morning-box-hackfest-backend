package routes

import (
	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/handler"
	"morning-box-hackfest-be/repository"
	"morning-box-hackfest-be/service"

	"github.com/gin-gonic/gin"
)

func AddFoodRoutes(r *gin.Engine) {
	foodGroup := r.Group("/foods")

	firestoreClient := config.GetFirestoreClient()

	repository := repository.NewFoodRepository(firestoreClient)
	service := service.NewFoodService(repository)
	handler := handler.NewFoodHandler(service)

	foodGroup.GET("", handler.GetAllFoods)
	foodGroup.GET("/:id", handler.GetFood)
	foodGroup.POST("", handler.CreateFood)
	foodGroup.PUT("/:id", handler.UpdateFood)
	foodGroup.DELETE("/:id", handler.DeleteFood)
}
