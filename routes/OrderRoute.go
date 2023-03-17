package routes

import (
	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/handler"
	"morning-box-hackfest-be/middleware"
	"morning-box-hackfest-be/repository"
	"morning-box-hackfest-be/service"

	"github.com/gin-gonic/gin"
)

func AddOrderRoutes(r *gin.Engine) {
	firestoreClient := config.GetFirestoreClient()

	orderRepo := repository.NewOrderRepository(firestoreClient)
	orderService := service.NewOrderService(orderRepo)

	packageRepo := repository.NewPackageRepository(firestoreClient)
	packageService := service.NewPackageService(packageRepo)

	foodRepo := repository.NewFoodRepository(firestoreClient)
	foodService := service.NewFoodService(foodRepo)

	userRepo := repository.NewUserRepository(firestoreClient)
	userService := service.NewUserService(userRepo)

	pointRepo := repository.NewPointRepository(firestoreClient)
	pointService := service.NewPointService(pointRepo)

	orderHandler := handler.NewOrderHandler(orderService, userService, packageService, foodService, pointService)

	orderGroup := r.Group("/orders")
	{
		orderGroup.GET("/", orderHandler.GetAllOrders)
		orderGroup.GET("/:id", orderHandler.GetOrder)
		orderGroup.POST("/", orderHandler.CreateOrder)
		orderGroup.PUT("/:id", orderHandler.UpdateOrder)
		orderGroup.DELETE("/:id", orderHandler.DeleteOrder)
		orderGroup.POST("/confirm", orderHandler.ConfirmOrderADayBefore)
		orderGroup.GET("/active", middleware.AuthMiddleware(), orderHandler.GetActiveOrder)
	}

}
