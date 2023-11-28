package entities

import "gitlab.com/project-quiz/internal/entities/base"

type Material struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Level string `json:"level"`
	base.Timestamp
}
