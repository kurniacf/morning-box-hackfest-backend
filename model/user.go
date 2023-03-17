package model

type User struct {
	Name               string `json:"name" firestore:"Name"`
	Email              string `json:"email" firestore:"Email"`
	Password           string `json:"password" firestore:"Password"`
	PhoneNumber        string `json:"phone_number" firestore:"PhoneNumber"`
	Address            string `json:"address" firestore:"Address"`
	City               string `json:"city" firestore:"City"`
	PostalCode         string `json:"postal_code" firestore:"Postal_code"`
	PastStrikePoint    int    `json:"past_strike_point" firestore:"PastStrikePoint"`
	CurrentStrikePoint int    `json:"current_strike_point" firestore:"CurrentStrikePoint"`
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
	City               string `json:"city"`
	PostalCode         string `json:"postal_code"`
	PastStrikePoint    int    `json:"past_strike_point"`
	CurrentStrikePoint int    `json:"current_strike_point"`
}
