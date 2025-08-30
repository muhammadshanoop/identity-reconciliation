package database

import (
	"log"
	"sync"
	"time"

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
			for i := 0; i < 5; i++ {
				db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
				if err == nil {
					log.Println("✅ Connected to database successfully!")
					break
				}
				log.Printf("⏳ Failed to connect (attempt %d/5): %v", i+1, err)
				time.Sleep(2 * time.Second)
			}
		}
		log.Println("Successfully connected to the database!")
		DB = db
	})
}
