package schemas

import (
	"time"

	"github.com/marcos-nsantos/adopet-backend/internal/entity"
)

type AdoptionResponse struct {
	ID        uint64    `json:"id"`
	TutorID   uint64    `json:"tutorId"`
	PetID     uint64    `json:"petId"`
	CreatedAt time.Time `json:"createdAt"`
}

func ToAdoptionResponse(adoption entity.Adoption) AdoptionResponse {
	return AdoptionResponse{
		ID:        adoption.ID,
		TutorID:   adoption.TutorID,
		PetID:     adoption.PetID,
		CreatedAt: adoption.CreatedAt,
	}
}
