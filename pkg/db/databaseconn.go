package db

import (
	"fmt"

	"github.com/skinkvi/cyrs/internal/config"
	"github.com/skinkvi/cyrs/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DataBase.Host,
		cfg.DataBase.User,
		cfg.DataBase.Password,
		cfg.DataBase.DBName,
		cfg.DataBase.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	// Автоматическая миграция
	err = db.AutoMigrate(&models.Book{},
		&models.Author{},
		&models.Reader{},
		&models.Loan{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate models: %w", err)
	}

	return db, nil
}
