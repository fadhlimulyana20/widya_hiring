package params

import (
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/entities/base"
	"gitlab.com/project-quiz/internal/params/generics"
)

type QuestionFilterParam struct {
	MaterialID      int    `json:"material_id" schema:"material_id"`
	Code            string `json:"code" schema:"code"`
	Show            string `json:"show" schema:"show"`
	ContributorID   int    `json:"contributor_id" schema:"contributor_id"`
	IncludePackOnly *bool  `json:"include_pack_only" schema:"include_pack_only"`
	generics.GenericFilter
}

type QuestionDetailResponse struct {
	ID              int                       `json:"id" gorm:"primaryKey"`
	Body            string                    `json:"body"`
	MaterialID      int                       `json:"material_id"`
	IsImage         bool                      `json:"is_image"`
	ImgPath         string                    `json:"img_path"`
	QuestionOptions []entities.QuestionOption `json:"question_options,omitempty"`
	IsAnswered      bool                      `json:"is_answered"`
	AnswerID        int                       `json:"answer_id"`
	IsAnswerTrue    bool                      `json:"is_answer_true"`
	TrueAnswerID    int                       `json:"true_answer_id"`
	IsMarked        bool                      `json:"is_marked"`
	QuestionTags    []entities.QuestionTag    `json:"tags,omitempty"`
	Code            string                    `json:"code"`
	ImgPlacementUrl string                    `json:"img_placement_url"`
	ContributorID   int                       `json:"contributor_id"`
	base.Timestamp
}

type QuestionCreate struct {
	Body            string                 `json:"body"`
	MaterialID      int                    `json:"material_id"`
	IsImage         bool                   `json:"is_image"`
	ImgPath         string                 `json:"img_path"`
	QuestionOptions []QuestionOptionCreate `json:"question_options,omitempty"`
	ImgPlacementUrl string                 `json:"img_placement_url"`
	ContributorID   int                    `json:"contributor_id"`
	IsPackOnly      *bool                  `json:"is_pack_only"`
}

type QuestionUpdate struct {
	ID              int                    `json:"id" gorm:"primaryKey"`
	Body            string                 `json:"body"`
	MaterialID      int                    `json:"material_id"`
	IsImage         bool                   `json:"is_image"`
	ImgPath         string                 `json:"img_path"`
	IsActive        *bool                  `json:"is_active"`
	QuestionOptions []QuestionOptionUpdate `json:"question_options,omitempty"`
	ContributorID   int                    `json:"contributor_id"`
	IsPackOnly      *bool                  `json:"is_pack_only"`
}

type QuestionListResponse struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	Body          string `json:"body"`
	MaterialID    int    `json:"material_id"`
	IsSubmitted   bool   `json:"is_submitted"`
	IsAnswered    bool   `json:"is_answered"`
	AnswerID      int    `json:"answer_id"`
	IsAnswerTrue  bool   `json:"is_answer_true"`
	IsMarked      bool   `json:"is_marked"`
	Code          string `json:"code"`
	ContributorID int    `json:"contributor_id"`
}

type QuestionMarkAddRemoveParam struct {
	QuestionID int    `json:"question_id" validate:"required"`
	UserID     int    `json:"user_id"`
	Action     string `json:"action" validate:"required"`
}

type QuestionAddTags struct {
	QuestionID int   `json:"question_id" validate:"required"`
	TagIDs     []int `json:"tag_ids" validate:"required"`
}

type QuestionRemoveTag struct {
	QuestionID int `json:"question_id" validate:"required"`
	TagID      int `json:"tag_id" validate:"required"`
}
