package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDbConnection(dbConnString string) (db *gorm.DB, err error) {
	var dbErr error
	db, dbErr = gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	return db, dbErr
}
