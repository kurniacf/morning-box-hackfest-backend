package handler

import (
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	signInService service.AuthServiceInterface
	signUpService service.UserServiceInterface
}

func NewAuthHandler(authService service.AuthServiceInterface, userService service.UserServiceInterface) *authHandler {
	return &authHandler{
		signInService: authService,
		signUpService: userService,
	}
}

func (h *authHandler) Signin(c *gin.Context) {
	var authSignin model.AuthSignin
	err := c.ShouldBindJSON(&authSignin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.signInService.Signin(authSignin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *authHandler) Signup(c *gin.Context) {
	var userRequest model.UserRequest
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := h.signUpService.CreateUser(userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := h.signUpService.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}
