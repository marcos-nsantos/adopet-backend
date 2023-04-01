package database

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func CreateTutor(tutor *entity.Tutor) error {
	return DB.Create(tutor).Error
}
