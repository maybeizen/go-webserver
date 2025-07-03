package routes

import (
	"net/http"

	"github.com/maybeizen/go-webserver/internal/controller"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) http.Handler {
	r.Get("/random", controller.RandomHandler)
	r.Get("/health", controller.HealthCheckController)
	r.Post("/say", controller.SayHandler)

	return r
}