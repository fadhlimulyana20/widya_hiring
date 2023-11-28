package params

import "gitlab.com/project-quiz/internal/params/generics"

type ProductFilterParam struct {
	generics.GenericFilter
}

type ProductCreateParam struct {
	Pname       string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ProductUpdateParam struct {
	ProductCreateParam
	ID int `json:"id"`
}
