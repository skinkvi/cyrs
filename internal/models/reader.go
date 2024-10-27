package models

import "gorm.io/gorm"

type Reader struct {
	gorm.Model
	Name           string `gorm:"size:255;not null"`
	Email          string `gorm:"size:255;not null"`
	MembershipDate string
}

