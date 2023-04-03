package tutorhandler

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

type UserResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
	City  string `json:"city"`
	About string `json:"about"`
}

func toUserResponse(tutor entity.User) UserResponse {
	return UserResponse{
		ID:    tutor.ID,
		Name:  tutor.Name,
		Email: tutor.Email,
		Phone: tutor.Phone,
		Photo: tutor.Photo,
		City:  tutor.City,
		About: tutor.About,
	}
}

type UsersResponse struct {
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
	Total int            `json:"total"`
	Users []UserResponse `json:"users"`
}

func toUsersResponse(users []entity.User, page, limit, total int) UsersResponse {
	var usersResponse []UserResponse
	for _, tutor := range users {
		usersResponse = append(usersResponse, toUserResponse(tutor))
	}

	return UsersResponse{
		Page:  page,
		Limit: limit,
		Total: total,
		Users: usersResponse,
	}
}
