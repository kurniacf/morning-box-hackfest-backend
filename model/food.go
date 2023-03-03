package model

type Food struct {
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
	ImageURL    string `json:"imageUrl" firestore:"imageUrl"`
	Calories    int    `json:"calories" firestore:"calories"`
	Type        string `json:"type" firestore:"type"`
}

type FoodResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
	Calories    int    `json:"calories"`
	Type        string `json:"type"`
}

type FoodRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageURL    string `json:"imageUrl" binding:"required"`
	Calories    int    `json:"calories" binding:"required"`
	Type        string `json:"type" binding:"required"`
}
