package database

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func CreateUser(user entity.User) (entity.User, error) {
	result := DB.Create(&user)
	return user, result.Error
}
