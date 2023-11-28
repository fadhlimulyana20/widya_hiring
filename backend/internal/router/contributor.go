package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/project-quiz/internal/handler"
)

func (rtr *router) ContributorRouterV1() http.Handler {
	router := chi.NewRouter()

	router.Mount("/question", rtr.questionContributorRouterV1())
	router.Mount("/material", rtr.materialContributorRouterV1())
	router.Mount("/question-tag", rtr.questionTagContributorRouterV1())

	return router
}

func (rtr *router) questionContributorRouterV1() http.Handler {
	question := handler.NewQuestionHandler(rtr.cfg.DB, rtr.cfg.Minio)
	router := chi.NewRouter()

	router.Get("/", question.GetListByContributor)
	router.Post("/", question.CreateByContributor)
	router.Put("/{id}", question.UpdateByContributor)
	router.Get("/{id}", question.GetDetailByContributor)

	router.Post("/option", question.AdminAddOption)
	router.Delete("/option/{id}", question.AdminDeleteOption)
	router.Put("/option/{id}", question.AdminUpdateOption)

	router.Post("/tags", question.AddTags)
	router.Post("/tags/remove", question.RemoveTag)

	return router
}

func (rtr *router) materialContributorRouterV1() http.Handler {
	material := handler.NewMaterialHandler(rtr.cfg.DB)
	router := chi.NewRouter()

	router.Get("/", material.GetListByContributor)

	return router
}

func (rtr *router) questionTagContributorRouterV1() http.Handler {
	qth := handler.NewQuestionTagHandler(rtr.cfg.DB)
	router := chi.NewRouter()

	router.Get("/", qth.ListByContributor)

	return router
}
