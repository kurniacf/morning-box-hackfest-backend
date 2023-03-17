// File: handler/PaymentHandler.go

package handler

import (
	"fmt"
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentControllerInterface interface {
	CreateTransaction(c *gin.Context)
	HandleMidtransWebhook(c *gin.Context)
}

type PaymentHandler struct {
	paymentService service.PaymentServiceInterface
	packageService service.PackageServiceInterface
}

func NewPaymentHandler(paymentService service.PaymentServiceInterface, packageService service.PackageServiceInterface) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService, packageService: packageService}
}

func (h *PaymentHandler) CreateTransaction(c *gin.Context) {
	var order model.OrderPaymentResponse
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid order data",
		})
		return
	}

	pkg, err := h.packageService.GetPackage(order.PackageId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	chargeReq := h.paymentService.GenerateSnapReq(order, pkg)

	redirectURL, err := h.paymentService.CreateTransaction(chargeReq)
	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("Failed to process transaction: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"redirect_url": redirectURL,
	})
}

// func (ph *PaymentHandler) HandleMidtransWebhook(c *gin.Context) {
// 	var notification midtrans.TransactionNotification
// 	if err := c.BindJSON(&notification); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success": false,
// 			"message": "Invalid webhook data",
// 		})
// 		return
// 	}

// 	err := ph.paymentService.ProcessWebhookNotification(notification)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"success": false,
// 			"message": "Failed to process webhook notification",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "Webhook notification processed",
// 	})
// }
