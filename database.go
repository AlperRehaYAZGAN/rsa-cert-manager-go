package main

import (
	// local packages
	databases "ttcertman/src/platform/database"

	// database
	"gorm.io/gorm"
)

var dbApp *gorm.DB
var connErr error

func InitDatabaseConnection(dbConnString string) {
	dbApp, connErr = databases.InitSqliteDbConnection(dbConnString)
}
