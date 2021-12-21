package services

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	datamanager "github.com/whitejokeer/clevertitest/database"
)

// GetAllBeers gets all the beers from the database
func GetAllBeers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	Beers, err := datamanager.BeerItemManager(db).GetAllBeerItems()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, Beers)
}

// CreateBeerItem creates a new beer in the database
func CreateBeerItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	beerItem := datamanager.BeerItems{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&beerItem); err != nil {
		respondError(w, http.StatusBadRequest, "Request invalida")
		return
	}
	defer r.Body.Close()

	if err := datamanager.BeerItemManager(db).CreateBeerItem(beerItem); err != nil {
		errMessage := err.Error()
		if errMessage == "pq: duplicate key value violates unique constraint \"beer_items_pkey\"" {
			respondError(w, http.StatusConflict, "El ID de la cerveza ya existe")
			return
		}
		if strings.Contains(errMessage, "pq: null value in column") {
			respondError(w, http.StatusBadRequest, "Request invalida")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, "Cerveza creada")
}

// GetBeerByID gets a beer by its ID
func GetBeerByID(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Request invalida")
		return
	}

	BeerItem, err := datamanager.BeerItemManager(db).GetBeerItemByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			respondError(w, http.StatusNotFound, "El id de la cerveza no existe")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, BeerItem)
}

// GetBeerBoxPrice -> Get the price of a beer box
func GetBeerBoxPrice(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Request invalida")
		return
	}

	BeerItem, err := datamanager.BeerItemManager(db).GetBeerItemByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			respondError(w, http.StatusNotFound, "El id de la cerveza no existe")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	quantityInt := 6
	quantity := r.FormValue("quantity")

	if quantity != "" {
		quantityInt, err = strconv.Atoi(quantity)
		if err != nil {
			respondError(w, http.StatusBadRequest, "Request invalida")
			return
		}
	}

	BeerItem.Price = BeerItem.Price * float64(quantityInt)

	currency := r.FormValue("currency")
	if currency != "" {
		BeerItem.Price = CurrencyConverter(currency, BeerItem.Price)
	}

	respondJSON(w, http.StatusOK, "Price Total: "+strconv.FormatFloat(BeerItem.Price, 'f', 2, 64))
}

// UpdateBeerItem updates a beer in the database
func UpdateBeerItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	beerItem := datamanager.BeerItems{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&beerItem); err != nil {
		respondError(w, http.StatusBadRequest, "Request invalida")
		return
	}
	defer r.Body.Close()

	if err := datamanager.BeerItemManager(db).UpdateBeerItem(beerItem); err != nil {
		if err.Error() == "record not found" {
			respondError(w, http.StatusNotFound, "El id de la cerveza no existe")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "Cerveza actualizada")
}

// DeleteBeerItem deletes a beer from the database
func DeleteBeerItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Request invalida")
		return
	}

	if err := datamanager.BeerItemManager(db).DeleteBeerItem(id); err != nil {
		if err.Error() == "record not found" {
			respondError(w, http.StatusNotFound, "El id de la cerveza no existe")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "Cerveza eliminada")
}
