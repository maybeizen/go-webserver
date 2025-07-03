package controller

import (
	"encoding/json"
	"net/http"

	"github.com/maybeizen/go-webserver/internal/types"
)

func ParseJSONBody(r *http.Request) (string, error) {
	var say types.Say
	err := json.NewDecoder(r.Body).Decode(&say)
	if err != nil {
		return "", err
	}

	return say.Message, nil
}

func SayHandler(w http.ResponseWriter, r *http.Request) {
	msg, err := ParseJSONBody(r)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(msg))
	
}