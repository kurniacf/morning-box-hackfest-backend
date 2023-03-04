package main

import (
	"fmt"
	"net/http"
	"os"

	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
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

	// Auth routes
	routes.AddAuthRoutes(r)

	// Food routes
	routes.AddFoodRoutes(r)

	// Package routes
	routes.AddPackageRoutes(r)

	// User routes
	routes.AddUserRoutes(r)

	// Point routes
	routes.AddPointRoutes(r)

	r.Run(port)
}

func getPort() string {
	if os.Getenv("APP_ENV") == "production" {
		return fmt.Sprintf(":%s", os.Getenv("PORT"))
	}
	return ":8080"
}
