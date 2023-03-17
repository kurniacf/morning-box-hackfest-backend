package routes

import (
	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/handler"
	"morning-box-hackfest-be/repository"
	"morning-box-hackfest-be/service"

	"github.com/gin-gonic/gin"
)

func AddSurveyRoutes(r *gin.Engine) {
	firestoreClient := config.GetFirestoreClient()
	surveyRepo := repository.NewSurveyRepository(firestoreClient)
	surveyService := service.NewSurveyService(surveyRepo)
	surveyHandler := handler.NewSurveyHandler(surveyService)
	surveyGroup := r.Group("/surveys")
	{
		surveyGroup.GET("", surveyHandler.GetAllSurveys)
		surveyGroup.GET("/:id", surveyHandler.GetSurvey)
		surveyGroup.POST("", surveyHandler.CreateSurvey)
		surveyGroup.DELETE("/:id", surveyHandler.DeleteSurvey)
	}
}
