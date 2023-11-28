package entities

import (
	"time"

	"gitlab.com/project-quiz/internal/entities/base"
)

type QuestionPackAttempt struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	UserID         int       `json:"user_id"`
	QuestionPackID int       `json:"question_pack_id"`
	IsFinish       bool      `json:"is_finish"`
	Score          float32   `json:"score"`
	StartedAt      time.Time `json:"started_at"`
	FinishedAt     time.Time `json:"finished_at"`
	base.Timestamp
}
