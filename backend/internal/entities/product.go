package entities

import "gitlab.com/project-quiz/internal/entities/base"

type Product struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Pname       string `json:"name"`
	Description string `json:"description"`
	base.Timestamp
}
