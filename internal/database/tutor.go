package database

import (
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
	"gorm.io/gorm"
)

func CreateTutor(tutor entity.Tutor) (entity.Tutor, error) {
	result := DB.Create(&tutor)
	return tutor, result.Error
}

func GetTutorByID(id uint64) (entity.Tutor, error) {
	var tutor entity.Tutor

	result := DB.Select("id", "name", "email", "phone", "photo", "city", "about").First(&tutor, id)
	if result.Error == gorm.ErrRecordNotFound {
		return tutor, entity.ErrTutorNotFound
	}

	return tutor, result.Error
}

func GetAllTutors(page, limit int) ([]entity.Tutor, int, error) {
	var tutors []entity.Tutor
	var total int64

	DB.Model(&entity.Tutor{}).Count(&total)

	offset := (page - 1) * limit
	result := DB.Select("id", "name", "email", "phone", "photo", "city", "about").
		Limit(limit).Offset(offset).Find(&tutors)

	return tutors, int(total), result.Error
}

func UpdateTutor(tutor *entity.Tutor) error {
	result := DB.Model(&tutor).Omit("id", "password").Updates(tutor)
	if result.RowsAffected == 0 {
		return entity.ErrTutorNotFound
	}

	return result.Error
}

func DeleteTutor(id uint64) error {
	result := DB.Delete(&entity.Tutor{}, id)
	if result.RowsAffected == 0 {
		return entity.ErrTutorNotFound
	}

	return result.Error
}

func GetIDAndPasswordByEmailFromTutor(email string) (uint64, string, error) {
	var tutor entity.Tutor
	result := DB.Select("id", "password").First(&tutor, "email = ?", email)
	return tutor.ID, tutor.Password, result.Error
}
