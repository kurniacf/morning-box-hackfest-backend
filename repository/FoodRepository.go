package repository

import (
	"context"
	"morning-box-hackfest-be/model"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type FoodRepositoryInterface interface {
	GetAllFoods() ([]model.FoodResponse, error)
	GetFood(id string) (model.FoodResponse, error)
	CreateFood(food model.Food) (string, error)
	UpdateFood(id string, food model.Food) error
	DeleteFood(id string) error
}

type foodRepository struct {
	client *firestore.Client
}

func NewFoodRepository(client *firestore.Client) *foodRepository {
	return &foodRepository{
		client: client,
	}
}

func (r *foodRepository) GetAllFoods() ([]model.FoodResponse, error) {
	var foods []model.FoodResponse

	iter := r.client.Collection("foods").Documents(context.Background())
	defer iter.Stop()

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var food model.FoodResponse
		if err = doc.DataTo(&food); err != nil {
			return nil, err
		}

		foods = append(foods, food)
	}

	return foods, nil
}

func (r *foodRepository) GetFood(id string) (model.FoodResponse, error) {
	doc, err := r.client.Collection("foods").Doc(id).Get(context.Background())
	if err != nil {
		return model.FoodResponse{}, err
	}

	var food model.FoodResponse
	if err = doc.DataTo(&food); err != nil {
		return model.FoodResponse{}, err
	}

	return food, nil
}

func (r *foodRepository) CreateFood(food model.Food) (string, error) {
	doc, _, err := r.client.Collection("foods").Add(context.Background(), food)
	if err != nil {
		return "", err
	}

	return doc.ID, nil
}

func (r *foodRepository) UpdateFood(id string, food model.Food) error {
	_, err := r.client.Collection("foods").Doc(id).Update(context.Background(), []firestore.Update{
		{
			Path:  "name",
			Value: food.Name,
		},
		{
			Path:  "description",
			Value: food.Description,
		},
		{
			Path:  "imageUrl",
			Value: food.ImageURL,
		},
		{
			Path:  "calories",
			Value: food.Calories,
		},
		{
			Path:  "type",
			Value: food.Type,
		},
	})
	// _, err := r.client.Collection("foods").Doc(id).Set(context.Background(), food)
	if err != nil {
		return err
	}
	return nil
}

func (r *foodRepository) DeleteFood(id string) error {
	_, err := r.client.Collection("foods").Doc(id).Delete(context.Background())
	if err != nil {
		return err
	}

	return nil
}
