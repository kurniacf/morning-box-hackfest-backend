package model

type StrikePoint struct {
	UserID   string `json:"user_id" firestore:"user_id"`
	Value    int    `json:"value" firestore:"value"`
	Duration int    `json:"duration" firestore:"duration"`
}

type StrikePointResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Value    int    `json:"value"`
	Duration int    `json:"duration"`
}

type StrikePointRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	Value    int    `json:"value" binding:"required"`
	Duration int    `json:"duration" binding:"required"`
}
