package models

import "gorm.io/gorm"

// Book представляет книгу в системе
// swagger:model Book
type Book struct {
	// Встроенная структура GORM
	gorm.Model `json:"-"`

	// Импортируем поля из GormModel для Swagger
	// @model models.GormModel

	// Title представляет название книги
	// example: Война и мир
	Title string `gorm:"size:255;not null" json:"title"`

	// AuthorID представляет ID автора книги
	// example: 1
	AuthorID uint `json:"author_id"`

	// PublishedDate представляет дату публикации книги
	// example: 1869-01-01
	PublishedDate string `json:"published_date"`

	// ISBN представляет международный стандартный книжный номер
	// example: 978-3-16-148410-0
	ISBN string `gorm:"size:255;not null" json:"isbn"`

	// Available указывает, доступна ли книга
	// example: true
	Available bool `json:"available"`

	// Author представляет автора книги
	Author Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"author"`
}
