package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqliteDbConnection(dbConnString string) (db *gorm.DB, err error) {
	var dbErr error
	db, dbErr = gorm.Open(sqlite.Open(dbConnString), &gorm.Config{})
	return db, dbErr
}
