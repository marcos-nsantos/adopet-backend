package database

import (
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
)

func CreateTutor(tutor entity.User) (entity.User, error) {
	result := DB.Create(&tutor)
	return tutor, result.Error
}

func GetTutorByID(id uint64) (entity.User, error) {
	var tutor entity.User
	result := DB.Select("id", "name", "email", "type", "phone", "photo", "city", "about").First(&tutor, id)
	return tutor, result.Error
}

func GetAllTutors(page, limit int) ([]entity.User, int, error) {
	var tutors []entity.User
	var total int64

	DB.Model(&entity.User{}).Count(&total)

	offset := (page - 1) * limit
	result := DB.Select("id", "name", "email", "type", "phone", "photo", "city", "about").
		Where("type = ?", entity.Tutor).Limit(limit).Offset(offset).Find(&tutors)
	return tutors, int(total), result.Error
}

func UpdateTutor(tutor *entity.User) error {
	return DB.Model(&tutor).Omit("id", "password").Save(tutor).Error
}

func DeleteTutor(id uint64) error {
	return DB.Delete(&entity.User{}, id).Error
}
