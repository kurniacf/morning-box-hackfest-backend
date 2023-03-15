// File: service/snap.go
package service

import (
	"errors"
	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/model"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentServiceInterface interface {
	CreateTransaction(order model.OrderResponse) (string, error)
	CreateTokenTransactionWithGateway(order model.OrderResponse) (string, error)
	CreateUrlTransactionWithGateway(order model.OrderResponse) (string, error)
}

type paymentService struct{}

func NewPaymentService() *paymentService {
	return &paymentService{}
}

func (s *paymentService) CreateTransaction(order model.OrderResponse) (string, error) {
	if order.Id == "" {
		return "", errors.New("invalid order")
	}

	chargeReq := GenerateSnapReq(order)

	response, err := config.MidtransClient.CreateTransaction(chargeReq)
	if err != nil {
		return "", err
	}

	return response.RedirectURL, nil
}

func (s *paymentService) CreateTokenTransactionWithGateway(order model.OrderResponse) (string, error) {
	if order.Id == "" {
		return "", errors.New("invalid order")
	}

	chargeReq := GenerateSnapReq(order)

	token, err := config.MidtransClient.CreateTransactionToken(chargeReq)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *paymentService) CreateUrlTransactionWithGateway(order model.OrderResponse) (string, error) {
	if order.Id == "" {
		return "", errors.New("invalid order")
	}

	chargeReq := GenerateSnapReq(order)

	url, err := config.MidtransClient.CreateTransactionUrl(chargeReq)
	if err != nil {
		return "", err
	}

	return url, nil
}

func GenerateSnapReq(order model.OrderResponse) *snap.Request {
	// Initiate Customer address
	custAddress := &midtrans.CustomerAddress{
		FName:       order.User.Name,
		Phone:       order.User.PhoneNumber,
		Address:     order.User.Address,
		City:        order.User.City,
		Postcode:    order.User.PostalCode,
		CountryCode: "IDN",
	}

	// Initiate Snap Request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  order.Id,
			GrossAmt: int64(order.Package.Price),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    order.User.Name,
			Email:    order.User.Email,
			Phone:    order.User.PhoneNumber,
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    order.Package.Id,
				Price: int64(order.Package.Price),
				Qty:   1,
				Name:  order.Package.Name,
			},
		},
	}
	return snapReq
}
