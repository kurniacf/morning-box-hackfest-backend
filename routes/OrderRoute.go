package routes

import (
	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/handler"
	"morning-box-hackfest-be/repository"
	"morning-box-hackfest-be/service"

	"github.com/gin-gonic/gin"
)

func AddOrderRoutes(r *gin.Engine) {
	firestoreClient := config.GetFirestoreClient()
	orderRepo := repository.NewOrderRepository(firestoreClient)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	orderGroup := r.Group("/orders")
	{
		orderGroup.GET("/", orderHandler.GetAllOrders)
		orderGroup.GET("/:id", orderHandler.GetOrder)
		orderGroup.POST("/", orderHandler.CreateOrder)
		orderGroup.PUT("/:id", orderHandler.UpdateOrder)
		orderGroup.DELETE("/:id", orderHandler.DeleteOrder)
		orderGroup.POST("/confirm/:id", orderHandler.ConfirmOrderADayBefore)
	}

}
