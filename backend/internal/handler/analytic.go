package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/usecase"
	"gitlab.com/project-quiz/utils/minio"
	"gitlab.com/project-quiz/utils/validator"
	"gorm.io/gorm"
)

type analytic struct {
	name             string
	analyticUsecase  usecase.AnalyticUsecase
	userPointUsecase usecase.UserPointUsecase
	handler          Handler
}

type AnalyticHandler interface {
	GetCreatorAnalytic(w http.ResponseWriter, r *http.Request)
	GetAttemptAnalytic(w http.ResponseWriter, r *http.Request)
	GetUserPoint(w http.ResponseWriter, r *http.Request)
	GetUserPointList(w http.ResponseWriter, r *http.Request)
}

func NewAnalyticHandler(db *gorm.DB, m minio.MinioStorageContract) AnalyticHandler {
	return &analytic{
		name:             "Analytic Handler",
		analyticUsecase:  usecase.NewAnaliticUsecase(db, m),
		userPointUsecase: usecase.NewUserPointUsecase(db),
	}
}

func (a *analytic) GetCreatorAnalytic(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	resp := a.analyticUsecase.GetCreatorAnalytic()

	a.handler.Response(w, resp, startTime, time.Now())
}

func (a *analytic) GetAttemptAnalytic(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	userID, _ := strconv.Atoi(r.Header.Get("user"))
	resp := a.analyticUsecase.GetTotalUserAttempt(userID)

	a.handler.Response(w, resp, startTime, time.Now())
}

func (a *analytic) GetUserPoint(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	userID, _ := strconv.Atoi(r.Header.Get("user"))
	resp := a.userPointUsecase.Get(userID)

	a.handler.Response(w, resp, startTime, time.Now())
}

func (a *analytic) GetUserPointList(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.UserPointFilterParam
	ctx := appctx.NewResponse()

	if err := decoder.Decode(&param, r.URL.Query()); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		a.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		a.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := a.userPointUsecase.GetList(param)

	a.handler.Response(w, resp, startTime, time.Now())
}
