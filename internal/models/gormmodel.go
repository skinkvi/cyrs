package models

import (
	"time"
)

// GormModel описывает поля из gorm.Model для Swagger
// swagger:model GormModel
type GormModel struct {
	// ID уникальный идентификатор
	// example: 1
	ID uint `json:"id"`

	// CreatedAt время создания записи
	// example: 2023-10-05T14:48:00Z
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt время обновления записи
	// example: 2023-10-05T14:48:00Z
	UpdatedAt time.Time `json:"updated_at"`

	// DeletedAt время удаления записи
	// example: null
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
