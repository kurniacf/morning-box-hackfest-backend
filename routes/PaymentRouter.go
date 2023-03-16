package routes

import (
	"morning-box-hackfest-be/handler"
	"morning-box-hackfest-be/service"

	"github.com/gin-gonic/gin"
)

func AddPaymentRoutes(r *gin.Engine) {
	paymentService := service.NewPaymentService()
	paymentHandler := handler.NewPaymentHandler(paymentService)

	payment := r.Group("/payment")
	{
		payment.POST("/create", paymentHandler.CreateTransaction)
		//payment.POST("/webhook", paymentHandler.HandleMidtransWebhook)
	}
}
