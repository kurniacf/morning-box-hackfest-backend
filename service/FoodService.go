package service

import (
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/repository"
)

type FoodServiceInterface interface {
	GetAllFoods() ([]model.FoodResponse, error)
	GetFood(id string) (model.FoodResponse, error)
	CreateFood(food model.FoodRequest) (string, error)
	UpdateFood(id string, food model.FoodRequest) error
	DeleteFood(id string) error
}

type foodService struct {
	repo repository.FoodRepositoryInterface
}

func NewFoodService(repo repository.FoodRepositoryInterface) *foodService {
	return &foodService{repo}
}

func (s *foodService) GetAllFoods() ([]model.FoodResponse, error) {
	foods, err := s.repo.GetAllFoods()
	if err != nil {
		return nil, err
	}

	return foods, nil
}

func (s *foodService) GetFood(id string) (model.FoodResponse, error) {
	food, err := s.repo.GetFood(id)
	if err != nil {
		return model.FoodResponse{}, err
	}

	return food, nil
}

func (s *foodService) CreateFood(food model.FoodRequest) (string, error) {
	newFood := model.Food{
		Name:        food.Name,
		Description: food.Description,
		ImageURL:    food.ImageURL,
		Calories:    food.Calories,
		Type:        food.Type,
	}

	newFoodRequest := model.FoodRequest{
		Name:        newFood.Name,
		Description: newFood.Description,
		ImageURL:    newFood.ImageURL,
		Calories:    newFood.Calories,
		Type:        newFood.Type,
	}

	foodId, err := s.repo.CreateFood(newFoodRequest)
	if err != nil {
		return "", err
	}

	return foodId, nil
}

func (s *foodService) UpdateFood(id string, food model.FoodRequest) error {
	updateFood := model.Food{
		Name:        food.Name,
		Description: food.Description,
		ImageURL:    food.ImageURL,
		Calories:    food.Calories,
		Type:        food.Type,
	}

	updateFoodRequest := model.FoodRequest{
		Name:        updateFood.Name,
		Description: updateFood.Description,
		ImageURL:    updateFood.ImageURL,
		Calories:    updateFood.Calories,
		Type:        updateFood.Type,
	}

	err := s.repo.UpdateFood(id, updateFoodRequest)
	if err != nil {
		return err
	}

	return nil
}

func (s *foodService) DeleteFood(id string) error {
	err := s.repo.DeleteFood(id)
	if err != nil {
		return err
	}

	return nil
}
