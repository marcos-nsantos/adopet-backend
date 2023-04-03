package mock

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func Pet() []entity.Pet {
	return []entity.Pet{
		{
			Name:        "Spike",
			Description: "Spike is a very friendly dog. He loves to play and is very energetic. He is very good with children and other dogs. He is a very good dog.",
			IsAdopt:     false,
			Age:         8,
			Photo:       "https://some-dog-photo.jpg",
			UF:          "SP",
			City:        "São Paulo",
		},
		{
			Name:        "Luna",
			Description: "A playful and energetic Labrador Retriever.",
			IsAdopt:     false,
			Age:         5,
			Photo:       "https://some-dog-photo.jpg",
			UF:          "MG",
			City:        "Belo Horizonte",
		},
		{
			Name:        "Buddy",
			Description: "A friendly and loyal Golden Retriever.",
			IsAdopt:     true,
			Age:         7,
			Photo:       "https://some-dog-photo.jpg",
			UF:          "RJ",
			City:        "Rio de Janeiro",
		},
		{
			Name:        "Simba",
			Description: "A majestic and independent Persian cat.",
			IsAdopt:     true,
			Age:         3,
			Photo:       "https://some-cat-photo.jpg",
			UF:          "AC",
			City:        "Rio Branco",
		},
		{
			Name:        "Max",
			Description: "A curious and adventurous Beagle.",
			IsAdopt:     false,
			Age:         2,
			Photo:       "https://some-dog-photo.jpg",
			UF:          "AC",
			City:        "Rio Branco",
		},
	}
}
