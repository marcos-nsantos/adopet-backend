package mock

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func Shelters() []entity.Shelter {
	return []entity.Shelter{
		{
			ID:       1,
			Name:     "Shelter One",
			Email:    "shelterone@email.com",
			Password: "secretPassword",
			Phone:    "123456781",
			Photo:    "https://go.dev/blog/gopher/header.jpg",
			City:     "São Paulo",
			About:    "It's shelter one",
		},
		{
			ID:       2,
			Name:     "Shelter Two",
			Email:    "sheltertwo@email.com",
			Password: "secretPassword",
			Phone:    "123456782",
			City:     "São Paulo",
			About:    "It's shelter two",
		},
		{
			ID:       3,
			Name:     "Shelter Three",
			Email:    "shelterthree@email.com",
			Password: "secretPassword",
			Phone:    "123456783",
			Photo:    "https://go.dev/blog/gopher/header.jpg",
			City:     "São Paulo",
			About:    "It's shelter three",
		},
		{
			ID:       4,
			Name:     "Shelter Four",
			Email:    "shelterfour@email.com",
			Password: "secretPassword",
			Phone:    "123456784",
			Photo:    "https://go.dev/blog/gopher/header.jpg",
			City:     "São Paulo",
			About:    "It's shelter four",
		},
		{
			ID:       5,
			Name:     "Shelter Five",
			Email:    "shelterfive@email.com",
			Password: "secretPassword",
			Phone:    "123456785",
			Photo:    "https://go.dev/blog/gopher/header.jpg",
			City:     "São Paulo",
			About:    "It's shelter five",
		},
	}
}
