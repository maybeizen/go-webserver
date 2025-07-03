package routes

import (
	"net/http"

	"github.com/maybeizen/go-webserver/internal/controller"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux, client *mongo.Client) http.Handler {
	r.Get("/random", controller.RandomHandler)
	r.Get("/health", controller.HealthCheckController)
	r.Post("/say", controller.SayHandler)
	r.Post("/user", controller.AddUserHandler(client))

	return r
}