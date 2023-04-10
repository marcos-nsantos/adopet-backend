package schemas

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

type ShelterResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
	City  string `json:"city"`
	About string `json:"about"`
}

func ToShelterResponse(tutor entity.Shelter) ShelterResponse {
	return ShelterResponse{
		ID:    tutor.ID,
		Name:  tutor.Name,
		Email: tutor.Email,
		Phone: tutor.Phone,
		Photo: tutor.Photo,
		City:  tutor.City,
		About: tutor.About,
	}
}

type SheltersResponse struct {
	Page     int               `json:"page"`
	Limit    int               `json:"limit"`
	Total    int               `json:"total"`
	Shelters []ShelterResponse `json:"shelters"`
}

func ToSheltersResponse(users []entity.Shelter, page, limit, total int) SheltersResponse {
	var sheltersResponse []ShelterResponse
	for _, tutor := range users {
		sheltersResponse = append(sheltersResponse, ToShelterResponse(tutor))
	}

	return SheltersResponse{
		Page:     page,
		Limit:    limit,
		Total:    total,
		Shelters: sheltersResponse,
	}
}
