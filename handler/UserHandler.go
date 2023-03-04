package handler

import (
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserServiceInterface
}

func NewUserHandler(service service.UserServiceInterface) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (h *userHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var (
		userRequest model.UserRequest
		err         error
	)
	err = c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUserID, err := h.service.CreateUser(userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data created successfully",
		"data":    createdUserID,
	})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var (
		userRequest model.UserRequest
		id          string
		err         error
	)
	id = c.Param("id")

	err = c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.UpdateUser(id, userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data updated successfully",
	})
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data deleted successfully",
	})
}
