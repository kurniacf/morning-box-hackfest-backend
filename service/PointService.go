package service

import (
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/repository"
)

type PointServiceInterface interface {
	GetAllPoints() ([]model.StrikePointResponse, error)
	GetPoint(id string) (model.StrikePointResponse, error)
	CreatePoint(point model.StrikePointRequest) (string, error)
	UpdatePoint(id string, point model.StrikePointRequest) error
	DeletePoint(id string) error
}

type pointService struct {
	repo repository.PointRepositoryInterface
}

func NewPointService(repo repository.PointRepositoryInterface) *pointService {
	return &pointService{repo}
}

func (s *pointService) GetAllPoints() ([]model.StrikePointResponse, error) {
	points, err := s.repo.GetAllPoints()
	if err != nil {
		return nil, err
	}
	return points, nil
}

func (s *pointService) GetPoint(id string) (model.StrikePointResponse, error) {
	point, err := s.repo.GetPoint(id)
	if err != nil {
		return model.StrikePointResponse{}, err
	}
	return point, nil
}

func (s *pointService) CreatePoint(point model.StrikePointRequest) (string, error) {
	newPoint := model.StrikePoint{
		UserID:   point.UserID,
		Value:    point.Value,
		Duration: point.Duration,
	}
	newPointRequest := model.StrikePointRequest{
		UserID:   newPoint.UserID,
		Value:    newPoint.Value,
		Duration: newPoint.Duration,
	}

	pointId, err := s.repo.CreatePoint(newPointRequest)
	if err != nil {
		return "", err
	}

	return pointId, nil
}

func (s *pointService) UpdatePoint(id string, point model.StrikePointRequest) error {
	updatePoint := model.StrikePoint{
		UserID:   point.UserID,
		Value:    point.Value,
		Duration: point.Duration,
	}
	updatePointRequest := model.StrikePointRequest{
		UserID:   updatePoint.UserID,
		Value:    updatePoint.Value,
		Duration: updatePoint.Duration,
	}

	err := s.repo.UpdatePoint(id, updatePointRequest)
	if err != nil {
		return err
	}

	return nil
}

func (s *pointService) DeletePoint(id string) error {
	err := s.repo.DeletePoint(id)
	if err != nil {
		return err
	}
	return nil
}
