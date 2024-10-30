package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title         string `gorm:"size:255;not null" json:"title"`
	AuthorID      uint   `json:"author_id"`
	PublishedDate string `json:"published_date"`
	ISBN          string `gorm:"size:13;not null" json:"isbn"`
	Available     bool   `json:"available"`
	Author        Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"author"`
}
