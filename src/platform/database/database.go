package database

import (
	// database
	"gorm.io/gorm"

	// local packages
	models "ttcertman/models"
)

var dbApp *gorm.DB
var connErr error

func InitDatabaseConnection(dbConnString string) {
	dbApp, connErr = InitSqliteDbConnection(dbConnString)
}

func InitMigrations() {
	dbApp.AutoMigrate(&models.Platform{})
	dbApp.AutoMigrate(&models.Resource{})
	dbApp.AutoMigrate(&models.Scan{})
	dbApp.AutoMigrate(&models.Service{})
	dbApp.AutoMigrate(&models.SSLPair{})
	dbApp.AutoMigrate(&models.SSHPair{})
}

func GetDbConn() *gorm.DB {
	return dbApp
}
