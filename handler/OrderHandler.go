package handler

import (
	"fmt"
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService   service.OrderServiceInterface
	userService    service.UserServiceInterface
	packageService service.PackageServiceInterface
	foodService    service.FoodServiceInterface
	pointService   service.PointServiceInterface
}

func NewOrderHandler(orderService service.OrderServiceInterface, userService service.UserServiceInterface, packageService service.PackageServiceInterface, foodService service.FoodServiceInterface, pointService service.PointServiceInterface) *orderHandler {
	return &orderHandler{orderService: orderService, userService: userService, packageService: packageService, foodService: foodService, pointService: pointService}
}

func (h *orderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	for _, order := range orders {
		order.Package, _ = h.packageService.GetPackage(order.Package.Id)
		order.User, _ = h.userService.GetUser(order.User.Id)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

func (h *orderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := h.orderService.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	order.Package, _ = h.packageService.GetPackage(order.Package.Id)
	order.User, _ = h.userService.GetUser(order.User.Id)

	order.Package.Foods.Monday, _ = h.foodService.GetFood(order.Package.Foods.Monday.Id)
	order.Package.Foods.Tuesday, _ = h.foodService.GetFood(order.Package.Foods.Tuesday.Id)
	order.Package.Foods.Wednesday, _ = h.foodService.GetFood(order.Package.Foods.Wednesday.Id)
	order.Package.Foods.Thursday, _ = h.foodService.GetFood(order.Package.Foods.Thursday.Id)
	order.Package.Foods.Friday, _ = h.foodService.GetFood(order.Package.Foods.Friday.Id)
	order.Package.Foods.Saturday, _ = h.foodService.GetFood(order.Package.Foods.Saturday.Id)
	order.Package.Foods.Sunday, _ = h.foodService.GetFood(order.Package.Foods.Tuesday.Id)

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

	id, err := h.orderService.CreateOrder(req)
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

	if err := h.orderService.UpdateOrder(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *orderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	err := h.orderService.DeleteOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete package"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package deleted successfully"})
}

func (h *orderHandler) ConfirmOrderADayBefore(c *gin.Context) {
	// hour := time.Now().Hour()

	// if !(hour > 18 && hour < 24) {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Confirmation can only be done between 18:00 to 24:00"})
	// 	return
	// }

	var req model.OrderConfirmRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get Order
	order, err := h.orderService.GetOrder(req.OrderId)
	if err != nil || order == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(order)

	if order.Status != model.BELUM_DIKONFIRMASI {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Order has been confirmed"})
		return
	}

	// Confirm order
	err = h.orderService.ConfirmOrderADayBefore(req.OrderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get User
	user, err := h.userService.GetUser(req.UserId)
	if err != nil || user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add strike point
	err = h.pointService.UpdatePoint(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order confirmed successfully"})
}

func (h *orderHandler) GetActiveOrder(c *gin.Context) {
	user, _ := c.Get("user")

	userId := user.(model.UserResponse).Id

	order, err := h.orderService.GetActiveOrder(userId)
	if err != nil || order == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get active order"})
		return
	}

	order.Package, _ = h.packageService.GetPackage(order.Package.Id)
	order.User, _ = h.userService.GetUser(order.User.Id)

	order.Package.Foods.Monday, _ = h.foodService.GetFood(order.Package.Foods.Monday.Id)
	order.Package.Foods.Tuesday, _ = h.foodService.GetFood(order.Package.Foods.Tuesday.Id)
	order.Package.Foods.Wednesday, _ = h.foodService.GetFood(order.Package.Foods.Wednesday.Id)
	order.Package.Foods.Thursday, _ = h.foodService.GetFood(order.Package.Foods.Thursday.Id)
	order.Package.Foods.Friday, _ = h.foodService.GetFood(order.Package.Foods.Friday.Id)
	order.Package.Foods.Saturday, _ = h.foodService.GetFood(order.Package.Foods.Saturday.Id)
	order.Package.Foods.Sunday, _ = h.foodService.GetFood(order.Package.Foods.Tuesday.Id)

	c.JSON(http.StatusOK, gin.H{"data": order})
}
