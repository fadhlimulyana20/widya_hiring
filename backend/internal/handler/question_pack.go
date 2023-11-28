package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/usecase"
	"gitlab.com/project-quiz/utils/boolpointer"
	"gitlab.com/project-quiz/utils/json"
	"gitlab.com/project-quiz/utils/validator"
	"gorm.io/gorm"
)

type questionPack struct {
	handler             Handler
	questionPackUsecase usecase.QuestionPackUsecase
	name                string
}

type QuestionPackHandler interface {
	// Get List of question pack
	GetList(w http.ResponseWriter, r *http.Request)
	// Get List of question pack for admin
	AdminGetList(w http.ResponseWriter, r *http.Request)
	// Get detail of question pack
	GetDetail(w http.ResponseWriter, r *http.Request)
	// Create question pack
	Create(w http.ResponseWriter, r *http.Request)
	// Update question pack
	Update(w http.ResponseWriter, r *http.Request)
	// Delete question pack
	Delete(w http.ResponseWriter, r *http.Request)
	// Add Question
	AddQuestion(w http.ResponseWriter, r *http.Request)
	// Delete Question
	DeleteQuestion(w http.ResponseWriter, r *http.Request)
	// Take Question for Basic Role
	BasicTakeQuestionPack(w http.ResponseWriter, r *http.Request)
	// Finish Question for Basic Role
	BasicFinishQuestionPack(w http.ResponseWriter, r *http.Request)
	// Get Question Pack attempt list for basic role
	BasicGetQuestionPackAttemptList(w http.ResponseWriter, r *http.Request)
}

func NewQuestionPackHandler(db *gorm.DB) QuestionPackHandler {
	return &questionPack{
		questionPackUsecase: usecase.NewQuestionPackUsecase(db),
		name:                "QUestion Pack Handler",
	}
}

func (q *questionPack) GetList(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.QuestionPackFilterParam
	ctx := appctx.NewResponse()

	param.IsActive = boolpointer.BoolPointer(true)
	if err := decoder.Decode(&param, r.URL.Query()); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	if len(ctx.Errors) > 0 {
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionPackUsecase.List(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionPack) AdminGetList(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.QuestionPackFilterParam
	ctx := appctx.NewResponse()

	if err := decoder.Decode(&param, r.URL.Query()); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	if len(ctx.Errors) > 0 {
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionPackUsecase.List(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionPack) GetDetail(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)

	resp := q.questionPackUsecase.Detail(idx)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionPack) Create(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.QuestionPackCreateParam
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

	resp := q.questionPackUsecase.Create(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionPack) Update(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.QuestionPackUpdateParam
	ctx := appctx.NewResponse()

	id := chi.URLParam(r, "id")
	param.ID, _ = strconv.Atoi(id)

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

	resp := q.questionPackUsecase.Update(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionPack) Delete(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)

	resp := q.questionPackUsecase.Delete(idx)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionPack) AddQuestion(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var param params.QuestionPackAddQuestionParam
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

	resp := q.questionPackUsecase.AddQuestions(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionPack) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var param params.QuestionPackAddQuestionParam
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

	resp := q.questionPackUsecase.DeleteQuestions(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionPack) BasicTakeQuestionPack(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var param params.QuestionPackAttemptTakeParam
	ctx := appctx.NewResponse()

	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.UserID = userID

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

	resp := q.questionPackUsecase.TakeQuestionPack(param.QuestionPackID, param.UserID)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionPack) BasicFinishQuestionPack(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var param params.QuestionPackAttemptTakeParam
	ctx := appctx.NewResponse()

	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.UserID = userID

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

	resp := q.questionPackUsecase.FinishQuestionPack(param.QuestionPackID, param.UserID)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionPack) BasicGetQuestionPackAttemptList(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var param params.QuestionPackAttemptFilterParam
	ctx := appctx.NewResponse()

	if err := decoder.Decode(&param, r.URL.Query()); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	if len(ctx.Errors) > 0 {
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.UserID = userID

	resp := q.questionPackUsecase.GetAttemptList(param)
	q.handler.Response(w, resp, startTime, time.Now())
}
