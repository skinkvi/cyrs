package models

import "gorm.io/gorm"

type Loan struct {
	gorm.Model
	ID         uint   `json:"id"`
	BookID     uint   `json:"book_id"`
	Book       Book   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"book"`
	ReaderID   uint   `json:"reader_id"`
	Reader     Reader `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reader"`
	LoanDate   string `json:"loan_date"`
	ReturnDate string `json:"return_date"`
}
