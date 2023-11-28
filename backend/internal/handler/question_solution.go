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
	"gitlab.com/project-quiz/utils/minio"
	"gitlab.com/project-quiz/utils/validator"
	"gorm.io/gorm"
)

type questionSolution struct {
	handler                 Handler
	questionSolutionUsecase usecase.QuestionSolutionUsecase
	name                    string
}

type QuestionSolutionUsecase interface {
	// Create a new role
	Create(w http.ResponseWriter, r *http.Request)
	CreateWithUploadFile(w http.ResponseWriter, r *http.Request)
	// Get list of roles
	// List(w http.ResponseWriter, r *http.Request)
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
}

func NewQuestionSolutionUsecase(db *gorm.DB, minio minio.MinioStorageContract) QuestionSolutionUsecase {
	return &questionSolution{
		name:                    "Question Solution Handler",
		questionSolutionUsecase: usecase.NewQuestionSolutionUsecase(db, minio),
	}
}

func (q *questionSolution) Create(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Create] is executed", q.name))
	startTime := time.Now()

	var param params.QuestionSolutionCreate
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

	resp := q.questionSolutionUsecase.Create(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionSolution) CreateWithUploadFile(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Create] is executed", q.name))
	startTime := time.Now()

	if err := r.ParseMultipartForm(1024); err != nil {
		d := appctx.NewResponse().WithErrors(err.Error())
		q.handler.Response(w, *d, startTime, time.Now())
		return
	}

	qIDform := r.FormValue("question_id")
	questionID, _ := strconv.Atoi(qIDform)
	param := params.QuestionSolutionWithFileUploadCreate{
		QuestionID:   questionID,
		SolutionType: r.FormValue("solution_type"),
	}

	if param.SolutionType == "img" {
		_, fileImg, err := r.FormFile("file")
		if err != nil {
			d := appctx.NewResponse().WithErrors(err.Error())
			q.handler.Response(w, *d, startTime, time.Now())
			return
		}
		param.SolutionImg = fileImg
	}

	if param.SolutionType == "pdf" {
		_, filePDF, err := r.FormFile("file")
		if err != nil {
			d := appctx.NewResponse().WithErrors(err.Error())
			q.handler.Response(w, *d, startTime, time.Now())
			return
		}
		param.PdfFile = filePDF
	}

	ctx := appctx.NewResponse()

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionSolutionUsecase.CreateWithFile(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

// func (q *questionSolution) List(w http.ResponseWriter, r *http.Request) {
// 	logrus.Info(fmt.Sprintf("[%s][List] is executed", q.name))
// 	startTime := time.Now()

// 	var param params.Q
// 	ctx := appctx.NewResponse()

// 	if err := decoder.Decode(&param, r.URL.Query()); err != nil {
// 		logrus.Error(err.Error())
// 		ctx = ctx.WithErrors(err.Error())
// 		q.handler.Response(w, *ctx, startTime, time.Now())
// 		return
// 	}

// 	if err := validator.Validate(param); err != nil {
// 		logrus.Error(err.Error())
// 		ctx = ctx.WithErrors(err.Error())
// 		q.handler.Response(w, *ctx, startTime, time.Now())
// 		return
// 	}

// 	resp := q.questionTagUsecase.List(param)
// 	q.handler.Response(w, resp, startTime, time.Now())
// }

func (q *questionSolution) Update(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Update] is executed", q.name))
	startTime := time.Now()

	var param params.QuestionSolutionUpdate
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

	resp := q.questionSolutionUsecase.Update(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionSolution) Delete(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Delete] is executed", q.name))
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	ID, _ := strconv.Atoi(id)

	resp := q.questionSolutionUsecase.Delete(ID)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *questionSolution) Detail(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Detail] is executed", q.name))
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	ID, _ := strconv.Atoi(id)

	resp := q.questionSolutionUsecase.Detail(ID)
	q.handler.Response(w, resp, startTime, time.Now())
}
