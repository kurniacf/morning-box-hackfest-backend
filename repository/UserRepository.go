package repository

import (
	"context"
	"morning-box-hackfest-be/model"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type UserRepositoryInterface interface {
	GetAllUsers() ([]model.UserResponse, error)
	GetUser(id string) (model.UserResponse, error)
	CreateUser(user model.UserRequest) (string, error)
	UpdateUser(id string, user model.UserRequest) error
	DeleteUser(id string) error
}

type userRepository struct {
	client *firestore.Client
}

func NewUserRepository(client *firestore.Client) *userRepository {
	return &userRepository{
		client: client,
	}
}

func (r *userRepository) GetAllUsers() ([]model.UserResponse, error) {
	var users []model.UserResponse
	iter := r.client.Collection("users").Documents(context.Background())
	defer iter.Stop()

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var user model.UserResponse
		if err = doc.DataTo(&user); err != nil {
			return nil, err
		}

		user.Id = doc.Ref.ID
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetUser(id string) (model.UserResponse, error) {
	doc, err := r.client.Collection("users").Doc(id).Get(context.Background())
	if err != nil {
		return model.UserResponse{}, err
	}
	var user model.UserResponse
	if err = doc.DataTo(&user); err != nil {
		return model.UserResponse{}, err
	}

	user.Id = doc.Ref.ID
	return user, nil
}

func (r *userRepository) CreateUser(user model.UserRequest) (string, error) {
	doc, _, err := r.client.Collection("users").Add(context.Background(), user)
	if err != nil {
		return "", err
	}
	return doc.ID, nil
}

func (r *userRepository) UpdateUser(id string, user model.UserRequest) error {
	_, err := r.client.Collection("users").Doc(id).Set(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(id string) error {
	_, err := r.client.Collection("users").Doc(id).Delete(context.Background())
	if err != nil {
		return err
	}
	return nil
}
