package api

import (
	"net/http"

	"github.com/whitejokeer/clevertitest/services"
)

// getBeers -> Consume the getAllBeers service
func (a *App) getBeers(w http.ResponseWriter, r *http.Request) {
	services.GetAllBeers(a.DB, w, r)
}

// createBeerItem -> Consume the createBeerItem service
func (a *App) createBeerItem(w http.ResponseWriter, r *http.Request) {
	services.CreateBeerItem(a.DB, w, r)
}


// getBeerByID -> Consume the getBeerByID service
func (a *App) getBeerByID(w http.ResponseWriter, r *http.Request) {
	services.GetBeerByID(a.DB, w, r)
}

