package models

import (
	"gorm.io/gorm"
)

// Author представляет автора в системе
// swagger:model Author
type Author struct {
	// Встроенная структура GORM
	gorm.Model `json:"-"` // игнорируем встроенную структуру для Swagger

	// Импортируем поля из GormModel для Swagger
	// @model models.GormModel

	// Name представляет имя автора
	// example: Лев Толстой
	Name string `gorm:"size:255;not null" json:"name"`

	// BirthDate представляет дату рождения автора
	// example: 1828-09-09
	BirthDate string `json:"birth_date"`

	// Nationality представляет национальность автора
	// example: Русский
	Nationality string `gorm:"size:255" json:"nationality"`
}
