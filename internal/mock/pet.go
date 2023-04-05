package mock

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func Pet() []entity.Pet {
	return []entity.Pet{
		{
			Name:        "Spike",
			Description: "Spike is a very friendly dog. He loves to play and is very energetic. He is very good with children and other dogs. He is a very good dog.",
			IsAdopted:   false,
			Age:         8,
			Photo:       "https://some-dog-photo.jpg",
			UF:          "SP",
			City:        "São Paulo",
			UserID:      1,
		},
		{
			Name:        "Luna",
			Description: "A playful and energetic Labrador Retriever.",
			IsAdopted:   false,
			Age:         5,
			Photo:       "https://some-dog-photo.jpg",
			UF:          "MG",
			City:        "Belo Horizonte",
			UserID:      1,
		},
		{
			Name:        "Buddy",
			Description: "A friendly and loyal Golden Retriever.",
			IsAdopted:   true,
			Age:         7,
			Photo:       "https://some-dog-photo.jpg",
			UF:          "RJ",
			City:        "Rio de Janeiro",
			UserID:      1,
		},
		{
			Name:        "Simba",
			Description: "A majestic and independent Persian cat.",
			IsAdopted:   true,
			Age:         3,
			Photo:       "https://some-cat-photo.jpg",
			UF:          "AC",
			City:        "Rio Branco",
			UserID:      1,
		},
		{
			Name:        "Max",
			Description: "A curious and adventurous Beagle.",
			IsAdopted:   false,
			Age:         2,
			Photo:       "https://some-dog-photo.jpg",
			UF:          "AC",
			City:        "Rio Branco",
			UserID:      1,
		},
	}
}
