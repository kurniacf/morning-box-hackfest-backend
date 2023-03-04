package model

import "time"

type Order struct {
	UserId          string    `json:"userId" firestore:"userId"`
	PackageId       string    `json:"packageId" firestore:"packageId"`
	OrderDate       time.Time `json:"orderDate" firestore:"orderDate"`
	Days            int       `json:"days" firestore:"days"`
	DeliveryAddress string    `json:"deliveryAddress" firestore:"deliveryAddress"`
	Status          string    `json:"status" firestore:"status"`
}

type OrderResponse struct {
	User            UserResponse    `json:"user"`
	Package         PackageResponse `json:"package"`
	OrderDate       time.Time       `json:"orderDate"`
	Days            int             `json:"days"`
	DeliveryAddress string          `json:"deliveryAddress"`
	Status          string          `json:"status"`
}

type OrderRequest struct {
	UserId          string `json:"userId" binding:"required"`
	PackageId       string `json:"packageId" binding:"required"`
	OrderDate       time.Time
	Days            int    `json:"days" binding:"required"`
	DeliveryAddress string `json:"deliveryAddress" binding:"required"`
	Status          string
}

const (
	ORDER_DIKONFIRMASI = "DIKONFIRMASI"
	ORDER_DISIAPKAN    = "DISIAPKAN"
	ORDER_DIANTAR      = "DIANTAR"
	ORDER_DITERIMA     = "DITERIMA"
)
