package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skinkvi/cyrs/internal/config"
	"github.com/skinkvi/cyrs/internal/handlers"
	"github.com/skinkvi/cyrs/internal/router"
	"github.com/skinkvi/cyrs/pkg/db"
	"go.uber.org/zap"
)

func main() {
	cfg := config.MustLoadBuPath("../../config/local.yaml")

	logger, err := zap.NewProduction()
	if err != nil {
		logger.Sugar().Fatalf("Failed to initialize logger: ", err)
	}
	defer logger.Sync()

	dbInstance, err := db.InitDB(*cfg)
	if err != nil {
		logger.Sugar().Fatalf("Database initialization failed: ", err)
	}

	h := handlers.NewHandler(dbInstance, logger)

	router := gin.Default()

	routes.SetupRoutes(router, h)

	if err := router.Run(cfg.Server.Port); err != nil {
		logger.Sugar().Fatalf("Failed to run server: ", err)
	}

}
