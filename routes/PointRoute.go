package routes

import (
	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/handler"
	"morning-box-hackfest-be/repository"
	"morning-box-hackfest-be/service"

	"github.com/gin-gonic/gin"
)

func AddPointRoutes(r *gin.Engine) {
	firestoreClient := config.GetFirestoreClient()
	pointRepo := repository.NewPointRepository(firestoreClient)
	pointService := service.NewPointService(pointRepo)
	pointHandler := handler.NewPointHandler(pointService)
	pointGroup := r.Group("/points")
	{
		pointGroup.GET("", pointHandler.GetAllPoints)
		pointGroup.GET("/:id", pointHandler.GetPoint)
		pointGroup.POST("", pointHandler.CreatePoint)
		pointGroup.PUT("/:id", pointHandler.UpdatePoint)
		pointGroup.DELETE("/:id", pointHandler.DeletePoint)
	}
}
