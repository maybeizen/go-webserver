package controller

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"

	"github.com/maybeizen/go-webserver/internal/types"
)

func GenerateRandomNumber(number int) int {
	return rand.IntN(number)
}

func RandomHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(types.Random{Value: GenerateRandomNumber(100)})
}
