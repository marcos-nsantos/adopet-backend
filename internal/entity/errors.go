package entity

import "errors"

var (
	ErrPetNotFound     = errors.New("pet not found")
	ErrShelterNotFound = errors.New("shelter not found")
	ErrTutorNotFound   = errors.New("tutor not found")
)
