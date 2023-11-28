package entities

import "gitlab.com/project-quiz/internal/entities/base"

type QuestionPack struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Questions []Question `json:"questions" gorm:"many2many:question_pack_items;foreginKey:ID;joinForeignKey:QuestionPackID;references:ID;joinReferences:QuestionID"`
	IsFree    *bool      `json:"is_free" gorm:"default:false"`
	IsActive  *bool      `json:"is_active" gorm:"default:true"`
	TimeLimit int        `json:"time_limit"`
	base.Timestamp
}
