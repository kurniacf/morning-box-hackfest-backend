package main

import (
	"fmt"
	"morning-box-hackfest-be/config"
	"net/http"
	"os"

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

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Hackfest 2023",
		})
	})

	r.Run(port)
}
