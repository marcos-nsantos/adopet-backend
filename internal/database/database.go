package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/marcos-nsantos/adopet-backend/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	databaseURL := os.Getenv("DATABASE_URL")
	var count uint8

	for {
		var err error
		DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
		if err != nil {
			fmt.Println("Failed to connect to database")
			fmt.Println("Retrying in 5 seconds")
			time.Sleep(5 * time.Second)
			count++

			if count == 5 {
				log.Println("Failed to connect to database")
				log.Fatal(err)
			}
		}
		break
	}
}

func Migrate() {
	DB.AutoMigrate(&entity.User{}, &entity.Pet{})
}

func DropTables() {
	DB.Migrator().DropTable(&entity.User{}, &entity.Pet{})
}
