package mock

import "github.com/marcos-nsantos/adopet-backend/internal/entity"

func Tutors() []entity.Tutor {
	return []entity.Tutor{
		{
			ID:       1,
			Name:     "Tutor One",
			Email:    "tutorone@email.com",
			Password: "secretPassword",
			Phone:    "123456789",
			Photo:    "https://www.google.com",
			City:     "São Paulo",
			About:    "Hello, I'm Tutor One",
		},
		{
			ID:       2,
			Name:     "Tutor Two",
			Email:    "tutortwo@email.com",
			Password: "secretPassword",
			Phone:    "123456789",
			Photo:    "https://go.dev",
			City:     "São Paulo",
			About:    "Hello, I'm Tutor Two",
		},
		{
			ID:       3,
			Name:     "Tutor Three",
			Email:    "tutorthree@email.com",
			Password: "secretPassword",
			Phone:    "123456789",
			Photo:    "https://pkg.go.dev",
			City:     "São Paulo",
			About:    "Hello, I'm Tutor Three",
		},
		{
			ID:       4,
			Name:     "Tutor Four",
			Email:    "tutorfour@email.com",
			Password: "secretPassword",
			Phone:    "123456789",
			Photo:    "https://go.dev/play/",
			City:     "São Paulo",
			About:    "Hello, I'm Tutor Four",
		},
		{
			ID:       5,
			Name:     "Tutor Five",
			Email:    "tutorfive@email.com",
			Password: "secretPassword",
			Phone:    "123456789",
			Photo:    "https://stackoverflow.com/",
			City:     "São Paulo",
			About:    "Hello, I'm Tutor Five",
		},
	}
}
