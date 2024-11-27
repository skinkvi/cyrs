package models

import "gorm.io/gorm"

// Reader представляет читателя в системе
// swagger:model Reader
type Reader struct {
	// Встроенная структура GORM
	gorm.Model `json:"-"`

	// Импортируем поля из GormModel для Swagger
	// @model docs.GormModel

	// Name представляет имя читателя
	// example: Иван Иванов
	Name string `gorm:"size:255;not null" json:"name"`

	// Email представляет электронную почту читателя
	// example: ivan.ivanov@example.com
	Email string `gorm:"size:255;not null" json:"email"`

	// MembershipDate представляет дату вступления читателя в библиотеку
	// example: 2020-01-15
	MembershipDate string `json:"membership_date"`
}
