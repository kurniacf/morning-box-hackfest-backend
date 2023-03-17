package service

import (
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/repository"
)

type PointServiceInterface interface {
	UpdatePoint(user model.UserResponse) error
}

type pointService struct {
	repo repository.PointRepositoryInterface
}

func NewPointService(repo repository.PointRepositoryInterface) *pointService {
	return &pointService{repo}
}

func (s *pointService) UpdatePoint(user model.UserResponse) error {
	err := s.repo.UpdatePoint(user)

	return err

}
