package service

import (
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/repository"
)

type FoodServiceInterface interface {
	GetAllFoods() ([]model.FoodResponse, error)
	GetFood(id string) (model.FoodResponse, error)
	CreateFood(food model.Food) (string, error)
	UpdateFood(id string, food model.Food) error
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

	return food, err
}

func (s *foodService) CreateFood(food model.Food) (string, error) {
	foodId, err := s.repo.CreateFood(food)

	return foodId, err
}

func (s *foodService) UpdateFood(id string, food model.Food) error {
	err := s.repo.UpdateFood(id, food)

	return err
}

func (s *foodService) DeleteFood(id string) error {
	err := s.repo.DeleteFood(id)

	return err
}
