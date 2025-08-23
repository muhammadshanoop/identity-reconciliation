package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muhammadshanoop/identity-reconciliation/database"
	"github.com/muhammadshanoop/identity-reconciliation/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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
