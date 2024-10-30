package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name        string `gorm:"size:255;not null" json:"name"`
	BirthDate   string `json:"birth_date"`
	Nationality string `gorm:"size:255" json:"nationality"`
}
