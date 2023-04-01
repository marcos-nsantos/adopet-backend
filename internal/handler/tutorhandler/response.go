package tutorhandler

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

type TutorResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
	City  string `json:"city"`
	About string `json:"about"`
}

func toTutorResponse(tutor *entity.Tutor) *TutorResponse {
	return &TutorResponse{
		ID:    tutor.ID,
		Name:  tutor.Name,
		Email: tutor.Email,
		Phone: tutor.Phone,
		Photo: tutor.Photo,
		City:  tutor.City,
		About: tutor.About,
	}
}
