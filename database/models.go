// This file contains the database table models information.
package datamanager

type BeerItems struct {
	Id       int     `gorm:"primaryKey;default:null"`
	Name     string  `gorm:"type:varchar(255);not null;default:null"`
	Brewery  string  `gorm:"type:varchar(255);not null;default:null"`
	Country  string  `gorm:"type:varchar(255);not null;default:null"`
	Price    float64 `gorm:"type:float;not null;default:null"`
	Currency string  `gorm:"type:varchar(32);not null;default:null"`
}
