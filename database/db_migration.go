package datamanager

import "gorm.io/gorm"

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(
		// TODO: Add models here
	)
	return db
}
