package main

import (
	"fmt"
	"log"
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
	var port string

	if os.Getenv("APP_ENV") == "production" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	} else {
		port = ":8080"
	}

	config.LoadEnv()
	authClient, firestoreClient, err := config.SetupFirebase()
	if err != nil {
		log.Fatalf("Failed to setup Firebase: %v", err)
	}

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("firebaseAuth", authClient)
		c.Set("firestoreClient", firestoreClient)
		c.Next()
	})

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

	r.Run(port)
}
