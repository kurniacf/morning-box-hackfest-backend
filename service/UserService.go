package service

import (
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/repository"
)

type UserServiceInterface interface {
	GetAllUsers() ([]model.UserResponse, error)
	GetUser(id string) (model.UserResponse, error)
	CreateUser(user model.UserRequest) (string, error)
	UpdateUser(id string, user model.UserRequest) error
	DeleteUser(id string) error
}

type userService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *userService {
	return &userService{repo}
}

func (s *userService) GetAllUsers() ([]model.UserResponse, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) GetUser(id string) (model.UserResponse, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return model.UserResponse{}, err
	}

	return user, nil
}

func (s *userService) CreateUser(user model.UserRequest) (string, error) {
	newUser := model.User{
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		City:        user.City,
		PostalCode:  user.PostalCode,
	}

	newUserRequest := model.UserRequest{
		Name:        newUser.Name,
		Email:       newUser.Email,
		Password:    newUser.Password,
		PhoneNumber: newUser.PhoneNumber,
		Address:     newUser.Address,
		City:        user.City,
		PostalCode:  user.PostalCode,
	}

	userId, err := s.repo.CreateUser(newUserRequest)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (s *userService) UpdateUser(id string, user model.UserRequest) error {
	updateUser := model.User{
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		City:        user.City,
		PostalCode:  user.PostalCode,
	}

	updateUserRequest := model.UserRequest{
		Name:        updateUser.Name,
		Email:       updateUser.Email,
		Password:    updateUser.Password,
		PhoneNumber: updateUser.PhoneNumber,
		Address:     updateUser.Address,
		City:        user.City,
		PostalCode:  user.PostalCode,
	}

	err := s.repo.UpdateUser(id, updateUserRequest)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(id string) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
