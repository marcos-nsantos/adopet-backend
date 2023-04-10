package database

import (
	"github.com/marcos-nsantos/adopet-backend/internal/entity"
	"gorm.io/gorm"
)

func Adopt(adoption *entity.Adoption) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(adoption)
		if result.Error != nil {
			return result.Error
		}

		result = tx.Model(&entity.Pet{}).Where("id = ?", adoption.PetID).Update("is_adopted", true)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	return err
}
