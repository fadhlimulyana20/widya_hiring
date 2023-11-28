package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"gitlab.com/project-quiz/internal/appctx"

	"github.com/gorilla/schema"
)

// type response func(w http.ResponseWriter, resp appctx.Response, startTime time.Time)

var decoder = schema.NewDecoder()

type Handler struct{}

func (h *Handler) Response(w http.ResponseWriter, resp appctx.Response, startTime time.Time, endTime time.Time) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	resp = *resp.WithProcessTime(startTime, endTime)
	d, _ := json.Marshal(resp)
	w.Write(d)
}
