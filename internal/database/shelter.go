package database

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func CreateShelter(shelter entity.Shelter) (entity.Shelter, error) {
	result := DB.Create(&shelter)
	return shelter, result.Error
}

func GetShelterByID(id uint64) (entity.Shelter, error) {
	var shelter entity.Shelter
	result := DB.Select("id", "name", "email", "phone", "photo", "city", "about").First(&shelter, id)
	return shelter, result.Error
}

func GetAllShelters(page, limit int) ([]entity.Shelter, int, error) {
	var shelters []entity.Shelter
	var total int64

	DB.Model(&entity.Shelter{}).Count(&total)

	offset := (page - 1) * limit
	result := DB.Select("id", "name", "email", "phone", "photo", "city", "about").
		Limit(limit).Offset(offset).Find(&shelters)
	return shelters, int(total), result.Error
}

func UpdateShelter(shelter *entity.Shelter) error {
	result := DB.Model(&shelter).Omit("id", "password").Updates(shelter)
	return result.Error
}

func DeleteShelter(id uint64) error {
	result := DB.Delete(&entity.Shelter{}, id)
	return result.Error
}

func GetIDAndPasswordByEmailFromShelter(email string) (uint64, string, error) {
	var shelter entity.Shelter
	result := DB.Select("id", "password").First(&shelter, "email = ?", email)
	return shelter.ID, shelter.Password, result.Error
}
