package schemas

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

type PetResponse struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsAdopt     bool   `json:"isAdopt"`
	Age         uint64 `json:"age"`
	Photo       string `json:"photo"`
	UF          string `json:"uf"`
	City        string `json:"city"`
	ShelterID   uint64 `json:"shelterId"`
}

func ToPetResponse(pet entity.Pet) PetResponse {
	return PetResponse{
		ID:          pet.ID,
		Name:        pet.Name,
		Description: pet.Description,
		IsAdopt:     pet.IsAdopted,
		Age:         pet.Age,
		Photo:       pet.Photo,
		UF:          pet.UF,
		City:        pet.City,
		ShelterID:   pet.ShelterID,
	}
}

type PetsResponse struct {
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
	Total int           `json:"total"`
	Pets  []PetResponse `json:"pets"`
}

func ToPetsResponse(pets []entity.Pet, page, limit, total int) PetsResponse {
	var petsResponse []PetResponse
	for _, pet := range pets {
		petsResponse = append(petsResponse, ToPetResponse(pet))
	}

	return PetsResponse{
		Page:  page,
		Limit: limit,
		Total: total,
		Pets:  petsResponse,
	}
}
