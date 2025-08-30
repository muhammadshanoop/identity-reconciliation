package testutils

import (
	"log"

	"github.com/muhammadshanoop/identity-reconciliation/database"
	"github.com/muhammadshanoop/identity-reconciliation/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to test database: %v", err)
	}

	err = db.AutoMigrate(&models.Contact{})
	if err != nil {
		log.Fatalf("failed to migrate test database: %v", err)
	}
	database.DB = db
	return db
}
