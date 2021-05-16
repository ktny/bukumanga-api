package config

import (
	"fmt"
	"os"
)

var (
	env = os.Getenv("ENV")
	dbUser = os.Getenv("POSTGRES_USER")
	dbPassword = os.Getenv("POSTGRES_PASSWORD")
	dbName = os.Getenv("POSTGRES_DB")
	dbHost = os.Getenv("POSTGRES_HOST")
	dbPort = "5432"
	dbType = "postgres"
)

func GetDBType() string {
	return dbType
}

func GetPostgresConnectionString() string {
	conString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		dbHost,
		dbPort,
		dbUser,
		dbName,
		dbPassword)
	if env == "production" {
		return conString
	} else {
		return conString + " sslmode=disable"
	}
}
