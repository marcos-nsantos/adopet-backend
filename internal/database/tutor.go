package database

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func CreateTutor(tutor *entity.Tutor) error {
	return DB.Create(tutor).Error
}

func GetTutorByID(id uint64) (entity.Tutor, error) {
	var tutor entity.Tutor
	result := DB.Omit("password", "deleted_at").First(&tutor, id)
	return tutor, result.Error
}

func GetAllTutors(page, limit int) ([]entity.Tutor, int, error) {
	var tutors []entity.Tutor
	var total int64

	DB.Model(&entity.Tutor{}).Count(&total)

	offset := (page - 1) * limit
	result := DB.Omit("password", "deleted_at").Limit(limit).Offset(offset).Find(&tutors)
	return tutors, int(total), result.Error
}

func UpdateTutor(tutor *entity.Tutor) error {
	return DB.Model(&tutor).Omit("id", "password").Save(tutor).Error
}

func DeleteTutor(id uint64) error {
	return DB.Delete(&entity.Tutor{}, id).Error
}
