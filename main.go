package main

import (
	// built-in
	"fmt"
	"log"
	"os"

	// web framework
	"github.com/gin-gonic/gin"

	// local packages
	config "ttcertman/config"
	db "ttcertman/src/platform/database"
)

func InitEnv() {
	// get current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// load env variables from
	config.LoadEnv(dir)

	// init database
	db.InitDatabaseConnection(dir + "/" + config.DBCONNSTR)
	// init migrations
	db.InitMigrations()
}

func run() error {
	// Init environment variables
	InitEnv()

	// Create a new gin router
	r := gin.Default()
	Routes(r)

	// start server on port APP_PORT
	if config.APP_PORT == "" {
		config.APP_PORT = "9090"
	}
	return r.Run(":" + config.APP_PORT)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
