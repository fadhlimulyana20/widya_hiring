package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/project-quiz/internal/handler"
)

func (rtr *router) BasicRouterV1() http.Handler {
	router := chi.NewRouter()

	router.Mount("/auth", rtr.basicAuthRouterV1())
	router.Mount("/question", rtr.basicQuestionRouterV1())
	router.Mount("/analytic", rtr.basicAnaylticRouterV1())
	router.Mount("/question-pack", rtr.questionPackBasicRouterV1())
	router.Mount("/product", rtr.productBasicRouterV1())

	return router
}

func (rtr *router) basicAuthRouterV1() http.Handler {
	router := chi.NewRouter()
	authHandler := handler.NewAuthHandler(rtr.cfg.DB, &rtr.cfg.SMTP, rtr.cfg.Secret, rtr.cfg.GoogleClientID)

	router.Get("/me", authHandler.GetAuthenticatedUser)
	router.Post("/update-password", authHandler.UpdatePassword)
	router.Post("/update-account", authHandler.UpdateAccount)

	return router
}

func (rtr *router) basicQuestionRouterV1() http.Handler {
	router := chi.NewRouter()
	questionHandler := handler.NewQuestionHandler(rtr.cfg.DB, rtr.cfg.Minio)
	attemptHandler := handler.NewUserQuestionAttemptHandler(rtr.cfg.DB)

	router.Get("/", questionHandler.GetList)
	router.Get("/{id}", questionHandler.GetDetail)
	router.Post("/mark", attemptHandler.MarkLatestAttempt)
	router.Post("/answer", attemptHandler.AnswerQuestion)
	router.Post("/clear-answer", attemptHandler.ClearAnswer)
	router.Post("/submit-answer", attemptHandler.SubmitAnswer)
	router.Get("/answer/{id}", attemptHandler.GetLatestAnswer)
	router.Put("/add-remove-mark", questionHandler.AddRemoveMark)
	router.Get("/solution/{id}", questionHandler.GetSolution)

	return router
}

func (rtr *router) basicAnaylticRouterV1() http.Handler {
	router := chi.NewRouter()
	analyticHandler := handler.NewAnalyticHandler(rtr.cfg.DB, rtr.cfg.Minio)

	router.Get("/attempt", analyticHandler.GetAttemptAnalytic)
	router.Get("/point", analyticHandler.GetUserPoint)
	router.Get("/point-list", analyticHandler.GetUserPointList)

	return router
}

func (rtr *router) questionPackBasicRouterV1() http.Handler {
	questionPackHandler := handler.NewQuestionPackHandler(rtr.cfg.DB)
	router := chi.NewRouter()

	router.Get("/", questionPackHandler.GetList)
	router.Get("/{id}", questionPackHandler.GetDetail)
	router.Post("/take", questionPackHandler.BasicTakeQuestionPack)
	router.Post("/finish", questionPackHandler.BasicFinishQuestionPack)
	router.Get("/attempt", questionPackHandler.BasicGetQuestionPackAttemptList)

	return router
}

func (rtr *router) productBasicRouterV1() http.Handler {
	productHandler := handler.NewProductHandler(rtr.cfg.DB)
	router := chi.NewRouter()

	router.Get("/", productHandler.GetList)
	router.Get("/{id}", productHandler.GetDetail)
	router.Post("/", productHandler.Create)
	router.Delete("/{id}", productHandler.Delete)
	router.Put("/{id}", productHandler.Update)

	return router
}
