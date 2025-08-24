package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muhammadshanoop/identity-reconciliation/database"
	"github.com/muhammadshanoop/identity-reconciliation/routes"
)

func init() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Println("Error loading .env file:", err)
		}
	}
}

func main() {
	database.GetConnect()
	database.Migrate()
	router := routes.SetupRouter()

	err := router.Run(":" + os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
