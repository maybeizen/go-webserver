package controller

import (
	"encoding/json"
	"net/http"

	"github.com/maybeizen/go-webserver/internal/types"
)

func HealthCheckController(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(types.Health{Status: "OK"})
}