package routes

import (
	"github.com/gin-gonic/gin"

	"morning-box-hackfest-be/handler"
)

func AddAuthRoutes(r *gin.Engine) {
	r.POST("/signin", handler.Signin)
}
