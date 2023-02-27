package main

import (
	"fmt"
	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
}

func main() {
	port := getPort()

	r := gin.Default()

	firebaseAuth := config.SetupFirebase()

	r.Use(func(c *gin.Context) {
		c.Set("firebaseAuth", firebaseAuth)
	})

	r.Use(middleware.AuthMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Hackfest 2023",
		})
	})

	r.Run(port)
}

func getPort() string {
	if os.Getenv("APP_ENV") == "production" {
		return fmt.Sprintf(":%s", os.Getenv("PORT"))
	}
	return ":8080"
}
