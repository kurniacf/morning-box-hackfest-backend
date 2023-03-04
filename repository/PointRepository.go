package repository

import (
	"context"
	"morning-box-hackfest-be/model"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type PointRepositoryInterface interface {
	GetAllPoints() ([]model.StrikePointResponse, error)
	GetPoint(id string) (model.StrikePointResponse, error)
	CreatePoint(point model.StrikePointRequest) (string, error)
	UpdatePoint(id string, point model.StrikePointRequest) error
	DeletePoint(id string) error
}

type pointRepository struct {
	client *firestore.Client
}

func NewPointRepository(client *firestore.Client) *pointRepository {
	return &pointRepository{
		client: client,
	}
}

func (r *pointRepository) GetAllPoints() ([]model.StrikePointResponse, error) {
	var points []model.StrikePointResponse
	iter := r.client.Collection("strike_points").Documents(context.Background())
	defer iter.Stop()
	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var point model.StrikePointResponse
		if err = doc.DataTo(&point); err != nil {
			return nil, err
		}

		point.ID = doc.Ref.ID
		points = append(points, point)
	}

	return points, nil
}

func (r *pointRepository) GetPoint(id string) (model.StrikePointResponse, error) {
	doc, err := r.client.Collection("strike_points").Doc(id).Get(context.Background())
	if err != nil {
		return model.StrikePointResponse{}, err
	}
	var point model.StrikePointResponse
	if err = doc.DataTo(&point); err != nil {
		return model.StrikePointResponse{}, err
	}
	point.ID = doc.Ref.ID
	return point, nil

}

func (r *pointRepository) CreatePoint(point model.StrikePointRequest) (string, error) {
	doc, _, err := r.client.Collection("strike_points").Add(context.Background(), point)
	if err != nil {
		return "", err
	}
	return doc.ID, nil
}

func (r *pointRepository) UpdatePoint(id string, point model.StrikePointRequest) error {
	_, err := r.client.Collection("strike_points").Doc(id).Set(context.Background(), point)
	if err != nil {
		return err
	}
	return nil
}

func (r *pointRepository) DeletePoint(id string) error {
	_, err := r.client.Collection("strike_points").Doc(id).Delete(context.Background())
	if err != nil {
		return err
	}
	return nil
}
