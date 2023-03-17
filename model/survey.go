package model

import "time"

type Survey struct {
	UserID    string    `json:"user_id" firestore:"user_id"`
	Survey1   int       `json:"survey_1" firestore:"survey_1"`
	Survey2   int       `json:"survey_2" firestore:"survey_2"`
	Survey3   int       `json:"survey_3" firestore:"survey_3"`
	Survey4   int       `json:"survey_4" firestore:"survey_4"`
	Survey5   int       `json:"survey_5" firestore:"survey_5"`
	Survey6   []string  `json:"survey_6" firestore:"survey_6"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
}

type SurveyResponse struct {
	Id        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Survey1   int       `json:"survey_1"`
	Survey2   int       `json:"survey_2"`
	Survey3   int       `json:"survey_3"`
	Survey4   int       `json:"survey_4"`
	Survey5   int       `json:"survey_5"`
	Survey6   []string  `json:"survey_6"`
	CreatedAt time.Time `json:"created_at"`
}

type SurveyRequest struct {
	UserID  string   `json:"user_id" binding:"required"`
	Survey1 int      `json:"survey_1" binding:"required"`
	Survey2 int      `json:"survey_2" binding:"required"`
	Survey3 int      `json:"survey_3" binding:"required"`
	Survey4 int      `json:"survey_4" binding:"required"`
	Survey5 int      `json:"survey_5" binding:"required"`
	Survey6 []string `json:"survey_6" binding:"required"`
}
