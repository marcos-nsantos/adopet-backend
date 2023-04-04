package database

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func CreatePet(pet entity.Pet) (entity.Pet, error) {
	result := DB.Create(&pet)
	return pet, result.Error
}
