package main

import (
	"log"
	"net/http"

	"github.com/maybeizen/go-webserver/internal/routes"
	"github.com/maybeizen/go-webserver/pkg/utils"

	"github.com/go-chi/chi/v5"
)

func main() {
	utils.LoadEnv()

	r := chi.NewRouter()
	r.Mount("/api", routes.Routes(r))

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
