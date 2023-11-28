package entities

import "gitlab.com/project-quiz/internal/entities/base"

type UserQuestionMark struct {
	ID         int `json:"id" gorm:"primaryKey"`
	QuestionID int `json:"question_id"`
	UserID     int `json:"user_id"`
	base.Timestamp
}
