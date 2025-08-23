package database

import (
	"log"

	"github.com/muhammadshanoop/identity-reconciliation/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.Contact{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database migration completed successfully!")
}
