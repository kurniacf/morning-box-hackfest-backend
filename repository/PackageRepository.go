package repository

import (
	"context"
	"morning-box-hackfest-be/model"

	"cloud.google.com/go/firestore"
)

type PackageRepository struct {
	client     *firestore.Client
	collection string
}

func NewPackageRepository(client *firestore.Client) *PackageRepository {
	return &PackageRepository{
		client:     client,
		collection: "packages",
	}
}

func (r *PackageRepository) GetAllPackages() ([]*model.PackageResponse, error) {
	var packages []*model.PackageResponse
	iter := r.client.Collection(r.collection).Documents(context.Background())

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var p model.Package
		doc.DataTo(&p)

		packages = append(packages, &model.PackageResponse{
			Id:    doc.Ref.ID,
			Name:  p.Name,
			Price: p.Price,
			Foods: model.FoodPackageResponse{
				Monday:    model.FoodResponse{Id: p.Foods.Monday},
				Tuesday:   model.FoodResponse{Id: p.Foods.Tuesday},
				Wednesday: model.FoodResponse{Id: p.Foods.Wednesday},
				Thursday:  model.FoodResponse{Id: p.Foods.Thursday},
				Friday:    model.FoodResponse{Id: p.Foods.Friday},
				Saturday:  model.FoodResponse{Id: p.Foods.Saturday},
				Sunday:    model.FoodResponse{Id: p.Foods.Sunday},
			},
		})
	}

	return packages, nil

}

func (r *PackageRepository) GetPackage(id string) (model.PackageResponse, error) {
	doc, err := r.client.Collection(r.collection).Doc(id).Get(context.Background())
	if err != nil {
		return model.PackageResponse{}, err
	}

	var p model.Package
	doc.DataTo(&p)

	return model.PackageResponse{
		Id:    doc.Ref.ID,
		Name:  p.Name,
		Price: p.Price,
		Foods: model.FoodPackageResponse{
			Monday:    model.FoodResponse{Id: p.Foods.Monday},
			Tuesday:   model.FoodResponse{Id: p.Foods.Tuesday},
			Wednesday: model.FoodResponse{Id: p.Foods.Wednesday},
			Thursday:  model.FoodResponse{Id: p.Foods.Thursday},
			Friday:    model.FoodResponse{Id: p.Foods.Friday},
			Saturday:  model.FoodResponse{Id: p.Foods.Saturday},
			Sunday:    model.FoodResponse{Id: p.Foods.Sunday},
		},
	}, nil
}

func (r *PackageRepository) CreatePackage(p model.Package) (string, error) {
	ref, _, err := r.client.Collection(r.collection).Add(context.Background(), p)
	if err != nil {
		return "", err
	}
	return ref.ID, nil
}

func (r *PackageRepository) UpdatePackage(id string, p model.Package) error {
	_, err := r.client.Collection(r.collection).Doc(id).Set(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

func (r *PackageRepository) DeletePackage(id string) error {
	_, err := r.client.Collection(r.collection).Doc(id).Delete(context.Background())
	if err != nil {
		return err
	}
	return nil
}
