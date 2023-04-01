package tutorhandler

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

type TutorCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone"`
	Photo    string `json:"photo"`
	City     string `json:"city"`
	About    string `json:"about"`
}

func (t *TutorCreateRequest) ToEntity() entity.Tutor {
	return entity.Tutor{
		Name:     t.Name,
		Email:    t.Email,
		Password: t.Password,
		Phone:    t.Phone,
		Photo:    t.Photo,
		City:     t.City,
		About:    t.About,
	}
}
