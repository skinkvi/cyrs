package models

import "gorm.io/gorm"

type Reader struct {
	gorm.Model
	Name           string `gorm:"size:255;not null" json:"name"`
	Email          string `gorm:"size:255;not null" json:"email"`
	MembershipDate string `json:"membership_date"`
}
