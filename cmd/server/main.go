package main

import "github.com/marcos-nsantos/adopet-backend/internal/database"

func main() {
	database.Init()
	database.Migrate()
}
