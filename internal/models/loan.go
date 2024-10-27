package models

import "gorm.io/gorm"

type Loan struct {
	gorm.Model
	BookID     uint
	ReaderID   uint
	LoanDate   string
	ReturnDate string
}

