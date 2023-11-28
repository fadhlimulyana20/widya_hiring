package router

import (
	"encoding/json"
	"net/http"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"gitlab.com/project-quiz/internal/appctx"
	m "gitlab.com/project-quiz/internal/middleware"
	mail "gitlab.com/project-quiz/utils/mailer"
	"gitlab.com/project-quiz/utils/minio"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type router struct {
	router *chi.Mux
	cfg    *RouterCfg
}

type RouterCfg struct {
	DB             *gorm.DB
	SMTP           mail.Mailer
	Minio          minio.MinioStorageContract
	Secret         string
	AesSecret      string
	GoogleClientID string
}

func NewRouter(r *RouterCfg) Router {
	return &router{
		router: chi.NewRouter(),
		cfg:    r,
	}
}

func (rtr *router) Route() http.Handler {
	rtr.router.Use(m.Cors(rtr.cfg.DB))
	rtr.router.Use(m.Logger)
	rtr.router.Use(m.Recovery)
	rtr.router.Use(m.Authorization(rtr.cfg.DB))
	rtr.router.Use(m.Pagination)

	// Sentry
	sentryMiddleware := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})
	rtr.router.Use(sentryMiddleware.Handle)

	rtr.router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		logrus.Error("Error 404 page not found")
		resp := *appctx.NewResponse().WithErrors("Page not found").WithCode(404)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.Code)
		d, _ := json.Marshal(resp)
		w.Write(d)
	})

	rtr.router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		logrus.Error("Error 405 method not allowed")
		resp := *appctx.NewResponse().WithErrors("method not allowed").WithCode(405)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.Code)
		d, _ := json.Marshal(resp)
		w.Write(d)
	})

	rtr.router.Mount("/hello", rtr.helloRouter())
	rtr.router.Mount("/public/v1", rtr.PublicRouterV1())
	rtr.router.Mount("/admin/v1", rtr.AdminRouterV1())
	rtr.router.Mount("/basic/v1", rtr.BasicRouterV1())
	rtr.router.Mount("/contributor/v1", rtr.ContributorRouterV1())

	return rtr.router
}
