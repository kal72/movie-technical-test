package usecase

import (
	"context"
	"movie-technical-test/internal/entity"
	"movie-technical-test/internal/model"
	"movie-technical-test/internal/repository"
	"movie-technical-test/internal/utils/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MovieUseCase struct {
	DB              *gorm.DB
	Log             *logger.Logger
	Validate        *validator.Validate
	MovieRepository *repository.MovieRepository
}

func NewMovieUseCase(db *gorm.DB, log *logger.Logger, validate *validator.Validate, movieRepository *repository.MovieRepository) *MovieUseCase {
	return &MovieUseCase{
		DB:              db,
		Log:             log,
		Validate:        validate,
		MovieRepository: movieRepository,
	}
}

func (u *MovieUseCase) AddNew(ctx context.Context, request model.MovieCreateRequest) (fiber.Map, error) {
	tx := u.DB.WithContext(ctx)
	var err error

	err = u.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	newMovie := entity.Movie{
		Title:       request.Title,
		Description: request.Description,
		Rating:      request.Rating,
		Image:       request.Image,
	}

	err = u.MovieRepository.Create(tx, &newMovie)
	if err != nil {
		u.Log.Error(ctx, err, "failed insert new movie")
		return nil, fiber.ErrInternalServerError
	}

	response := fiber.Map{"id": newMovie.ID}
	return response, nil
}