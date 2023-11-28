package router

import (
	"net/http"

	"gitlab.com/project-quiz/internal/handler"

	"github.com/go-chi/chi/v5"
)

// Router with auth with RBAC admin
func (rtr *router) AdminRouterV1() http.Handler {
	router := chi.NewRouter()

	router.Mount("/user", rtr.userAdminRouterV1())
	router.Mount("/role", rtr.roleAdminRouterV1())
	router.Mount("/question", rtr.questionAdminRouterV1())
	router.Mount("/material", rtr.materialAdminRouterV1())
	router.Mount("/question-tag", rtr.questionTagAdminRouterV1())
	router.Mount("/question-solution", rtr.questionSolutionRouterV1())
	router.Mount("/analytic", rtr.analyticAdminRouterV1())
	router.Mount("/question-pack", rtr.questionPackAdminRouterV1())

	return router
}

func (rtr *router) userAdminRouterV1() http.Handler {
	userHandler := handler.NewUserHandler(rtr.cfg.DB)
	router := chi.NewRouter()

	router.Post("/", userHandler.Create)
	router.Get("/", userHandler.List)
	router.Get("/{id}", userHandler.Get)
	router.Put("/{id}", userHandler.Update)

	return router
}

func (rtr *router) roleAdminRouterV1() http.Handler {
	roleHandler := handler.NewRoleHandler(rtr.cfg.DB)
	router := chi.NewRouter()

	router.Post("/", roleHandler.Create)
	router.Get("/", roleHandler.Read)
	router.Get("/{id}", roleHandler.Detail)
	router.Put("/{id}", roleHandler.Update)
	router.Delete("/{id}", roleHandler.Delete)
	router.Post("/assign", roleHandler.AssignRole)
	router.Post("/revoke", roleHandler.RevokeRole)

	return router
}

func (rtr *router) questionAdminRouterV1() http.Handler {
	questionHandler := handler.NewQuestionHandler(rtr.cfg.DB, rtr.cfg.Minio)
	router := chi.NewRouter()

	router.Get("/", questionHandler.AdminGetList)
	router.Post("/", questionHandler.Create)
	router.Get("/{id}", questionHandler.AdminGetDetail)
	router.Put("/{id}", questionHandler.Update)

	router.Post("/option/", questionHandler.AdminAddOption)
	router.Delete("/option/{id}", questionHandler.AdminDeleteOption)
	router.Put("/option/{id}", questionHandler.AdminUpdateOption)

	router.Post("/add-tags", questionHandler.AddTags)
	router.Post("/remove-tag", questionHandler.RemoveTag)

	router.Get("/solution/{id}", questionHandler.GetSolution)

	router.Post("/question-image", questionHandler.UploadImagePlacement)

	return router
}

func (rtr *router) materialAdminRouterV1() http.Handler {
	materialRouter := handler.NewMaterialHandler(rtr.cfg.DB)
	router := chi.NewRouter()

	router.Get("/", materialRouter.GetList)
	router.Get("/{id}", materialRouter.GetDetail)
	router.Post("/", materialRouter.Create)
	router.Put("/{id}", materialRouter.Update)
	router.Delete("/{id}", materialRouter.Delete)

	return router
}

func (rtr *router) questionTagAdminRouterV1() http.Handler {
	questionTagRouter := handler.NewQuestionTagHandler(rtr.cfg.DB)
	router := chi.NewRouter()

	router.Get("/", questionTagRouter.List)
	router.Get("/{id}", questionTagRouter.Detail)
	router.Post("/", questionTagRouter.Create)
	router.Put("/{id}", questionTagRouter.Update)
	router.Delete("/{id}", questionTagRouter.Delete)

	return router
}

func (rtr *router) questionSolutionRouterV1() http.Handler {
	questionSolution := handler.NewQuestionSolutionUsecase(rtr.cfg.DB, rtr.cfg.Minio)
	router := chi.NewRouter()

	router.Post("/", questionSolution.Create)
	router.Post("/create-with-file", questionSolution.CreateWithUploadFile)
	router.Put("/{id}", questionSolution.Update)
	router.Get("/{id}", questionSolution.Detail)
	router.Delete("/{id}", questionSolution.Delete)

	return router
}

func (rtr *router) analyticAdminRouterV1() http.Handler {
	analyticHandler := handler.NewAnalyticHandler(rtr.cfg.DB, rtr.cfg.Minio)
	router := chi.NewRouter()

	router.Get("/", analyticHandler.GetCreatorAnalytic)

	return router
}

func (rtr *router) questionPackAdminRouterV1() http.Handler {
	questionPackHandler := handler.NewQuestionPackHandler(rtr.cfg.DB)
	router := chi.NewRouter()

	router.Post("/", questionPackHandler.Create)
	router.Get("/", questionPackHandler.AdminGetList)
	router.Get("/{id}", questionPackHandler.GetDetail)
	router.Put("/{id}", questionPackHandler.Update)
	router.Delete("/{id}", questionPackHandler.Delete)
	router.Post("/add-question", questionPackHandler.AddQuestion)
	router.Post("/delete-question", questionPackHandler.DeleteQuestion)

	return router
}
