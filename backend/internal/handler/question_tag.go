package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/usecase"
	"gitlab.com/project-quiz/utils/json"
	"gitlab.com/project-quiz/utils/validator"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type questionTag struct {
	handler            Handler
	questionTagUsecase usecase.QuestionTagUsecase
	name               string
}

type QuestionTagHandler interface {
	// Create a new role
	Create(w http.ResponseWriter, r *http.Request)
	// Get list of roles
	List(w http.ResponseWriter, r *http.Request)
	// Update a role
	Update(w http.ResponseWriter, r *http.Request)
	// Delete a role
	Delete(w http.ResponseWriter, r *http.Request)
	// Get detaile of Role
	Detail(w http.ResponseWriter, r *http.Request)
	// // Assign Role to User
	// AssignRole(w http.ResponseWriter, r *http.Request)
	// // Revoke Role from User
	// RevokeRole(w http.ResponseWriter, r *http.Request)
	// Get list of question tags by contributor
	ListByContributor(w http.ResponseWriter, r *http.Request)
}

func NewQuestionTagHandler(db *gorm.DB) QuestionTagHandler {
	return &questionTag{
		name:               "Question Tag Handler",
		questionTagUsecase: usecase.NewQuestionTagUsecase(db),
	}
}

func (q *questionTag) Create(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Create] is executed", q.name))
	startTime := time.Now()

	var param params.QuestionTagCreateParam
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionTagUsecase.Create(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionTag) List(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][List] is executed", q.name))
	startTime := time.Now()

	var param params.QuestionTagFilter
	ctx := appctx.NewResponse()

	if err := decoder.Decode(&param, r.URL.Query()); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionTagUsecase.List(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionTag) Update(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Update] is executed", q.name))
	startTime := time.Now()

	var param params.QuestionTagUpdateParam
	ctx := appctx.NewResponse()

	id := chi.URLParam(r, "id")
	param.ID, _ = strconv.Atoi(id)

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionTagUsecase.Update(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionTag) Delete(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Delete] is executed", q.name))
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	ID, _ := strconv.Atoi(id)

	resp := q.questionTagUsecase.Delete(ID)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionTag) Detail(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Detail] is executed", q.name))
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	ID, _ := strconv.Atoi(id)

	resp := q.questionTagUsecase.Detail(ID)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionTag) ListByContributor(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][List] is executed", q.name))
	startTime := time.Now()

	var param params.QuestionTagFilter
	ctx := appctx.NewResponse()

	if err := decoder.Decode(&param, r.URL.Query()); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionTagUsecase.List(param)
	q.handler.Response(w, resp, startTime, time.Now())
}
