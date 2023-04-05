package database

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func CreatePet(pet entity.Pet) (entity.Pet, error) {
	result := DB.Create(&pet)
	return pet, result.Error
}

func GetPetByID(id uint64) (entity.Pet, error) {
	var pet entity.Pet
	result := DB.Select("id", "name", "description", "is_adopted", "age", "photo", "uf", "city", "user_id").
		Where("is_adopted = ?", false).First(&pet, id)
	return pet, result.Error
}

func GetAllPets(page, limit int) ([]entity.Pet, int, error) {
	var pets []entity.Pet
	var total int64

	DB.Model(&entity.Pet{}).Where("is_adopted = ?", false).Count(&total)

	offset := (page - 1) * limit
	result := DB.Select("id", "name", "description", "is_adopted", "age", "photo", "uf", "city", "user_id").
		Where("is_adopted = ?", false).Limit(limit).Offset(offset).Find(&pets)
	return pets, int(total), result.Error
}

func UpdatePet(pet entity.Pet) error {
	result := DB.Model(&pet).Select("name", "description", "is_adopted", "age", "photo", "uf", "city").Updates(pet)
	return result.Error
}

func DeletePet(id uint64) error {
	return DB.Delete(&entity.Pet{}, id).Error
}
