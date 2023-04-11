package main

import (
	"log"
	"os"

	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/router"
)

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and JWT token
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
