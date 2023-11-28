package entities

import (
	"time"

	"gitlab.com/project-quiz/internal/entities/base"
)

type PremiumPackage struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Token       string    `json:"token"`
	IsActive    bool      `json:"is_active"`
	ActiveUntil time.Time `json:"active_until"`
	UserID      int       `json:"user_id"`
	User        User      `json:"user"`
	LongPeriod  int       `json:"long_period"`
	base.Timestamp
}
