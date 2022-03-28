package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Env variables config
var (
	APP_PORT      = os.Getenv("APP_PORT")
	DBCONNSTR     = os.Getenv("DBCONNSTR")
	MQCONNSTR     = os.Getenv("MQCONNSTR")
	CACHECONNSTR  = os.Getenv("CACHECONNSTR")
	S3_REGION     = os.Getenv("S3_REGION")
	S3_BUCKET     = os.Getenv("S3_BUCKET")
	S3_ACCESS_KEY = os.Getenv("S3_ACCESS_KEY")
	S3_SECRET_KEY = os.Getenv("S3_SECRET_KEY")
	S3_ENDPOINT   = os.Getenv("S3_ENDPOINT")
	S3_PUBLIC     = os.Getenv("S3_PUBLIC")
)

func LoadEnv(dirname string) {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// get env variables
	APP_PORT = os.Getenv("APP_PORT")
	DBCONNSTR = os.Getenv("DBCONNSTR")
	MQCONNSTR = os.Getenv("MQCONNSTR")
	CACHECONNSTR = os.Getenv("CACHECONNSTR")
	S3_PUBLIC = os.Getenv("S3_BUCKET")
	S3_REGION = os.Getenv("S3_REGION")
	S3_BUCKET = os.Getenv("S3_BUCKET")
	S3_ACCESS_KEY = os.Getenv("S3_ACCESS_KEY")
	S3_SECRET_KEY = os.Getenv("S3_SECRET_KEY")
	S3_ENDPOINT = os.Getenv("S3_ENDPOINT")

}
