package handlers

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewHandler(db *gorm.DB, logger *zap.Logger) *Handler {
	return &Handler{
		DB:     db,
		Logger: logger,
	}
}
