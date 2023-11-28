package params

import "gitlab.com/project-quiz/internal/params/generics"

type MaterialFilterParam struct {
	Level string `json:"level"`
	generics.GenericFilter
}

type MaterialCreateParam struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}

type MaterialEditParam struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level string `json:"level"`
}
