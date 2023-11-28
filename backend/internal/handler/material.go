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
	"gitlab.com/project-quiz/utils/json"
	"gitlab.com/project-quiz/utils/validator"
	"gorm.io/gorm"
)

type material struct {
	handler         Handler
	materialUsecase usecase.MaterialUsecase
	name            string
}

type MaterialHandler interface {
	// Handler to get list of material
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetListByContributor(w http.ResponseWriter, r *http.Request)
}

func NewMaterialHandler(db *gorm.DB) MaterialHandler {
	return &material{
		materialUsecase: usecase.NewMaterialUsecase(db),
		name:            "Material Handler",
	}
}

func (m *material) GetList(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.MaterialFilterParam
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
		m.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := m.materialUsecase.List(param)
	m.handler.Response(w, resp, startTime, time.Now())
}

func (m *material) GetDetail(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)

	resp := m.materialUsecase.Detail(idx)
	m.handler.Response(w, resp, startTime, time.Now())
}

func (m *material) Delete(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)

	resp := m.materialUsecase.Delete(idx)
	m.handler.Response(w, resp, startTime, time.Now())
}

func (m *material) Create(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.MaterialCreateParam
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		m.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		m.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := m.materialUsecase.Create(param)
	m.handler.Response(w, resp, startTime, time.Now())
}

func (m *material) Update(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.MaterialEditParam
	ctx := appctx.NewResponse()

	id := chi.URLParam(r, "id")
	param.ID, _ = strconv.Atoi(id)

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		m.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		m.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := m.materialUsecase.Update(param)
	m.handler.Response(w, resp, startTime, time.Now())
}

func (m *material) GetListByContributor(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.MaterialFilterParam
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
		m.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := m.materialUsecase.List(param)
	m.handler.Response(w, resp, startTime, time.Now())
}
