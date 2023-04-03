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

func toTutorResponse(tutor entity.User) TutorResponse {
	return TutorResponse{
		ID:    tutor.ID,
		Name:  tutor.Name,
		Email: tutor.Email,
		Phone: tutor.Phone,
		Photo: tutor.Photo,
		City:  tutor.City,
		About: tutor.About,
	}
}

type TutorsResponse struct {
	Page   int             `json:"page"`
	Limit  int             `json:"limit"`
	Total  int             `json:"total"`
	Tutors []TutorResponse `json:"tutors"`
}

func toTutorsResponse(tutors []entity.User, page, limit, total int) TutorsResponse {
	var tutorsResponse []TutorResponse
	for _, tutor := range tutors {
		tutorsResponse = append(tutorsResponse, toTutorResponse(tutor))
	}

	return TutorsResponse{
		Page:   page,
		Limit:  limit,
		Total:  total,
		Tutors: tutorsResponse,
	}
}
