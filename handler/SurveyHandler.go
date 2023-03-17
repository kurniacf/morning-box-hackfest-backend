package handler

import (
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type surveyHandler struct {
	service service.SurveyServiceInterface
}

func NewSurveyHandler(service service.SurveyServiceInterface) *surveyHandler {
	return &surveyHandler{service}
}

func (h *surveyHandler) GetAllSurveys(c *gin.Context) {
	surveys, err := h.service.GetAllSurveys()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": surveys,
	})
}

func (h *surveyHandler) GetSurvey(c *gin.Context) {
	id := c.Param("id")

	survey, err := h.service.GetSurvey(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": survey,
	})
}

func (h *surveyHandler) CreateSurvey(c *gin.Context) {
	var (
		surveyRequest model.SurveyRequest
		err           error
	)
	err = c.ShouldBindJSON(&surveyRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdSurveyID, err := h.service.CreateSurvey(surveyRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data created successfully",
		"data":    createdSurveyID,
	})
}

func (h *surveyHandler) DeleteSurvey(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteSurvey(id)
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
