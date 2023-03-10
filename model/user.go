package model

type User struct {
	Name        string `json:"name" firestore:"name"`
	Email       string `json:"email" firestore:"email"`
	Password    string `json:"password" firestore:"password"`
	PhoneNumber string `json:"phone_number" firestore:"phone_number"`
	Address     string `json:"address" firestore:"address"`
	Point       int    `json:"point" firestore:"point"`
}

type UserResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Point       int    `json:"point"`
}

type UserRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Point       int    `json:"point"`
}
