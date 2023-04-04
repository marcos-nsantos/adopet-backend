package database

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func CreatePet(pet entity.Pet) (entity.Pet, error) {
	result := DB.Create(&pet)
	return pet, result.Error
}

func GetPetByID(id uint64) (entity.Pet, error) {
	var pet entity.Pet
	result := DB.Select("id", "name", "description", "is_adopt", "age", "photo", "uf", "city").First(&pet, id)
	return pet, result.Error
}
