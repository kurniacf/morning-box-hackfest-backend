package repository

import (
	"context"
	"errors"
	"fmt"
	"morning-box-hackfest-be/model"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type AuthRepositoryInterface interface {
	GetUserByEmail(email string) (*model.UserQuery, error)
}

type authRepository struct {
	client     *firestore.Client
	collection string
}

func NewAuthRepository(client *firestore.Client) *authRepository {
	return &authRepository{
		client:     client,
		collection: "users",
	}
}

func (r *authRepository) GetUserByEmail(email string) (*model.UserQuery, error) {
	iter := r.client.Collection(r.collection).Where("Email", "==", email).Documents(context.Background())

	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.New("Error retreiving user")

		}

		var user *model.UserQuery
		err = doc.DataTo(&user)
		if err != nil {
			return nil, errors.New("Error retreiving user")
		}

		return user, nil
	}

	fmt.Println("qwe")
	// return nil, errors.New("User has not been registered")
	return nil, errors.New("User has not been registered")
}
