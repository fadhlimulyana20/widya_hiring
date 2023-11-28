package handler

import (
	"net/http"
	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/usecase/hello"
	"time"
)

type helloHandler struct {
	handler      Handler
	helloUsecase hello.HelloUsecase
}

func NewHelloHandler() *helloHandler {
	return &helloHandler{
		helloUsecase: hello.NewHelloUsecase(),
	}
}

func (h *helloHandler) Hello(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	d := appctx.Data{
		Request: r,
	}
	resp := h.helloUsecase.Hello(d)
	h.handler.Response(w, resp, startTime, time.Now())
}