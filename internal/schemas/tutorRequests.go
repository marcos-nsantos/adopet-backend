package schemas

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

type TutorCreationRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone"`
	Photo    string `json:"photo"`
	City     string `json:"city"`
	About    string `json:"about"`
}

func (r *TutorCreationRequest) ToEntity() entity.Tutor {
	return entity.Tutor{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		Phone:    r.Phone,
		Photo:    r.Photo,
		City:     r.City,
		About:    r.About,
	}
}

type TutorUpdateRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone" binding:"required"`
	Photo string `json:"photo" binding:"required"`
	City  string `json:"city" binding:"required"`
	About string `json:"about" binding:"required"`
}

func (t *TutorUpdateRequest) ToEntity() entity.Tutor {
	return entity.Tutor{
		Name:  t.Name,
		Email: t.Email,
		Phone: t.Phone,
		Photo: t.Photo,
		City:  t.City,
		About: t.About,
	}
}
