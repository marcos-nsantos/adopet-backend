package schemas

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

type PetCreateRequests struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsAdopted   bool   `json:"isAdopted"`
	Age         uint64 `json:"age" binding:"required"`
	Photo       string `json:"photo" binding:"required,uri"`
	UF          string `json:"uf" binding:"required"`
	City        string `json:"city" binding:"required"`
	ShelterID   uint64 `json:"shelterId" binding:"required"`
}

func (p *PetCreateRequests) ToEntity() entity.Pet {
	return entity.Pet{
		Name:        p.Name,
		Description: p.Description,
		IsAdopted:   p.IsAdopted,
		Age:         p.Age,
		Photo:       p.Photo,
		UF:          p.UF,
		City:        p.City,
		ShelterID:   p.ShelterID,
	}
}

type PetUpdateRequests struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Age         uint64 `json:"age" binding:"required"`
	Photo       string `json:"photo" binding:"required,uri"`
	UF          string `json:"uf" binding:"required"`
	City        string `json:"city" binding:"required"`
}

func (p *PetUpdateRequests) ToEntity() entity.Pet {
	return entity.Pet{
		Name:        p.Name,
		Description: p.Description,
		Age:         p.Age,
		Photo:       p.Photo,
		UF:          p.UF,
		City:        p.City,
	}
}

type UpdateIsAdoptPetRequests struct {
	IsAdopted bool `json:"isAdopted" binding:"required"`
}
