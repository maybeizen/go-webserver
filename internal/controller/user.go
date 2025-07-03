package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/maybeizen/go-webserver/internal/types"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AddUserToDatabase(client *mongo.Client, user types.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Database("golang-test").Collection("users")

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	log.Printf("Added user %s to database", user.Name)
	return nil
}

func AddUserHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user types.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := AddUserToDatabase(client, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}