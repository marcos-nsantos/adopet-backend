package database

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func GetShelterByID(id uint64) (entity.User, error) {
	var shelter entity.User
	result := DB.Select("id", "name", "email", "phone", "photo", "city", "about").Where("type = ?", entity.Shelter).First(&shelter, id)
	return shelter, result.Error
}

func GetAllShelters(page, limit int) ([]entity.User, int, error) {
	var shelters []entity.User
	var total int64

	DB.Model(&entity.User{}).Count(&total)

	offset := (page - 1) * limit
	result := DB.Select("id", "name", "email", "phone", "photo", "city", "about").
		Where("type = ?", entity.Shelter).Limit(limit).Offset(offset).Find(&shelters)
	return shelters, int(total), result.Error
}
