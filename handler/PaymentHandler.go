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
}

func NewPaymentHandler(paymentService service.PaymentServiceInterface) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (pc *PaymentHandler) CreateTransaction(c *gin.Context) {
	var order model.OrderResponse
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid order data",
		})
		return
	}

	response, err := pc.paymentService.CreateTransaction(order)
	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to process transaction",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
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
