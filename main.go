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
	mq "ttcertman/src/platform/brokers"
	db "ttcertman/src/platform/database"
)

func InitEnv(dir string) {
	// load env variables from
	config.LoadEnv(dir)
}

func run() error {
	// get current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Init environment variables
	InitEnv(dir)

	// init database
	db.InitDatabaseConnection(dir + "/" + config.DBCONNSTR)
	// init migrations
	db.InitMigrations()

	// init mq connection for queue
	mq.InitMQConnection(config.MQCONNSTR)

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
