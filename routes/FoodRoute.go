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

	foodRepo := repository.NewFoodRepository(firestoreClient)
	foodService := service.NewFoodService(foodRepo)
	foodHandler := handler.NewFoodHandler(foodService)

	foodGroup.GET("", foodHandler.GetAllFoods)
	foodGroup.GET("/:id", foodHandler.GetFood)
	foodGroup.POST("", foodHandler.CreateFood)
	foodGroup.PUT("/:id", foodHandler.UpdateFood)
	foodGroup.DELETE("/:id", foodHandler.DeleteFood)
}
