package handler

import (
	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/repository"
	"morning-box-hackfest-be/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	service service.OrderServiceInterface
}

func NewOrderHandler(service service.OrderServiceInterface) *orderHandler {
	return &orderHandler{service: service}
}

func (h *orderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	firestoreClient := config.GetFirestoreClient()

	packageRepo := repository.NewPackageRepository(firestoreClient)
	packageService := service.NewPackageService(packageRepo)

	userRepo := repository.NewUserRepository(firestoreClient)
	userService := service.NewUserService(userRepo)

	for _, order := range orders {
		order.Package, _ = packageService.GetPackage(order.Package.Id)
		order.User, _ = userService.GetUser(order.User.Id)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

func (h *orderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := h.service.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	firestoreClient := config.GetFirestoreClient()

	packageRepo := repository.NewPackageRepository(firestoreClient)
	packageService := service.NewPackageService(packageRepo)

	userRepo := repository.NewUserRepository(firestoreClient)
	userService := service.NewUserService(userRepo)

	foodRepo := repository.NewFoodRepository(firestoreClient)
	foodService := service.NewFoodService(foodRepo)

	order.Package, _ = packageService.GetPackage(order.Package.Id)
	order.User, _ = userService.GetUser(order.User.Id)

	order.Package.Foods.Monday, _ = foodService.GetFood(order.Package.Foods.Monday.Id)
	order.Package.Foods.Tuesday, _ = foodService.GetFood(order.Package.Foods.Tuesday.Id)
	order.Package.Foods.Wednesday, _ = foodService.GetFood(order.Package.Foods.Wednesday.Id)
	order.Package.Foods.Thursday, _ = foodService.GetFood(order.Package.Foods.Thursday.Id)
	order.Package.Foods.Friday, _ = foodService.GetFood(order.Package.Foods.Friday.Id)
	order.Package.Foods.Saturday, _ = foodService.GetFood(order.Package.Foods.Saturday.Id)
	order.Package.Foods.Sunday, _ = foodService.GetFood(order.Package.Foods.Tuesday.Id)

	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	var req model.OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.CreateOrder(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *orderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	var req model.OrderUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateOrder(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *orderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete package"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package deleted successfully"})
}

func (h *orderHandler) ConfirmOrderADayBefore(c *gin.Context) {
	hour := time.Now().Hour()

	if !(hour > 18 && hour < 24) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Confirmation can only be done between 18:00 to 24:00"})
		return
	}

	id := c.Param("id")

	err := h.service.ConfirmOrderADayBefore(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Confirmation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order confirmed successfully"})
}
