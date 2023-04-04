package schemas

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

type PetResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsAdopt     bool   `json:"isAdopt"`
	Age         uint64 `json:"age"`
	Photo       string `json:"photo"`
	UF          string `json:"uf"`
	City        string `json:"city"`
}

func ToPetResponse(pet entity.Pet) PetResponse {
	return PetResponse{
		Name:        pet.Name,
		Description: pet.Description,
		IsAdopt:     pet.IsAdopt,
		Age:         pet.Age,
		Photo:       pet.Photo,
		UF:          pet.UF,
		City:        pet.City,
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
