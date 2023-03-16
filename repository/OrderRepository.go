package repository

import (
	"context"
	"fmt"
	"morning-box-hackfest-be/model"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type OrderRepositoryInterface interface {
	GetAllOrders() ([]*model.OrderResponse, error)
	GetOrder(id string) (model.OrderResponse, error)
	CreateOrder(order model.Order) (string, error)
	UpdateOrder(id string, order model.Order) error
	DeleteOrder(id string) error
	ConfirmOrderADayBefore(id string) error
	GetActiveOrder(userId string) (*model.OrderResponse, error)
}

type orderRepository struct {
	client     *firestore.Client
	collection string
}

func NewOrderRepository(client *firestore.Client) *orderRepository {
	return &orderRepository{
		client:     client,
		collection: "orders",
	}
}

func (r *orderRepository) GetAllOrders() ([]*model.OrderResponse, error) {
	var orders []*model.OrderResponse

	iter := r.client.Collection(r.collection).Documents(context.Background())
	defer iter.Stop()

	for {
		fmt.Println("loop")
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var order *model.Order
		if err = doc.DataTo(&order); err != nil {
			return nil, err
		}

		orders = append(orders, &model.OrderResponse{
			User: model.UserResponse{
				Id: order.UserId,
			},
			Package: model.PackageResponse{
				Id: order.PackageId,
			},
			EndDate:         order.EndDate,
			DeliveryAddress: order.DeliveryAddress,
			Status:          order.Status,
		})
	}

	return orders, nil
}

func (r *orderRepository) GetOrder(id string) (model.OrderResponse, error) {
	doc, err := r.client.Collection(r.collection).Doc(id).Get(context.Background())
	if err != nil {
		return model.OrderResponse{}, err
	}

	type orderQueryModel struct {
		UserId          string    `json:"userId"`
		PackageId       string    `json:"packageId"`
		OrderDate       time.Time `json:"orderDate"`
		Days            int       `json:"days"`
		DeliveryAddress string    `json:"deliveryAddress"`
		Status          string    `json:"status"`
	}

	var orderQuery orderQueryModel
	if err = doc.DataTo(&orderQuery); err != nil {
		return model.OrderResponse{}, err
	}

	order := model.OrderResponse{
		User: model.UserResponse{
			Id: orderQuery.UserId,
		},
		Package: model.PackageResponse{
			Id: orderQuery.PackageId,
		},
		EndDate:         orderQuery.OrderDate,
		DeliveryAddress: orderQuery.DeliveryAddress,
		Status:          orderQuery.Status,
	}

	return order, nil
}

func (r *orderRepository) CreateOrder(order model.Order) (string, error) {
	doc, _, err := r.client.Collection(r.collection).Add(context.Background(), order)
	if err != nil {
		return "", err
	}

	return doc.ID, nil
}
func (r *orderRepository) UpdateOrder(id string, order model.Order) error {
	_, err := r.client.Collection(r.collection).Doc(id).Set(context.Background(), order)
	if err != nil {
		return err
	}
	return nil
}

func (r *orderRepository) DeleteOrder(id string) error {
	_, err := r.client.Collection(r.collection).Doc(id).Delete(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *orderRepository) ConfirmOrderADayBefore(id string) error {
	_, err := r.client.Collection(r.collection).Doc(id).Update(context.Background(), []firestore.Update{
		{
			Path:  "status",
			Value: model.ORDER_DIKONFIRMASI,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) GetActiveOrder(userId string) (*model.OrderResponse, error) {
	iter := r.client.Collection(r.collection).Where("UserId", "==", userId).Where("EndDate", ">=", time.Now()).Documents(context.Background())

	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var order *model.Order
		if err = doc.DataTo(&order); err != nil {
			return nil, err
		}

		return &model.OrderResponse{
			User: model.UserResponse{
				Id: order.UserId,
			},
			Package: model.PackageResponse{
				Id: order.PackageId,
			},
			EndDate:         order.EndDate,
			DeliveryAddress: order.DeliveryAddress,
			Status:          order.Status,
		}, nil
	}

	return nil, nil
}
