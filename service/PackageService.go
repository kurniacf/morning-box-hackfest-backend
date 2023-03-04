package service

import (
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/repository"
)

type PackageService struct {
	repo *repository.PackageRepository
}

func NewPackageService(repo *repository.PackageRepository) *PackageService {
	return &PackageService{repo: repo}
}

func (s *PackageService) GetAllPackages() ([]*model.PackageResponse, error) {
	packages, err := s.repo.GetAllPackages()
	if err != nil {
		return nil, err
	}
	return packages, nil
}

func (s *PackageService) GetPackage(id string) (model.PackageResponse, error) {
	p, err := s.repo.GetPackage(id)
	if err != nil {
		return model.PackageResponse{}, err
	}
	return p, nil
}

func (s *PackageService) CreatePackage(p model.PackageRequest) (string, error) {
	packageToCreate := model.Package{
		Name:  p.Name,
		Price: p.Price,
		Foods: model.FoodPackage{
			Monday:    p.Foods.Monday,
			Tuesday:   p.Foods.Tuesday,
			Wednesday: p.Foods.Wednesday,
			Thursday:  p.Foods.Thursday,
			Friday:    p.Foods.Friday,
			Saturday:  p.Foods.Saturday,
			Sunday:    p.Foods.Sunday,
		},
	}
	id, err := s.repo.CreatePackage(packageToCreate)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *PackageService) UpdatePackage(id string, p model.PackageRequest) error {
	_, err := s.GetPackage(id)
	if err != nil {
		return err
	}
	packageToUpdate := model.Package{
		Name:  p.Name,
		Price: p.Price,
		Foods: model.FoodPackage{
			Monday:    p.Foods.Monday,
			Tuesday:   p.Foods.Tuesday,
			Wednesday: p.Foods.Wednesday,
			Thursday:  p.Foods.Thursday,
			Friday:    p.Foods.Friday,
			Saturday:  p.Foods.Saturday,
			Sunday:    p.Foods.Sunday,
		},
	}

	err = s.repo.UpdatePackage(id, packageToUpdate)
	if err != nil {
		return err
	}

	return nil
}

func (s *PackageService) DeletePackage(id string) error {
	_, err := s.GetPackage(id)
	if err != nil {
		return err
	}
	err = s.repo.DeletePackage(id)
	if err != nil {
		return err
	}

	return nil
}
