package database

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func GetShelterByID(id uint64) (entity.User, error) {
	var shelter entity.User
	result := DB.Select("id", "name", "email", "phone", "photo", "city", "about").Where("type = ?", entity.Shelter).First(&shelter, id)
	return shelter, result.Error
}
