// @title           Library API
// @version         1.0
// @description     API для управления библиотекой
// @host            localhost:8080
// @BasePath        /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/skinkvi/cyrs/docs"
	"github.com/skinkvi/cyrs/internal/config"
	"github.com/skinkvi/cyrs/internal/handlers"
	routes "github.com/skinkvi/cyrs/internal/router"
	"github.com/skinkvi/cyrs/pkg/db"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title Library API
// @version 1.0
// @description API для управления библиотекой
// @host localhost:8080
// @BasePath /api
func main() {
	cfg := config.MustLoadBuPath("config/local.yaml")

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(cfg.Server.Port); err != nil {
		logger.Sugar().Fatalf("Failed to run server: ", err)
	}

}
