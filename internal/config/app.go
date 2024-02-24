package config

import (
	"movie-technical-test/internal/delivery/http/handler"
	"movie-technical-test/internal/delivery/http/middleware"
	"movie-technical-test/internal/delivery/http/router"
	"movie-technical-test/internal/repository"
	"movie-technical-test/internal/usecase"
	"movie-technical-test/internal/utils/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppConfig struct {
	FiberApp *fiber.App
	Log      *logger.Logger
	DB       *gorm.DB
	Validate *validator.Validate
	Config   *viper.Viper
}

func Container(config *AppConfig) {
	// setup repositories
	movieRepo := repository.NewMovieRepository()

	// setup use cases
	moviewUseCase := usecase.NewMovieUseCase(config.DB, config.Log, config.Validate, movieRepo)

	// setup handler
	movieHandler := handler.NewMovieHandler(moviewUseCase)

	// setup middleware
	loggingMiddleware := middleware.HandleReqLogging(config.Log)

	route := router.Route{
		App:           config.FiberApp,
		LogMiddleware: loggingMiddleware,
		MovieHandler:  movieHandler,
	}
	route.Setup()
}
