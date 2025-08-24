package config

import (
	"fmt"
	"log"
	"os"
)

var (
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
)

func GetDSN() string {
	sslmode := "disable"
	if os.Getenv("APP_ENV") == "production" {
		sslmode = "require"
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Kolkata",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
		sslmode,
	)

	if dsn == "" {
		log.Fatal("Database connection string is empty — check environment variables")
	}
	return dsn
}
