package repository

import (
	"context"
	"morning-box-hackfest-be/model"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type SurveyRepositoryInterface interface {
	GetAllSurveys() ([]model.SurveyResponse, error)
	GetSurvey(id string) (model.SurveyResponse, error)
	CreateSurvey(survey model.SurveyRequest) (string, error)
	DeleteSurvey(id string) error
}

type surveyRepository struct {
	client *firestore.Client
}

func NewSurveyRepository(client *firestore.Client) *surveyRepository {
	return &surveyRepository{
		client: client,
	}
}

func (r *surveyRepository) GetAllSurveys() ([]model.SurveyResponse, error) {
	var surveys []model.SurveyResponse

	iter := r.client.Collection("surveys").Documents(context.Background())
	defer iter.Stop()

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var survey model.SurveyResponse
		if err = doc.DataTo(&survey); err != nil {
			return nil, err
		}

		survey.Id = doc.Ref.ID
		surveys = append(surveys, survey)
	}

	return surveys, nil
}

func (r *surveyRepository) GetSurvey(id string) (model.SurveyResponse, error) {
	doc, err := r.client.Collection("surveys").Doc(id).Get(context.Background())
	if err != nil {
		return model.SurveyResponse{}, err
	}

	var survey model.SurveyResponse
	if err = doc.DataTo(&survey); err != nil {
		return model.SurveyResponse{}, err
	}

	survey.Id = doc.Ref.ID
	return survey, nil
}

func (r *surveyRepository) CreateSurvey(survey model.SurveyRequest) (string, error) {
	doc, _, err := r.client.Collection("surveys").Add(context.Background(), survey)
	if err != nil {
		return "", err
	}

	return doc.ID, nil
}

func (r *surveyRepository) DeleteSurvey(id string) error {
	_, err := r.client.Collection("surveys").Doc(id).Delete(context.Background())
	if err != nil {
		return err
	}

	return nil
}
