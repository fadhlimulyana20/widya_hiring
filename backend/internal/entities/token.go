package entities

import "gitlab.com/project-quiz/internal/entities/base"

type Token struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	UserID      int    `json:"user_id"`
	Code        string `json:"code"`
	TokenType   string `json:"token_type"`
	IsCompleted bool   `json:"is_completed" gorm:"dafault:false"`
	ValidUntil  int64  `json:"valid_until"`
	base.Timestamp
}
