package service

import (
	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/repository"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentServiceInterface interface {
	CreateTransaction(order model.OrderResponse) (string, error)
	CreateTokenTransactionWithGateway(order model.OrderResponse) (string, error)
	CreateUrlTransactionWithGateway(order model.OrderResponse) (string, error)
	//ProcessWebhookNotification(notification *midtrans.TransactionNotification) error
}

type paymentService struct {
	pkgRepo  repository.PackageRepositoryInterface
	userRepo repository.UserRepositoryInterface
}

func NewPaymentService(pkgRepo repository.PackageRepositoryInterface, userRepo repository.UserRepositoryInterface) *paymentService {
	return &paymentService{
		pkgRepo:  pkgRepo,
		userRepo: userRepo,
	}
}

func (s *paymentService) CreateTransaction(chargeReq *snap.Request) (string, error) {
	response, err := config.MidtransClient.CreateTransaction(chargeReq)
	if err != nil {
		return "", err
	}

	return response.RedirectURL, nil
}

func (s *paymentService) CreateTokenTransactionWithGateway(chargeReq *snap.Request) (string, error) {
	token, err := config.MidtransClient.CreateTransactionToken(chargeReq)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *paymentService) CreateUrlTransactionWithGateway(chargeReq *snap.Request) (string, error) {
	url, err := config.MidtransClient.CreateTransactionUrl(chargeReq)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (r *paymentService) GenerateSnapReq(order model.OrderPaymentResponse, pkg model.PackageResponse) *snap.Request {
	// Get user details based on UserId
	user, err := r.userRepo.GetUser(order.UserId)
	if err != nil {
		return nil // handle the error appropriately
	}

	// Initiate Snap Request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  order.Id,
			GrossAmt: int64(pkg.Price),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
			Phone: user.PhoneNumber,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    order.PackageId,
				Price: int64(pkg.Price),
				Qty:   1,
				Name:  pkg.Name,
			},
		},
	}
	return snapReq
}
