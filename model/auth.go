package model

type AuthSignin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
