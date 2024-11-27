package models

import "gorm.io/gorm"

// Loan представляет заем в системе
// swagger:model Loan
type Loan struct {
	// Встроенная структура GORM
	gorm.Model `json:"-"`

	// Импортируем поля из GormModel для Swagger
	// @model docs.GormModel

	// BookID представляет ID книги, которую берут в заем
	// example: 123
	BookID uint `json:"book_id"`

	// Book представляет информацию о книге
	Book Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"book"`

	// ReaderID представляет ID читателя, берущего книгу
	// example: 456
	ReaderID uint `json:"reader_id"`

	// Reader представляет информацию о читателе
	Reader Reader `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reader"`

	// LoanDate представляет дату выдачи книги
	// example: 2024-04-01
	LoanDate string `json:"loan_date"`

	// ReturnDate представляет дату возврата книги
	// example: 2024-04-15
	ReturnDate string `json:"return_date"`
}
