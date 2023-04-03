package entity

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	ID          uint64 `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"not null"`
	IsAdopt     bool   `gorm:"not null"`
	Age         uint64 `gorm:"not null"`
	Photo       string `gorm:"not null"`
	UF          string `gorm:"type:varchar(2);not null"`
	City        string `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
