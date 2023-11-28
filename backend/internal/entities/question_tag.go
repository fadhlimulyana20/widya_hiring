package entities

import "gitlab.com/project-quiz/internal/entities/base"

type QuestionTag struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	base.Timestamp
}

func (QuestionTag) TableName() string {
	return "tags"
}
