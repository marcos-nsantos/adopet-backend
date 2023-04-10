package entity

import (
	"time"

	"gorm.io/gorm"
)

type Shelter struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"not null"`
	Phone     string `gorm:"type:varchar(15)"`
	Photo     string
	City      string `gorm:"type:varchar(255)"`
	About     string
	Pet       []Pet
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
