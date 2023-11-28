package entities

import "gitlab.com/project-quiz/internal/entities/base"

type UserQuestionAttempt struct {
	ID               int  `json:"id" gorm:"primaryKey"`
	QuestionID       int  `json:"question_id"`
	QuestionOptionID *int `json:"option_id"`
	UserID           int  `json:"user_id"`
	AttemptValue     bool `json:"attempt_value"`
	IsMarked         bool `json:"is_marked"`
	IsSubmitted      bool `json:"is_submitted"`
	base.Timestamp
}
