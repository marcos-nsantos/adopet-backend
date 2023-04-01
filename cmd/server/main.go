package main

import (
	"log"
	"os"

	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/router"
)

func main() {
	database.Init()
	database.Migrate()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := router.SetupRoutes()
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalln("Error running server: ", err)
	}
}
