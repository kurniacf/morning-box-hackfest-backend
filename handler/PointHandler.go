package handler

import (
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pointHandler struct {
	service service.PointServiceInterface
}

func NewPointHandler(service service.PointServiceInterface) *pointHandler {
	return &pointHandler{service}
}

func (h *pointHandler) GetAllPoints(c *gin.Context) {
	points, err := h.service.GetAllPoints()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": points,
	})
}

func (h *pointHandler) GetPoint(c *gin.Context) {
	id := c.Param("id")
	point, err := h.service.GetPoint(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": point,
	})
}

func (h *pointHandler) CreatePoint(c *gin.Context) {
	var (
		pointRequest model.StrikePointRequest
		err          error
	)
	err = c.ShouldBindJSON(&pointRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdPointID, err := h.service.CreatePoint(pointRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data created successfully",
		"data":    createdPointID,
	})
}

func (h *pointHandler) UpdatePoint(c *gin.Context) {
	var (
		pointRequest model.StrikePointRequest
		id           string
		err          error
	)
	id = c.Param("id")
	err = c.ShouldBindJSON(&pointRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.UpdatePoint(id, pointRequest)
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

func (h *pointHandler) DeletePoint(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeletePoint(id)
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
