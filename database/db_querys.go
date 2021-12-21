// This file contains the database querys functions.

package datamanager

import "github.com/jinzhu/gorm"

type DataManager interface {
	GetAllBeerItems() ([]BeerItems, error)
	GetBeerItemByID(id int) (BeerItems, error)
	CreateBeerItem(beerItem BeerItems) error
	UpdateBeerItem(beerItem BeerItems) error
	DeleteBeerItem(id int) error
}

// BeerItems CRUD interface
type BeerItemsQuerys struct {
	DB *gorm.DB
}

func BeerItemManager(db *gorm.DB) DataManager {
	return &BeerItemsQuerys{
		DB: db,
	}
}

// GetAllBeerItems -> Get all BeerItems
func (db *BeerItemsQuerys) GetAllBeerItems() ([]BeerItems, error) {
	var beerItems []BeerItems
	err := db.DB.Find(&beerItems).Error
	return beerItems, err
}

// GetBeerItemByID -> Get BeerItem by ID
func (db *BeerItemsQuerys) GetBeerItemByID(id int) (BeerItems, error) {
	beerItem := BeerItems{}
	err := db.DB.First(&beerItem, BeerItems{Id: id}).Error
	return beerItem, err
}

// CreateBeerItem -> Create BeerItem
func (db *BeerItemsQuerys) CreateBeerItem(beerItem BeerItems) error {
	err := db.DB.Create(&beerItem).Error
	return err
}

// UpdateBeerItem -> Update BeerItem
func (db *BeerItemsQuerys) UpdateBeerItem(beerItem BeerItems) error {
	err := db.DB.Save(&beerItem).Error
	return err
}

// DeleteBeerItem -> Delete BeerItem
func (db *BeerItemsQuerys) DeleteBeerItem(id int) error {
	err := db.DB.Delete(&BeerItems{}, id).Error
	return err
}
