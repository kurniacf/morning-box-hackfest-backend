package main

import (
	"fmt"
	"net/http"
	"os"

	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/repository"
	"morning-box-hackfest-be/routes"
	"morning-box-hackfest-be/service"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.InitMidtrans()
}

func main() {
	port := getPort()

	r := gin.Default()

	// Home route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Hackfest 2023",
		})
	})

	firestoreClient := config.GetFirestoreClient()
	packageRepo := repository.NewPackageRepository(firestoreClient)
	userRepo := repository.NewUserRepository(firestoreClient)

	paymentService := service.NewPaymentService(packageRepo, userRepo)
	packageService := service.NewPackageService(packageRepo)

	// Auth routes
	routes.AddAuthRoutes(r)

	// Food routes
	routes.AddFoodRoutes(r)

	// Package routes
	routes.AddPackageRoutes(r)

	// User routes
	routes.AddUserRoutes(r)

	// Order routes
	routes.AddOrderRoutes(r)

	// Payment routes
	routes.AddPaymentRoutes(r, paymentService, packageService)

	r.Run(port)
}

func getPort() string {
	if os.Getenv("PORT") != "" {
		return fmt.Sprintf(":%s", os.Getenv("PORT"))
	}
	return ":8080"
}
