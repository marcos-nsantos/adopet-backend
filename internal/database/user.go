package database

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func CreateUser(user entity.User) (entity.User, error) {
	result := DB.Create(&user)
	return user, result.Error
}

func UpdateUser(tutor *entity.User) error {
	return DB.Model(&tutor).Omit("id", "password").Updates(tutor).Error
}

func DeleteUser(id uint64) error {
	return DB.Delete(&entity.User{}, id).Error
}
