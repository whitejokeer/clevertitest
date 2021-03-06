package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	datamanager "github.com/whitejokeer/clevertitest/database"
	"gorm.io/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize the app with predefined configuration
func (a *App) Initialize() {
	var err error
	a.DB, err = datamanager.DatabaseInitialization()
	if err != nil {
		log.Fatal("Could not connect database")
	}
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	a.Router.Queries([]string{"currency", "quantity"}...)
	// Routing for handling the projects
	a.Get("/beers", a.getBeers)
	a.Get("/beers/{beerID}", a.getBeerByID)
	a.Get("/beers/{beerID}/boxprice", a.getBeerBoxPrice)
	a.Post("/beers", a.createBeerItem)

	a.Put("/beers", a.updateBeerItem)
	a.Delete("/beers/{beerID}", a.deleteBeerItem)
}

// Get -> Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET", "OPTIONS")
}

// Post -> Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST", "OPTIONS")
}

// Put -> Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete -> Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Printf("Listening on port %s", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
