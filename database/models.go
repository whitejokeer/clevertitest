package datamanager

// This file contains the database table models information.

import "github.com/jinzhu/gorm"

type BeerItems struct {
	gorm.Model
	Id       int     `gorm:"primaryKey"`
	Name     string  `gorm:"type:varchar(255);not null"`
	Brewery  string  `gorm:"type:varchar(255);not null"`
	Country  string  `gorm:"type:varchar(255);not null"`
	Price    float64 `gorm:"type:float;not null"`
	Currency string  `gorm:"type:varchar(32);not null"`
}
