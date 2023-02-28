package model

type Food struct {
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
	ImageURL    string `json:"imageUrl" firestore:"imageUrl"`
}
