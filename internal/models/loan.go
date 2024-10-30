package models

import "gorm.io/gorm"

type Loan struct {
	gorm.Model
	BookID     uint
	Book       Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReaderID   uint
	Reader     Reader `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LoanDate   string
	ReturnDate string
}
