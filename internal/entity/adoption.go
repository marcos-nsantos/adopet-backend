package entity

import (
	"time"

	"gorm.io/gorm"
)

type Adoption struct {
	ID        uint64 `gorm:"primaryKey"`
	TutorID   uint64
	Tutor     Tutor
	PetID     uint64
	Pet       Pet
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
