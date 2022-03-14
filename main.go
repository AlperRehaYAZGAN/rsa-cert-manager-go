package main

import (
	// built-in
	"fmt"
	"log"
	"os"

	// web framework
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// local packages
	db "ttcertman/src/platform/database"
)

// Env variables config
var (
	APP_PORT      = os.Getenv("APP_PORT")
	DBCONNSTR     = os.Getenv("DBCONNSTR")
	NATS_URL      = os.Getenv("NATS_URL")
	REDISCONNSTR  = os.Getenv("REDISCONNSTR")
	S3_REGION     = os.Getenv("S3_REGION")
	S3_BUCKET     = os.Getenv("S3_BUCKET")
	S3_ACCESS_KEY = os.Getenv("S3_ACCESS_KEY")
	S3_SECRET_KEY = os.Getenv("S3_SECRET_KEY")
	S3_ENDPOINT   = os.Getenv("S3_ENDPOINT")
	S3_PUBLIC     = os.Getenv("S3_PUBLIC")
)

func LoadEnv() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// get env variables
	APP_PORT = os.Getenv("APP_PORT")
	DBCONNSTR = os.Getenv("DBCONNSTR")
	NATS_URL = os.Getenv("NATS_URL")
	REDISCONNSTR = os.Getenv("REDISCONNSTR")
	S3_PUBLIC = os.Getenv("S3_BUCKET")
	S3_REGION = os.Getenv("S3_REGION")
	S3_BUCKET = os.Getenv("S3_BUCKET")
	S3_ACCESS_KEY = os.Getenv("S3_ACCESS_KEY")
	S3_SECRET_KEY = os.Getenv("S3_SECRET_KEY")
	S3_ENDPOINT = os.Getenv("S3_ENDPOINT")

	// init database
	db.InitDatabaseConnection(dir + "/" + DBCONNSTR)
	// init migrations
	db.InitMigrations()
}

func run() error {
	// Load environment variables
	LoadEnv()

	// Create a new gin router
	r := gin.Default()
	Routes(r)

	// start server on port APP_PORT
	if APP_PORT == "" {
		APP_PORT = "9090"
	}
	return r.Run(":" + APP_PORT)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
