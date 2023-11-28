package entities

import "gitlab.com/project-quiz/internal/entities/base"

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	base.Timestamp
}
