package datamanager

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// dbMigration -> Create and migrate the database tables with the required relationships
func dbMigration(db *gorm.DB) *gorm.DB {
	err := db.AutoMigrate(&BeerItems{})
	if err != nil {
		log.Fatal("Could not migrate database")
	}
	return db
}

// CreateConnectionString -> Create the connection string for the database
func CreateConnectionString(config *DBConfig) string {
	return "host=" + config.Host + " port=" + config.Port + " user=" + config.Username + " dbname=" + config.Name + " password=" + config.Password
}

// DatabaseInitialization -> Create the database connection (Should run once only)
func DatabaseInitialization() (*gorm.DB, error) {
	config := GetConfig()
	dbURI := CreateConnectionString(config.DB)

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect database")
	}

	db = dbMigration(db)

	return db, nil
}
