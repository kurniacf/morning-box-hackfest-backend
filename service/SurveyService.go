package service

import (
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/repository"
)

type SurveyServiceInterface interface {
	GetAllSurveys() ([]model.SurveyResponse, error)
	GetSurvey(id string) (model.SurveyResponse, error)
	CreateSurvey(survey model.SurveyRequest) (string, error)
	DeleteSurvey(id string) error
}

type surveyService struct {
	surveyRepo repository.SurveyRepositoryInterface
}

func NewSurveyService(surveyRepo repository.SurveyRepositoryInterface) *surveyService {
	return &surveyService{
		surveyRepo: surveyRepo,
	}
}

func (s *surveyService) GetAllSurveys() ([]model.SurveyResponse, error) {
	return s.surveyRepo.GetAllSurveys()
}

func (s *surveyService) GetSurvey(id string) (model.SurveyResponse, error) {
	return s.surveyRepo.GetSurvey(id)
}

func (s *surveyService) CreateSurvey(survey model.SurveyRequest) (string, error) {
	return s.surveyRepo.CreateSurvey(survey)
}

func (s *surveyService) DeleteSurvey(id string) error {
	return s.surveyRepo.DeleteSurvey(id)
}
