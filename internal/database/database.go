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
	var err error
	var count uint8
	databaseURL := os.Getenv("DATABASE_URL")

	for {
		DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{
			SkipDefaultTransaction: true,
		})

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

func InitGoogleCloudSQL() {
	var err error
	databaseURL := fmt.Sprintf("user=%s password=%s database=%s host=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("INSTANCE_UNIX_SOCKET"))
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func Migrate() {
	DB.AutoMigrate(&entity.Tutor{}, &entity.Shelter{}, &entity.Pet{}, &entity.Adoption{})
}

func DropTables() {
	DB.Migrator().DropTable(&entity.Tutor{}, &entity.Shelter{}, &entity.Pet{}, &entity.Adoption{})
}
