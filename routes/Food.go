package routes

import (
	"morning-box-hackfest-be/handler"

	"github.com/gin-gonic/gin"
)

func AddFoodRoutes(r *gin.Engine) {
	foodGroup := r.Group("/foods")

	foodGroup.GET("", handler.GetAllFoods)
	foodGroup.GET("/:id", handler.GetFood)
	foodGroup.POST("", handler.CreateFood)
	foodGroup.PUT("/:id", handler.UpdateFood)
	foodGroup.DELETE("/:id", handler.DeleteFood)
}
