package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name        string `gorm:"size:255;not null"`
	BirthDate   string
	Nationality string `gorm:"size:255"`
}

