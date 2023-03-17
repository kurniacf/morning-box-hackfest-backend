package service

import (
	"errors"
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/repository"
)

type AuthServiceInterface interface {
	Signin(authSignin model.AuthSignin) (*model.UserResponse, error)
}

type authService struct {
	repo repository.AuthRepositoryInterface
}

func NewAuthService(repo repository.AuthRepositoryInterface) *authService {
	return &authService{repo}
}

func (s *authService) Signin(authSignIn model.AuthSignin) (*model.UserResponse, error) {
	// get user
	user, err := s.repo.GetUserByEmail(authSignIn.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("User has not been registered")
	}

	// validate password
	if user.Password != authSignIn.Password {
		return nil, errors.New("Username or password invalid!")
	}

	userResponse := &model.UserResponse{
		Id:                 user.Id,
		Name:               user.Name,
		Email:              user.Email,
		PhoneNumber:        user.PhoneNumber,
		Address:            user.Address,
		City:               user.City,
		PostalCode:         user.PostalCode,
		PastStrikePoint:    user.PastStrikePoint,
		CurrentStrikePoint: user.CurrentStrikePoint,
	}

	return userResponse, nil
}
