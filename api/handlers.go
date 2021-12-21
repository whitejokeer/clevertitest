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

// getBeerBoxPrice -> Consume the getBeerBoxPrice service
func (a *App) getBeerBoxPrice(w http.ResponseWriter, r *http.Request) {
	services.GetBeerBoxPrice(a.DB, w, r)
}

// updateBeerItem -> Consume the updateBeerItem service
func (a *App) updateBeerItem(w http.ResponseWriter, r *http.Request) {
	services.UpdateBeerItem(a.DB, w, r)
}

// deleteBeerItem -> Consume the deleteBeerItem service
func (a *App) deleteBeerItem(w http.ResponseWriter, r *http.Request) {
	services.DeleteBeerItem(a.DB, w, r)
}



