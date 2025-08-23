package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USERNAME = "myuser"
	DB_PASSWORD = "mypassword"
	DB_DATABASE = "mydb"
)

func loadConfig() {
	// DB_HOST = "localhost"
	// DB_PORT = "5432"
	// DB_USERNAME = "myuser"
	// DB_PASSWORD = "mypassword"
	// DB_DATABASE = "mydb"

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using default configuration")
		return
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_DATABASE = os.Getenv("DB_DATABASE")

}

func GetDSN() string {
	loadConfig()
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		DB_HOST, DB_USERNAME, DB_PASSWORD, DB_DATABASE, DB_PORT,
	)
}
