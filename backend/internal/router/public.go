package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/project-quiz/internal/handler"
)

func (rtr *router) PublicRouterV1() http.Handler {
	router := chi.NewRouter()

	router.Mount("/auth", rtr.publicAuthRouterV1())
	router.Mount("/material", rtr.publicMaterialRouterV1())

	return router
}

func (rtr *router) publicAuthRouterV1() http.Handler {
	authHandler := handler.NewAuthHandler(rtr.cfg.DB, &rtr.cfg.SMTP, rtr.cfg.Secret, rtr.cfg.GoogleClientID)
	router := chi.NewRouter()

	router.Post("/registration", authHandler.Register)
	router.Post("/login", authHandler.Login)
	router.Post("/refresh", authHandler.Refresh)
	router.Post("/reset-password/request", authHandler.RequestResetPassword)
	router.Post("/reset-password/update", authHandler.ResetPassword)
	router.Post("/email-validation/request", authHandler.RequestValidationEmail)
	router.Post("/email-validation/validate", authHandler.ValidateEmail)
	router.Post("/oauth/google", authHandler.AuthWithGoogle)

	return router
}

func (rtr *router) publicMaterialRouterV1() http.Handler {
	materialHandler := handler.NewMaterialHandler(rtr.cfg.DB)
	router := chi.NewRouter()

	router.Get("/", materialHandler.GetList)
	router.Get("/{id}", materialHandler.GetDetail)

	return router
}
