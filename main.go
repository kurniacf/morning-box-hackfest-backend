package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	var port string

	if os.Getenv("APP_ENV") == "production" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	} else {
		port = ":8080"
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Hackfest 2023",
		})
	})

	r.Run(port)
}
