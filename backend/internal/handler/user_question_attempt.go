package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/usecase"
	"gitlab.com/project-quiz/utils/json"
	"gitlab.com/project-quiz/utils/validator"
	"gorm.io/gorm"
)

type userQuestionAttempt struct {
	handler        Handler
	attemptUsecase usecase.UserQuestionAttemptUsecase
	name           string
}

type UserQuestionAttemptHandler interface {
	MarkLatestAttempt(w http.ResponseWriter, r *http.Request)
	AnswerQuestion(w http.ResponseWriter, r *http.Request)
	ClearAnswer(w http.ResponseWriter, r *http.Request)
	SubmitAnswer(w http.ResponseWriter, r *http.Request)
	GetLatestAnswer(w http.ResponseWriter, r *http.Request)
}

func NewUserQuestionAttemptHandler(db *gorm.DB) UserQuestionAttemptHandler {
	return &userQuestionAttempt{
		name:           "User Question Attempt Handler",
		attemptUsecase: usecase.NewUserQuestionAttemptUsecase(db),
	}
}

func (u *userQuestionAttempt) MarkLatestAttempt(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.AttemptSubmitAnswerQuestionParam
	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.UserID = userID
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", u.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", u.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	if len(ctx.Errors) > 0 {
		u.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := u.attemptUsecase.MarkAttempt(param)
	u.handler.Response(w, resp, startTime, time.Now())
}

func (u *userQuestionAttempt) AnswerQuestion(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.AttemptAnswerQuestionParam
	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.UserID = userID
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", u.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", u.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	if len(ctx.Errors) > 0 {
		u.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := u.attemptUsecase.AnswerQuestion(param)
	u.handler.Response(w, resp, startTime, time.Now())
}

func (u *userQuestionAttempt) ClearAnswer(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.AttemptClearAnswerQuestionParam
	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.UserID = userID
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", u.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", u.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	if len(ctx.Errors) > 0 {
		u.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := u.attemptUsecase.ClearAnswer(param)
	u.handler.Response(w, resp, startTime, time.Now())
}

func (u *userQuestionAttempt) SubmitAnswer(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.AttemptSubmitAnswerQuestionParam
	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.UserID = userID
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", u.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", u.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	if len(ctx.Errors) > 0 {
		u.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := u.attemptUsecase.SubmitAnswer(param)
	u.handler.Response(w, resp, startTime, time.Now())
}

func (u *userQuestionAttempt) GetLatestAnswer(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.AttemptClearAnswerQuestionParam
	userID, _ := strconv.Atoi(r.Header.Get("user"))
	questionID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	param.UserID = userID
	param.QuestionID = questionID
	ctx := appctx.NewResponse()

	if err := validator.Validate(param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", u.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	if len(ctx.Errors) > 0 {
		u.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := u.attemptUsecase.GetLatestAnswer(param)
	u.handler.Response(w, resp, startTime, time.Now())
}
