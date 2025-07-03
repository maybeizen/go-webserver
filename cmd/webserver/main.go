package main

import (
	"context"
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

	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	log.Printf("Connected to MongoDB")

	defer func() {
		if err := db.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Starting server on port 8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
