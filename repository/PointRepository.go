package repository

import (
	"context"
	"morning-box-hackfest-be/model"

	"cloud.google.com/go/firestore"
)

type PointRepositoryInterface interface {
	UpdatePoint(user model.UserResponse) error
}

type pointRepository struct {
	client     *firestore.Client
	collection string
}

func NewPointRepository(client *firestore.Client) *pointRepository {
	return &pointRepository{
		client:     client,
		collection: "users",
	}
}

func (r *pointRepository) UpdatePoint(user model.UserResponse) error {
	_, err := r.client.Collection(r.collection).Doc(user.Id).Update(context.Background(), []firestore.Update{
		{
			Path:  "past_strike_point",
			Value: user.PastStrikePoint,
		},
		{
			Path:  "current_strike_point",
			Value: user.CurrentStrikePoint + 1,
		},
	})

	return err
}
