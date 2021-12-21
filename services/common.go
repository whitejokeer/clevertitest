package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	respondJSON(w, code, map[string]string{"error": message})
}

func getIDFromRequest(r *http.Request) (int, error) {
	beerID := mux.Vars(r)

	if beerID["beerID"] == "" {
		err := errors.New("no se encontro el ID de la cerveza")
		return 0, err
	}

	id, err := strconv.Atoi(beerID["beerID"])
	if err != nil {
		return 0, err
	}
	return id, nil
}
