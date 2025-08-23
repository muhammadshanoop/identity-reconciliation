package database

import (
	"log"
	"sync"

	"github.com/muhammadshanoop/identity-reconciliation/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func GetConnect() {
	once.Do(func() {
		dsn := config.GetDSN()
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}
		log.Println("Successfully connected to the database!")
		DB = db
	})
}
