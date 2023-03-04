package service

import (
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/repository"
	"time"
)

type OrderServiceInterface interface {
	GetAllOrders() ([]*model.OrderResponse, error)
	GetOrder(id string) (model.OrderResponse, error)
	CreateOrder(order model.OrderRequest) (string, error)
	UpdateOrder(id string, order model.OrderRequest) error
	DeleteOrder(id string) error
}

type orderService struct {
	repo repository.OrderRepositoryInterface
}

func NewOrderService(repo repository.OrderRepositoryInterface) *orderService {
	return &orderService{repo}
}

func (s *orderService) GetAllOrders() ([]*model.OrderResponse, error) {
	orders, err := s.repo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *orderService) GetOrder(id string) (model.OrderResponse, error) {
	order, err := s.repo.GetOrder(id)
	if err != nil {
		return model.OrderResponse{}, err
	}

	return order, nil
}

func (s *orderService) CreateOrder(order model.OrderRequest) (string, error) {
	orderRequest := model.Order{
		UserId:          order.UserId,
		PackageId:       order.PackageId,
		OrderDate:       time.Now(),
		Days:            order.Days,
		DeliveryAddress: order.DeliveryAddress,
		Status:          "",
	}

	orderId, err := s.repo.CreateOrder(orderRequest)
	if err != nil {
		return "", err
	}

	return orderId, nil
}

func (s *orderService) UpdateOrder(id string, order model.OrderRequest) error {
	_, err := s.repo.GetOrder(id)
	if err != nil {
		return err
	}

	orderRequest := model.OrderRequest{
		UserId:          order.UserId,
		PackageId:       order.PackageId,
		OrderDate:       order.OrderDate,
		Days:            order.Days,
		DeliveryAddress: order.DeliveryAddress,
		Status:          order.Status,
	}

	err = s.UpdateOrder(id, orderRequest)
	if err != nil {
		return err
	}

	return nil
}

func (s *orderService) DeleteOrder(id string) error {
	err := s.repo.DeleteOrder(id)
	if err != nil {
		return err
	}

	return nil
}
