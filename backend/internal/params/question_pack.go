package params

import "gitlab.com/project-quiz/internal/params/generics"

type QuestionPackCreateParam struct {
	Name      string `json:"name"`
	TimeLimit int    `json:"time_limit"`
}

type QuestionPackUpdateParam struct {
	ID       int   `json:"id"`
	IsFree   *bool `json:"is_free"`
	IsActive *bool `json:"is_active"`
	QuestionPackCreateParam
}

type QuestionPackFilterParam struct {
	IsActive *bool `json:"is_active"`
	generics.GenericFilter
}

type QuestionPackAddQuestionParam struct {
	ID          int   `json:"id"`
	QuestionIDs []int `json:"question_ids"`
}
