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
