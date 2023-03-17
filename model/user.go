package model

type User struct {
	Name               string `json:"name" firestore:"name"`
	Email              string `json:"email" firestore:"email"`
	Password           string `json:"password" firestore:"password"`
	PhoneNumber        string `json:"phone_number" firestore:"phone_number"`
	Address            string `json:"address" firestore:"address"`
	City               string `json:"city" firestore:"city"`
	PostalCode         string `json:"postal_code" firestore:"postal_code"`
	PastStrikePoint    int    `json:"past_strike_point" firestore:"past_strike_point"`
	CurrentStrikePoint int    `json:"current_strike_point" firestore:"current_strike_point"`
}

type UserResponse struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phone_number"`
	Address            string `json:"address"`
	City               string `json:"city"`
	PostalCode         string `json:"postal_code"`
	PastStrikePoint    int    `json:"past_strike_point"`
	CurrentStrikePoint int    `json:"current_strike_point"`
}

type UserRequest struct {
	Name               string `json:"name" binding:"required"`
	Email              string `json:"email" binding:"required"`
	Password           string `json:"password" binding:"required"`
	PhoneNumber        string `json:"phone_number" binding:"required"`
	Address            string `json:"address" binding:"required"`
	City               string `json:"city" binding:"required"`
	PostalCode         string `json:"postal_code" binding:"required"`
	PastStrikePoint    int    `json:"past_strike_point"`
	CurrentStrikePoint int    `json:"current_strike_point"`
}

type UserQuery struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	PhoneNumber        string `json:"phone_number"`
	Address            string `json:"address"`
	Point              int    `json:"point"`
	PastStrikePoint    int    `json:"past_strike_point"`
	CurrentStrikePoint int    `json:"current_strike_point"`
}
