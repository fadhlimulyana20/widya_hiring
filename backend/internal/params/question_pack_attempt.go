package params

import "gitlab.com/project-quiz/internal/params/generics"

type QuestionPackAttemptFilterParam struct {
	generics.GenericFilter
	UserID         int   `json:"user_id" schema:"user_id"`
	QuestionPackID int   `json:"question_pack_id" schema:"question_pack_id"`
	IsFinish       *bool `json:"is_finish" shcema:"is_finish"`
}

type QuestionPackAttemptTakeParam struct {
	QuestionPackID int `json:"question_pack_id" validate:"required"`
	UserID         int `json:"user_id"`
}

type QuestionPackAttemptFinishParam struct {
	QuestionPackAttemptID int `json:"question_pack_attempt_id" validate:"required"`
	UserID                int `json:"user_id"`
}
