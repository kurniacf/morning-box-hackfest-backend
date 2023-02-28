package model

type Food struct {
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
	ImageURL    string `json:"imageUrl" firestore:"imageUrl"`
}

type FoodResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
}

type FoodRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageURL    string `json:"imageUrl" binding:"required"`
}
