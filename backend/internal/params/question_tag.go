package params

import "gitlab.com/project-quiz/internal/params/generics"

type QuestionTagFilter struct {
	Name string `json:"name"`
	generics.GenericFilter
}

type QuestionTagCreateParam struct {
	Name string `json:"name" validate:"required"`
}

type QuestionTagUpdateParam struct {
	ID int `json:"id" validate:"required"`
	QuestionTagCreateParam
}
