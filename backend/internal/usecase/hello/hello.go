package hello

import (
	"gitlab.com/project-quiz/internal/appctx"
)

type hello struct{}

type HelloUsecase interface {
	Hello(appctx.Data) appctx.Response
}

func NewHelloUsecase() HelloUsecase {
	return &hello{}
}

func (h *hello) Hello(appctx.Data) appctx.Response {
	return appctx.Response{
		Code: 200,
		Data: map[string]string{
			"hello": "world",
		},
	}
}
