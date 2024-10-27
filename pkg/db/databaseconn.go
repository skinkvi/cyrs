package db

import (
	"fmt"

	"github.com/skinkvi/cyrs/internal/config"
	"github.com/skinkvi/cyrs/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg config.Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DataBase.Host,
		cfg.DataBase.User,
		cfg.DataBase.Password,
		cfg.DataBase.User,
		cfg.DataBase.Port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		cfg.Logger.Sugar().Fatalf("Failed to connect to db: %v", err)
	}

	err = DB.AutoMigrate(&models.Book{},
		&models.Author{},
		&models.Reader{},
		&models.Loan{})
	if err != nil {
		cfg.Logger.Sugar().Fatalf("Failed to migrate models: %v", err)
	}

}
