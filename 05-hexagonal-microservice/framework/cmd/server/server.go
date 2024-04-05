package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/pedro-hca/golang-studies/05-rest-api/framework/database"
)

var db database.DbConfig

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db.DbType = os.Getenv("DB_TYPE")
	db.Dsn = os.Getenv("DSN")
	db.Debug, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatalf("Error parsing bool")
	}
	os.Getenv("ENV")
}
