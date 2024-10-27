package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title         string `gorm:"size:255;not null"`
	AuthorID      uint
	PublishedDate string
	ISBN          string `gorm:"size:13;not null"`
	Available     bool
}

