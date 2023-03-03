package handler

import "github.com/gin-gonic/gin"

func Signin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login sukses",
	})
}
