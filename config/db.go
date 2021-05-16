package config

import (
	"fmt"
	"os"
)

var (
	DBUser = os.Getenv("POSTGRES_USER")
	DBPassword = os.Getenv("POSTGRES_PASSWORD")
	DBName = os.Getenv("POSTGRES_DB")
	DBHost = os.Getenv("POSTGRES_HOST")
	DBPort = "5432"
	DBType = "postgres"
)

func GetDBType() string {
	return DBType
}

func GetPostgresConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)
}
