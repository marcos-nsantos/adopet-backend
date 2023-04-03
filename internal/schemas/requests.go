package schemas

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone"`
	Photo    string `json:"photo"`
	City     string `json:"city"`
	About    string `json:"about"`
}

func (r *UserCreateRequest) ToEntity() entity.User {
	return entity.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		Phone:    r.Phone,
		Photo:    r.Photo,
		City:     r.City,
		About:    r.About,
	}
}

type UserUpdateRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
	City  string `json:"city"`
	About string `json:"about"`
}

func (t *UserUpdateRequest) ToEntity() entity.User {
	return entity.User{
		Name:  t.Name,
		Email: t.Email,
		Phone: t.Phone,
		Photo: t.Photo,
		City:  t.City,
		About: t.About,
	}
}
