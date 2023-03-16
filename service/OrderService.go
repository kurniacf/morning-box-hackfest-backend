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
	UpdateOrder(id string, order model.OrderUpdateRequest) error
	DeleteOrder(id string) error
	ConfirmOrderADayBefore(id string) error
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
		EndDate:         time.Now().AddDate(0, 0, order.WeekOrder*7),
		DeliveryAddress: order.DeliveryAddress,
		Status:          model.BELUM_DIKONFIRMASI,
	}

	orderId, err := s.repo.CreateOrder(orderRequest)
	if err != nil {
		return "", err
	}

	return orderId, nil
}

func (s *orderService) UpdateOrder(id string, order model.OrderUpdateRequest) error {
	_, err := s.repo.GetOrder(id)
	if err != nil {
		return err
	}

	orderRequest := model.Order{
		UserId:    order.UserId,
		PackageId: order.PackageId,

		EndDate:         order.EndDate,
		DeliveryAddress: order.DeliveryAddress,
		Status:          order.Status,
	}

	err = s.repo.UpdateOrder(id, orderRequest)
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

func (s *orderService) ConfirmOrderADayBefore(id string) error {
	err := s.repo.ConfirmOrderADayBefore(id)
	if err != nil {
		return err
	}

	return nil
}
