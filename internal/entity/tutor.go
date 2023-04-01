package entity

import "gorm.io/gorm"

type Tutor struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"not null"`
	Phone    string `gorm:"type:varchar(15)"`
	Photo    string
	City     string `gorm:"type:varchar(255)"`
	About    string
}
