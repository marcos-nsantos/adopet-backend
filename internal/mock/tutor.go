package mock

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func Tutors() []entity.Tutor {
	return []entity.Tutor{
		{
			ID:       1,
			Name:     "User One",
			Email:    "tutorone@email.com",
			Password: "secretPassword",
			Phone:    "123456789",
			Photo:    "https://www.google.com",
			City:     "São Paulo",
			About:    "Hello, I'm User One",
		},
		{
			ID:       2,
			Name:     "User Two",
			Email:    "tutortwo@email.com",
			Password: "secretPassword",
			Phone:    "123456789",
			Photo:    "https://go.dev",
			City:     "São Paulo",
			About:    "Hello, I'm User Two",
		},
		{
			ID:       3,
			Name:     "User Three",
			Email:    "tutorthree@email.com",
			Password: "secretPassword",
			Phone:    "123456789",
			Photo:    "https://pkg.go.dev",
			City:     "São Paulo",
			About:    "Hello, I'm User Three",
		},
		{
			ID:       4,
			Name:     "User Four",
			Email:    "tutorfour@email.com",
			Password: "secretPassword",
			Phone:    "123456789",
			Photo:    "https://go.dev/play/",
			City:     "São Paulo",
			About:    "Hello, I'm User Four",
		},
		{
			ID:       5,
			Name:     "User Five",
			Email:    "tutorfive@email.com",
			Password: "secretPassword",
			Phone:    "123456789",
			Photo:    "https://stackoverflow.com/",
			City:     "São Paulo",
			About:    "Hello, I'm User Five",
		},
	}
}
