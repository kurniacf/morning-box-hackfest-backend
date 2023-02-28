package handler

import (
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type foodHandler struct {
	service service.FoodServiceInterface
}

func NewFoodHandler(service service.FoodServiceInterface) *foodHandler {
	return &foodHandler{service}
}

func (h *foodHandler) GetAllFoods(c *gin.Context) {
	foods, err := h.service.GetAllFoods()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": foods,
	})
}

func (h *foodHandler) GetFood(c *gin.Context) {
	id := c.Param("id")

	food, err := h.service.GetFood(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": food,
	})
}

func (h *foodHandler) CreateFood(c *gin.Context) {
	var (
		foodRequest model.FoodRequest
		food        model.Food
		err         error
	)

	err = c.ShouldBindJSON(&foodRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	food = convertToFoodModel(foodRequest)

	createdFoodID, err := h.service.CreateFood(food)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data created successfully",
		"data":    createdFoodID,
	})
}

func (h *foodHandler) UpdateFood(c *gin.Context) {
	var (
		foodRequest model.FoodRequest
		food        model.Food
		id          string
		err         error
	)

	id = c.Param("id")

	err = c.ShouldBindJSON(&foodRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	food = convertToFoodModel(foodRequest)

	err = h.service.UpdateFood(id, food)
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

func (h *foodHandler) DeleteFood(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteFood(id)
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

func convertToFoodModel(food model.FoodRequest) model.Food {
	return model.Food{
		Name:        food.Name,
		Description: food.Description,
		ImageURL:    food.ImageURL,
	}
}
