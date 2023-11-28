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

type product struct {
	handler        Handler
	productUsecase usecase.ProductUsecase
	name           string
}

type ProductHandler interface {
	// Handler to get list of material
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewProductHandler(db *gorm.DB) ProductHandler {
	return &product{
		productUsecase: usecase.NewProductUsecase(db),
		name:           "Product Handler",
	}
}

func (p *product) GetList(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.ProductFilterParam
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
		p.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := p.productUsecase.List(param)
	p.handler.Response(w, resp, startTime, time.Now())
}

func (p *product) GetDetail(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)

	resp := p.productUsecase.Detail(idx)
	p.handler.Response(w, resp, startTime, time.Now())
}

func (p *product) Delete(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)

	resp := p.productUsecase.Delete(idx)
	p.handler.Response(w, resp, startTime, time.Now())
}

func (p *product) Create(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.ProductCreateParam
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		p.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		p.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := p.productUsecase.Create(param)
	p.handler.Response(w, resp, startTime, time.Now())
}

func (p *product) Update(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.ProductUpdateParam
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		p.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		p.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	id := chi.URLParam(r, "id")
	param.ID, _ = strconv.Atoi(id)

	resp := p.productUsecase.Update(param)
	p.handler.Response(w, resp, startTime, time.Now())
}
