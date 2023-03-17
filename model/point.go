package model

type StrikePoint struct {
	PastStrikePoint    int `json:"past_strike_point" firestore:"past_strike_point"`
	CurrentStrikePoint int `json:"current_strike_point" firestore:"current_strike_point"`
}

type StrikePointRequest struct {
	UserID             string `json:"user_id" binding:"required"`
	PastStrikePoint    int    `json:"past_strike_point"`
	CurrentStrikePoint int    `json:"current_strike_point"`
}
