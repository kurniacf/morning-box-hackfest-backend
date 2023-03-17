package model

import "time"

type Order struct {
	UserId          string    `json:"userId" firestore:"userId"`
	PackageId       string    `json:"packageId" firestore:"packageId"`
	EndDate         time.Time `json:"endDate" firestore:"endDate"`
	DeliveryAddress string    `json:"deliveryAddress" firestore:"deliveryAddress"`
	Status          string    `json:"status" firestore:"status"`
}

type OrderResponse struct {
	Id              string          `json:"id"`
	User            *UserResponse   `json:"user"`
	Package         PackageResponse `json:"package"`
	EndDate         time.Time       `json:"endDate" firestore:"endDate"`
	DeliveryAddress string          `json:"deliveryAddress"`
	Status          string          `json:"status"`
}

type OrderPaymentResponse struct {
	Id              string    `json:"id"`
	UserId          string    `json:"userId" firestore:"userId"`
	PackageId       string    `json:"packageId" firestore:"packageId"`
	EndDate         time.Time `json:"endDate" firestore:"endDate"`
	DeliveryAddress string    `json:"deliveryAddress"`
	Status          string    `json:"status"`
}

type OrderRequest struct {
	UserId          string `json:"userId" binding:"required"`
	PackageId       string `json:"packageId" binding:"required"`
	WeekOrder       int    `json:"weekOrder" binding:"required"`
	DeliveryAddress string `json:"deliveryAddress" binding:"required"`
	Status          string
}

type OrderUpdateRequest struct {
	UserId          string    `json:"userId" binding:"required"`
	PackageId       string    `json:"packageId" binding:"required"`
	EndDate         time.Time `json:"endDate" firestore:"endDate"`
	DeliveryAddress string    `json:"deliveryAddress" binding:"required"`
	Status          string
}

type OrderConfirmRequest struct {
	UserId  string `json:"userId" binding:"required"`
	OrderId string `json:"orderId" binding:"required"`
}

const (
	BELUM_DIKONFIRMASI = "BELUM_DIKONFIRMASI"
	ORDER_DIKONFIRMASI = "DIKONFIRMASI"
	ORDER_DISIAPKAN    = "DISIAPKAN"
	ORDER_DIANTAR      = "DIANTAR"
	ORDER_DITERIMA     = "DITERIMA"
)
